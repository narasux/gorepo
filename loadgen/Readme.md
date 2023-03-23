# goloadgen (Golang PayLoad Generator)

A simple HTTP server that supports quickly increasing CPU/memory usage by calling APIs, generally used for stress or k8s HPA scaling testing. Please do not use it for malicious purposes!

## How to use

```shell
# build goloadgen docker image
make docker

# run goloadgen on docker
docker run -p 8080:8080 --name goloadgen -d goloadgen:1.0.0

# curl http://127.0.0.1:8080/api/cpu to improve cpu usage
# Parameter `size` means to calculate Fibonacci sequence [fib(30)] N times on each CPU
curl http://127.0.0.1:8080/api/cpu?size=500

# curl http://127.0.0.1:8080/api/mem to improve memory usage
# Parameter `size` indicates how much memory will be allocated, in MB
curl http://127.0.0.1:8080/api/mem?size=1024
```
