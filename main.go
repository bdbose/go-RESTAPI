package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Detail []struct {
		Source    string        `json:"source"`
		NewsURL   string        `json:"newsUrl"`
		ImgURL    string        `json:"imgUrl"`
		Title     string        `json:"title"`
		MetaExtra []string      `json:"metaExtra"`
		MetaTag   []interface{} `json:"metaTag"`
		Descp     string        `json:"descp"`
	} `json:"detail"`
}


func main() {
	r := gin.Default()

	response,err:= http.Get("https://lit-woodland-31639.herokuapp.com/world")
	if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
	var data Response
	json.Unmarshal(responseData,&data)

	r.GET("/",func(c *gin.Context) {
		c.JSON(200,data)
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/",func (c *gin.Context)  {
		res,err:= ioutil.ReadAll(c.Request.Body)
		if err!=nil{
			c.JSON(200,gin.H{
				"message": err.Error(),
			})
		}
		var Response interface{}
		json.Unmarshal(res,&Response)
		c.JSON(200,Response)
	})
	r.Run(":5000")
}