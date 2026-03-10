package models

type Technology struct {
	Name  string
	Color string // blue, green, purple, yellow, red, indigo
}

type Project struct {
	Title        string
	Description  string
	Technologies []Technology
	Link         string // GitHub URL (optional)
	WebsiteLink  string // Website URL (optional)
}
