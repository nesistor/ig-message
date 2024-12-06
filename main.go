package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"ig-message/graph"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Set up GraphQL server and playground routes using Gin
	r.GET("/playground", func(c *gin.Context) {
		playground.Handler("GraphQL Playground", "/graphql").ServeHTTP(c.Writer, c.Request)
	})

	// Create and serve GraphQL server using gqlgen's generated executable schema
	r.POST("/graphql", func(c *gin.Context) {
		// Use the generated executable schema
		srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
		srv.ServeHTTP(c.Writer, c.Request)
	})

	// WebSocket setup to listen for new Instagram Likes (stubbed as WebSocket service)
	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()

		for {
			// Simulate listening for a "like" event
			msgType, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				break
			}

			// Send notification with image from the images folder
			if string(p) == "like" {
				fileName := fmt.Sprintf("images/%d.jpg", time.Now().Unix())
				conn.WriteMessage(msgType, []byte(fmt.Sprintf("New Like received! Check out the image: %s", fileName)))
			}
		}
	})

	// Start the Gin server on port 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
