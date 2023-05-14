package main

import (
  "net/http"
  "github.com/firecub/fib"
  "github.com/gin-gonic/gin"
  "strconv"
)

func main() {
    routerGroup := setUpRouterGroup()
    routerGroup.Run("localhost:8080")
}

func setUpRouterGroup() *gin.Engine {
    routerGroup := gin.Default()
    routerGroup.GET("/fib/number/:index", getFibonacciNumberWithIndex)
    routerGroup.GET("/fib/index/:number", getFloorIndexOfFibonacciNumber)
    return routerGroup
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

