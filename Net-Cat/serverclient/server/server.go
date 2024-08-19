package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

//TODO faire les logs
var (
	clients     = make(map[net.Conn]string)
	clientsLock sync.Mutex
	PORT        = "8989"
)

func main() {
	arguments := os.Args

	switch len(arguments) {
	case 2:
		for _, c := range arguments[1] {
			if c < '0' || c > '9' {
				fmt.Println("USAGE]: ./TCPChat $port")
				return
			}
		}
		PORT = arguments[1]
	case 1:
	default:
		fmt.Println("USAGE]: ./TCPChat $port")
		return
	}

	ln, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		fmt.Println("USAGE]: ./TCPChat $port")
		return
	}
	defer ln.Close()
	fmt.Println("Serveur en attente de connexions... sur le port:", PORT)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Erreur lors de la connexion d'un client:", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	/* 	var log []string
	 */
	// Demander et stocker le pseudo de l'utilisateur une seule fois
	welcome, err := os.ReadFile("pingouin.txt")
	if err != nil {
		fmt.Println("pas de pingouin", err)
	}
	conn.Write([]byte("[Serveur] : Welcome to TCP-Chat\n"))
	conn.Write([]byte(welcome))
	conn.Write([]byte("[Serveur] :Enter Username: "))
	username, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Erreur lors de la lecture du nom d'utilisateur:", err)
		return
	}
	username = username[:len(username)-1] // Retirer le caractère de nouvelle ligne

	// Stocker le client avec son pseudo
	clientsLock.Lock()
	clients[conn] = username
	User := clients[conn]
	clientsLock.Unlock()

	fmt.Printf("Nouveau client connecté: %s\n", username)
	messEntr := "has joined our chat..."
	broadcastMessageEntryorLeft(User, messEntr)

	// Boucle pour la réception et la diffusion des messages
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()

		// Récupérer le pseudo du client
		clientsLock.Lock()
		sender := clients[conn]
		clientsLock.Unlock()

		// Diffuser le message à tous les autres clients
		broadcastMessage(sender, message)
	}

	// Enlever le client lors de la déconnexion
	clientsLock.Lock()
	delete(clients, conn)
	clientsLock.Unlock()

	fmt.Printf("Client déconnecté: %s\n", username)
	messLeft := "has left the chat..."
	broadcastMessageEntryorLeft(User, messLeft)
}

func broadcastMessage(sender, message string) {
	clientsLock.Lock()
	defer clientsLock.Unlock()

	for clientConn, username := range clients {
		if username != sender {
			clientConn.Write([]byte(fmt.Sprintf("[%s] [%s]: %s\n", currentTime(), sender, message)))
		}
	}
}

func currentTime() string {
	return time.Now().Format(time.Stamp)
}

func broadcastMessageEntryorLeft(sender, message string) {
	clientsLock.Lock()
	defer clientsLock.Unlock()

	for clientConn, username := range clients {
		if username != sender {
			clientConn.Write([]byte(fmt.Sprintf("%s %s\n", sender, message)))
		}
	}
}
