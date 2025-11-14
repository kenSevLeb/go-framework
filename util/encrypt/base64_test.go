package encrypt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	source := "some thing"
	encoded := Base64Encode(source)
	assert.Equal(t, source, Base64Decode(encoded))

}
