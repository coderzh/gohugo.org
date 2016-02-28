package main

import (
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestUpdateOriginHugoRepo(t *testing.T) {
	updateOriginHugoRepo()

	_, err := os.Stat("../originhugo/README.md")
	assert.Equal(t, nil, err)
}

func TestConvertContentText(t *testing.T) {
	content := "overview content\nurl: /overview/start.md\nlink: /theme/go.md"
	docSubDirs := []string{"overview", "theme"}

	convertedContent := convertContentText(docSubDirs, content)
	assert.Equal(t,
		"overview content\nurl: /doc/overview/start.md\nlink: /doc/theme/go.md",
		convertedContent)
}

func TestConvertAllDocs(t *testing.T) {
	err := convertAllDocs()
	assert.Equal(t, nil, err)

	metadata_en, content_en, err := getMetadata("content/doc/commands/hugo_new_en.md")
	assert.Equal(t, nil, err)
	docs_en, err := cast.ToStringSliceE(metadata_en["doc"])
	assert.Equal(t, nil, err)
	assert.Equal(t, "commands_en", docs_en[0])
	assert.NotEqual(t, 0, len(content_en))

	metadata, content, err := getMetadata("content/doc/commands/hugo_new.md")
	assert.Equal(t, nil, err)
	docs, err := cast.ToStringSliceE(metadata["doc"])
	assert.Equal(t, nil, err)
	assert.Equal(t, "commands", docs[0])
	assert.NotEqual(t, 0, len(content))

	// exclude showcase dir
	_, err = os.Stat("/content/doc/showcase/spf13.md")
	assert.True(t, os.IsNotExist(err))
}
