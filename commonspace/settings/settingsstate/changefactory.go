package settingsstate

import (
	"github.com/anytypeio/any-sync/commonspace/spacesyncproto"
	"time"
)

type ChangeFactory interface {
	CreateObjectDeleteChange(id string, state *State, isSnapshot bool) (res []byte, err error)
	CreateSpaceDeleteChange(t time.Time, state *State, isSnapshot bool) (res []byte, err error)
	CreateSpaceRestoreChange(state *State, isSnapshot bool) (res []byte, err error)
}

func NewChangeFactory() ChangeFactory {
	return &changeFactory{}
}

type changeFactory struct {
}

func (c *changeFactory) CreateObjectDeleteChange(id string, state *State, isSnapshot bool) (res []byte, err error) {
	content := &spacesyncproto.SpaceSettingsContent_ObjectDelete{
		ObjectDelete: &spacesyncproto.ObjectDelete{Id: id},
	}
	change := &spacesyncproto.SettingsData{
		Content: []*spacesyncproto.SpaceSettingsContent{
			{Value: content},
		},
		Snapshot: nil,
	}
	// TODO: add snapshot logic
	res, err = change.Marshal()
	return
}

func (c *changeFactory) CreateSpaceDeleteChange(t time.Time, state *State, isSnapshot bool) (res []byte, err error) {
	content := &spacesyncproto.SpaceSettingsContent_SpaceDelete{
		SpaceDelete: &spacesyncproto.SpaceDelete{Timestamp: t.UnixNano()},
	}
	change := &spacesyncproto.SettingsData{
		Content: []*spacesyncproto.SpaceSettingsContent{
			{Value: content},
		},
		Snapshot: nil,
	}
	// TODO: add snapshot logic
	res, err = change.Marshal()
	return
}

func (c *changeFactory) CreateSpaceRestoreChange(state *State, isSnapshot bool) (res []byte, err error) {
	content := &spacesyncproto.SpaceSettingsContent_SpaceDelete{
		SpaceDelete: &spacesyncproto.SpaceDelete{},
	}
	change := &spacesyncproto.SettingsData{
		Content: []*spacesyncproto.SpaceSettingsContent{
			{Value: content},
		},
		Snapshot: nil,
	}
	// TODO: add snapshot logic
	res, err = change.Marshal()
	return
}
