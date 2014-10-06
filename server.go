package boxcars

import (
//	"fmt"
	"net/http"
)

func Listen(port string) {
	debug("Starting at %d", port)
	http.HandleFunc("/", OnRequest)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		debug("Fatal: %v", err)
	}
}
