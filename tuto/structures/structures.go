package structures

import "fmt"

type Personnage struct {
	nom        string
	vie        int
	puissance  int
	mort       bool
	inventaire []string
}

// Déclaration d'une méthode affichage() lié à la structure Personnage
func (p Personnage) Affichage() {
	fmt.Println("--------------------------------------------------")
	fmt.Println("Vie du personnage", p.nom, ":", p.vie)
	fmt.Println("Puissance du personnage", p.nom, ":", p.puissance)

	if p.mort {
		fmt.Println(p.nom, "est mort")
	} else {
		fmt.Println(p.nom, "est vivant")
	}

	fmt.Println("\nLe personnage", p.nom, "possède dans son inventaire :")

	for _, item := range p.inventaire {
		fmt.Println("-", item)
	}
}

func New(Nom string, Vie int, Puissance int, Mort bool, Inventaire [3]string) Personnage {
	personnage := Personnage{Nom, Vie, Puissance, Mort, Inventaire[:]}
	return personnage
}

func (p *Personnage) Soigner() {
	println(p.nom, " se soigne. Il récupère ", p.puissance/10, " points de vie.")
	p.vie = p.vie + p.puissance/10
	println("Il a maintenant ", p.vie, " point de vie.")
}

func (p *Personnage) Attaquer(p1 *Personnage) {
	fmt.Println(p.nom, " attaque ", p1.nom, "!!")
	fmt.Println("Il enlève ", p.puissance, " points de vie")
	p1.vie = p1.vie - p.puissance
	fmt.Println(p1.nom, " a maintenant ", p1.vie, " points de vie")
}

func (p *Personnage) SetMort() {
	p.mort = p.vie <= 0
}

func (p *Personnage) GetMort() bool {
	return p.mort
}

func (p *Personnage) GetVie() int {
	return p.vie
}
