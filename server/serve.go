package server

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
)

var tmpl *template.Template

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func displayInfo(port string) {
	var buffer bytes.Buffer
	buffer.WriteString("\n    You are listening on:\n")

	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		// HACK! to get ipv4 address
		ip := addr.String()
		if strings.ContainsRune(ip, '.') {
			ip = strings.Split(ip, "/")[0]
			buffer.WriteString("       • " + ip + ":" + port + "\n")
		}
	}
	fmt.Println(buffer.String())
}

func handler(c *gin.Context) {
	filepath := c.Param("filepath")
	filepath = filepath[1:len(filepath)]

	if filepath == "" {
		c.File("index.html")
		return
	}

	c.File(filepath)
}

func Serve(port string) {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/*filepath", handler)
	router.HEAD("/*filepath", handler)

	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	displayInfo(port)
	log.Fatal(router.Run(addr))
}
