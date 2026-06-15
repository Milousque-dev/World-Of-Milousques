package display

import (
	"bufio"
	"fmt"
	"os"
	"world-of-milousques/classes"
)

func AfficherCombatInteractif(combat *classes.Combat) {
	scanner := bufio.NewScanner(os.Stdin)

	for !combat.EstTermine() {
		fmt.Printf("\n%s : %d PV\n", combat.Joueur.GetNom(), combat.Joueur.GetVie())
		fmt.Printf("%s : %d PV\n", combat.Adversaire.GetNom(), combat.Adversaire.GetVie())
		fmt.Println("\n1. Attaquer")
		fmt.Println("2. Utiliser une capacite")
		fmt.Println("3. Consommer un objet")
		fmt.Println("4. Fuir")
		fmt.Print("Votre choix : ")

		scanner.Scan()
		choix := scanner.Text()

		var action classes.TypeAction
		switch choix {
		case "1":
			action = classes.Attaque
		case "2":
			action = classes.UtiliserCapacite
		case "3":
			action = classes.Consommer
		case "4":
			action = classes.Fuir
		default:
			fmt.Println("Choix invalide, vous attaquez.")
			action = classes.Attaque
		}

		evenements := combat.ExecuterTourJoueur(action)
		for _, e := range evenements {
			if e.MessageSpecial != "" {
				fmt.Println(e.MessageSpecial)
			} else {
				fmt.Printf("%s attaque %s pour %d degats.\n", e.Attaquant, e.Cible, e.Degats)
			}
		}

		if action == classes.Fuir {
			break
		}

		if !combat.EstTermine() {
			e := combat.ExecuterTourAdversaire()
			fmt.Printf("%s attaque %s pour %d degats.\n", e.Attaquant, e.Cible, e.Degats)
		}
	}

	if !combat.EstTermine() {
		fmt.Println("\nVous avez fui le combat.")
	} else {
		fmt.Printf("\nCombat termine ! Vainqueur : %s\n", combat.GetVainqueur())
	}
}
