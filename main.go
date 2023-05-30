package main

import HelloController "bxb_home/internal/server/http"

func main() {
	hello := HelloController.New()
	hello.Start()
	//addr := os.Getenv("ADDR")
	//
	//mux := http.NewServeMux()
	//helloController := Server.New()
	//mux.HandleFunc("/v1/hello", helloController.HelloHandler)
	//
	//log.Printf("server is listening at %s", addr)
	//log.Fatal(http.ListenAndServe(":3000", mux))
}
