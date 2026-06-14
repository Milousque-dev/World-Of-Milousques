package classes

import "fmt"

type SortMage struct {
	Nom         string
	Coût        int
	Dégâts      int
	Description string
}

type Mage struct {
	Combattant
	Intel   int
	Mana    int
	ManaMax int
	Sorts   []SortMage
}

func (m *Mage) Attaquer() int {

	return m.Intel
}

func (m *Mage) Défendre(dégats int) {

	dégatsRéels := dégats - m.Armure

	if dégatsRéels < 0 {
		dégatsRéels = 0
	}

	m.Vie -= dégatsRéels
}

func (m *Mage) EstMort() bool {

	return m.Vie <= 0
}

func (m *Mage) GetNom() string {

	return m.Nom
}

func (m *Mage) GetVie() int {

	return m.Vie
}

func (m *Mage) LancerSort(sort SortMage) (int, error) {

	if m.Mana < sort.Coût {
		return 0, fmt.Errorf("mana insuffisant (%d/%d)", m.Mana, sort.Coût)
	}

	m.Mana -= sort.Coût
	return sort.Dégâts, nil
}

func NewMage(nom string) *Mage {
	return &Mage{
		Combattant: Combattant{
			Nom:        nom,
			VieMax:     100,
			Vie:        100,
			Armure:     10,
			Inventaire: NewInventaire(20),
			Equipement: make(map[string]Objet),
		},
		Intel:   15,
		Mana:    50,
		ManaMax: 50,
		Sorts:   []SortMage{},
	}
}
