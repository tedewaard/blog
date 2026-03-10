package services

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/tedewaard/folio/src/models"
)

// BlogService handles blog post operations
type BlogService struct {
	contentDir string
}

// NewBlogService creates a new blog service
func NewBlogService(contentDir string) *BlogService {
	return &BlogService{
		contentDir: contentDir,
	}
}

// ListPosts returns all blog posts sorted by date (descending)
func (s *BlogService) ListPosts(ctx context.Context) ([]models.BlogPost, error) {
	files, err := filepath.Glob(filepath.Join(s.contentDir, "*.adoc"))
	if err != nil {
		return nil, fmt.Errorf("failed to list blog posts: %w", err)
	}

	log.Printf("Found %d blog post files in %s", len(files), s.contentDir)

	posts := make([]models.BlogPost, 0, len(files))
	for _, file := range files {
		post, err := s.parsePost(file)
		if err != nil {
			log.Printf("Error parsing post %s: %v", file, err)
			continue
		}
		posts = append(posts, *post)
	}

	log.Printf("Successfully parsed %d blog posts", len(posts))

	// Sort by date descending (newest first)
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Metadata.Date.After(posts[j].Metadata.Date)
	})

	return posts, nil
}

// GetPost retrieves a single blog post by slug
func (s *BlogService) GetPost(ctx context.Context, slug string) (*models.BlogPost, error) {
	filePath := filepath.Join(s.contentDir, slug+".adoc")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("post not found: %s", slug)
	}

	return s.parsePost(filePath)
}

// parsePost reads and parses an AsciiDoc file with frontmatter
func (s *BlogService) parsePost(filePath string) (*models.BlogPost, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// Parse frontmatter
	var metadata models.BlogMetadata
	rest, err := frontmatter.Parse(strings.NewReader(string(content)), &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to parse frontmatter in %s: %w", filePath, err)
	}

	// Convert AsciiDoc to HTML
	html, err := ParseAsciiDoc(rest)
	if err != nil {
		return nil, fmt.Errorf("failed to parse AsciiDoc in %s: %w", filePath, err)
	}

	// Calculate reading time
	wordCount := countWords(string(rest))
	readingMinutes := int(math.Ceil(float64(wordCount) / 200.0)) // Average reading speed: 200 words/min
	metadata.ReadingTime = fmt.Sprintf("%d min read", readingMinutes)

	// Extract slug from filename
	slug := strings.TrimSuffix(filepath.Base(filePath), ".adoc")

	return &models.BlogPost{
		Slug:       slug,
		Metadata:   metadata,
		Content:    html,
		RawContent: rest,
	}, nil
}

// countWords counts the number of words in text
func countWords(text string) int {
	fields := strings.Fields(text)
	return len(fields)
}
