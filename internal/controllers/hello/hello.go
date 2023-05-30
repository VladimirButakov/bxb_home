package hello

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	fmt.Fprintf(w, `{"message":"Hello, %s!"}`, name)
	//logger := loggerServer.New()
	//logger.Log(request)
}
