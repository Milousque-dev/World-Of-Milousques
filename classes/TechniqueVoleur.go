package classes

type TechniqueVoleur struct {
	Nom          string
	Cout         int
	Dégats       int
	Description  string
	Energie      int
	EnergieMax   int
	Techniques   []TechniqueVoleur
}