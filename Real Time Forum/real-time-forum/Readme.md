# Project README

## Project Overview

This project is an upgraded version of a previous forum application, incorporating several new features including private messaging and real-time actions. The main objective is to develop a sophisticated forum using Golang, SQLite, and JavaScript, with a single-page application (SPA) approach.

## Objectives

The project focuses on the following key areas:

1. **Registration and Login**
2. **Creation of Posts**
3. **Commenting on Posts**
4. **Private Messages**

## Project Structure

The new forum will consist of five main parts:

1. **SQLite**: Used for data storage.
2. **Golang**: Handles data processing and WebSockets for real-time interactions (Backend).
3. **JavaScript**: Manages frontend events and client-side WebSockets.
4. **HTML**: Organizes page elements.
5. **CSS**: Styles page elements.

## Single Page Application (SPA)

The application will have only one HTML file. Any page changes will be managed through JavaScript, ensuring a seamless single-page application experience.

## Features

### Registration and Login

- **Registration Form**: Users must fill in their nickname, age, gender, first name, last name, email, and password.
- **Login**: Users can log in using either their nickname or email combined with their password.
- **Logout**: Users can log out from any page on the forum.

### Posts and Comments

- **Create Posts**: Users can create posts with categories similar to the first forum.
- **Comment on Posts**: Users can comment on posts.
- **Feed Display**: Posts are displayed in a feed.
- **View Comments**: Comments are visible only when a post is clicked.

### Private Messages

- **Chat Section**: Displays who is online/offline and available for chat. Organized by the last message sent, or alphabetically for new users.
- **Send Messages**: Users can send private messages to online users.
- **Persistent Chat Section**: The chat section is always visible.
- **Message History**: Clicking on a user reloads past messages. Previous messages are displayed, with the last 10 messages loaded initially. More messages load when scrolling up.
- **Message Format**: Messages include the date sent and the sender's username.
- **Real-Time Messaging**: Messages are delivered in real-time using WebSockets, without page refresh.

## Allowed Packages

- All standard Go packages.
- Gorilla WebSocket
- SQLite3
- Bcrypt
- UUID

**Note**: Do not use frontend libraries or frameworks such as React, Angular, Vue, etc.

## Learning Outcomes

This project will help you learn about:

- Basic web technologies: HTML, HTTP, Sessions and Cookies, CSS
- Backend and Frontend integration
- DOM manipulation
- Go routines and Go channels
- WebSockets in Go and JavaScript
- SQL language and database manipulation

## Getting Started

### Prerequisites

Ensure you have the following installed:

- Go
- SQLite3
- Node.js (for running JavaScript)
- A web browser

### Installation

1. **Clone the repository**:
    ```sh
    git clone <repository_url>
    ```

2. **Navigate to the project directory**:
    ```sh
    cd <project_directory>
    ```

3. **Install Go dependencies**:
    ```sh
    go get ./...
    ```

4. **Run the application**:
    ```sh
    go run main.go
    ```

5. **Open the application in your browser**:
    ```sh
    http://localhost:8080
    ```

## Contributing

Please read the contributing guidelines in the `CONTRIBUTING.md` file for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

---

Feel free to reach out for any queries or further information.

Happy coding!