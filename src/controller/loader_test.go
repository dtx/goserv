package controller

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

func TestPossiblePaths(t *testing.T){
	fmt.Println("Getting possible paths")
	assert := assert.New(t)
	requiredPaths := mapset.NewSet()
	obtainedPaths := mapset.NewSet()
	requiredPaths.Add("/usr/local/opt/go/libexec/src/muxmapping/")
	requiredPaths.Add("/Users/dsanghani/personalGit/goserv/src/muxmapping/")
	requiredPaths.Add("/Users/dsanghani/go/src/muxmapping/")
	possible_mux_paths := possiblepaths()
	for _, v := range possible_mux_paths{
		obtainedPaths.Add(v)
	}
	assert.True(requiredPaths.Equal(obtainedPaths))
}