

export function Istyping(from, to, socket) {

    let data = {
        from: from,
        to: to,
        type: "typing"
    }

    try {
        socket.send(JSON.stringify(data));
    } catch (error) {
        console.error("Error sending typing message:", error);
    }
}



export function Typing(to) {
    let typingTimeout;

    let ChatDiv = document.getElementById('chat')

    if (document.getElementById('typingdiv')) {
        document.getElementById('typingdiv').remove()
    }
    let typingdiv = document.createElement('div')
    typingdiv.id = 'typingdiv'
    typingdiv.classList.add('typing')
    typingdiv.innerHTML = `
        <p id="typer" class="typer">${to} est en train d'écrire</p>
        <div id="animation-type" class="animation>
            <span class="dot">.</span>
            <span class="dot">.</span>
            <span class="dot">.</span>
        </div>
    `
    ChatDiv.appendChild(typingdiv)

    if (typingTimeout) {
        clearTimeout(typingTimeout);
    }

    // Set a new timeout to remove the typing indicator after a certain period of inactivity
    typingTimeout = setTimeout(() => {
        if (document.getElementById('typingdiv')) {
            document.getElementById('typingdiv').remove();
        }
    }, 10000);
}