package repositories

import types "github.com/reizzao/backendgo/entities/example1/types"

var JsonRepositoryExample1 = types.RepositoryExample1{

	Collection: []types.Example1Model{},

	Save: func(model types.Example1Model) types.Example1Model {
		// todo: implementar - salvar na base de dados - Json

		return model
	},
}
