package route

import (
	"application/todo/common"

	"application/todo/tododatabase"

	"github.com/gin-gonic/gin"
)

func GetTodo(c *gin.Context) {
	var Alltodo []common.Newtodo

	rows, err := tododatabase.DB.Query("SELECT id, title, description,status,priority,created_at FROM todo")

	if err != nil {
		c.JSON(400, gin.H{"Error2": err})
		return
	}
	for rows.Next() {
		var todo common.Newtodo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.Priority, &todo.CreatedAt); err != nil {
			c.JSON(400, gin.H{"Error3": err})
			return
		}
		Alltodo = append(Alltodo, todo)

	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		c.JSON(400, gin.H{"Error4": err})
	}
	c.JSON(200, gin.H{"Message": "succesfully found", "data": Alltodo})
}

func AddTodo(c *gin.Context) {
	var todo common.Newtodo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"Error4": err})
		return
	}
	if todo.Title == "" {
		c.JSON(400, gin.H{"Error": "Need valid Title"})
		return
	}
	if todo.Description != "" {
		_, err := tododatabase.DB.Exec(`INSERT INTO todo(title, description, status, priority)
VALUES(?, ?, ?, ?)`, todo.Title, todo.Description, "pending", todo.Priority)
		if err != nil {
			c.JSON(400, gin.H{"Error5": err})
			return
		}
		c.JSON(200, gin.H{"List": todo})
		return
	}
	_, err := tododatabase.DB.Exec(`INSERT INTO todo(title, description, status, priority)
VALUES(?, ?, ?, ?)`, todo.Title, "", "pending", todo.Priority)
	if err != nil {
		c.JSON(400, gin.H{"Error5": err})
		return
	}
	c.JSON(200, gin.H{"Message": "Successfully created", "data": todo})
}
