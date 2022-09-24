package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        //"strings"
)

func main() {

	arguments := os.Args
	if len(arguments) < 3 { // On vérifie si les ports et l'adresse a était renseigné
			fmt.Println("Please provide host:port.")
			return
	}
	
	CONNECT := arguments[1] + ":" + arguments[2] // Format AdressIP : Port


	c, err := net.Dial("tcp", CONNECT) // On essaie de se connecter
	if err != nil { // Si on a une erreur on l'affiche
			fmt.Println(err)
			return
	}

}

func Config(){
	for i := true; i;{
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		if text[0] != '/'{
			fmt.Print(c, "-> Invalide input, please use : /..commande.."+"\n")
		}else {
			if text[:5] == "/stop"{
				fmt.Print(c, "Stop"+"\n")
				i = false
			}
			fmt.Fprintf(c, text+"\n")
		}
	}
}
