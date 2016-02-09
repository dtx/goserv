package app1

import(
	"net/http"
	"fmt"
)

func SayHello(r http.ResponseWriter, w *http.Request){
	//do something
	fmt.Println("Hello World!")
}

func TryMe(r http.ResponseWriter, w *http.Request){
	//do something
	fmt.Println("Trying everything, still learning all languages!")
}
