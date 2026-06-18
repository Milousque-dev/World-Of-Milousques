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
			{Texte: "Que sont les Milousques",                          Suite: "milousques"},
			{Texte: "Ou suis-je?",                                      Suite: "lieu"},
			{Texte: "Je part de suite à la conquêtes des milousques !", Suite: "depart"},
		},
	},

	"lieu": {
		ID: "lieu",
		Texte: "Tu es a Incarmite, le point de départ de tout les chasseurs de milousques !",
		Choix: []Choix{
			{Texte: "Qui etes-vous?",                                   Suite: "identite"},
			{Texte: "Que sont les Milousques",                          Suite: "milousques"},
			{Texte: "Je part de suite à la conquêtes des milousques !", Suite: "depart"},
		},
	},

	"milousques": {
		ID: "milousques",
		Texte: "Les Milousques sont de puissantes créatures chimères qui de grands pouvoir a ceux qui arrivent a les dompters",
		Choix: []Choix{
			{Texte: "Qui etes-vous?",                                   Suite: "identite"},
			{Texte: "Ou suis-je?",                                      Suite: "lieu"},
			{Texte: "Je part de suite à la conquêtes des milousques !", Suite: "depart"},
		},
	},

	"depart": {
		ID:  "depart",
		Texte: "Pas si vite ! Tu dois d'abord passer 2 épreuves avant de pouvoir commencer ta quêtes",
		Choix: []Choix{
			{Texte: "Quels sont ses épreuves ?", Suite: "epreuves"},
			{Texte: "Je suis prêt ! Allons-y",   Suite: "Tuto"},
		},
	},

	"epreuves": {
		ID: "epreuves",
		Texte: "Tu va devoir créer ta première arme et vaincre le gardien de la porte d'Astrab afin de prouver ta valeurs",
		Choix: []Choix{
			{Texte: "Je suis prêt ! Allons-y",   Suite: "Tuto"},
		},
	},
}