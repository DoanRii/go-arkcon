package main

import (
	"net/http"
	"strconv"

	"github.com/DoanRii/go-arkrcon"
	"github.com/gin-gonic/gin"
)

/*
Default Ark Command
*/
func Slomo(c *gin.Context) {
	ark, err := arkrcon.NewARKRconConnection(c.Param("svr"), c.Param("key"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid server info"})
		return
	}
	val, err := strconv.Atoi(c.Param("val"))
	ark.Slomo(val)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

/*
Plugins Command
*/
func loadPlugins(c *gin.Context) {
	ark, err := arkrcon.NewARKRconConnection(c.Param("svr"), c.Param("key"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid server info"})
		return
	}
	plugins := c.Param("plugins")
	ark.pluginsLoad(plugins)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func main() {
	router := gin.Default()
	router.GET("/gamespeed/:svr/:key/:val", Slomo)
	router.GET("/load/:svr/:key/:plugins", loadPlugins)
	router.Run(":9222")
}
