package short_encoding_of_words

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, 10, minimumLengthEncoding([]string{"time", "me", "bell"}))
	assert.Equal(t, 2, minimumLengthEncoding([]string{"t"}))
}
