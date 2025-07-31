package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/Caffeino/social/internal/store"
)

var usernames = []string{
	"alice", "andy", "john", "bob", "carol", "dave", "eve", "frank", "grace", "heidi",
	"ivan", "judy", "karen", "leo", "mike", "nina", "oscar", "peggy", "quinn", "ruth",
	"sam", "trudy", "ursula", "victor", "wendy", "xander", "yara", "zane", "bruce", "claire",
	"diana", "edgar", "fiona", "george", "harry", "ingrid", "jack", "kate", "louis", "maria",
	"nathan", "olga", "paul", "queen", "ron", "susan", "tom", "uma", "vince", "will",
}

var lastNames = []string{
	"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez",
	"Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson", "Thomas", "Taylor", "Moore", "Jackson", "Martin",
	"Lee", "Perez", "Thompson", "White", "Harris", "Sanchez", "Clark", "Ramirez", "Lewis", "Robinson",
	"Walker", "Young", "Allen", "King", "Wright", "Scott", "Torres", "Nguyen", "Hill", "Flores",
	"Green", "Adams", "Nelson", "Baker", "Hall", "Rivera", "Campbell", "Mitchell", "Carter", "Roberts",
}

var titles = []string{
	"Getting Started with Go",
	"Why Simplicity Matters",
	"Top 5 Debugging Tips",
	"Understanding Goroutines",
	"REST APIs in Go",
	"Writing Clean Code",
	"Handling Errors Gracefully",
	"Go vs Python: A Comparison",
	"Concurrency Made Easy",
	"Testing in Go",
	"Structs and Interfaces",
	"Managing Dependencies",
	"Optimizing Go Code",
	"Intro to Channels",
	"Build a CLI Tool",
	"Logging Best Practices",
	"Working with JSON",
	"Deploying Go Apps",
	"Go Modules Explained",
	"Performance Tips for Go",
}

var contents = []string{
	"Go is a statically typed language designed for simplicity and performance.",
	"Concurrency in Go is handled through goroutines and channels, making it lightweight.",
	"Error handling in Go favors explicit over implicit, helping you write more robust code.",
	"Using the `net/http` package, you can build a web server in just a few lines of Go.",
	"Go’s standard library provides everything you need to get started with real-world apps.",
	"Interfaces in Go are satisfied implicitly, leading to very flexible code designs.",
	"The Go tooling ecosystem is minimal but powerful — from `go fmt` to `go test`.",
	"Channels help synchronize goroutines, ensuring safe data sharing.",
	"Go modules make dependency management straightforward and reproducible.",
	"With Go, you get fast compile times and efficient memory usage.",
	"Testing in Go is simple with the built-in `testing` package.",
	"Go favors composition over inheritance, which keeps designs clean and modular.",
	"You can use goroutines to handle thousands of concurrent connections efficiently.",
	"A well-structured Go project separates concerns using packages and interfaces.",
	"JSON encoding and decoding is seamless with Go's `encoding/json` package.",
	"Go compiles to a single binary, making deployment incredibly easy.",
	"Use `defer` for clean resource management, especially with files and locks.",
	"The Go Playground is a great way to try out snippets and share them.",
	"Avoid global state in Go programs to keep code easy to test and reason about.",
	"Benchmarking in Go is supported right out of the box via `go test -bench`.",
}

var tags = []string{
	"golang", "webdev", "api", "tutorial", "backend",
	"concurrency", "coding", "testing", "devtools", "json",
	"programming", "performance", "cli", "best-practices", "deployment",
	"clean-code", "rest", "opensource", "productivity", "errors",
}

var comments = []string{
	"Great article, very informative!",
	"This really helped me understand goroutines better.",
	"Thanks for sharing this. Looking forward to more posts!",
	"I had trouble with this topic before — now it makes sense.",
	"Can you write more about error handling next?",
	"Nice and simple explanation. Bookmarked!",
	"I think you could also mention context cancellation here.",
	"This is exactly what I was looking for, thank you!",
	"Awesome write-up. Clean and to the point.",
	"Helpful tips. Keep up the good work!",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)
	tx, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Println("Error creating user:", err)
			return
		}
	}

	tx.Commit()

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		name := usernames[i%len(usernames)]

		users[i] = &store.User{
			FirstName: strings.ToUpper(name[:1]) + strings.ToLower(name[1:]),
			LastName:  lastNames[i%len(lastNames)],
			Username:  name + fmt.Sprintf("%d", i),
			Email:     name + fmt.Sprintf("%d", i) + "@example.com",
			Role: store.Role{
				Name: "user",
			},
		}

		users[i].Password.Set("123123")
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  posts[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return cms
}
