
import { Data } from '../../Api/Data.js';
import { PersonalData } from './Datacart.js';

// Function who create the header
/**
 * 
 * @param {HTMLDivElement} cont 
 * @param {Jwt} token 
 */
export async function Header(cont, token) {
    let Header = CreateHearderDiv(cont)
    let button = LogoutBtn(Header)
    let buttonInfos = await CreateUserInfos(Header, token)
    let main = document.getElementById('Main');
    let title = document.getElementById('Title');


    buttonInfos.addEventListener('click', () => {
        main.style.transition = 'margin-right 0.5s';
        main.style.width = '70%'
        title.style.fontSize = '70px'
        main.padding = '0px 0xp 0px 0px'
        PersonalData(token, '30%')

    })

    button.addEventListener('click', () => {
        localStorage.removeItem('remember_token')
        window.location.reload()
    })

}

// Function who create the header div
/**
 * 
 * @param {HTMLDivElement} cont 
 * @returns {HTMLDivElement} Header
 */
function CreateHearderDiv(cont) {
    let Header = document.createElement('div');
    Header.id = 'Header';
    Header.style.height = '100px';
    Header.style.width = '100%';
    Header.style.backgroundColor = 'rgb(225,225,225)';
    Header.style.color = 'white';
    Header.style.padding = '10px';
    Header.style.display = 'flex';
    Header.style.flexDirection = 'row';
    Header.style.alignContent = 'center';
    Header.style.position = 'sticky';
    Header.style.top = '';
    Header.style.zIndex = '3';
    Header.style.background = 'linear-gradient(rgba(221,221,221,1) 0%, rgba(255,255,255,0) 100%)';
    cont.appendChild(Header);
    return Header
}

// Function who create the logout button
/**
 * 
 * @param {HTMLDivElement} Header 
 * @returns {HTMLButtonElement} button
 */
function LogoutBtn(Header) {
    let button = document.createElement('button');
    button.setAttribute('id', 'logout');
    button.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="black" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-log-out"><path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/><polyline points="16 17 21 12 16 7"/><line x1="21" x2="9" y1="12" y2="12"/></svg>';
    button.style.right = '10px';
    button.style.position = 'absolute';
    button.style.right = '20px';
    button.style.top = '25px';

    Header.appendChild(button);

    return button

}

// Function who create the user infos button
/**
 * 
 * @param {HTMLDivElement} Header 
 * @param {Jwt} token 
 * @returns {HTMLButtonElement} buttonInfos
 */
async function CreateUserInfos(Header, token) {
    let Result = await Data(token)
    let buttonInfos = document.createElement('button');
    buttonInfos.setAttribute('id', 'UserInfos');
    buttonInfos.innerHTML = `${Result.data.user[0].attrs.firstName}<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="black" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-user-pen"><path d="M11.5 15H7a4 4 0 0 0-4 4v2"/><path d="M21.378 16.626a1 1 0 0 0-3.004-3.004l-4.01 4.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z"/><circle cx="10" cy="7" r="4"/></svg>`;
    buttonInfos.style.display = 'flex';
    buttonInfos.style.flexDirection = 'row';
    buttonInfos.style.justifyContent = 'space-around';
    buttonInfos.style.alignItems = 'center';
    buttonInfos.style.gap = '20px';
    buttonInfos.style.right = '100px';
    buttonInfos.style.color = 'black';
    buttonInfos.style.marginRight = '10px';
    buttonInfos.style.alignItems = 'center';
    buttonInfos.style.position = 'absolute';
    buttonInfos.style.top = '25px';
    Header.appendChild(buttonInfos);

    return buttonInfos
}
