// internal/app/app.go
package app

import (
	"fmt"
	"millennium/web"
)

// Start initializes and starts the application.
func Start() {
    fmt.Println("POS system started!")

    // Initialize and start the web server
    web.InitializeServer()
}
