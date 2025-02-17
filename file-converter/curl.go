// Lancer le serveur sur le port 8080; pour tester en curl

package main

import (
	"fmt"
	"file-converter/api"
	"file-converter/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("🚀 Démarrage du serveur API...")
	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
