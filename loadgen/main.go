package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/narasux/goloadgen/payload"
)

func main() {
	http.HandleFunc("/cpu", IncrCPULoad)
	http.HandleFunc("/mem", IncrMemoryLoad)

	fmt.Println("Starting server on port 8080...")
	_ = http.ListenAndServe(":8080", nil)
}

// IncrCPULoad ...
func IncrCPULoad(w http.ResponseWriter, req *http.Request) {
	hostname, _ := os.Hostname()
	startAt := time.Now().Unix()

	size, err := strconv.Atoi(req.URL.Query().Get("size"))
	if err != nil {
		size = payload.DefaultCPURunTimes
	}
	payload.New(payload.TypeCPU).Incr(size)

	_, _ = fmt.Fprintf(w, "Run fib(30) %d times in %s, cost time %ds", size, hostname, time.Now().Unix()-startAt)
}

// IncrMemoryLoad ...
func IncrMemoryLoad(w http.ResponseWriter, req *http.Request) {
	hostname, _ := os.Hostname()
	startAt := time.Now().Unix()

	size, err := strconv.Atoi(req.URL.Query().Get("size"))
	if err != nil {
		size = payload.DefaultMemAllocateSize
	}
	payload.New(payload.TypeMem).Incr(size)

	_, _ = fmt.Fprintf(w, "Allocate %d MB memory in %s, cost time %ds", size, hostname, time.Now().Unix()-startAt)
}
