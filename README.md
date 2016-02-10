GoServ 
===============
### A simple server for declarative HTTP servers.

This is still under construction, just publishing for posterity and comments if any.

##How it works

GoServ is a pattern over Go's HTTP module and leverages the handler and multiplexor deafault interfaces
that come with the module.

GoServ needs 2 user inputs to work:
- A set of multiplexor mappings, each set is a declaration of an HTTP server. These mappings lie in the 
$GOPATH/muxmapping directory on your system.
These are maps of <<routes>,<method nicknames>>
```
/ sayhello_v1
/try try
```
After a while, I release a version 2 api with new endpoints, does not mean I have to give up on my previosu version endpoints.
```
/ sayhello_v2
/old sayhello_v1
/spanish sayhola
/india sayhindi
/whoami try
```
- You just import GoServ, teach it names of your routes in mappings and start the server.
```go
package main

import(
    "net/http"
    //Import all the handlers required for these servers
    "github.com/dsanghani/goserv/sample_handlers/app1"
    "github.com/dsanghani/goserv/sample_handlers/app1_ver2"
    "github.com/dsanghani/goserv"
)

func main(){
    var funcmap map[string]func(http.ResponseWriter, *http.Request)
    funcmap = make(map[string]func(http.ResponseWriter, *http.Request))

    //create a map of method nicks to actual func's
    //Note: management and maintenance of this mapping will have an overhead with time.
    funcmap["sayhello_v1"]=app1.SayHello
    funcmap["sayhello_v2"]=app1_ver2.SayHelloV2
    funcmap["try"]=app1.TryMe
    funcmap["sayhola"]=app1_ver2.SayHolaV1
    funcmap["sayhindi"]=app1_ver2.SayNamaste

    goserv := goserv.GoServ{}
    goserv.LearnNames(funcmap)
    goserv.StartServe()
}
```
As of now GoServ chooses random ports, which will be changed in the future, and starts all the mappings 
you provided as a HTTP server. Thus, giving you a concurrent set of HTTP server serving your applications.

##Uses
A variety of uses come in mind if all your application entry points satisfy the Go HTTP handler signatures
i.e. `func(w http.ResponseWriter, r * Request)`
* Profiling
* A/B testing
* Algorithm Correctness
* Multiple Backend poc's

The idea is to not re-create a whole HTTP application for some routes that need tweaking, with GoServ you can
mix and match different handlers in a declarative model and serve them over HTTP.

## License

[BSD license](http://opensource.org/licenses/bsd-license.php)

