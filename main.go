package main

import (
  "net/http"
  "github.com/firecub/fib"
  "github.com/gin-gonic/gin"
  "strconv"
)

func main() {
    routerGroup := gin.Default()
    routerGroup.GET("/fib/number/:index", getFibonacciNumberWithIndex)
    routerGroup.GET("/fib/index/:number", getFloorIndexOfFibonacciNumber)
    routerGroup.Run("localhost:8080")
}

func getFibonacciNumberWithIndex(c *gin.Context) {
    indexAsString := c.Param("index")
    index, indexErr := strconv.Atoi(indexAsString)
    if indexErr != nil {
        c.Status(http.StatusBadRequest)
        return
    }
    number := fib.FibonacciNumberWithIndex(index)
    c.String(http.StatusOK, "%d\n", number)
}

func getFloorIndexOfFibonacciNumber(c *gin.Context) {
    numberAsString := c.Param("number")
    number, numberErr := strconv.ParseInt(numberAsString, 10, 64)
    if numberErr != nil {
        c.Status(http.StatusBadRequest)
        return
    }
    index := fib.FibonacciFloorIndexFor(number)
    c.String(http.StatusOK, "%d\n", index)
}

