package classes

import (
	"fmt"
	"world-of-milousques/inventaire"
)

type SortMage struct {
	Nom               string
	Cout              int
	Degats            int
	Description       string
	NiveauNecessaire  int
	Effet             Effet
}

type Mage struct {
	Combattant
	Intel    int
	Mana     int
	ManaMax  int
	Sorts    []SortMage
}

func (m *Mage) Attaquer() int {
	return m.Intel
}

func (m *Mage) Defendre(degats int) {
	degatsReels := degats - m.Armure
	if degatsReels < 0 {
		degatsReels = 0
	}
	m.Vie -= degatsReels
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

func sortsMage() []SortMage {
	return []SortMage{
		{
			Nom:               "Boule de feu",
			Cout:              20,
			Degats:            35,
			Description:       "Un sort de feu basique",
			NiveauNecessaire:  1,
			Effet:             DegatsDirect,
		},
		{
			Nom:               "Metamorphose",
			Cout:              20,
			Degats:            0,
			Description:       "Transforme l'ennemi en mouton pour 1 tour",
			NiveauNecessaire:  1,
			Effet:             Stun,
		},
		{
			Nom:               "Soin",
			Cout:              30,
			Degats:            50,
			Description:       "Te soigne de 50pv",
			NiveauNecessaire:  3,
			Effet:             Soin,
		},
	}
}

func (m *Mage) LancerSort(sort SortMage) (int, error) {
	if m.Mana < sort.Cout {
		return 0, fmt.Errorf("mana insuffisant (%d/%d)", m.Mana, sort.Cout)
	}
	m.Mana -= sort.Cout
	return sort.Degats, nil
}

func NewMage(nom string) *Mage {
	return &Mage{
		Combattant: Combattant{
			Nom:         nom,
			VieMax:      100,
			Vie:         100,
			Armure:      10,
			Inventaire:  inventaire.NewInventaire(20),
			Equipement:  inventaire.NewEquipement(),
		},
		Intel:    15,
		Mana:     50,
		ManaMax:  50,
		Sorts:    sortsMage(),
	}
}
