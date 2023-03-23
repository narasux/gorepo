package payload

import (
	"fmt"
	"math/rand"
	"runtime"
	"runtime/debug"
	"sync"
	"time"

	// https://gaocegege.com/Blog/maxprocs-cpu
	_ "go.uber.org/automaxprocs"
)

// CPUBurner provides a way to quickly increase cpu usage
type CPUBurner struct{}

// Incr increase cpu usage
func (C *CPUBurner) Incr(size int) {
	numCPU := runtime.NumCPU()
	fmt.Printf("Using %d CPU cores\n", numCPU)
	fmt.Printf("Run fib(30) %d times on each cpu core...\n", size)
	runtime.GOMAXPROCS(numCPU)

	var wg sync.WaitGroup
	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < size; j++ {
				fib(30)
			}
		}()
	}
	wg.Wait()
	fmt.Println("Done!")
}

var _ LoadGenerator = &CPUBurner{}

// MemoryHog provides a way to quickly increase memory usage
type MemoryHog struct {
	mem [][]byte
}

// Incr increase memory usage
func (m *MemoryHog) Incr(size int) {
	fmt.Printf("Allocate %d MB memory...\n", size)

	for i := 0; i < size; i++ {
		// 1 MB every time
		b := make([]byte, 1024*1024)
		for j := range b {
			b[j] = byte(rand.Intn(256))
		}
		m.mem = append(m.mem, b)
		// wait 50 ms for allocate memory
		time.Sleep(time.Millisecond * 10)
	}
	fmt.Println("Allocated!")
	m.free()
}

// free allocated memory
func (m *MemoryHog) free() {
	fmt.Println("Start free allocated memory...")
	for i := range m.mem {
		m.mem[i] = nil
	}
	m.mem = nil

	// debug.FreeOSMemory() free allocated memory faster than runtime.GC()
	debug.FreeOSMemory()
	fmt.Println("Freed!")
}

var _ LoadGenerator = &MemoryHog{}

// New according to payload type return LoadGenerator
func New(t PayLoadType) LoadGenerator {
	switch t {
	case TypeCPU:
		return &CPUBurner{}
	case TypeMem:
		return &MemoryHog{}
	}
	return nil
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
