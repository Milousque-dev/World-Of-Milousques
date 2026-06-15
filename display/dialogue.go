package display

type Choix struct {
	Texte  string
	Suite  string
}

type Dialogue struct {
	ID     string
	Texte  string
	Choix  []Choix
}

var IntroDialogues = map[string]Dialogue{
	"debut": {
		ID: "debut",
		Texte: "Oh ! Un nouvel aventurier.\n Es tu prêt a commencer ton apprentissage ?",
		Choix: []Choix{
			{Texte: "Qui etes-vous?",                  Suite: "identite"},
			{Texte: "Ou suis-je?",                     Suite: "lieu"},
			{Texte: "Je n'ai pas de temps a perdre.",  Suite: "depart"},
		},
	},

	"identite": {
		ID: "identite",
		Texte: "Je suis Ganymède, le gardien des chasseurs de Milousques.\n Je suis là pour aider les aventuriers comme toi a bien débuter leur quête des Milousques",
		Choix: []Choix{
			{Texte: "Que sont les Milousques",                          Suite: "Milousques"},
			{Texte: "Ou suis-je?",                                      Suite: "lieu"},
			{Texte: "Je part de suite à la conquêtes des milousques !", Suite: "depart"},
		},
	},

	"lieu": {
		ID: "lieu",
		Texte: "Tu es a Incarmite, le point de départ de tout les chasseurs de milousques !"
		Choix: []Choix{
			{Texte: "Qui etes-vous?", Suite}
		}
	}
}