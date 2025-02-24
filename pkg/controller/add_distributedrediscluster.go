package controller

import (
	"github.com/truongnh28/redis-cluster-operator/pkg/controller/distributedrediscluster"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, distributedrediscluster.Add)
}
