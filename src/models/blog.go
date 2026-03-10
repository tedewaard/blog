package models

import (
	"html/template"
	"time"
)

// BlogMetadata holds the frontmatter metadata for a blog post
type BlogMetadata struct {
	Title       string    `yaml:"title"`
	Date        time.Time `yaml:"date"`
	Tags        []string  `yaml:"tags"`
	Excerpt     string    `yaml:"excerpt"`
	ReadingTime string    // Calculated field
}

// BlogPost represents a complete blog post with metadata and content
type BlogPost struct {
	Slug       string
	Metadata   BlogMetadata
	Content    template.HTML // Rendered HTML content
	RawContent []byte        // Original AsciiDoc content
}
