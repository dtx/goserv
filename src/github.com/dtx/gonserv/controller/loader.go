package controller

import(
	"io/ioutil"
	"os"
	"fmt"
	"strings"
	//"strconv"
)

func add_to_path(path string) string{
	path_components := strings.Split(path, "/")
	if path_components[len(path_components)-1] == "src"{
		return path+"/muxmapping/"
	}
	return path+"/src/muxmapping/"
}
//Take all gopaths and create possible muxmapping dir paths
func possiblepaths() []string{
	gopath := strings.Split(os.Getenv("GOPATH"), ":")
	var possible_paths []string
	for _, path := range gopath{
		 possible_paths = append(possible_paths, add_to_path(path))
	}
	return possible_paths
}

//returns a list of .txt mapping files found in muxmapping dir across gopaths
func findcontent(paths []string) map[string]string{
	var muxlocations map[string]string
	muxlocations = make(map[string]string)
	for _,path := range paths {
		files, _ := ioutil.ReadDir(path)
		if len(files) == 0{
			continue
		}
		for i:=0; i<len(files); i++{
			muxlocations[files[i].Name()] = path
		}
	}
	return muxlocations
}

func Loadmappings() map[string]string{
	//read all the files in this directory and expose them
	//fmt.Println(muxpath)
	muxlocations := findcontent(possiblepaths())

	for k, v := range muxlocations{
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("--")
	}
	return muxlocations
}
