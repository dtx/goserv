package goserv

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
			fmt.Printf("Invalid path provided %s, ignoring and continuing\n", path)
			continue
		}
		if len(files) == 0{
			fmt.Printf("No file found in path %s, ignoring and continuing\n", path)
			continue
		}
		for i:=0; i<len(files); i++{
			if check_extension(files[i].Name(), "mux"){
				fmt.Printf("Multiplexer mapping (.mux) %s found in path %s\n", files[i].Name(), path)
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
	return muxlocations
}

//return a map of keys as routes and values as method names, from a muxmapping file path
func Readmapping(muxmap_path string) map[string]string{
	mapping_data, err := ioutil.ReadFile(muxmap_path)
	mapping := make(map[string]string)
	if err != nil{
		fmt.Printf("There was an error in reading from file: %s\n", muxmap_path)
	}
	mapping_data_lines := strings.Split(string(mapping_data), "\n")
	for _,line := range mapping_data_lines{
		route_method := strings.Split(line, " ")
		//if route method was succesfully split into 2 parts
		if len(route_method) == 2 {
			mapping[route_method[0]] = route_method[1]
			continue
		}
		//if not a valid split, then dont break
	}
	return mapping
}

func Readallmapping() []map[string]string{
	fmt.Println("Reading all mappings")
	op := Loadmappings()
	var all_muxmappings []map[string]string
	for file,dir := range op{
		//join k+v to form a full path
		a_muxmapping := Readmapping(dir+file)
		all_muxmappings = append(all_muxmappings, a_muxmapping)
	}
	return all_muxmappings
}
