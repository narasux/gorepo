# gowebserver

A simple web server that supports accessing files in a specified directory through a browser.

## How to use (Docker)

```shell
# build gowebserver docker image
make docker

# run gowebserver service via docker, mount host directory to serve files
docker run -p 8080:8080 -v /home/assets:/data/workspace/assets/ -e RANDOM_PATH_PREFIX=ae2168cd --name gowebserver -d gowebserver:1.0.0

# access http://127.0.0.1:8080/ae2168cd/static/ on browser to access the files in directory
curl http://127.0.0.1:8080/ae2168cd/static/
```
