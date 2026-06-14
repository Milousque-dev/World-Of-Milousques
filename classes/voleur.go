package classes

import "fmt"

type TechniqueVoleur struct {
	Nom         string
	Coût        int
	Dégâts      int
	Description string
}

type Voleur struct {
	Combattant
	Agilité    int
	Energie    int
	EnergieMax int
	Techniques []TechniqueVoleur
}

func (v *Voleur) Attaquer() int {

	return v.Agilité
}

func (v *Voleur) Défendre(dégats int) {

	dégatsRéels := dégats - v.Armure

	if dégatsRéels < 0 {
		dégatsRéels = 0
	}

	v.Vie -= dégatsRéels
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

func (v *Voleur) UtiliserTechnique(tech TechniqueVoleur) (int, error) {

	if v.Energie < tech.Coût {
		return 0, fmt.Errorf("énergie insuffisante (%d/%d)", v.Energie, tech.Coût)
	}

	v.Energie -= tech.Coût
	return tech.Dégâts, nil
}

func NewVoleur(nom string) *Voleur {
	return &Voleur{
		Combattant: Combattant{
			Nom:        nom,
			VieMax:     150,
			Vie:        150,
			Armure:     10,
			Inventaire: NewInventaire(20),
			Equipement: make(map[string]Objet),
		},
		Agilité:    15,
		Energie:    0,
		EnergieMax: 100,
		Techniques: []TechniqueVoleur{},
	}
}
