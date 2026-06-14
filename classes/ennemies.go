package classes

type Ennemi struct {
	Nom    string
	Vie    int
	VieMax int
	Armure int
	Atq    int
}

func (e *Ennemi) Attaquer() int {
	return e.Atq
}

func (e *Ennemi) Defendre(degats int) {
	degatsReels := degats - e.Armure
	if degatsReels < 0 {
		degatsReels = 0
	}
	e.Vie -= degatsReels
}

func (e *Ennemi) EstMort() bool {
	return e.Vie <= 0
}

func (e *Ennemi) GetNom() string {
	return e.Nom
}

func (e *Ennemi) GetVie() int {
	return e.Vie
}

func NewEnnemi(nom string) *Ennemi {
	return &Ennemi{
		Nom:    nom,
		VieMax: 150,
		Vie:    150,
		Armure: 10,
		Atq:    10,
	}
}
