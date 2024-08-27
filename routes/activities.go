package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Activity struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Athlete  string    `json:"athlete"`
	Distance float64   `json:"distance"`
	Time     float64   `json:"time"`
	Date     time.Time `json:"date"`
}

var activities = []Activity{
	{ID: "1", Title: "Blue Train", Athlete: "John Coltrane", Distance: 56.99, Time: 56.99, Date: time.Now()},
	{ID: "2", Title: "Jeru", Athlete: "Gerry Mulligan", Distance: 17.99, Time: 56.99, Date: time.Now()},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Athlete: "Sarah Vaughan", Distance: 39.99, Time: 56.99, Date: time.Now()},
}

func GetActivities(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, activities)
}
