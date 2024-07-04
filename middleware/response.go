package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type responseWriter struct {
	gin.ResponseWriter
	body []byte
}

func (w *responseWriter) Write(data []byte) (int, error) {
	w.body = append(w.body, data...)
	return w.ResponseWriter.Write(data)
}

func ResponseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := os.OpenFile("logfile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		defer file.Close()

		customWriter := &responseWriter{ResponseWriter: c.Writer, body: make([]byte, 0)}
		c.Writer = customWriter

		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		executionTime := endTime.Sub(startTime)

		responseData := gin.H{}
		if response, ok := c.Get("Response"); ok {
			responseData["Response"] = response
		}
		if response, ok := c.Get("Error"); ok {
			responseData["Error"] = response
		}

		route := c.FullPath()
		clientIP := c.ClientIP()
		httpStatusCode := c.Writer.Status()

		logData := gin.H{
			"ExecutionTime":  executionTime.String(),
			"Route":          route,
			"HttpStatusCode": httpStatusCode,
			"IP":             clientIP,
		}
		combinedData := gin.H{
			"LogData": logData,
			"Data":    responseData,
		}

		logString := fmt.Sprintf("ExecutionTime: %s, Route: %s, HttpStatusCode: %d, IP: %s\n",
			logData["ExecutionTime"], logData["Route"], logData["HttpStatusCode"], logData["IP"])

		if _, err := file.WriteString(logString); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		responseJSON, err := json.Marshal(combinedData)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		// Envia a resposta apenas se ainda não tiver sido enviada
		if !c.Writer.Written() {
			c.Header("Content-Type", "application/json")
			c.Writer.WriteHeader(httpStatusCode)
			c.Writer.Write(responseJSON)
		}
	}
}
