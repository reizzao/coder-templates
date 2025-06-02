package repositories

import types "github.com/reizzao/dotclones/coder-templates/GO_Backend_Example1/entities/example1/types"

var JsonRepositoryExample1 = types.RepositoryExample1{

	Collection: []types.Example1Model{},

	Save: func(model types.Example1Model) types.Example1Model {
		// todo: implementar - salvar na base de dados - Json

		return model
	},
}
