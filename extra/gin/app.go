package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
)

type Image struct {
	Url string `form:"url"`
}

func getImg(url string) (err error, name string) {
	name = filepath.Join("download", uuid.New())
	fmt.Println(name)
	out, err := os.Create(name)
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err, ""
	}
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	_, err = io.Copy(out, bytes.NewReader(pix))
	return
}

func get(c *gin.Context) {
	var image Image
	if c.ShouldBind(&image) == nil {
		url := image.Url
		err, path := getImg(url)
		if err != nil {
			c.JSON(200, gin.H{
				"message": "error",
			})
		} else {
			c.JSON(200, gin.H{
				"message": path,
			})
		}
	}
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func flag(c *gin.Context) {
	ip := c.ClientIP()
	if ip == "127.0.0.1" {
		c.JSON(200, gin.H{
			"flag": "flag{}",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "your ip no 127.0.0.1?",
		})
	}
}

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Static("/download", "./download")
	r.Static("/fonts", "./static/fonts")
	r.Static("/images", "./static/images")
	r.Static("/stylesheets", "./static/stylesheets")
	r.Static("/javascripts", "./static/javascripts")

	r.GET("/", index)
	r.POST("/get", get)
	r.GET("/flag", flag)
	r.Run(":8088")
}
