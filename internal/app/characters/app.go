package characters

import (
	"fmt"
	"go-micro/internal/pkg/server"
)

func Run() {
	fmt.Println("Running characters")
	app := server.CreateServer("Characters")
	ApplyRoutes(app)
	app.Listen(":3000")
}
