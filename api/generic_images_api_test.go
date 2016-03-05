package api

import (
	"github.com/flexiant/concerto/testdata"
	"testing"
)

func TestGetGenericImageList(t *testing.T) {
	genericImagesIn := testdata.GetGenericImageData()
	GetGenericImageListMocked(t, genericImagesIn)
}
