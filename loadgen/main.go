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
	size, err := strconv.Atoi(req.URL.Query().Get("size"))
	if err != nil {
		size = payload.DefaultCPURunTimes
	}
	_, _ = fmt.Fprintf(w, "Run fib(30) %d times in %s at %s", size, hostname, time.Now().Format(time.RFC1123Z))

	payload.New(payload.TypeCPU).Incr(size)

}

// IncrMemoryLoad ...
func IncrMemoryLoad(w http.ResponseWriter, req *http.Request) {
	hostname, _ := os.Hostname()
	size, err := strconv.Atoi(req.URL.Query().Get("size"))
	if err != nil {
		size = payload.DefaultMemAllocateSize
	}
	_, _ = fmt.Fprintf(w, "Allocate %d MB memory in %s at %s", size, hostname, time.Now().Format(time.RFC1123Z))

	payload.New(payload.TypeMem).Incr(size)
}
