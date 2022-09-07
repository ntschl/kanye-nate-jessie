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

type Quote struct {
	Quote string
}

func main() {
	r := gin.Default()
	r.GET("/kanye", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Inspirational Wisdom": getQuote(),
			"Team":                 "Gold Team Rules!",
		})
	})
	r.Run() //default's on 8080
}

func getQuote() string {
	response, err := http.Get("https://api.kanye.rest")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	q := Quote{}

	err = json.Unmarshal(responseData, &q)
	if err != nil {
		panic(err)
	}

	return q.Quote
	// return string(responseData)
}
