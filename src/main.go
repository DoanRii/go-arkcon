package main

import (
  "github.com/DoanRii/go-arkrcon"
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

func ArkShopAddPoints(c *gin.Context) {
  svr := c.Param("svr")
  key := c.Param("key")
  steamID, err := strconv.Atoi(c.Param("steamID"))
  amount, err := strconv.Atoi(c.Param("amount"))
  ark, err := arkrcon.NewARKRconConnection(svr, key)
  if err != nil {
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Auth fail!"})
    return
  }
  ark.ArkShopAddPoints(steamID, amount)
  c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Ok!"})
}

func main() {
  router := gin.Default()
    router.GET("/slomo/:svr/:key/:val", Slomo)
    router.GET("/arkshopaddpoints/:svr/:key/:steamID/:amount", ArkShopAddPoints)
    router.Run(":9222")
}