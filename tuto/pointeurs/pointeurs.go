package pointeurs

import "fmt"

func MainPointeurs() {
	var a int = 20

	//fmt.Printf("Adresse de la variable a: %p\n", &a)

	var ap *int // creation du pointeur
	ap = &a     // stockage de l'adresse mémoire de la varialble dans le pointeur
	fmt.Printf("Adresse de a: %p\n", &a)
	fmt.Printf("Valeur du pointeur ap: %p\n", ap)
	fmt.Printf("Valeur de l'adresse %p: %d\n", ap, *ap)

	var ptr *int

	if ptr == nil { // est ce que le pointeur est vide ?
		fmt.Println("Aucune adresse mémoire")
	} else {
		fmt.Printf("Votre adresse mémoire est %p\n", ptr)
	}

	var b int = 20
	var pb *int = &b

	if pb == nil {
		fmt.Println("Aucune adresse mémoire")
	} else {
		fmt.Printf("Votre adresse mémoire est %p\n", pb)
	}

	*pb = 30

	fmt.Printf("Nouvelle valeur de b: %d\n", b)
}
