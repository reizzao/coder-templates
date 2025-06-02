package create

import (
	adm "github.com/reizzao/dotclones/coder-templates/GO_Backend_Example1/entities/example1/adm"
	types "github.com/reizzao/dotclones/coder-templates/GO_Backend_Example1/entities/example1/types"
)

func Create(model types.Example1Model) types.Example1Model {
	res := prepareCreate(model.Request, adm.UseAdmExample1)
	return res
}
