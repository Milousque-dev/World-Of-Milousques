package classes

import (
	"fmt"
	"world-of-milousques/inventaire"
)

type CapaciteGuerrier struct {
	Nom              string
	Cout             int
	Degats           int
	Description      string
	NiveauNecessaire int
	Effet            Effet
}

type Guerrier struct {
	Combattant
	Force     int
	Rage      int
	RageMax   int
	Capacites []CapaciteGuerrier
}

func (g *Guerrier) Attaquer() int {
	return g.Force
}

func (g *Guerrier) Defendre(degats int) {
	degatsReels := degats - g.Armure
	if degatsReels < 0 {
		degatsReels = 0
	}
	g.Vie -= degatsReels
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

func capacitesGuerrier() []CapaciteGuerrier {
	return []CapaciteGuerrier{
		{
			Nom:              "Frappe puissante",
			Cout:             20,
			Degats:           35,
			Description:      "Un coup violent qui ecrase l'ennemi",
			NiveauNecessaire: 1,
			Effet:            DegatsDirect,
		},
		{
			Nom:              "Cri de guerre",
			Cout:             10,
			Degats:           0,
			Description:      "Esquive les degats pour 1 tour",
			NiveauNecessaire: 1,
			Effet:            Bouclier,
		},
		{
			Nom:              "Coup de bouclier",
			Cout:             30,
			Degats:           15,
			Description:      "Etourdit l'ennemi pour 1 tour",
			NiveauNecessaire: 3,
			Effet:            Stun,
		},
	}
}

func (g *Guerrier) UtiliserCapacite(cap CapaciteGuerrier) (int, error) {
	if g.Rage < cap.Cout {
		return 0, fmt.Errorf("rage insuffisante (%d/%d)", g.Rage, cap.Cout)
	}
	g.Rage -= cap.Cout
	return cap.Degats, nil
}

func NewGuerrier(nom string) *Guerrier {
	return &Guerrier{
		Combattant: Combattant{
			Nom:        nom,
			VieMax:     150,
			Vie:        150,
			Armure:     20,
			Inventaire: inventaire.NewInventaire(20),
			Equipement: make(map[string]inventaire.Objet),
		},
		Force:     15,
		Rage:      0,
		RageMax:   100,
		Capacites: capacitesGuerrier(),
	}
}
