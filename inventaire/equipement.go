package inventaire

import "fmt"

type Slot int

const (
	SlotArme Slot = iota
	SlotTete
	SlotTorse
	SlotJambes
	SlotPieds
	SlotMains
)

type Equipement struct {
	slots map[Slot]*Objet
}

func NewEquipement() *Equipement {
	return &Equipement{
		slots: make(map[Slot]*Objet),
	}
}

func (e *Equipement) Equiper(slot Slot, objet *Objet) error {
	if e.slots[slot] != nil {
		return fmt.Errorf("Vous avez déja un objet équiper a cet emplacement")
	}

	e.slots[slot] = objet
	return nil
}

func (e *Equipement) Desequiper(slot Slot) error {
	if e.slots[slot] == nil  {
		return fmt.Errorf("Vous n'avez rien d'équiper a cet emplacement")
	}

	e.slots[slot] = nil
	return nil
}

func (e *Equipement) GetStats() StatsObjet {
    total := StatsObjet{}
    for _, objet := range e.slots {
        if objet != nil {
			total.Force       +=  objet.Stats.Force
			total.Intel       +=  objet.Stats.Intel
			total.Agilite     +=  objet.Stats.Agilite
			total.Defense     +=  objet.Stats.Defense
			total.Vitesse     +=  objet.Stats.Vitesse
			total.VieMax      +=  objet.Stats.VieMax
			total.ManaMax     +=  objet.Stats.ManaMax
			total.RageMax     +=  objet.Stats.RageMax
			total.EnergieMax  +=  objet.Stats.EnergieMax
        }
    }
    return total
}