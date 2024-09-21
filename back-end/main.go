package main

import (
	"database/sql"

	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Ares1605/casual-chess-backend/env"
	_ "github.com/mattn/go-sqlite3"

)
func main() {
	fmt.Println(uuid.New())

	router := gin.Default()

  router.GET("/create/user", func(c *gin.Context) {
		db, err := openDB()
		if err != nil {
			c.JSON(500, gin.H{"success": false, "error": "failed to open DB"});
			return
		}
		defer db.Close()

		rows, err := db.Query("SELECT id, name, age FROM users")
		if err != nil {
			c.JSON(500, gin.H{"success": false, "error": "failed to query DB"});
			return
		}
		defer rows.Close()

		// Fetch and print the results
		for rows.Next() {
			var id int
			var name string
			var age int
			err := rows.Scan(&id, &name, &age)
			if err != nil {
				c.JSON(500, gin.H{"success": false, "error": "failed to scan row"});
				return;
			}
		}

    // c.JSON(200, gin.H{"success": true, "data": id})
  })
	router.GET("/user", func(c *gin.Context) {
		var data struct {
			Name string `json:"name"`
			UUID string `json:"uuid"`
		}
		
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"error": gin.H{
					"type": "Validation",
					"message": "Invalid JSON format in request body",
				},
			})
			return
		}

		if data.Name == "" && data.UUID == "" {
			c.JSON(400, gin.H{
				"success": false,
				"erorr": gin.H{
					"type": "Validation",
					"message": "name or uuid must be passed in request body",
				},
			})
			return
		}

		db, err := openDB()
		if err != nil { panic(err) }

		if data.Name != "" { // if name is empty, use uuid
			stmt, err := db.Prepare("SELECT id, name, uuid FROM users WHERE name=? LIMIT 1")
			if err != nil {
				panic(err)
			}
			_, err = stmt.Exec(data.Name)
		} else { // use name
			stmt, err := db.Prepare("SELECT id, name, uuid FROM users WHERE uuid=? LIMIT 1")
			if err != nil {
				panic(err)
			}
			_, err = stmt.Exec(data.UUID)
		}
	})
	router.POST("/user/:name", func(c *gin.Context) {
		name := c.Param("name")

    userUUID := uuid.New()

    db, err := openDB()
    if err != nil { panic(err) }
    stmt, err := db.Prepare("INSERT INTO users (name, uuid) VALUES (?, ?)")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(name, userUUID)
		if err != nil {
			panic(err)
		}

    c.JSON(200, gin.H{
    	"success": true,
      "data": userUUID,
    })
	})
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
func openDB() (*sql.DB, error) {
	return sql.Open("sqlite3", env.Get("DB_PATH"))
}
