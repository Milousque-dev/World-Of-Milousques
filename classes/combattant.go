package classes

import "world-of-milousques/inventaire"

type Combattant struct {
	Nom         string
	VieMax      int
	Vie         int
	Armure      int
	Mort        bool
	Inventaire  *inventaire.Inventaire
	Equipement  *inventaire.Equipement
}
