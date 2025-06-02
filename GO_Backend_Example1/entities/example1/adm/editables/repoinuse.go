package editables

import (
	"github.com/reizzao/backendgo/entities/example1/repositories"
	types "github.com/reizzao/backendgo/entities/example1/types"
)

type Example1OptionsRepoinuse struct {
	Json   types.RepositoryExample1
	Memory types.RepositoryExample1
}

var optionsRepoinuseExample1 = Example1OptionsRepoinuse{
	Json:   repositories.JsonRepositoryExample1,
	Memory: repositories.MemoryRepositoryExample1,
}

// editable repoinuse # important
var RepoinuseExample1 types.RepositoryExample1 = optionsRepoinuseExample1.Memory
