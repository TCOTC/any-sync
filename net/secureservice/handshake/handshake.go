package handshake

import (
	"encoding/binary"
	"errors"
	"github.com/anytypeio/any-sync/net/secureservice/handshake/handshakeproto"
	"github.com/libp2p/go-libp2p/core/sec"
	"golang.org/x/exp/slices"
	"io"
	"sync"
)

const headerSize = 5 // 1 byte for type + 4 byte for uint32 size

const (
	msgTypeCred = byte(1)
	msgTypeAck  = byte(2)
)

type handshakeError struct {
	e handshakeproto.Error
}

func (he handshakeError) Error() string {
	return he.e.String()
}

var (
	ErrUnexpectedPayload       = handshakeError{handshakeproto.Error_UnexpectedPayload}
	ErrDeadlineExceeded        = handshakeError{handshakeproto.Error_DeadlineExceeded}
	ErrInvalidCredentials      = handshakeError{handshakeproto.Error_InvalidCredentials}
	ErrPeerDeclinedCredentials = errors.New("remote peer declined the credentials")
	ErrSkipVerifyNotAllowed    = handshakeError{handshakeproto.Error_SkipVerifyNotAllowed}
	ErrUnexpected              = handshakeError{handshakeproto.Error_Unexpected}

	ErrGotNotAHandshakeMessage = errors.New("go not a handshake message")
)

var handshakePool = &sync.Pool{New: func() any {
	return &handshake{
		remoteCred: &handshakeproto.Credentials{},
		remoteAck:  &handshakeproto.Ack{},
		localAck:   &handshakeproto.Ack{},
		buf:        make([]byte, 0, 1024),
	}
}}

type CredentialChecker interface {
	MakeCredentials(sc sec.SecureConn) *handshakeproto.Credentials
	CheckCredential(sc sec.SecureConn, cred *handshakeproto.Credentials) (err error)
}

func OutgoingHandshake(sc sec.SecureConn, cc CredentialChecker) (err error) {
	h := newHandshake()
	defer h.release()
	h.conn = sc
	localCred := cc.MakeCredentials(sc)
	if err = h.writeCredentials(localCred); err != nil {
		h.tryWriteErrAndClose(err)
		return err
	}
	msg, err := h.readMsg()
	if err != nil {
		h.tryWriteErrAndClose(err)
		return err
	}
	if msg.ack != nil {
		if msg.ack.Error == handshakeproto.Error_InvalidCredentials {
			return ErrPeerDeclinedCredentials
		}
		return handshakeError{e: msg.ack.Error}
	}

	if err = cc.CheckCredential(sc, msg.cred); err != nil {
		h.tryWriteErrAndClose(err)
		return err
	}

	if err = h.writeAck(handshakeproto.Error_Null); err != nil {
		h.tryWriteErrAndClose(err)
		return err
	}

	msg, err = h.readMsg()
	if err != nil {
		h.tryWriteErrAndClose(err)
		return err
	}
	if msg.ack == nil {
		err = ErrUnexpectedPayload
		h.tryWriteErrAndClose(err)
		return err
	}
	if msg.ack.Error == handshakeproto.Error_Null {
		return nil
	} else {
		_ = h.conn.Close()
		return handshakeError{e: msg.ack.Error}
	}
}

func IncomingHandshake(sc sec.SecureConn, cc CredentialChecker) (err error) {
	h := newHandshake()
	defer h.release()
	h.conn = sc

	msg, err := h.readMsg()
	if err != nil {
		h.tryWriteErrAndClose(err)
		return err
	}
	if msg.ack != nil {
		return ErrUnexpectedPayload
	}
	if err = cc.CheckCredential(sc, msg.cred); err != nil {
		h.tryWriteErrAndClose(err)
		return err
	}

	if err = h.writeCredentials(cc.MakeCredentials(sc)); err != nil {
		h.tryWriteErrAndClose(err)
		return err
	}

	msg, err = h.readMsg()
	if err != nil {
		h.tryWriteErrAndClose(err)
		return err
	}
	if msg.ack == nil {
		err = ErrUnexpectedPayload
		h.tryWriteErrAndClose(err)
		return err
	}
	if msg.ack.Error != handshakeproto.Error_Null {
		if msg.ack.Error == handshakeproto.Error_InvalidCredentials {
			return ErrPeerDeclinedCredentials
		}
		return handshakeError{e: msg.ack.Error}
	}
	if err = h.writeAck(handshakeproto.Error_Null); err != nil {
		h.tryWriteErrAndClose(err)
		return err
	}
	return
}

