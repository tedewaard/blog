package models

type TimelineItem struct {
	Title       string
	Date        string
	Description string
	Position    string // "left" or "right"
}
