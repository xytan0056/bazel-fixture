package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitize(t *testing.T) {
	assert.Equal(t, "you and me at home", Sanitize("you & me @ home"))
}

func TestTitleWords(t *testing.T) {
	assert.Equal(t, "Hello World", TitleWords("hello world"))
}

func TestGreet(t *testing.T) {
	assert.Equal(t, "Hola", Greet("es"))
	assert.Equal(t, "Hello", Greet("xx"))
}

func TestReverse(t *testing.T) {
	assert.Equal(t, "cba", Reverse("abc"))
}
