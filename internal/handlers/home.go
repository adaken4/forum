package handlers

import (
	"fmt"
	"forum/internal/auth"
	"forum/internal/db"
	"forum/internal/models"
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	currentUserID := auth.GetCurrentUserID(r)
	fmt.Println(currentUserID)

	query := `
		SELECT p.post_id, p.title, p.content, u.username, u.user_id, c.name AS category, p.created_at,
		       (SELECT COUNT(*) FROM likes WHERE post_id = p.post_id AND comment_id IS NULL AND like_type = 'like') AS like_count,
		       (SELECT COUNT(*) FROM likes WHERE post_id = p.post_id AND comment_id IS NULL AND like_type = 'dislike') AS dislike_count
		FROM posts p
		JOIN users u ON p.user_id = u.user_id
		JOIN categories c ON p.category_id = c.category_id
		ORDER BY p.created_at DESC`

	rows, err := db.DB.Query(query)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to fetch posts", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// // Check if there are any rows
	// if !rows.Next() {
	// 	// No rows returned
	// 	http.Error(w, "No posts found", http.StatusNotFound)
	// 	return
	// }

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.PostID, &post.Title, &post.Content, &post.Username, &post.UserID, &post.Category, &post.CreatedAt, &post.LikeCount, &post.DislikeCount)
		if err != nil {
			http.Error(w, "Error scanning posts", http.StatusInternalServerError)
			return
		}

		// Fetch comments for each post
		commentQuery := `
			SELECT c.comment_id, c.post_id, c.content, u.username, u.user_id, c.created_at,
				(SELECT COUNT(*) FROM likes WHERE comment_id = c.comment_id AND like_type = 'like') AS like_count,
				(SELECT COUNT(*) FROM likes WHERE comment_id = c.comment_id AND like_type = 'dislike') AS dislike_count
			FROM comments c
			JOIN users u ON c.user_id = u.user_id
			WHERE c.post_id = ?
			ORDER BY c.created_at ASC`
		commentRows, err := db.DB.Query(commentQuery, post.PostID)
		if err != nil {
			log.Println(err)
			http.Error(w, "Unable to fetch comments", http.StatusInternalServerError)
			return
		}

		var comments []models.Comment
		for commentRows.Next() {
			var comment models.Comment
			err := commentRows.Scan(&comment.CommentID, &comment.PostID, &comment.Content, &comment.Username, &comment.UserID, &comment.CreatedAt, &comment.LikeCount, &comment.DislikeCount)
			if err != nil {
				http.Error(w, "Error scanning comments", http.StatusInternalServerError)
				return
			}
			comments = append(comments, comment)
		}
		commentRows.Close()

		post.Comments = comments
		posts = append(posts, post)
	}

	data := struct {
		Posts         []models.Post
		CurrentUserID int
	}{
		Posts:         posts,
		CurrentUserID: currentUserID,
	}

	tmpl := template.Must(template.ParseFiles("web/templates/layout.html", "web/templates/home.html"))

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		return
	}
}
