package repositories

import types "github.com/reizzao/backendgo/entities/example1/types"

var MemoryRepositoryExample1 = types.RepositoryExample1{

	Collection: []types.Example1Model{},

	Save: func(model types.Example1Model) types.Example1Model {

		this := types.RepositoryExample1{} // recebendo o obj pra usar props
		_ = append(this.Collection, model) // nao usando o retorno ??? todo

		return model
	},
}
