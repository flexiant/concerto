package api

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
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

// GetServerListFailErrMocked test mocked function
func GetServerListFailErrMocked(t *testing.T, serversIn *[]types.Server) *[]types.Server {

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
	cs.On("Get", "/v1/cloud/servers").Return(dIn, 200, fmt.Errorf("Mocked error"))
	serversOut, err := ds.GetServerList()
	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serversOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &serversOut
}

// GetServerListFailStatusMocked test mocked function
func GetServerListFailStatusMocked(t *testing.T, serversIn *[]types.Server) *[]types.Server {

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
	cs.On("Get", "/v1/cloud/servers").Return(dIn, 499, nil)
	serversOut, err := ds.GetServerList()
	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serversOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &serversOut
}

// GetServerListFailJSONMocked test mocked function
func GetServerListFailJSONMocked(t *testing.T, serversIn *[]types.Server) *[]types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/cloud/servers").Return(dIn, 200, nil)
	serversOut, err := ds.GetServerList()
	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serversOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// GetServerFailErrMocked test mocked function
func GetServerFailErrMocked(t *testing.T, server *types.Server) *types.Server {

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
	cs.On("Get", fmt.Sprintf("/v1/cloud/servers/%s", server.Id)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	serverOut, err := ds.GetServer(server.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return serverOut
}

// GetServerFailStatusMocked test mocked function
func GetServerFailStatusMocked(t *testing.T, server *types.Server) *types.Server {

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
	cs.On("Get", fmt.Sprintf("/v1/cloud/servers/%s", server.Id)).Return(dIn, 499, nil)
	serverOut, err := ds.GetServer(server.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverOut
}

// GetServerFailJSONMocked test mocked function
func GetServerFailJSONMocked(t *testing.T, server *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/v1/cloud/servers/%s", server.Id)).Return(dIn, 200, nil)
	serverOut, err := ds.GetServer(server.Id)
	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// CreateServerFailErrMocked test mocked function
func CreateServerFailErrMocked(t *testing.T, serverIn *types.Server) *types.Server {

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
	cs.On("Post", "/v1/cloud/servers/", mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	serverOut, err := ds.CreateServer(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return serverOut
}

// CreateServerFailStatusMocked test mocked function
func CreateServerFailStatusMocked(t *testing.T, serverIn *types.Server) *types.Server {

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
	cs.On("Post", "/v1/cloud/servers/", mapIn).Return(dOut, 499, nil)
	serverOut, err := ds.CreateServer(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverOut
}

// CreateServerFailJSONMocked test mocked function
func CreateServerFailJSONMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/v1/cloud/servers/", mapIn).Return(dIn, 200, nil)
	serverOut, err := ds.CreateServer(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// UpdateServerFailErrMocked test mocked function
func UpdateServerFailErrMocked(t *testing.T, serverIn *types.Server) *types.Server {

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
	cs.On("Put", fmt.Sprintf("/v1/cloud/servers/%s", serverIn.Id), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	serverOut, err := ds.UpdateServer(mapIn, serverIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return serverOut
}

// UpdateServerFailStatusMocked test mocked function
func UpdateServerFailStatusMocked(t *testing.T, serverIn *types.Server) *types.Server {

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
	cs.On("Put", fmt.Sprintf("/v1/cloud/servers/%s", serverIn.Id), mapIn).Return(dOut, 499, nil)
	serverOut, err := ds.UpdateServer(mapIn, serverIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverOut
}

// UpdateServerFailJSONMocked test mocked function
func UpdateServerFailJSONMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/v1/cloud/servers/%s", serverIn.Id), mapIn).Return(dIn, 200, nil)
	serverOut, err := ds.UpdateServer(mapIn, serverIn.Id)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// BootServerFailErrMocked test mocked function
func BootServerFailErrMocked(t *testing.T, serverIn *types.Server) *types.Server {

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
	cs.On("Put", fmt.Sprintf("/v1/cloud/servers/%s/boot", serverIn.Id), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	serverOut, err := ds.BootServer(mapIn, serverIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return serverOut
}

// BootServerFailStatusMocked test mocked function
func BootServerFailStatusMocked(t *testing.T, serverIn *types.Server) *types.Server {

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
	cs.On("Put", fmt.Sprintf("/v1/cloud/servers/%s/boot", serverIn.Id), mapIn).Return(dOut, 499, nil)
	serverOut, err := ds.BootServer(mapIn, serverIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverOut
}

// BootServerFailJSONMocked test mocked function
func BootServerFailJSONMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/v1/cloud/servers/%s/boot", serverIn.Id), mapIn).Return(dIn, 200, nil)
	serverOut, err := ds.BootServer(mapIn, serverIn.Id)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// RebootServerFailErrMocked test mocked function
func RebootServerFailErrMocked(t *testing.T, serverIn *types.Server) *types.Server {

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
	cs.On("Put", fmt.Sprintf("/v1/cloud/servers/%s/reboot", serverIn.Id), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	serverOut, err := ds.RebootServer(mapIn, serverIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return serverOut
}

// RebootServerFailStatusMocked test mocked function
func RebootServerFailStatusMocked(t *testing.T, serverIn *types.Server) *types.Server {

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
	cs.On("Put", fmt.Sprintf("/v1/cloud/servers/%s/reboot", serverIn.Id), mapIn).Return(dOut, 499, nil)
	serverOut, err := ds.RebootServer(mapIn, serverIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverOut
}

// RebootServerFailJSONMocked test mocked function
func RebootServerFailJSONMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/v1/cloud/servers/%s/reboot", serverIn.Id), mapIn).Return(dIn, 200, nil)
	serverOut, err := ds.RebootServer(mapIn, serverIn.Id)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// ShutdownServerFailErrMocked test mocked function
func ShutdownServerFailErrMocked(t *testing.T, serverIn *types.Server) *types.Server {

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
	cs.On("Put", fmt.Sprintf("/v1/cloud/servers/%s/shutdown", serverIn.Id), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	serverOut, err := ds.ShutdownServer(mapIn, serverIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return serverOut
}

// ShutdownServerFailStatusMocked test mocked function
func ShutdownServerFailStatusMocked(t *testing.T, serverIn *types.Server) *types.Server {

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
	cs.On("Put", fmt.Sprintf("/v1/cloud/servers/%s/shutdown", serverIn.Id), mapIn).Return(dOut, 499, nil)
	serverOut, err := ds.ShutdownServer(mapIn, serverIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverOut
}

// ShutdownServerFailJSONMocked test mocked function
func ShutdownServerFailJSONMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/v1/cloud/servers/%s/shutdown", serverIn.Id), mapIn).Return(dIn, 200, nil)
	serverOut, err := ds.ShutdownServer(mapIn, serverIn.Id)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// OverrideServerFailErrMocked test mocked function
func OverrideServerFailErrMocked(t *testing.T, serverIn *types.Server) *types.Server {

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
	cs.On("Put", fmt.Sprintf("/v1/cloud/servers/%s/override", serverIn.Id), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	serverOut, err := ds.OverrideServer(mapIn, serverIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return serverOut
}

// OverrideServerFailStatusMocked test mocked function
func OverrideServerFailStatusMocked(t *testing.T, serverIn *types.Server) *types.Server {

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
	cs.On("Put", fmt.Sprintf("/v1/cloud/servers/%s/override", serverIn.Id), mapIn).Return(dOut, 499, nil)
	serverOut, err := ds.OverrideServer(mapIn, serverIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverOut
}

// OverrideServerFailJSONMocked test mocked function
func OverrideServerFailJSONMocked(t *testing.T, serverIn *types.Server) *types.Server {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*serverIn)
	assert.Nil(err, "Server test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/v1/cloud/servers/%s/override", serverIn.Id), mapIn).Return(dIn, 200, nil)
	serverOut, err := ds.OverrideServer(mapIn, serverIn.Id)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// DeleteServerFailErrMocked test mocked function
func DeleteServerFailErrMocked(t *testing.T, serverIn *types.Server) {

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
	cs.On("Delete", fmt.Sprintf("/v1/cloud/servers/%s", serverIn.Id)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	err = ds.DeleteServer(serverIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")
}

// DeleteServerFailStatusMocked test mocked function
func DeleteServerFailStatusMocked(t *testing.T, serverIn *types.Server) {

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
	cs.On("Delete", fmt.Sprintf("/v1/cloud/servers/%s", serverIn.Id)).Return(dIn, 499, nil)
	err = ds.DeleteServer(serverIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

//======= DNS ==========

// GetDNSListMocked test mocked function
func GetDNSListMocked(t *testing.T, serverIn *types.Server, dnssIn *[]types.Dns) *[]types.Dns {

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

// GetDNSListFailErrMocked test mocked function
func GetDNSListFailErrMocked(t *testing.T, serverIn *types.Server, dnssIn *[]types.Dns) *[]types.Dns {

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
	cs.On("Get", fmt.Sprintf("/v1/cloud/servers/%s/records", serverIn.Id)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	dnssOut, err := ds.GetDNSList(serverIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(dnssOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &dnssOut
}

// GetDNSListFailStatusMocked test mocked function
func GetDNSListFailStatusMocked(t *testing.T, serverIn *types.Server, dnssIn *[]types.Dns) *[]types.Dns {

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
	cs.On("Get", fmt.Sprintf("/v1/cloud/servers/%s/records", serverIn.Id)).Return(dIn, 499, nil)
	dnssOut, err := ds.GetDNSList(serverIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(dnssOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &dnssOut
}

// GetDNSListFailJSONMocked test mocked function
func GetDNSListFailJSONMocked(t *testing.T, serverIn *types.Server, dnssIn *[]types.Dns) *[]types.Dns {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load dns service")
	assert.NotNil(ds, "Dns service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/v1/cloud/servers/%s/records", serverIn.Id)).Return(dIn, 200, nil)
	dnssOut, err := ds.GetDNSList(serverIn.Id)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(dnssOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return &dnssOut
}

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

// GetServerEventListFailErrMocked test mocked function
func GetServerEventListFailErrMocked(t *testing.T, eventsIn *[]types.Event, serverID string) *[]types.Event {

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
	cs.On("Get", fmt.Sprintf("/v1/cloud/servers/%s/events", serverID)).Return(evIn, 200, fmt.Errorf("Mocked error"))
	evOut, err := ds.GetEventsList(serverID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(evOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &evOut
}

// GetServerEventListFailStatusMocked test mocked function
func GetServerEventListFailStatusMocked(t *testing.T, eventsIn *[]types.Event, serverID string) *[]types.Event {

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
	cs.On("Get", fmt.Sprintf("/v1/cloud/servers/%s/events", serverID)).Return(evIn, 499, nil)
	evOut, err := ds.GetEventsList(serverID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(evOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &evOut
}

// GetServerEventListFailJSONMocked test mocked function
func GetServerEventListFailJSONMocked(t *testing.T, eventsIn *[]types.Event, serverID string) *[]types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// wrong json
	evIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/v1/cloud/servers/%s/events", serverID)).Return(evIn, 200, nil)
	evOut, err := ds.GetEventsList(serverID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(evOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return &evOut
}

// GetOperationalScriptListMocked test mocked function
func GetOperationalScriptListMocked(t *testing.T, scriptsIn *[]types.ScriptChar, serverID string) *[]types.ScriptChar {

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
	scriptsOut, err := ds.GetOperationalScriptsList(serverID)
	assert.Nil(err, "Error getting operational script list")
	assert.Equal(*scriptsIn, scriptsOut, "GetOperationalScriptList returned different operational scripts")

	return &scriptsOut
}

// GetOperationalScriptFailErrMocked test mocked function
func GetOperationalScriptFailErrMocked(t *testing.T, scriptsIn *[]types.ScriptChar, serverID string) *[]types.ScriptChar {

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
	cs.On("Get", fmt.Sprintf("/v1/cloud/servers/%s/operational_scripts", serverID)).Return(oscIn, 200, fmt.Errorf("Mocked error"))
	scriptsOut, err := ds.GetOperationalScriptsList(serverID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(scriptsOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &scriptsOut
}

// GetOperationalScriptFailStatusMocked test mocked function
func GetOperationalScriptFailStatusMocked(t *testing.T, scriptsIn *[]types.ScriptChar, serverID string) *[]types.ScriptChar {

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
	cs.On("Get", fmt.Sprintf("/v1/cloud/servers/%s/operational_scripts", serverID)).Return(oscIn, 499, nil)
	scriptsOut, err := ds.GetOperationalScriptsList(serverID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(scriptsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &scriptsOut
}

// GetOperationalScriptFailJSONMocked test mocked function
func GetOperationalScriptFailJSONMocked(t *testing.T, scriptsIn *[]types.ScriptChar, serverID string) *[]types.ScriptChar {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerService(cs)
	assert.Nil(err, "Couldn't load server service")
	assert.NotNil(ds, "Server service not instanced")

	// wrong json
	oscIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/v1/cloud/servers/%s/operational_scripts", serverID)).Return(oscIn, 200, nil)
	scriptsOut, err := ds.GetOperationalScriptsList(serverID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(scriptsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return &scriptsOut
}
