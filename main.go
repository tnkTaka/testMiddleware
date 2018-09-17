package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/makki0205/gojwt"
	"fmt"
)

func main() {
	r := gin.Default()
	r.GET("/login", AuthMiddleware(), func(c *gin.Context) {
		fmt.Println(c.Query("token"))
		c.JSON(200, gin.H{"message": "OK!"})
	})

	r.GET("/add/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if id <= 0 {
			c.JSON(400, gin.H{"message": "invalid id"})
		}

		Jwt(id)
		c.JSON(200, gin.H{"message": "Hello! USER"+strconv.Itoa(id)})
	})


	r.Run(":8000")
}

func Jwt(id int)  {
	jwt.SetSalt("D79998A7-3F2B-4505-BE2B-6E68500AAE37")
	jwt.SetExp(60 * 60 * 24)

	claims := map[string]string{
		"user": "user"+strconv.Itoa(id),
	}
	token := jwt.Generate(claims)
	fmt.Println(token)
}

func JwtAuthentication(token string)  error {
	_, err := jwt.Decode(token)
	if err != nil{
		return err
	}
	return nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		err := JwtAuthentication(token)
		if err != nil {
			c.JSON(400, gin.H{"message": "Authentication Failure" })
			c.Abort()
		}

	}
}