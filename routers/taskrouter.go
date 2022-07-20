package routers

import (
	"net/http"

	jwtAuth "todo/authentication"
	"todo/models"

	"github.com/gin-gonic/gin"
)

type work struct {
	ID          string `json:"id"`
	Events      string `json:"events"`
	Description string `json:"description"`
}

var DB = models.Init()

func RegisterRoutes() *gin.Engine {
	router := gin.Default()

	// routes
	router.GET("/task", getTasks)
	router.GET("/task/:id", getTaskById)
	router.POST("/task", postTasks)
	router.POST("/signup", jwtAuth.SignUp)
	router.DELETE("/task/:id", deleteTask)
	router.PUT("/task/:id", updateTask)

	return router
}

func getTasks(c *gin.Context) {
	//c.IndentedJSON(http.StatusOK, job)
	// id := c.Param("id")
	var allTask []work
	sqlStmnt := `SELECT * from assignment`
	out, err := DB.Query(sqlStmnt)
	if err != nil {
		panic(err)
	}
	for out.Next() {
		var id string
		var events string
		var description string
		err = out.Scan(&id, &events, &description)
		if err != nil {
			panic(err)
		}
		allTask = append(allTask, work{ID: id, Events: events, Description: description})

	}
	c.IndentedJSON(http.StatusOK, allTask)
	/*for _, a := range todos {
	  if a.ID == id {
	    c.IndentedJSON(http.StatusOK, a)
	    return
	  }
	}*/
	//c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
func postTasks(c *gin.Context) {
	var newtodo work
	if err := c.BindJSON(&newtodo); err != nil {
		return
	}

	sqlStmnt := `INSERT INTO assignment (id,events,description) VALUES ($1,$2,$3)`
	_, err := DB.Exec(sqlStmnt, newtodo.ID, newtodo.Events, newtodo.Description)
	if err != nil {
		panic(err)
	}
	// job = append(job, newtodo)
	c.IndentedJSON(http.StatusCreated, newtodo)
}

func getTaskById(c *gin.Context) {
	id := c.Param("id")
	var allTask []work

	sqlStmnt := `SELECT * from Assignment WHERE ID IN ($1)`
	out, err := DB.Query(sqlStmnt, id)
	if err != nil {
		panic(err)
	}

	for out.Next() {
		var id string
		var events string

		var description string
		err = out.Scan(&id, &events, &description)
		if err != nil {
			panic(err)
		}
		allTask = append(allTask, work{ID: id, Events: events, Description: description})

	}
	c.IndentedJSON(http.StatusOK, allTask)
}
func deleteTask(c *gin.Context) {
	id := c.Param("id")
	sqlStatement := `DELETE FROM Assignment WHERE id = $1;`
	_, err := DB.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}
}

func updateTask(c *gin.Context) {
	id := c.Param("id")
	var update work

	if err := c.BindJSON(&update); err != nil {
		return
	}

	sqlStmnt := `UPDATE Assignment SET task=$1, description=$2 WHERE id=$3`
	_, err := DB.Exec(sqlStmnt, update.Events, update.Description, id)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "updated successfully",
	})

}
