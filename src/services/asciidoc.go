package services

import (
	"bytes"
	"html/template"

	"github.com/bytesparadise/libasciidoc"
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
)

// ParseAsciiDoc converts AsciiDoc content to HTML
func ParseAsciiDoc(content []byte) (template.HTML, error) {
	config := configuration.NewConfiguration(
		configuration.WithBackEnd("html5"),
	)

	var buf bytes.Buffer
	_, err := libasciidoc.Convert(bytes.NewReader(content), &buf, config)
	if err != nil {
		return "", err
	}

	return template.HTML(buf.String()), nil
}
