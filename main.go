package main

import (
	"log"
	"time"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	api := app.Group("/api", logger.New())

	// Define a route for the GET method on the root path '/'
	api.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	// Login route
	api.Post("/login", login)

	// JWT Middleware
	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	// }))

	// Restricted Routes
	api.Get("/restricted", Protected(), restricted)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3001"))
}

func Protected() func(fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
		ErrorHandler: func(c fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})

			} else {
				c.Status(fiber.StatusUnauthorized)
				return c.JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
			}
		},
	})
}

type User struct {
	User     string `json:"user"`
	Password string `json:"pass"`
}

func login(c fiber.Ctx) error {
	var user User
	if err := c.Bind().Body(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Throws Unauthorized error
	if user.User != "john" || user.Password != "doe" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"token": t})
}

func restricted(c fiber.Ctx) error {
	// user := c.Locals("USER").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// name := claims["name"].(string)
	//
	user := jwtware.FromContext(c)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	return c.JSON(fiber.Map{"msg": "Welcome " + name})
}
