package classes

type Personnage interface {

	Attaquer() int
	Défendre(dégats int) 
	EstMort() bool
	GetNom() string
	GetVie() int
}