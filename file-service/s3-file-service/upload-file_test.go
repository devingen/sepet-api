package s3fs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFilePath(t *testing.T) {
	assert.Equal(t,
		"folder-for-70l9ad53s5/0.0.2/images/selimiye.jpg",
		GetFilePath("folder-for-70l9ad53s5", "0.0.2", "images/selimiye.jpg"),
		"S3 path is not valid",
	)
}

func TestGenerateSepetCDNURL(t *testing.T) {
	assert.Equal(t,
		"https://70l9ad53s5.sepet.devingen.io/images/selimiye.jpg",
		GenerateSepetCDNURL("sepet.devingen.io", "https", "70l9ad53s5", "images/selimiye.jpg"),
		"CDN URL is not valid",
	)
}
