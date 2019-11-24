package api

import (
	"database/sql"
	"fmt"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

// GinRouter Creates and returns a Gin Router
func GinRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.GET("/employees", func(c *gin.Context) {
		emps, err := GetAllEmployees(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, emps)
	})

	router.GET("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		emp, err := GetEmployee(db, uint(idInt))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if (Employee{}) == emp {
			c.JSON(http.StatusNotFound, gin.H{})
			return
		}
		c.JSON(http.StatusOK, emp)
	})

	router.DELETE("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		_, err = DeleteEmployee(db, uint(idInt))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	// To update the whole resource a!
	router.PUT("/employees/:id", func(c *gin.Context) {
		var form Employee
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err = c.ShouldBindJSON(&form)
		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
			return
		}

		_, err = GetEmployee(db, uint(idInt))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		form.ID = uint(idInt)
		rows, err := EditEmployee(db, &form)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if rows > 1 {
			c.JSON(http.StatusOK, gin.H{"error": fmt.Sprintf("%d rows updated", rows)})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.POST("/employees", func(c *gin.Context) {
		var form Employee
		err := c.ShouldBind(&form)
		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{"message": err.Error()})
			return
		}

		id, _, err := CreateEmployee(db, &form)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id})
	})

	return router
}
