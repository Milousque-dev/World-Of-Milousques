package classes

type EvenementCombat struct {
	Attaquant      string
	Cible          string
	Degats         int
	MessageSpecial string
}

type ResultatCombat struct {
	Vainqueur  string
	NbTours    int
	Evenements []EvenementCombat
}

type Combat struct {
	Joueur     Personnage
	Adversaire Personnage
}

func NewCombat(joueur Personnage, adversaire Personnage) *Combat {
	return &Combat{
		Joueur:     joueur,
		Adversaire: adversaire,
	}
}

type TypeAction int

const (
	Attaque TypeAction = iota
	UtiliserCapacite
	Consommer
	Fuir
)

func (c *Combat) ExecuterTourJoueur(action TypeAction) []EvenementCombat {
	var evenements []EvenementCombat

	switch action {
	case Attaque:
		degats := c.Joueur.Attaquer()
		c.Adversaire.Defendre(degats)
		evenements = append(evenements, EvenementCombat{
			Attaquant: c.Joueur.GetNom(),
			Cible:     c.Adversaire.GetNom(),
			Degats:    degats,
		})

	case Fuir:
		evenements = append(evenements, EvenementCombat{
			MessageSpecial: "Vous fuyez",
		})

	case UtiliserCapacite:
		evenements = append(evenements, EvenementCombat{
			MessageSpecial: "non implemente",
		})

	case Consommer:
		evenements = append(evenements, EvenementCombat{
			MessageSpecial: "non implemente",
		})

	default:
		degats := c.Joueur.Attaquer()
		c.Adversaire.Defendre(degats)
		evenements = append(evenements, EvenementCombat{
			Attaquant: c.Joueur.GetNom(),
			Cible:     c.Adversaire.GetNom(),
			Degats:    degats,
		})
	}

	return evenements
}

func (c *Combat) ExecuterTourAdversaire() EvenementCombat {
	degats := c.Adversaire.Attaquer()
	c.Joueur.Defendre(degats)
	return EvenementCombat{
		Attaquant: c.Adversaire.GetNom(),
		Cible:     c.Joueur.GetNom(),
		Degats:    degats,
	}
}

func (c *Combat) EstTermine() bool {
	return c.Joueur.EstMort() || c.Adversaire.EstMort()
}

func (c *Combat) GetVainqueur() string {
	if c.Joueur.EstMort() {
		return c.Adversaire.GetNom()
	}
	return c.Joueur.GetNom()
}
