package middleware

import (
	"final-task-pbi-rakamin-fullstack-MuhammadArsyadBirohmatillah/database"
	"final-task-pbi-rakamin-fullstack-MuhammadArsyadBirohmatillah/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
    // Cookie
    tokenString, err := c.Cookie("Authorization")

    if err!= nil {
        c.Redirect(http.StatusFound, "/login")
        return
        // c.AbortWithStatus(http.StatusUnauthorized)
    }

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC);!ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }

        return []byte(os.Getenv("SECRET")), nil
    })

    if err!= nil {
        c.Redirect(http.StatusFound, "/login")
        return
        // log.Fatal(err)
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        // Cek expiration
        if float64(time.Now().Unix()) > claims["exp"].(float64) {
            c.Redirect(http.StatusFound, "/login")
            return
            // c.AbortWithStatus(http.StatusUnauthorized)
        }

        // Cari user id
        var user models.User
        resultSearch := database.DB.First(&user, claims["sub"])

        if resultSearch.RowsAffected == 0 {
            c.Redirect(http.StatusFound, "/login")
            return
        }

        if resultSearch.Error!= nil {
            c.Redirect(http.StatusFound, "/login")
            return
        }

        c.Set("user", user)

        c.Next()
    } else {
        c.Redirect(http.StatusFound, "/login")
        return
        // c.AbortWithStatus(http.StatusUnauthorized)
    }
}