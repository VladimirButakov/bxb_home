package main

import HelloController "bxb_home/internal/server/http"

func main() {
	hello := HelloController.New()
	hello.Start()
}
