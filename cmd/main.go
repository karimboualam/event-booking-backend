// Entrée principale de l'application
package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "os"
)

func main() {
    r := gin.Default()

    // Routes ici (exemple)
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Listening on port %s...", port)
    r.Run(":" + port)
}
