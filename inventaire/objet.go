package inventaire

type TypeObjet int

const (
	Arme TypeObjet = iota
	Armure
	Consommable
	Ingredient
	ObjetQuete
	Misc
)

type StatsObjet struct {
	Force       int
	Intel       int
	Agilite     int
	Defense     int
	Vitesse     int
	VieMax      int
	ManaMax     int
	RageMax     int
	EnergieMax  int
}

type Objet struct {
	Nom          string
	Poids        int
	Valeur       int
	Description  string
	Type         TypeObjet
	Stats        StatsObjet
}

