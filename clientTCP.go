package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
		"os/exec"
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
	fmt.Print("====> Move <===="+"\n")
	fmt.Print("You can move div with up left right down key's"+"\n")

	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	
    var b []byte = make([]byte, 1)
	
	for i := true; i;{
        os.Stdin.Read(b)

		if string(b) == "e"{
			exec.Command("reset").Run()
			Menu(c)
			i = false
		}

		if string(b) == "e"{
			exec.Command("reset").Run()
			Menu(c)
			i = false
		}
        fmt.Println("I got the byte", b, "("+string(b)+")")
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
		if text[0] == '1' {
			Move(c)
			i = false
		} else if text[0] == '2' {
			Config(c)
			i = false
		} else if text[0] == '3' {
			i = false
		} else {
			fmt.Print("Please choose a valide option.." + "\n")
		}
	}
}
