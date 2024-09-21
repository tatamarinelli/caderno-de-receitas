package main

import (
	"fmt"
)

type (
	receita struct {
		nome         string
		ingredientes string
		tempo        string
	}
)

func main() {
	brigadeiro := receita{
		nome:         "brigadeiro",
		ingredientes: "leite condensado, cacau em pó, manteiga",
		tempo:        "15 minutos",
	}

	miojo := receita{
		nome:         "miojo",
		ingredientes: "miojo e água",
		tempo:        "7 minutos",
	}

	ovoFrito := receita{
		nome:         "ovo frito",
		ingredientes: "ovo e azeite",
		tempo:        "5 minutos",
	}

	fmt.Println(brigadeiro)
	fmt.Println(miojo)
	fmt.Println(ovoFrito)

}
