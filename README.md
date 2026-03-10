# Folio - Go Website

A modern portfolio website built with Go, Templ, and Datastar.

## Tech Stack

- **Go 1.26.1** - Backend language
- **Templ** - Type-safe HTML templating
- **Tailwind CSS** - Utility-first CSS framework
- **Datastar** - Hypermedia-driven interactivity
- **chi** - HTTP router
- **AsciiDoc** - Blog content format
- **Air** - Hot reload for development

## Prerequisites

- Go 1.26.1 or later
- templ CLI: `go install github.com/a-h/templ/cmd/templ@latest`
- air CLI: `go install github.com/air-verse/air@latest`

## Project Structure

```
go/
├── src/
│   ├── main.go              # Application entry point
│   ├── config/              # Configuration management
│   ├── handlers/            # HTTP handlers
│   ├── models/              # Data structures
│   ├── services/            # Business logic
│   ├── middleware/          # HTTP middleware
│   ├── templates/           # Templ templates
│   └── static/
│       └── css/
│           ├── input.css    # Tailwind input (directives)
│           └── output.css   # Generated Tailwind CSS
├── content/
│   └── blog/                # AsciiDoc blog posts
├── bin/                     # Compiled binaries
├── tmp/                     # Air temporary files
├── tailwindcss              # Tailwind CLI binary
├── tailwind.config.js       # Tailwind configuration
├── go.mod                   # Go dependencies
└── .air.toml                # Hot reload config
```

## Getting Started

### Installation

1. Install dependencies:
```bash
cd /home/tedewaard/repo/folio/go
go mod download
```

2. Generate Templ files:
```bash
cd src
templ generate
```

### Development

Run with hot reload (watches .go, .templ, and .adoc files):
```bash
cd /home/tedewaard/repo/folio/go
air
```

The server will start on http://localhost:3000 (or use `PORT=8080 air` for a different port).

### Production Build

```bash
# Generate templates
cd src
templ generate

# Build binary
cd ..
go build -o bin/server ./src

# Run
./bin/server
```

Or set custom configuration:
```bash
PORT=8080 ENVIRONMENT=production ./bin/server
```

## Styling with Tailwind CSS

This project uses Tailwind CSS for styling. The workflow is fully automated with Air's hot reload.

### How It Works

1. **Input CSS** (`src/static/css/input.css`): Contains Tailwind directives
2. **Tailwind CLI**: Scans templates for class names and generates CSS
3. **Output CSS** (`src/static/css/output.css`): Optimized CSS served to browsers

### Adding Custom Styles

Edit `src/static/css/input.css`:

```css
@tailwind base;
@tailwind components;
@tailwind utilities;

/* Your custom CSS here */
@layer components {
  .btn-custom {
    @apply px-4 py-2 bg-blue-500 text-white rounded;
  }
}
```

### Tailwind Configuration

Customize Tailwind in `tailwind.config.js`:

```javascript
module.exports = {
  content: [
    "./src/templates/**/*.templ",
    "./src/templates/**/*.go",
  ],
  theme: {
    extend: {
      colors: {
        brand: '#your-color',
      },
    },
  },
}
```

### Manual Rebuild

If you need to manually rebuild Tailwind CSS:

```bash
./tailwindcss -i ./src/static/css/input.css -o ./src/static/css/output.css --minify
```

Air automatically rebuilds Tailwind when you change `.templ`, `.go`, or `.css` files.

## Adding Blog Posts

1. Create a new `.adoc` file in `content/blog/`:

```bash
touch content/blog/my-new-post.adoc
```

2. Add frontmatter and content:

```asciidoc
---
title: "My New Blog Post"
date: 2026-03-07T10:00:00Z
tags: ["go", "web"]
excerpt: "A brief description of the post"
---
= My New Blog Post

Content goes here...

== Section Heading

More content...

[source,go]
----
// Code example
package main
----
```

3. The post will automatically appear on the blog listing page.

## Configuration

Environment variables:

- `PORT` - Server port (default: 3000)
- `ENVIRONMENT` - Environment mode (default: development)
- `CONTENT_DIR` - Blog content directory (default: ../content/blog)
- `STATIC_DIR` - Static files directory (default: ./static)

## Architecture

### Server-Side Rendering

All pages are rendered on the server using Templ templates. No separate frontend build step is required.

### Datastar Integration

Datastar enables hypermedia-driven interactivity by swapping HTML fragments (not JSON APIs). When users interact with the page, Datastar makes requests to Go endpoints that return HTML snippets, which are then swapped into the DOM.

### Routes

- `GET /` - Home page
- `GET /blog` - Blog listing
- `GET /blog/{slug}` - Individual blog post
- `GET /static/*` - Static files (CSS, JS, images)

## Development Tips

- Templates are in `src/templates/*.templ`
- Run `templ generate` after editing `.templ` files (or use Air which does this automatically)
- All templates must be in the same package
- Blog posts use AsciiDoc format with YAML frontmatter
- Hot reload watches for changes and rebuilds automatically

## Troubleshooting

**Port already in use:**
```bash
PORT=8080 air
```

**Templ generation errors:**
```bash
cd src
templ generate
```

**Build errors:**
Make sure all dependencies are installed:
```bash
go mod download
```

## Next Steps

This scaffolding provides:
- ✅ Server-side rendering with Templ
- ✅ Blog with AsciiDoc support
- ✅ Hot reload for development
- ✅ Type-safe templates
- ✅ Syntax highlighting for code blocks
- ✅ Responsive design

Future enhancements could include:
- Datastar fragment handlers for interactive features
- RSS/Atom feed generation
- Tag filtering on blog listing
- Full-text search
- Image optimization
- Sitemap generation
