package model

//Pata representa uma ave do tipo Galinha
type Galinha interface {
	Cacareja() string
}

//Pata representa uma ave do tipo Pato
type Pata interface {
	Quack() string
}

type Ave struct {
	Nome string
}

//retorna o som da galinha
func (a Ave) Cacareja() string {
	return "Cócórico..."
}

// returna o som do pato
func (a Ave) Quack() string {
	return "Quack, quack..."
}
