package create

import (
	"github.com/reizzao/backendgo/entities/example1/adm"
	types "github.com/reizzao/backendgo/entities/example1/types"
)

func prepareCreate(request types.RequestCreateExample1, admRepoinuse adm.AdmExample1) types.Example1Model {
	model := types.Example1Model{
		ID:      "todo make ID",
		Request: request,
	}

	saved := admRepoinuse.Repoinuse.Save(model)

	return saved
}
