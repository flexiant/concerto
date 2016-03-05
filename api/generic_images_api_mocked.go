package api

import (
	"encoding/json"
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
