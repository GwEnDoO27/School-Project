package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Veuillez fournir un numéro de port")
		return
	}
	PORT := ":" + arguments[1]

	conn, err := net.Dial("tcp", PORT)
	if err != nil {
		fmt.Println("Erreur lors de la connexion au serveur:", err)
		return
	}
	defer conn.Close()

	// Demander et stocker le pseudo de l'utilisateur une seule fois
	fmt.Print("Veuillez entrer votre pseudo: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	pseudo := scanner.Text()

	// Utilisation d'une goroutine pour la saisie utilisateur
	wg.Add(1)
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			message := fmt.Sprintf("[%s][%s]: %s\n", time.Now().Format(time.Stamp), pseudo, text)
			_, err := conn.Write([]byte(message))
			if err != nil {
				fmt.Println("Erreur lors de l'envoi du message:", err)
				break
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Erreur lors de la saisie utilisateur:", err)
		}
	}()

	// Goroutine pour la réception des messages du serveur
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Erreur lors de la réception du message du serveur:", err)
				break
			}
			fmt.Print(message)
		}
	}()

	// Attendez la fin des goroutines avant de quitter
	wg.Wait()
}
