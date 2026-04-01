package main

import (
	"github.com/gin-gonic/gin"
)

type Pizza struct {
	ID    int
	nome  string
	preco float64
}

func main() {
	router := gin.Default()
	router.GET("/pizzas", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"pizzas":  "Toscana, Calabresa",
		})
	})
	router.Run()
	//	var pizzas []Pizza{
	//		{ID:1,nome:"toscana",49,50},
	//		{ID:2,nome:"marguerita",79,50},
	//		{ID:3,nome:"atum com queijo",69,50}
	//	}
	//
	// var toscana Pizza
	// var marguerita Pizza
	// var atumComQueijo Pizza
	//
	//	fmt.Printf(toscana)
}
