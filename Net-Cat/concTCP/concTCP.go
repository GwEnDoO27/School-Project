package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
	/* "github.com/jroimartin/gocui" */)

//var count = 0

func handleConnections(c net.Conn) {
	var wg sync.WaitGroup

	fmt.Printf("[username] : ")
	fmt.Println("Welcome to TCP-Chat!")
	fmt.Println("ENTER YOUR NAME: ")
	// permet d'avoir l'input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	User := scanner.Text()
	/* fmt.Println(text) */

	/* 	// permet d'avoir le temps
	   	t := time.Now()
	   	myTime := t.Format(time.Stamp)
	   	// permet de join tout
	   	s := []string{text, myTime, ": "}
	   	//message envoyer dans le terminal
	   	c.Write([]byte(strings.Join(s, " ")))

	   	//Affchage dans le server
	   	for {
	   		netData, err := bufio.NewReader(c).ReadString('\n')
	   		log.Printf("[%s][%s] : %s ", text, t, netData)
	   		if err != nil {
	   			fmt.Println(err)
	   			return
	   		}
	   		if strings.TrimSpace(string(netData)) == "STOP" {
	   			fmt.Printf("[%s] Exiting TCP server !", text)
	   			return
	   		}
	   	} */

	wg.Add(2)

	go func() { //Utilisateur
		defer wg.Done()

		for {
			reader := bufio.NewReader((os.Stdin))
			text, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			c.Write([]byte(text))
		}
	}()

	go func() { // reception msg serveur
		defer wg.Done()
		for {
			t := time.Now()
			UsTime := t.Format(time.Stamp)
			message, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				panic(err)

			}
			/* s := []string{User, UsTime, ":", message} */
			fmt.Printf("[%s][%s] : %s", User, UsTime, message)
		}
	}()
	/* c.Close() */
	wg.Wait()
}

func read()

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}
	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
		//TODO: Erreur
	}
	/* defer l.Close() */
	var clients []net.Conn

	for {
		c, err := l.Accept()
		if err == nil {
			clients = append(clients, c)
		}
		if err != nil {
			panic(err)
		}
		fmt.Println("Un client est connecté", c.RemoteAddr())

		go func() {
			buf := bufio.NewReader(c)
			for {
				name, err := buf.ReadString('\n')
				if err != nil {
					fmt.Printf("Client Deconnecte \n")
					break
				}
				for _, c := range clients {
					c.Write([]byte(name))
				}
			}
		}()
		go handleConnections(c)
	}

}

//TODO:boucle range pour chopper le nom
//TODO:Afficher tout sur les memes terminaux

/* fmt.Println(
		    " _nnnn_",
		    "dGGGGMMb",
	        "@p~qp~~qMb",
	        "M|@||@) M|",
	        "@,----.JM|",
	       "JS^\__/  qKL",
	      "dZP        qKRb",
	     "dZP          qKKb",
	    "fZP            SMMb",
	    "HZM            MMMM",
	    "FqM            MMMM",
	 "__| ".        |\dS"qML",
	"  |    `.       | `' \Zq ",
	"_)      \.___.,|     .' ",
	"\____   )MMMMMP|   .'",
	      "`-'       `--'",
) */

/* 	for {
	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	go handleConnections(c)
} */
