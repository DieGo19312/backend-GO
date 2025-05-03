package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(c *fiber.Ctx) error {
    authHeader := c.Get("Authorization")
    if authHeader == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "No se proporcionó un token de autorización",
        })
    }
    tokenString := authHeader[len("Bearer "):]

    secretKey := os.Getenv("JWT_SECRET") 
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fiber.NewError(fiber.StatusUnauthorized, "Método de firma inválido")
        }
        return []byte(secretKey), nil
    })
    if err != nil || !token.Valid {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "Token inválido o expirado",
        })
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "No se pudieron extraer las reclamaciones del token",
        })
    }

    roleID, ok := claims["roleID"].(float64) 
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "Invalid role ID in token",
            "log":     fmt.Sprintf("Invalid role ID in token: %v", claims["role_id"]),
        })
    }

    c.Locals("userID", claims["id"])
    c.Locals("email", claims["email"])
    c.Locals("roleID", int(roleID)) 

    return c.Next()
}
