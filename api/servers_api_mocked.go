package api

import (
	"encoding/json"
	"fmt"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO exclude from release compile

// GetServerListMocked test mocked function
func GetServerListMocked(t *testing.T, serversIn *[]types.Server) *[]types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	dIn, err := json.Marshal(serversIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Get", "/v1/cloud/servers").Return(dIn, 200, nil)
	serversOut, err := ds.GetServerList()
	assert.Nil(err, "Error getting server list")
	assert.Equal(*serversIn, serversOut, "GetServerList returned different servers")

	return &serversOut
}

// GetServerMocked test mocked function
func GetServerMocked(t *testing.T, server *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	dIn, err := json.Marshal(server)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/cloud/servers/%s", server.Id)).Return(dIn, 200, nil)
	serverOut, err := ds.GetServer(server.Id)
	assert.Nil(err, "Error getting server")
	assert.Equal(*server, *serverOut, "GetServer returned different servers")

	return serverOut
}

// CreateServerMocked test mocked function
func CreateServerMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Post", "/v1/cloud/servers/", mapIn).Return(dOut, 200, nil)
	serverOut, err := ds.CreateServer(mapIn)
	assert.Nil(err, "Error creating server list")
	assert.Equal(serverIn, serverOut, "CreateServer returned different servers")

	return serverOut
}

// UpdateServerMocked test mocked function
func UpdateServerMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/cloud/servers/%s", serverIn.Id), mapIn).Return(dOut, 200, nil)
	serverOut, err := ds.UpdateServer(mapIn, serverIn.Id)
	assert.Nil(err, "Error updating server list")
	assert.Equal(serverIn, serverOut, "UpdateServer returned different servers")

	return serverOut
}

// BootServerMocked test mocked function
func BootServerMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/cloud/servers/%s/boot", serverIn.Id), mapIn).Return(dOut, 200, nil)
	serverOut, err := ds.BootServer(mapIn, serverIn.Id)
	assert.Nil(err, "Error updating server list")
	assert.Equal(serverIn, serverOut, "BootServer returned different servers")

	return serverOut
}

// RebootServerMocked test mocked function
func RebootServerMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/cloud/servers/%s/reboot", serverIn.Id), mapIn).Return(dOut, 200, nil)
	serverOut, err := ds.RebootServer(mapIn, serverIn.Id)
	assert.Nil(err, "Error updating server list")
	assert.Equal(serverIn, serverOut, "RebootServer returned different servers")

	return serverOut
}

// ShutdownServerMocked test mocked function
func ShutdownServerMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/cloud/servers/%s/shutdown", serverIn.Id), mapIn).Return(dOut, 200, nil)
	serverOut, err := ds.ShutdownServer(mapIn, serverIn.Id)
	assert.Nil(err, "Error updating server list")
	assert.Equal(serverIn, serverOut, "ShutdownServer returned different servers")

	return serverOut
}

// OverrideServerMocked test mocked function
func OverrideServerMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// to json
	dOut, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/cloud/servers/%s/override", serverIn.Id), mapIn).Return(dOut, 200, nil)
	serverOut, err := ds.OverrideServer(mapIn, serverIn.Id)
	assert.Nil(err, "Error updating server list")
	assert.Equal(serverIn, serverOut, "OverrideServer returned different servers")

	return serverOut
}

// DeleteServerMocked test mocked function
func DeleteServerMocked(t *testing.T, serverIn *types.Server) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	dIn, err := json.Marshal(serverIn)
	assert.Nil(err, "Server test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/cloud/servers/%s", serverIn.Id)).Return(dIn, 200, nil)
	err = ds.DeleteServer(serverIn.Id)
	assert.Nil(err, "Error deleting server")
}

//======= DNS ==========
// GetDnsListMocked test mocked function
func GetDnsListMocked(t *testing.T, serverIn *types.Server, dnssIn *[]types.Dns) *[]types.Dns {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load dns service")
	assert.NotNil(ds, "Dns service not instanced")

	// to json
	dIn, err := json.Marshal(dnssIn)
	assert.Nil(err, "Dns test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/cloud/servers/%s/records", serverIn.Id)).Return(dIn, 200, nil)
	dnssOut, err := ds.GetDNSList(serverIn.Id)
	assert.Nil(err, "Error getting dns list")
	assert.Equal(*dnssIn, dnssOut, "GetDNSList returned different dnss")

	return &dnssOut
}

//======= Events ==========
// GetServerEventListMocked test mocked function
func GetServerEventListMocked(t *testing.T, eventsIn *[]types.Event, serverID string) *[]types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	evIn, err := json.Marshal(eventsIn)
	assert.Nil(err, "Server event test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/cloud/servers/%s/events", serverID)).Return(evIn, 200, nil)
	evOut, err := ds.GetEventsList(serverID)
	assert.Nil(err, "Error getting server event list")
	assert.Equal(*eventsIn, evOut, "GetServerEventList returned different server events")

	return &evOut
}

//======= Operational Scripts ==========
// GetOperationalScriptListMocked test mocked function
func GetOperationalScriptListMocked(t *testing.T, scriptsIn *[]types.ScriptChar, serverID string) *[]types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// to json
	oscIn, err := json.Marshal(scriptsIn)
	assert.Nil(err, "Server operational scripts test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/cloud/servers/%s/operational_scripts", serverID)).Return(oscIn, 200, nil)
	scriptsOut, err := ds.GetEventsList(serverID)
	assert.Nil(err, "Error getting operational script list")
	assert.Equal(*scriptsIn, scriptsOut, "GetOperationalScriptList returned different operational scripts")

	return &scriptsOut
}
