package goserv

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/deckarep/golang-set"
)

func TestLoadmappings(t *testing.T){
	op := Loadmappings()
	fmt.Println(op)
	// Output: MOOO!
}

func TestPossible_paths(t *testing.T){
	fmt.Println("Getting possible paths")
	assert := assert.New(t)
	emptySet := mapset.NewSet()
	obtainedPaths := mapset.NewSet()
	possible_mux_paths := possible_paths()
	for _, v := range possible_mux_paths{
		obtainedPaths.Add(v)
	}
	assert.False(emptySet.Equal(obtainedPaths))
}

func TestReadmapping(t *testing.T){
	fmt.Println("reading data from test files")
	mappings := Readmapping("/Users/dsanghani/personalGit/goserv/src/muxmapping/app1ver2.mux")
	assert.Equal(t, len(mappings), 5)
}

func TestReadallmapping(t *testing.T){
	fmt.Println("reading data from test files")
	Readallmapping()
}
