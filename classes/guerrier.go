package classes

import "fmt"

type CapacitéGuerrier struct {
	Nom         string
	Coût        int
	Dégâts      int
	Description string
}

type Guerrier struct {
	Combattant
	Force     int
	Rage      int
	RageMax   int
	Capacités []CapacitéGuerrier
}

func (g *Guerrier) Attaquer() int {

	return g.Force
}

func (g *Guerrier) Défendre(dégats int) {

	dégatsRéels := dégats - g.Armure

	if dégatsRéels < 0 {
		dégatsRéels = 0
	}

	g.Vie -= dégatsRéels
}

func (g *Guerrier) EstMort() bool {

	return g.Vie <= 0
}

func (g *Guerrier) GetNom() string {

	return g.Nom
}

func (g *Guerrier) GetVie() int {

	return g.Vie
}

func (g *Guerrier) UtiliserCapacité(cap CapacitéGuerrier) (int, error) {

	if g.Rage < cap.Coût {
		return 0, fmt.Errorf("rage insuffisante (%d/%d)", g.Rage, cap.Coût)
	}

	g.Rage -= cap.Coût
	return cap.Dégâts, nil
}

func NewGuerrier(nom string) *Guerrier {
	return &Guerrier{
		Combattant: Combattant{
			Nom:        nom,
			VieMax:     150,
			Vie:        150,
			Armure:     20,
			Inventaire: NewInventaire(20),
			Equipement: make(map[string]Objet),
		},
		Force:     15,
		Rage:      0,
		RageMax:   100,
		Capacités: []CapacitéGuerrier{},
	}
}
