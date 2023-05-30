package middleware

import (
	"log"
	"net/http"
	"time"
)

//type Logger struct {
//}

//func New() *Logger {
//	return &Logger{}
//}
//
//func (l Logger) Log(request *http.Request) {
//	start := time.Now()
//	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
//	infoLog.Printf("%s %s %s", request.Method, request.RequestURI, time.Since(start))
//}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Println(req.Method, req.RequestURI, time.Since(start))
	})
}
