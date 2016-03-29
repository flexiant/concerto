package api

import (
	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewServerServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewServerService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetServerList(t *testing.T) {
	serversIn := testdata.GetServerData()
	GetServerListMocked(t, serversIn)
}

func TestGetServer(t *testing.T) {
	serversIn := testdata.GetServerData()
	for _, serverIn := range *serversIn {
		GetServerMocked(t, &serverIn)
	}
}

func TestCreateServer(t *testing.T) {
	serversIn := testdata.GetServerData()
	for _, serverIn := range *serversIn {
		CreateServerMocked(t, &serverIn)
	}
}

func TestUpdateServer(t *testing.T) {
	serversIn := testdata.GetServerData()
	for _, serverIn := range *serversIn {
		UpdateServerMocked(t, &serverIn)
	}
}

func TestBootServer(t *testing.T) {
	serversIn := testdata.GetServerData()
	for _, serverIn := range *serversIn {
		BootServerMocked(t, &serverIn)
	}
}

func TestRebootServer(t *testing.T) {
	serversIn := testdata.GetServerData()
	for _, serverIn := range *serversIn {
		RebootServerMocked(t, &serverIn)
	}
}

func TestShutdownServer(t *testing.T) {
	serversIn := testdata.GetServerData()
	for _, serverIn := range *serversIn {
		ShutdownServerMocked(t, &serverIn)
	}
}

func TestOverrideServer(t *testing.T) {
	serversIn := testdata.GetServerData()
	for _, serverIn := range *serversIn {
		OverrideServerMocked(t, &serverIn)
	}
}

func TestDeleteServer(t *testing.T) {
	serversIn := testdata.GetServerData()
	for _, serverIn := range *serversIn {
		DeleteServerMocked(t, &serverIn)
	}
}

//======= DNS ==========v
func TestGetDnsList(t *testing.T) {
	serversIn := testdata.GetServerData()
	dnssIn := testdata.GetDnsData()
	for _, serverIn := range *serversIn {
		GetDnsListMocked(t, &serverIn, dnssIn)
	}
}

//======= Events ==========
func TestGetEventsList(t *testing.T) {
	serversIn := testdata.GetServerData()
	eventsIn := testdata.GetEventData()
	for _, serverIn := range *serversIn {
		GetServerEventListMocked(t, eventsIn, serverIn.Id)
	}
}

//======= Operational Scripts ==========
func TestGetOperationalScriptList(t *testing.T) {
	serversIn := testdata.GetServerData()
	scriptsIn := testdata.GetScriptCharData()
	for _, serverIn := range *serversIn {
		GetOperationalScriptListMocked(t, scriptsIn, serverIn.Id)
	}
}
