package main

import (
  "github.com/DoanRii/go-pkg-ark/arkrcon"
  "net/http"
  "github.com/gin-gonic/gin"
  "strconv"
)

func Slomo(c *gin.Context) {
  svr := c.Param("svr")
  key := c.Param("key")
  val, err := strconv.Atoi(c.Param("val"))
  ark, err := arkrcon.NewARKRconConnection(svr, key)
  if err != nil {
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Auth fail!"})
    return
  }
  ark.Slomo(val)
  c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Ok!"})
}

func main() {
  
  router := gin.Default()
    router.GET("/slomo/:svr/:key/:val", Slomo)

    router.Run("localhost:9222")
  
}