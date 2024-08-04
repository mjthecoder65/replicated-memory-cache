package cache

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	cache      = make(map[string]string)
	cacheMutex = &sync.RWMutex{}
)

var (
	nodes = []string{"cache-node-1:5000", "cache-node-2:5000", "cache-node-3:5000"}
)

func NotifyNodes(key, value string) {
	var wg sync.WaitGroup

	for _, node := range nodes {
		wg.Add(1)
		go func(node string) {
			defer wg.Done()
			url := fmt.Sprintf("http://%s/sync", node)
			_, err := http.Post(url, "application/json", generateRequestBody(key, value))
			if err != nil {
				log.Printf("Failed to notify node %s: %v", node, err)
			}
		}(node)
	}

	wg.Wait()
}

func generateRequestBody(key, value string) *strings.Reader {
	requestData := map[string]string{"key": key, "value": value}
	jsonData, _ := json.Marshal(requestData)
	return strings.NewReader(string(jsonData))
}

func GetKeyHandler(c *gin.Context) {
	key := c.Param("key")
	cacheMutex.RLock()
	value, exists := cache[key]
	cacheMutex.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"value": "Not Found"})
	}

	c.JSON(http.StatusOK, gin.H{"value": value})
}

func SetKeyHandler(c *gin.Context) {
	var requestData map[string]string

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	key, exists := requestData["key"]
	value, existsValue := requestData["value"]

	if !exists || existsValue {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
	}

	cacheMutex.Lock()
	cache[key] = value
	cacheMutex.Unlock()

	// Notifiy other nodes to udpate the data.
	NotifyNodes(key, value)

	// Return response
	c.JSON(http.StatusOK, gin.H{"status": "Success"})
}

func SyncDataHandler(c *gin.Context) {
	var payload map[string]string

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key, exists := payload["key"]
	value, valueExists := payload["value"]

	if !exists || !valueExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
	}

	cacheMutex.Lock()
	cache[key] = value
	cacheMutex.Unlock()

	c.JSON(http.StatusOK, gin.H{"status": "Sync Data Successfully"})
}

func AddNode(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Added node successfully."})
}
