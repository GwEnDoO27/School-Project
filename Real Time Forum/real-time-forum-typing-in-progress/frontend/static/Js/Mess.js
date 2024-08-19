import { NewCSS } from "./Newcss.js";
import { username } from "./Cookie.js";
import { LoadOldMessage, fetchAllUsers } from "./Fetching.js";
import { throttle } from "./Throttle.js";
import { Home } from "./Home.js";
import { Istyping, Typing } from "./istyping.js";


document.getElementById('nav').style.display = "block"


// mùain func for the message page
export async function Mess() {
    NewCSS(['/static/Css/Mess.css']);
    initializeMenu();
    const connectedUsersDiv = document.getElementById("connectedUsers");
    connectedUsersDiv.classList.add('connectedUsersDiv');

    const allUsers = await fetchAllUsers();
    const socket = new WebSocket("ws://localhost:8080/ws");

    socket.onopen = () => {
        console.log("WebSocket Status: Connected");
        socket.send(JSON.stringify({ type: "request_connected_users" }));
    };

    socket.onmessage = (e) => {
        const message = JSON.parse(e.data);
        if (message.type === "connected_users") {
            updateConnectedUsers(message.users, connectedUsersDiv);
        } else if (message.type === "typing") {
            Typing(message.from)
        } else {
            handleMessage(message);
        }
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };

    window.onbeforeunload = () => {
        socket.close();
    };

    document.getElementById('Homing').addEventListener('click', () => {
        Home()
    })

    displayUsers(allUsers, [], connectedUsersDiv, socket);
}

// Init the menu
function initializeMenu() {
    const regdiv = document.getElementsByClassName('container')[0];
    regdiv.classList.add('container');
    regdiv.innerHTML = `
        <h1>WebSocket Chat</h1>
        <div id="connectedUsers"></div>
    `;
}

// displaying the users
function displayUsers(allUsers, connectedUsers, connectedUsersDiv, socket) {
    connectedUsersDiv.innerHTML = '<strong>Users:</strong><br>';
    allUsers.forEach(({ user, notif }) => {
        const userElement = createUserElement(user, notif, connectedUsers);
        connectedUsersDiv.appendChild(userElement);
        userElement.addEventListener('click', () => {
            startChat(user, socket);
        });
    });
}

// UCreate the users div
function createUserElement(user, notif, connectedUsers) {
    const divbtn = document.createElement('div');
    divbtn.classList.add('divbtn');
    const userElement = document.createElement('button');
    userElement.classList.add('Link');
    userElement.textContent = user;
    userElement.id = user;
    userElement.style.color = connectedUsers.includes(user) ? "orange" : "white";

    if (notif > 0) {
        const Notifdiv = document.createElement('div');
        Notifdiv.classList.add('divbtn');
        Notifdiv.textContent = notif;
        divbtn.appendChild(Notifdiv);
        userElement.appendChild(divbtn);
    }

    return userElement;
}

// Create the button for the online and offline users
function updateConnectedUsers(connectedUsers, connectedUsersDiv) {
    const userElements = connectedUsersDiv.getElementsByTagName('button');
    Array.from(userElements).forEach(userElement => {
        userElement.style.color = connectedUsers.includes(userElement.id) ? "orange" : "white";
    });
}

async function startChat(user, socket) {
    //deletin notification 
    await delNotif(username, user);

    // fetch and sort older messages
    const oldmess = await LoadOldMessage(username, user);
    if (oldmess) {
        oldmess.sort((a, b) => new Date(a.Time) - new Date(b.Time));
    }
    //Check the older messages
    const totalMessages = oldmess ? oldmess.length : 0;
    let start = totalMessages ? totalMessages - 1 : 0;
    let end = Math.max(0, start - 10);

    //Create the chat
    const regdiv = document.getElementsByClassName('container')[0];
    regdiv.innerHTML = `
        <h1>Chat with ${user}</h1>
        <button id="backbtn">🔙</button>
        <div id="chat">
            <div id="messages" class="oldmessages"></div>
            <div id="messagesfields" class="messagesfields">
                <input type="text" id="messageInput" class="messageInput" placeholder="Type a message...">
                <button id="sendButton">Send</button>
            </div>
        </div>
    `;

    const messagesDiv = document.getElementById('messages');
    messagesDiv.style.overflow = 'scroll';
    messagesDiv.style.maxHeight = "150px";
    messagesDiv.style.display = 'flex';
    messagesDiv.style.flexDirection = 'column-reverse';

    loadMessages(start, end, messagesDiv, oldmess);

    //Manage the scrolling for displaying older message
    const handleScroll = throttle(() => {
        if (7 <= Math.abs(messagesDiv.scrollTop) + messagesDiv.clientHeight - 150 && start !== 0) {
            start = Math.max(0, start - 10);
            end = Math.max(0, start - 10);
            loadMessages(start, end, messagesDiv, oldmess);
        }
    }, 1000);

    messagesDiv.addEventListener('scroll', handleScroll);

    //Sending message
    document.getElementById('sendButton').addEventListener('click', () => {
        sendMessage(user, socket);
    });

    document.getElementById('messageInput').addEventListener('keypress', () => {
        Istyping(user, username, socket)
    });
    //Go back to message page
    document.getElementById('backbtn').addEventListener('click', () => {
        console.log("bye");
        socket.close();
        Mess();
    });
}

//Displaying the old messages
function loadMessages(start, end, messagesDiv, oldmess) {
    const initialScrollHeight = messagesDiv.scrollHeight;
    for (let i = start; i > end && i >= 0; i--) {
        const messdiv = document.createElement('p');
        messdiv.className = 'oldmessage';
        messdiv.id = oldmess[i].from;
        messdiv.innerHTML = `
            <strong>${oldmess[i].from}:</strong> ${oldmess[i].Oldmessages}
            <small>${oldmess[i].Time}</small>
        `;
        messagesDiv.appendChild(messdiv);
    }
    messagesDiv.scrollTop = messagesDiv.scrollHeight - initialScrollHeight;
}

//Send messsage througth the socket
function sendMessage(user, socket) {
    const messageInput = document.getElementById("messageInput");
    if (messageInput.value.trim() !== "") {
        const message = {
            from: username,
            to: user,
            message: messageInput.value
        };
        socket.send(JSON.stringify(message));
        messageInput.value = "";
    }
}

// Create the chat messages 
function handleMessage(message) {
    if (message.type === "chat_message") {
        const messagesDiv = document.getElementById("messages");
        const messageElement = document.createElement('div');
        messageElement.innerHTML = `<strong>${message.from}:</strong> ${message.message} ${message.Time}`;
        messagesDiv.insertBefore(messageElement, messagesDiv.firstChild);
    }
}

//Send to bakc for deleting notfs
async function delNotif(from, to) {
    const data = {
        from: from,
        to: to
    };
    try {
        await fetch('http://localhost:8080/delete-notif', {
            method: 'POST',
            body: JSON.stringify(data),
            headers: {
                'Content-Type': 'application/json'
            }
        });
    } catch (error) {
        console.error(error);
    }
}
