package handlers

import (
	"log"
	"net/http"

	"github.com/tedewaard/folio/src/services"
	"github.com/tedewaard/folio/src/templates"
)

// BlogHandler handles blog listing requests
type BlogHandler struct {
	blogService *services.BlogService
}

// NewBlogHandler creates a new blog handler
func NewBlogHandler(blogService *services.BlogService) *BlogHandler {
	return &BlogHandler{
		blogService: blogService,
	}
}

// HandleList renders the blog listing page
func (h *BlogHandler) HandleList(w http.ResponseWriter, r *http.Request) {
	posts, err := h.blogService.ListPosts(r.Context())
	if err != nil {
		log.Printf("Error listing posts: %v", err)
		http.Error(w, "Failed to load blog posts", http.StatusInternalServerError)
		return
	}

	component := templates.BlogList(posts)
	component.Render(r.Context(), w)
}
