package adm

import (
	editables "github.com/reizzao/dotclones/coder-templates/GO_Backend_Example1/entities/example1/adm/editables"
	types "github.com/reizzao/dotclones/coder-templates/GO_Backend_Example1/entities/example1/types"
)

type AdmExample1 struct {
	Repoinuse types.RepositoryExample1
}

// obj administraveis em uso

var UseAdmExample1 = AdmExample1{
	Repoinuse: editables.RepoinuseExample1,
}
