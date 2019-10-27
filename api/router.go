package api

import (
	"database/sql"
	"fmt"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Creates and returns a Gin Router
func GinRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.GET("/employees", func(c *gin.Context) {
		emps, err := GetAllEmployees(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, emps)
	})

	router.GET("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		emp, err := GetEmployee(db, uint(idInt))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, emp)
	})

	router.DELETE("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		_, err = DeleteEmployee(db, uint(idInt))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	// To update a single property
	router.PATCH("/employees/:id", func(c *gin.Context) {
		var id uint
		if err := c.ShouldBindUri(&id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	})

	// To update the whole resource
	router.PUT("/employees/:id", func(c *gin.Context) {
		var form Employee
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		err = c.ShouldBindJSON(&form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		_, err = GetEmployee(db, uint(idInt))
		if err != nil {
			fmt.Println("here")
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		form.ID = uint(idInt)
		rows, err := EditEmployee(db, &form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if rows > 1 {
			c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Expected only 1 row to be updated, got %d", rows)})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.POST("/employees", func(c *gin.Context) {
		var form Employee
		err := c.Bind(&form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
		}

		id, _, err := CreateEmployee(db, &form)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		}

		c.JSON(http.StatusOK, id)
	})

	return router
}
