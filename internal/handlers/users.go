package handlers

import (
	"database/sql"
	"forum/internal/db"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "web/templates/login.html")
	} else if r.Method == http.MethodPost {
		// login logic
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Serve the register page
		http.ServeFile(w, r, "web/templates/register.html")
	} else if r.Method == http.MethodPost {
		// Parse form data
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		// Extract form fields
		username := strings.TrimSpace(r.FormValue("username"))
		email := strings.TrimSpace(r.FormValue("email"))
		password := r.FormValue("password")

		// Validate form data
		if username == "" || email == "" || password == "" {
			http.Error(w, "Please fill in all fields", http.StatusBadRequest)
			return
		}
		if len(password) < 6 {
			http.Error(w, "Password must be at least 6 characters long", http.StatusBadRequest)
			return
		}

		// Check if email or username already exists
		var exists bool
		query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = ? OR username = ?)`
		err := db.DB.QueryRow(query, email, username).Scan(&exists)
		if err != nil && err != sql.ErrNoRows {
			http.Error(w, "Server error", http.StatusInternalServerError)
			log.Printf("Database qeuery error: %v\n", err)
			return
		}
		if exists {
			http.Error(w, "Email or Username already in use", http.StatusBadRequest)
			return
		}

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			log.Printf("Password hashing error: %v\n", err)
			return
		}

		// Insert the new user into the database
		insertQuery := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`
		_, err = db.DB.Exec(insertQuery, username, email, hashedPassword)
		if err != nil {
			http.Error(w, "Failed to register user", http.StatusInternalServerError)
			log.Printf("Database insert error: %v\n", err)
			return
		}

		// Redirect to the login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
