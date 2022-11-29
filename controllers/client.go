package controllers

import (
	"../database"
	"github.com/gin-gonic/gin"
)

type Post struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"body"`
}

func Read(c *gin.Context) {

	db := database.DBConn()
	rows, err := db.Query("SELECT id, title, body FROM posts WHERE id = " + c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "Story not found",
		})
	}

	post := Post{}

	for rows.Next() {
		var id int
		var title, body string

		err = rows.Scan(&id, &title, &body)
		if err != nil {
			panic(err.Error())
		}

		post.Id = id
		post.Title = title
		post.Content = body
	}

	c.JSON(200, post)
	defer db.Close() // Hoãn lại việc close database connect cho đến khi hàm Read() thực hiệc xong
}
