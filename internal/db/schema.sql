CREATE TABLE IF NOT EXISTS users (
	user_id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL UNIQUE,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	profile_picture TEXT,
	bio TEXT,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS categories (
	category_id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL UNIQUE,
	description TEXT,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS posts (
	post_id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER NOT NULL,
	category_id INTEGER NOT NULL,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
	FOREIGN KEY (category_id) REFERENCES categories(category_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS comments (
	comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
	post_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	content TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS likes (
	like_id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER NOT NULL,
	post_id INTEGER,
	comment_id INTEGER,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
	FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
	FOREIGN KEY (comment_id) REFERENCES comments(comment_id) ON DELETE CASCADE,
	CONSTRAINT check_post_or_comment CHECK (
		(post_id IS NOT NULL AND comment_id IS NULL) OR 
		(post_id IS NULL AND comment_id IS NOT NULL)
	)
);