package main

import (
	"github.com/code-tanks/golang-api/pkg/tanks"
	"encoding/json"
	"log"
	"net/http"
	// "sync"
)

func main() {
	tank := tanks.CreateTank()

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	http.HandleFunc("/request_commands", func(w http.ResponseWriter, r *http.Request) {
		tank.Run()

		tank.Mu.Lock()
		defer tank.Mu.Unlock()

		commands := make([]int, len(tank.Commands))
		copy(commands, tank.Commands)

		tank.Commands = nil

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(commands)
	})

	http.HandleFunc("/request_commands_by_event", func(w http.ResponseWriter, r *http.Request) {
		var event interface{}
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tank.OnEvent(event)

		tank.Mu.Lock()
		defer tank.Mu.Unlock()

		commands := make([]int, len(tank.Commands))
		copy(commands, tank.Commands)

		tank.Commands = nil

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(commands)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// package main

// import (
// 	"github.com/gin-gonic/gin"
// 	"net/http"
// 	// "bytes"
// 	"io/ioutil"
// 	"sync"
// 	// "time"
// )

// type Tank struct {
// 	commands []string
// 	mutex    sync.Mutex
// }

// func (t *Tank) run() {
// 	// Simulate tank running logic
// 	// ...
// }

// func (t *Tank) onEvent(event string) {
// 	// Simulate tank event handling logic
// 	// ...
// }

// var bot = &Tank{}

// func ping(c *gin.Context) {
// 	c.String(http.StatusOK, "pong")
// }

// func requestCommands(c *gin.Context) {
// 	bot.Mutex.Lock()
// 	defer bot.Mutex.Unlock()

// 	bot.run()
// 	commands := append([]string(nil), bot.Commands...)
// 	bot.Commands = nil

// 	c.JSON(http.StatusOK, commands)
// }

// func requestCommandsByEvent(c *gin.Context) {
// 	bot.Mutex.Lock()
// 	defer bot.Mutex.Unlock()

// 	var event string
// 	body, err := ioutil.ReadAll(c.Request.Body)
// 	if err != nil {
// 		c.String(http.StatusBadRequest, "Error reading request body")
// 		return
// 	}
// 	event = string(body)

// 	bot.onEvent(event)
// 	commands := append([]string(nil), bot.Commands...)
// 	bot.Commands = nil

// 	c.JSON(http.StatusOK, commands)
// }

// func main() {
// 	r := gin.Default()

// 	r.GET("/ping", ping)
// 	r.GET("/request_commands", requestCommands)
// 	r.POST("/request_commands_by_event", requestCommandsByEvent)

// 	r.Run(":8080")
// }
