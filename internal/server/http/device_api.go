package http

import "github.com/gin-gonic/gin"

func (s *Server) GetDevice(c *gin.Context) {
	deviceId := c.Query("device_id")
	s.App.Device.GetDevice(c, deviceId)
}
