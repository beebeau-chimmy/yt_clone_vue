package main

import (
    f "fmt"
    // Fiber Web Framework
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get( "/", func(c *fiber.Ctx) error {
        return c.SendString("Hello World!")
    } )

    serve_err := app.Listen(":3000")
    if serve_err != nil { f.Println(serve_err) } else { f.Println(serve_err) }
}
