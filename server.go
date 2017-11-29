package main
 
import (
	"fmt"
	"net/http"
	"runtime"
)
 
// request/response handling
func indexHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Running on %s with an %s CPU.\n", 
		runtime.GOOS,     // operating system
		runtime.GOARCH)   // machine architecture
	fmt.Fprintf(res, "User agent for request is %s.\n", req.UserAgent())
}
 
func main(){
	http.HandleFunc("/", indexHandler) // routing: URI / dispatched to indexHandler
	fmt.Printf("Awaiting requests for URI / on port 8080...\n")
	http.ListenAndServe(":8080", nil)  // nil = no error handler
}
