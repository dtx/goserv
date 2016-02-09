package goserv

import(
	"io"
	"fmt"
	"sync"
	"errors"
	"net/http"
	"math/rand"
)


type GoServ struct{
	rosettaStone map[string]func(http.ResponseWriter, *http.Request)
	wg sync.WaitGroup
}

func (gs *GoServ) LearnNames(dictionary map[string]func(http.ResponseWriter, *http.Request)){
	gs.rosettaStone = dictionary
}

func (gs* GoServ) doIKnowToRead() bool{
	if gs.rosettaStone != nil{
		return true
	}
	return false
}

func (gs *GoServ) translateMappingToMethods(raw_muxmappings []map[string]string) ([]map[string]func(http.ResponseWriter, *http.Request), error){
	if (gs.doIKnowToRead() == false){
		return nil, errors.New("Before starting servers, please teach me the names you gave to your routes!")

	}
	var muxmappings []map[string]func(http.ResponseWriter, *http.Request)
	var mux map[string]func(http.ResponseWriter, *http.Request)
	for _,raw_muxmapping := range raw_muxmappings{
		for route,nickname := range raw_muxmapping{
			//translate v to hold func() obj
			method := gs.rosettaStone[nickname]
			mux[route] = method
		}
		muxmappings = append(muxmappings, mux)
	}
	return muxmappings, nil
}

func (gs *GoServ) startServer(p map[string]func(http.ResponseWriter, *http.Request)){
	defer gs.wg.Done()
	//find a 'random' port between 9000 and 65535
	port := rand.Intn(65535-9000) + 9000
	//todo(darshan): need to add a check if a port number is already used
	MyServer := http.Server{
	Addr: ":"+string(port),
	Handler: &myHandlers{
		muxmapping : p,
		},
	}
	fmt.Printf("Starting a server on port %d\n", port)
	MyServer.ListenAndServe()
}
func (gs *GoServ) StartServe() int{
	//read muxmapping directory
	//create a mux map for every mapping in an []
	//create a server for every muxmapping in an []
	//init the map and adding a default example value
	raw_muxmappings := Readallmapping()
	translated_muxmappingsgs, err := gs.translateMappingToMethods(raw_muxmappings)
	if err != nil {
		fmt.Println("Failure in reading multiplexor names.")
		return 0
	}
	for _,mux := range translated_muxmappingsgs{
		//start a thread for each server here
		gs.wg.Add(1)
		go gs.startServer(mux)
	}
	gs.wg.Wait()
	return 1
}


//a custom handler that holds a route mapping
type myHandlers struct{
	muxmapping map[string]func(http.ResponseWriter, *http.Request)
}

func (p *myHandlers) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if h, ok := p.muxmapping[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, r.URL.String()+" does not exist")
}

