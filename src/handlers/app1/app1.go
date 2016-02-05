package app1

import(
	"net/http"
	"fmt"
)

func SayHello(r http.ResponseWriter, w *http.Request){
	fmt.Println("Hello World!")
}