func newHandshake() *handshake {
	return handshakePool.Get().(*handshake)
}

type handshake struct {
	conn       sec.SecureConn
	remoteCred *handshakeproto.Credentials
	remoteAck  *handshakeproto.Ack
	localAck   *handshakeproto.Ack
	buf        []byte
}

func (h *handshake) writeCredentials(cred *handshakeproto.Credentials) (err error) {
	h.buf = slices.Grow(h.buf, cred.Size()+headerSize)[:cred.Size()+headerSize]
	n, err := cred.MarshalToSizedBuffer(h.buf[headerSize:])
	if err != nil {
		return err
	}
	return h.writeData(msgTypeCred, n)
}

func (h *handshake) tryWriteErrAndClose(err error) {
	if err == ErrGotNotAHandshakeMessage {
		// if we got unexpected message - just close the connection
		_ = h.conn.Close()
		return
	}
	var ackErr handshakeproto.Error
	if he, ok := err.(handshakeError); ok {
		ackErr = he.e
	} else {
		ackErr = handshakeproto.Error_Unexpected
	}
	_ = h.writeAck(ackErr)
	_ = h.conn.Close()
}

func (h *handshake) writeAck(ackErr handshakeproto.Error) (err error) {
	h.localAck.Error = ackErr
	h.buf = slices.Grow(h.buf, h.localAck.Size()+headerSize)[:h.localAck.Size()+headerSize]
	n, err := h.localAck.MarshalTo(h.buf[headerSize:])
	if err != nil {
		return err
	}
	return h.writeData(msgTypeAck, n)
}

func (h *handshake) writeData(tp byte, size int) (err error) {
	h.buf[0] = tp
	binary.LittleEndian.PutUint32(h.buf[1:headerSize], uint32(size))
	_, err = h.conn.Write(h.buf[:size+headerSize])
	return err
}

type message struct {
	cred *handshakeproto.Credentials
	ack  *handshakeproto.Ack
}

func (h *handshake) readMsg() (msg message, err error) {
	h.buf = slices.Grow(h.buf, headerSize)[:headerSize]
	if _, err = io.ReadFull(h.conn, h.buf[:headerSize]); err != nil {
		return
	}
	tp := h.buf[0]
	if tp != msgTypeCred && tp != msgTypeAck {
		err = ErrGotNotAHandshakeMessage
		return
	}
	size := binary.LittleEndian.Uint32(h.buf[1:headerSize])
	h.buf = slices.Grow(h.buf, int(size))[:size]
	if _, err = io.ReadFull(h.conn, h.buf[:size]); err != nil {
		return
	}
	switch tp {
	case msgTypeCred:
		if err = h.remoteCred.Unmarshal(h.buf[:size]); err != nil {
			return
		}
		msg.cred = h.remoteCred
	case msgTypeAck:
		if err = h.remoteAck.Unmarshal(h.buf[:size]); err != nil {
			return
		}
		msg.ack = h.remoteAck
	}
	return
}

func (h *handshake) release() {
	h.buf = h.buf[:0]
	h.conn = nil
	h.localAck.Error = 0
	h.remoteAck.Error = 0
	h.remoteCred.Type = 0
	h.remoteCred.Payload = h.remoteCred.Payload[:0]
	handshakePool.Put(h)
}
