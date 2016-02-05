package app1_ver2

import(
	"net/http"
	"fmt"
)

func SayHelloV2(r http.ResponseWriter, w *http.Request){
	fmt.Println("Hello World in a new fashion!")
}
