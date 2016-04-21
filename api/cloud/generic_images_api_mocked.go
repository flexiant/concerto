package cloud

import (
	"encoding/json"
	"fmt"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO exclude from release compile

// GetGenericImageListMocked test mocked function
func GetGenericImageListMocked(t *testing.T, genericImagesIn *[]types.GenericImage) *[]types.GenericImage {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewGenericImageService(cs)
	assert.Nil(err, "Couldn't load genericImage service")
	assert.NotNil(ds, "GenericImage service not instanced")

	// to json
	dIn, err := json.Marshal(genericImagesIn)
	assert.Nil(err, "GenericImage test data corrupted")

	// call service
	cs.On("Get", "/v1/cloud/generic_images").Return(dIn, 200, nil)
	genericImagesOut, err := ds.GetGenericImageList()
	assert.Nil(err, "Error getting genericImage list")
	assert.Equal(*genericImagesIn, genericImagesOut, "GetGenericImageList returned different genericImages")

	return &genericImagesOut
}

// GetGenericImageListFailErrMocked test mocked function
func GetGenericImageListFailErrMocked(t *testing.T, genericImagesIn *[]types.GenericImage) *[]types.GenericImage {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewGenericImageService(cs)
	assert.Nil(err, "Couldn't load genericImage service")
	assert.NotNil(ds, "GenericImage service not instanced")

	// to json
	dIn, err := json.Marshal(genericImagesIn)
	assert.Nil(err, "GenericImage test data corrupted")

	// call service
	cs.On("Get", "/v1/cloud/generic_images").Return(dIn, 200, fmt.Errorf("Mocked error"))
	genericImagesOut, err := ds.GetGenericImageList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(genericImagesOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &genericImagesOut
}

// GetGenericImageListFailStatusMocked test mocked function
func GetGenericImageListFailStatusMocked(t *testing.T, genericImagesIn *[]types.GenericImage) *[]types.GenericImage {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewGenericImageService(cs)
	assert.Nil(err, "Couldn't load genericImage service")
	assert.NotNil(ds, "GenericImage service not instanced")

	// to json
	dIn, err := json.Marshal(genericImagesIn)
	assert.Nil(err, "GenericImage test data corrupted")

	// call service
	cs.On("Get", "/v1/cloud/generic_images").Return(dIn, 499, nil)
	genericImagesOut, err := ds.GetGenericImageList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(genericImagesOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &genericImagesOut
}

// GetGenericImageListFailJSONMocked test mocked function
func GetGenericImageListFailJSONMocked(t *testing.T, genericImagesIn *[]types.GenericImage) *[]types.GenericImage {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewGenericImageService(cs)
	assert.Nil(err, "Couldn't load genericImage service")
	assert.NotNil(ds, "GenericImage service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/cloud/generic_images").Return(dIn, 200, nil)
	genericImagesOut, err := ds.GetGenericImageList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(genericImagesOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return &genericImagesOut
}
