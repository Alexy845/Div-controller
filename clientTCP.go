package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
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

	Menu(c)

}

func Config(c net.Conn){
	fmt.Print("====> Configuration <===="+"\n")
	for i := true; i;{
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		if text[0] != '/'{
			fmt.Print(c, "-> Invalide input, please use : /..commande.."+"\n")
		}else {
			if text[:5] == "/stop"{
				Menu(c)
				i = false
			}
			fmt.Fprintf(c, text+"\n")
		}
	}
}

func Move(c net.Conn){
	left := 0
	top := 0
	command := ""
	fmt.Print("====> Move <===="+"\n")
	fmt.Print("You can move div with : z[up] | q[left] | [d]right | s[down] or use e[exit]"+"\n")
	
	for i := true; i;{
		
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')

		if string(text[0]) == "e"{
			Menu(c)
			i = false
		}
		if string(text[0]) == "s"{
			top += 10
			command = "/marginTop " + IntegerToString(top)
			fmt.Fprintf(c, command+"\n")
		}else if string(text[0]) == "z"{
			top -= 10
			command = "/marginTop " + IntegerToString(top)
			fmt.Fprintf(c, command+"\n")
		}else if string(text[0]) == "d"{
			left += 20
			command = "/marginLeft " + IntegerToString(top)
			fmt.Fprintf(c, command+"\n")
		}else if string(text[0]) == "q"{
			left -= 20
			command = "/marginLeft " + IntegerToString(top)
			fmt.Fprintf(c, command+"\n")
		}
    }
	
}

func Menu(c net.Conn){
	fmt.Print("======> Menu <======"+"\n")
	fmt.Print("[1] : Moving div"+"\n")
	fmt.Print("[2] : Config div"+"\n")
	fmt.Print("[3] : Exit"+"\n")
	fmt.Print("-- Choose an option --"+"\n")

	for i := true; i;{

		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')

		if string(text[0]) == "1" {
			Move(c)
			i = false
		} else if string(text[0]) == "2" {
			Config(c)
			i = false
		} else if string(text[0]) == "3" {
			i = false
		} else {
			fmt.Print("Please choose a valide option.." + "\n")
		}
	}
}

func IntegerToString(nbr int) string {
	snbr := ""
	var allunit []int
	if nbr == 0 {
		return "0"
	} else {
		for nbr > 0 {
			allunit = append(allunit, nbr%10)
			nbr = nbr / 10
		}
	}
	for c := len(allunit) - 1; c >= 0; c-- {
		snbr += string(rune(allunit[c] + 48))
	}
	return snbr
}
