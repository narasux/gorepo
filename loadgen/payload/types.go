package payload

type LoadGenerator interface {
	// Incr according to parameter size, quickly increase payload
	Incr(size int)
}

// PayLoadType describes an available resource in container
type PayLoadType string

// TypeCPU ...
const TypeCPU PayLoadType = "cpu"

// TypeMem ...
const TypeMem PayLoadType = "mem"

// DefaultCPURunTimes means run fib(30) 10 times when size unset
const DefaultCPURunTimes = 10

// DefaultCPURunTimes means allocate 5MB memory when size unset
const DefaultMemAllocateSize = 5
