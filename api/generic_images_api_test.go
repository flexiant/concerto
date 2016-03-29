package api

import (
	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGenericImageServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewGenericImageService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetGenericImageList(t *testing.T) {
	genericImagesIn := testdata.GetGenericImageData()
	GetGenericImageListMocked(t, genericImagesIn)
}
