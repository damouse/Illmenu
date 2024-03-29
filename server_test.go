package illmenu

// Uses internal testing tools from go-json-rest

import (
	"testing"

	"github.com/ant0ine/go-json-rest/rest/test"
	"github.com/stretchr/testify/assert"
)

func TestWork(t *testing.T) {

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	// assert.Nil(t, object)

	// // assert for not nil (good when you expect something)
	// if assert.NotNil(t, object) {

	// 	// now we know that object isn't nil, we are safe to make
	// 	// further assertions without causing any errors
	// 	assert.Equal(t, "Something", object.Value)

	// }
}

func TestSimpleRequest(t *testing.T) {
	s := NewServer(nil)

	recorded := test.RunRequest(t, s.Api.MakeHandler(), test.MakeSimpleRequest("GET", "http://1.2.3.4/query", nil))
	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
}
