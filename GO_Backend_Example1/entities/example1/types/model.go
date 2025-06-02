package types

type Example1Model struct {
	ID      string
	Request RequestCreateExample1
}

type RequestCreateExample1 struct {
	Nome      string
	SobreNome string
}

