package classes

type CapacitéGuerrier struct {
	Nom          string
	Cout         int
	Dégats       int
	Description  string
	Rage         int
	RageMax      int
	Capacités    []CapacitéGuerrier
}