package create

import (
	"fmt"

	data "github.com/reizzao/dotclones/coder-templates/GO_Backend_Example1/entities/example1/data"
)

func TesterCreate() {
	fmt.Println(
		Create(data.SeedCreateExample1),
	)
}
