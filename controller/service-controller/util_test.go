package srvcont

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFilePath(t *testing.T) {
	assert.Equal(t,
		"",
		GetFilePath("/devingen", "devingen", true),
		"incorrect file path",
	)

	assert.Equal(t,
		"",
		GetFilePath("/devingen/", "devingen", true),
		"incorrect file path",
	)

	assert.Equal(t,
		"assets",
		GetFilePath("/devingen/assets", "devingen", true),
		"incorrect file path",
	)

	assert.Equal(t,
		"assets",
		GetFilePath("/devingen/assets/", "devingen", true),
		"incorrect file path",
	)

	assert.Equal(t,
		"assets",
		GetFilePath("/devingen/assets", "devingen", false),
		"incorrect file path",
	)

	assert.Equal(t,
		"assets/icon",
		GetFilePath("/devingen/assets/icon", "devingen", true),
		"incorrect file path",
	)
}
