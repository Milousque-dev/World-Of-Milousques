package classes

import (
	"fmt"
	"world-of-milousques/inventaire"
)

type TechniqueVoleur struct {
	Nom               string
	Cout              int
	Degats            int
	Description       string
	NiveauNecessaire  int
	Effet              Effet
}

type Voleur struct {
	Combattant
	Agilite     int
	Energie     int
	EnergieMax  int
	Techniques  []TechniqueVoleur
}

func (v *Voleur) Attaquer() int {
	return v.Agilite
}

func (v *Voleur) Defendre(degats int) {
	degatsReels := degats - v.Armure
	if degatsReels < 0 {
		degatsReels = 0
	}
	v.Vie -= degatsReels
}

func (v *Voleur) EstMort() bool {
	return v.Vie <= 0
}

func (v *Voleur) GetNom() string {
	return v.Nom
}

func (v *Voleur) GetVie() int {
	return v.Vie
}

func techniquesVoleur() []TechniqueVoleur {
	return []TechniqueVoleur{
		{
			Nom:               "Attaque sournoise",
			Cout:              20,
			Degats:            35,
			Description:       "Une attaque dans le dos de l'ennemi",
			NiveauNecessaire:  1,
			Effet:             DegatsDirect,
		},
		{
			Nom:               "Empoisonnement",
			Cout:              15,
			Degats:            10,
			Description:       "Empoisonne l'ennemi pour 3 tours",
			NiveauNecessaire:  1,
			Effet:             Poison,
		},
		{
			Nom:               "Technique pernicieuse",
			Cout:              40,
			Degats:            25,
			Description:       "Etourdit l'ennemi pour 1 tour",
			NiveauNecessaire:  3,
			Effet:             Stun,
		},
	}
}

func (v *Voleur) UtiliserTechnique(tech TechniqueVoleur) (int, error) {
	if v.Energie < tech.Cout {
		return 0, fmt.Errorf("energie insuffisante (%d/%d)", v.Energie, tech.Cout)
	}
	v.Energie -= tech.Cout
	return tech.Degats, nil
}

func NewVoleur(nom string) *Voleur {
	return &Voleur{
		Combattant: Combattant{
			Nom:         nom,
			VieMax:      150,
			Vie:         150,
			Armure:      10,
			Inventaire:  inventaire.NewInventaire(20),
			Equipement:  inventaire.NewEquipement(),
		},
		Agilite:     15,
		Energie:     0,
		EnergieMax:  100,
		Techniques:  techniquesVoleur(),
	}
}
