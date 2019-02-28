package fake

import (
	mock "github.com/stretchr/testify/mock"

	common "github.com/visual42/demoinfocs-golang/common"
)

// Participants is a mock for of demoinfocs.IParticipants.
type Participants struct {
	mock.Mock
}

// ByUserID is a mock-implementation of IParticipants.ByUserID().
func (ptcp *Participants) ByUserID() map[int]*common.Player {
	return ptcp.Called().Get(0).(map[int]*common.Player)
}

// ByEntityID is a mock-implementation of IParticipants.ByEntityID().
func (ptcp *Participants) ByEntityID() map[int]*common.Player {
	return ptcp.Called().Get(0).(map[int]*common.Player)
}

// All is a mock-implementation of IParticipants.All().
func (ptcp *Participants) All() []*common.Player {
	return ptcp.Called().Get(0).([]*common.Player)
}

// Playing is a mock-implementation of IParticipants.Playing().
func (ptcp *Participants) Playing() []*common.Player {
	return ptcp.Called().Get(0).([]*common.Player)
}

// TeamMembers is a mock-implementation of IParticipants.TeamMembers().
func (ptcp *Participants) TeamMembers(team common.Team) []*common.Player {
	return ptcp.Called().Get(0).([]*common.Player)
}

// FindByHandle is a mock-implementation of IParticipants.FindByHandle().
func (ptcp *Participants) FindByHandle(handle int) *common.Player {
	return ptcp.Called().Get(0).(*common.Player)
}
