package repository

import (
	"database/sql"

	"github.com/henok/blog-app/internal/models" // Import the models package
	"github.com/henok/blog-app/pkg/database"
)

// CreateBlog inserts a new blog into the database
func CreateBlog(blog models.Blog) (int64, error) {
	result, err := database.DB.Exec("INSERT INTO blogs (user_id, title, body) VALUES (?, ?, ?)", blog.UserID, blog.Title, blog.Body)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// DeleteBlog deletes a blog from the database by its ID
func DeleteBlog(blogID int64) error {
	_, err := database.DB.Exec("DELETE FROM blogs WHERE id = ?", blogID)
	return err
}

// GetAllBlogs retrieves all blogs from the database
func GetAllBlogs() ([]models.Blog, error) {
	rows, err := database.DB.Query("SELECT * FROM blogs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []models.Blog
	for rows.Next() {
		var blog models.Blog
		if err := rows.Scan(&blog.ID, &blog.UserID, &blog.Title, &blog.Body, &blog.ViewCount, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}
	return blogs, nil
}

// GetBlogsByUser retrieves all blogs for a specific user from the database
func GetBlogsByUser(userID int64) ([]models.Blog, error) {
	rows, err := database.DB.Query("SELECT * FROM blogs WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []models.Blog
	for rows.Next() {
		var blog models.Blog
		if err := rows.Scan(&blog.ID, &blog.UserID, &blog.Title, &blog.Body, &blog.ViewCount, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}
	return blogs, nil
}

// GetBlogByID retrieves a single blog by its ID from the database
func GetBlogByID(blogID int64) (models.Blog, error) {
	var blog models.Blog
	err := database.DB.QueryRow("SELECT * FROM blogs WHERE id = ?", blogID).
		Scan(&blog.ID, &blog.UserID, &blog.Title, &blog.Body, &blog.ViewCount, &blog.CreatedAt, &blog.UpdatedAt)
	if err == sql.ErrNoRows {
		return blog, nil
	} else if err != nil {
		return blog, err
	}
	return blog, nil
}
