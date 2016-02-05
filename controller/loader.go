package controller

import(
	"io/ioutil"
	"os"
	"fmt"
	"strings"
	//"strconv"
)

//Go through path and concatenate required for a muxmapping directory
func add_to_path(path string) string{
	path_components := strings.Split(path, "/")
	if path_components[len(path_components)-1] == "src"{
		return path+"/muxmapping/"
	}
	return path+"/src/muxmapping/"
}
//Take all gopaths and create possible muxmapping dir paths
func possible_paths() []string{
	gopath := strings.Split(os.Getenv("GOPATH"), ":")
	var possible_paths []string
	for _, path := range gopath{
		 possible_paths = append(possible_paths, add_to_path(path))
	}
	return possible_paths
}

//returns a list of .txt mapping files found in muxmapping dir across gopaths
func find_content(paths []string) map[string]string{
	var muxlocations map[string]string
	muxlocations = make(map[string]string)
	for _,path := range paths {
		files, err := ioutil.ReadDir(path)
		if err!=nil{
			fmt.Println("Invalid path provided %s, ignoring and continuing", path)
			continue
		}
		if len(files) == 0{
			fmt.Println("No file found in path %s, ignoring and continuing", path)
			continue
		}
		for i:=0; i<len(files); i++{
			if check_extension(files[i].Name(), "mux"){
				muxlocations[files[i].Name()] = path
			}
		}
	}
	return muxlocations
}

func check_extension(filename string, extension string) bool{
	parts := strings.Split(filename, ".")
	this_extension := parts[len(parts)-1]
	return extension == this_extension
}

func Loadmappings() map[string]string{
	//read all the files in this directory and expose them
	//fmt.Println(muxpath)
	muxlocations:= find_content(possible_paths())

	for k, v := range muxlocations{
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("--")
	}
	return muxlocations
}
