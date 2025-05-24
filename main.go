package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ThiagoLincoln/curso-go/checksite"
)

func main() {
  var sites []string
  reader := bufio.NewReader(os.Stdin)

  fmt.Println("Digite os sites que deseja verificar (digite 0 para finalizar): ")

  for {
    fmt.Print("Site: ")
    entrada, _ := reader.ReadString('\n')
		site := strings.TrimSpace(entrada)

		if site == "0" {
			break
		}

		sites = append(sites, site)
  }

  fmt.Println("\n🔍 Verificando sites...\n")
	for _, site := range sites {
		checksite.CheckSite(site)
	}
}
