package inventaire

import "fmt"

type Objet struct {
	Nom         string
	Poids       int
	Valeur      int
	Equipable   bool
	Consommable bool
	Stats       []int
}

type Inventaire struct {
	objets   []Objet
	Capacite int
}

func NewInventaire(capacite int) *Inventaire {
	return &Inventaire{
		objets:   make([]Objet, 0),
		Capacite: capacite,
	}
}

func (inv *Inventaire) Add(objet Objet) error {
	if len(inv.objets) >= inv.Capacite {
		return fmt.Errorf("inventaire plein (%d/%d)", len(inv.objets), inv.Capacite)
	}
	inv.objets = append(inv.objets, objet)
	return nil
}

func (inv *Inventaire) Remove(nom string) error {
	for i := 0; i < len(inv.objets); i++ {
		if inv.objets[i].Nom == nom {
			for j := i; j < len(inv.objets)-1; j++ {
				inv.objets[j] = inv.objets[j+1]
			}
			inv.objets = inv.objets[:len(inv.objets)-1]
			return nil
		}
	}
	return fmt.Errorf("objet '%s' introuvable", nom)
}
