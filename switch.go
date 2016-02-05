package main

import(
	"net/http"
	//Import all the handlers required for these servers
	"handlers/app1"
	"handlers/app1_ver2"
)

func main(){
	var funcmap map[string]func(http.ResponseWriter, *http.Request)
	funcmap = make(map[string]func(http.ResponseWriter, *http.Request))

	//create a map of method nicks to actual func's
	//Note: management and maintenance of this mapping will have an overhead with time.
	funcmap["sayhello_v1"]=app1.SayHello
	funcmap["sayhello_v2"]=app1_ver2.SayHelloV2
}

