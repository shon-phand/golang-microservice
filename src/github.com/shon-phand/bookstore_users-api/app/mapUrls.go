package app

import "github.com/shon-phand/bookstore_users-api/controllers/user"

func mapUrls() {
	r.GET("/ping", user.Ping())

	r.GET("/users/:user_id", user.GetUser())
	r.POST("/users", user.CreateUser())
}
