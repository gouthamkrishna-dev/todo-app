package route

import (
	"application/todo/common"
	"sort"

	"application/todo/tododatabase"

	"strconv"

	"github.com/gin-gonic/gin"
)

func priorityValue(p string) int {
	switch p {
	case "high":
		return 3
	case "medium":
		return 2
	case "low":
		return 1
	default:
		return 0
	}
}

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

	sort.Slice(Alltodo, func(i, j int) bool {
		return priorityValue(Alltodo[i].Priority) > priorityValue(Alltodo[j].Priority)
	})
	c.IndentedJSON(200, gin.H{"Message": "succesfully found", "data": Alltodo})
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
		c.IndentedJSON(200, gin.H{"Message": "Successfully created", "data": todo})
		return
	}
	_, err := tododatabase.DB.Exec(`INSERT INTO todo(title, description, status, priority)
VALUES(?, ?, ?, ?)`, todo.Title, "", "pending", todo.Priority)
	if err != nil {
		c.JSON(400, gin.H{"Error5": err})
		return
	}
	c.IndentedJSON(200, gin.H{"Message": "Successfully created", "data": todo})
}

func GetTodobyId(c *gin.Context) {
	data := c.Param("id")
	val, err := strconv.Atoi(data)

	if err != nil {
		c.JSON(400, gin.H{"Error6": err})
		return
	}
	var GetTodobyId common.Newtodo
	err = tododatabase.DB.QueryRow("SELECT * FROM todo WHERE id = ?", val).Scan(&GetTodobyId.ID, &GetTodobyId.Title, &GetTodobyId.Description, &GetTodobyId.Status, &GetTodobyId.Priority, &GetTodobyId.CreatedAt)
	if err != nil {
		c.JSON(400, gin.H{"Error7": err})
		return
	}
	c.IndentedJSON(200, gin.H{"Message": "Successfully retrived by Id", "data": GetTodobyId})
}

func DeleteaTodobyId(c *gin.Context) {
	data := c.Param("deleteId")
	val, err := strconv.Atoi(data)

	if err != nil {
		c.JSON(400, gin.H{"Error6": err})
		return
	}
	result, err := tododatabase.DB.Exec(`DELETE FROM todo WHERE id = ?`, val)

	if err != nil {
		c.JSON(400, gin.H{"Error7": err})
		return
	}
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(400, gin.H{"Error7": err})
		return
	}

	c.IndentedJSON(200, gin.H{"Message:": "successfully deleted by Id", "data": rowsaffected})

}
