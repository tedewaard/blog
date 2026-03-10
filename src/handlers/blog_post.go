package handlers

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tedewaard/folio/src/templates"
)

// HandlePost renders an individual blog post
func (h *BlogHandler) HandlePost(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	if slug == "" {
		http.Error(w, "Blog post not found", http.StatusNotFound)
		return
	}

	post, err := h.blogService.GetPost(r.Context(), slug)
	if err != nil {
		log.Printf("Error getting post %s: %v", slug, err)
		http.Error(w, "Blog post not found", http.StatusNotFound)
		return
	}

	component := templates.BlogPost(*post)
	component.Render(r.Context(), w)
}
