package models

// Blog represents the structure of a blog post in the application
type Blog struct {
	ID        int64  `json:"id"`         // Unique identifier for the blog
	UserID    int64  `json:"user_id"`    // ID of the user who created the blog
	Title     string `json:"title"`      // Title of the blog post
	Body      string `json:"body"`       // Content of the blog post
	ViewCount int    `json:"view_count"` // Number of times the blog has been viewed
	CreatedAt string `json:"created_at"` // Timestamp of when the blog was created
	UpdatedAt string `json:"updated_at"` // Timestamp of when the blog was last updated
}
