import { Auth } from "../Auth/Auth.js";

//function that launch first for the login page
export async function Home() {
    try {

        CreateLeftContainer()
        let RightContainer = CreateRightContainer()
        let loginField = CreateLoginField()
        let login = document.createElement('ul');

        //Create the username input
        UsernameInput(login)
        //Create the password input
        PasswordInput(login)

        loginField.appendChild(login);
        CreateButton(loginField)

        RightContainer.appendChild(loginField);

        //Add event listener to the login button
        document.getElementById('loginButton').addEventListener('click', async () => {

            let username = document.getElementById('UsernameInput').value
            let password = document.getElementById('passwordInput').value;
            //Call the Auth function to authenticate the user
            Auth(username, password)
        });
    }
    catch (e) {
        console.log(e)
    }
}
Home();

//Function that create the left container
/**
 * 
 * @returns {HTMLDivElement} LeftContainer
 */
function CreateLeftContainer() {
    let Body = document.querySelector('body')
    let LeftContainer = document.createElement('div');
    LeftContainer.id = 'LeftContainer';
    LeftContainer.style.padding = '100px 0px 125px 125px';
    LeftContainer.style.width = '50%';
    LeftContainer.style.height = '100%';
    LeftContainer.style.right = '0';
    LeftContainer.style.minHeight = '100vh';
    LeftContainer.style.zIndex = '2';
    LeftContainer.style.position = 'relative';
    LeftContainer.style.display = 'flex';
    LeftContainer.style.flexDirection = 'column';
    LeftContainer.style.alignItems = 'center';
    LeftContainer.style.justifyContent = 'center';

    Body.appendChild(LeftContainer);
    return LeftContainer
}

//Function that create the right container
/**
 * 
 * @returns {HTMLDivElement} RightContainer
 */
function CreateRightContainer() {
    let Body = document.querySelector('body')
    let RightContainer = document.createElement('div');
    RightContainer.id = 'RightContainer';
    RightContainer.style.padding = '100px 100px 125px 0px';
    RightContainer.style.width = '50%';
    RightContainer.style.left = '0';
    RightContainer.style.minHeight = '100vh';
    RightContainer.style.zIndex = '2';
    RightContainer.style.position = 'relative';
    RightContainer.style.display = 'flex';
    RightContainer.style.flexDirection = 'column';
    RightContainer.style.alignContent = 'center';
    RightContainer.style.justifyContent = 'center';
    RightContainer.style.fontFamily = 'IBM Plex Mono, monospace';

    Body.appendChild(RightContainer);
    return RightContainer
}

//Function that create the password input
/**
 * 
 * @param {HTMLDivElement} login 
 */
function PasswordInput(login) {
    let passwordField = document.createElement('li');
    passwordField.setAttribute('id', 'passwordField');


    let passwordlabel = document.createElement('label');
    passwordlabel.setAttribute('for', 'passwordlabel');
    passwordlabel.innerText = 'Password';
    passwordField.appendChild(passwordlabel);

    let passwordInput = document.createElement('input');
    passwordInput.setAttribute('type', 'password');
    passwordInput.setAttribute('id', 'passwordInput');
    passwordInput.setAttribute('class', 'passwordInput');
    passwordField.appendChild(passwordInput);

    login.appendChild(passwordField);

}

//Function that create the username input
/**
 * 
 * @param {HTMLDivElement} login 
 */
function UsernameInput(login) {
    let usernameField = document.createElement('li');
    usernameField.setAttribute('id', 'usernameField');

    let UsernameLabel = document.createElement('label');
    UsernameLabel.setAttribute('for', 'usernameField');
    UsernameLabel.innerText = 'Email ou Nom d\'utilisateur';
    usernameField.appendChild(UsernameLabel);

    let usernameInput = document.createElement('input');
    usernameInput.setAttribute('type', 'text');
    usernameInput.setAttribute('id', 'UsernameInput');
    usernameInput.setAttribute('class', 'UsernameInput');
    usernameField.appendChild(usernameInput);

    login.appendChild(usernameField);

}

//Function that create the login field
/**
 * 
 * @returns {HTMLDivElement} loginField
 */

function CreateLoginField() {
    let loginField = document.createElement('div');
    loginField.setAttribute('class', 'loginField');
    loginField.setAttribute('id', 'loginField');
    loginField.style.display = 'flex';
    loginField.style.width = '100%';
    loginField.style.flexDirection = 'column';
    return loginField
}

//Function that create the login button
/**
* 
* @param {HTMLDivElement} loginField 
*/
function CreateButton(loginField) {
    let loginButton = document.createElement('button');
    loginButton.setAttribute('id', 'loginButton');
    loginButton.setAttribute('class', 'loginButton');
    loginButton.innerHTML = 'Login';
    loginButton.style.hover = 'cursor:pointer, background-color: rgb(76, 0, 250)';
    loginButton.style.width = '15%';
    loginButton.style.alignItems = 'center';
    loginButton.style.justifyContent = 'center';
    loginButton.style.borderRadius = '10px';
    loginButton.style.marginLeft = '40%';

    loginField.appendChild(loginButton);

}
