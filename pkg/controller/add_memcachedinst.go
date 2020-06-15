package controller

import (
	"github.com/bjoydeep/memcached-operator-instrumented/pkg/controller/memcachedinst"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, memcachedinst.Add)
}
