package classes

type Personnage interface {
	Attaquer() int
	Defendre(degats int)
	EstMort() bool
	GetNom() string
	GetVie() int
}
