package util

import "github.com/gin-gonic/gin"

// Response struct
type Response struct {
	Status bool        `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Token  string      `json:"token"`
}

// ErrorJSON : json error response function
func ErrorJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"status": statusCode, "msg": data})
}

// SuccessJSON : json error response function
func SuccessJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"status": true, "msg": data})
}
