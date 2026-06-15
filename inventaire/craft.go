package inventaire

import "fmt"

type Recette struct {
	Nom          string
	Ingredients  []Objet
	Resultat     Objet
}

func (inv *Inventaire) Contient(nom string) bool {
	for _, o := range inv.objets {
		if o.Nom == nom {
			return true
		}
	}
	return false
}

func Crafter(inv *Inventaire, recette Recette) error {
	for _, ingredient := range recette.Ingredients {
		if !inv.Contient(ingredient.Nom) {
			return fmt.Errorf("ingredient manquant : %s", ingredient.Nom)
		}
	}
	for _, ingredient := range recette.Ingredients {
		inv.Retirer(ingredient.Nom)
	}
	return inv.Ajouter(recette.Resultat)
}
