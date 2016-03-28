package format

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultFormatter(t *testing.T) {

	assert := assert.New(t)

	formatter = nil
	f := GetFormatter()
	assert.NotNil(f, "Formatter shouldn't be nil")
}
