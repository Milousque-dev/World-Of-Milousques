package classes

type Combattant struct{
	Nom        string
	VieMax     int
	Vie        int
	Armure     int
	Mort       bool
	Inventaire *Inventaire
	Equipement map[string]Objet
}

