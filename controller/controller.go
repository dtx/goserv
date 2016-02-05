package controller

import(
	"io"
	"net/http"
	"handlers/app1"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

func startServe(){
	//read muxmapping directory
	//create a mux map for every mapping in an []
	//create a server for every muxmapping in an []
	MyServer := http.Server{
		Addr: ":8000",
		Handler: &myHandler{
			name: "Darshan",
		},
	}

	//init the map and adding a default example value
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = app1.SayHello
	MyServer.ListenAndServe()
}


//defining an empty handler.
//see if we can add things to it in the future.
type myHandler struct{
	name string
}

func (p *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, p.name+" "+r.URL.String()+" does not exist")
}

