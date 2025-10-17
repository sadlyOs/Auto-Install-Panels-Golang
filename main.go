package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		ip := c.Query("ip")
		port := c.Query("port")
		login := c.Query("login")
		pass := c.Query("pass")

		if len(ip) <= 0 && len(port) <= 0 && len(login) <= 0 && len(pass) <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Put all the params: ip, port, login, pass",
			})
			return
		}

		sendRun(ip, port, login, pass)

		content, err := os.ReadFile("install.log")
		if err != nil {
			if os.IsNotExist(err) {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "Error read file log",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to read log file",
			})
			fmt.Fprintf(os.Stderr, "Error read file", err)
			return
		}
		c.Header("Content-Type", "text/plain; charset=utf-8")
		c.String(http.StatusOK, string(content))
	})

	r.Run()
}
