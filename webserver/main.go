package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// For convenience, this webserver does not add authentication, please do not include any
	// sensitive information in the provided file content. The random prefix can be used to ensure that
	// the service is "not easy" to be accessed by unexpected people. If it is leaked,
	// please replace RANDOM_PATH_PREFIX in time for security
	randomPathPrefix, _ := os.LookupEnv("RANDOM_PATH_PREFIX")
	assetsDir, _ := os.LookupEnv("ASSETS_DIR")

	path := fmt.Sprintf("/%s/static/", randomPathPrefix)
	fmt.Println("UrlPath:", path, "\nAssetsDir:", assetsDir)

	fs := http.FileServer(http.Dir(assetsDir))
	http.Handle(path, http.StripPrefix(path, fs))

	fmt.Println("Starting server on port 8080...")
	_ = http.ListenAndServe(":8080", nil)
}
