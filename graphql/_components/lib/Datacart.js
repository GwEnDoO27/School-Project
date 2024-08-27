
import { Data } from '../../Api/Data.js';
import { CreateBackDiv } from './BackDiv.js';

// Function who create the personal data div
/**
 * 
 * @param {Jwt} token 
 * @param {string} width 
 */
export async function PersonalData(token, width) {
    let Result = await Data(token)

    let PersonalData = CreateBackDiv('PersonalData', '100%', width, 'Body');
    PersonalData.style.justifyContent = 'auto';

    let div = SidenavDivSys(PersonalData);
    let main = document.getElementById('Main');
    let title = document.getElementById('Title');


    let CloseButton = document.createElement('button');

    SidenavDivSysClose(CloseButton, div)

    InfInsideDiv(Result, div)



    CloseButton.addEventListener('click', () => {
        PersonalData.style.top = '100%';
        PersonalData.style.transition = '0.5s';
        setTimeout(() => {
            PersonalData.remove()
            main.style.transition = 'margin-right 0.5s';
            main.style.width = '100%'
            title.style.fontSize = '100px'
        }, 500)
    })
}


// Function who display the personal data div
/**
 * 
 * @param {HTMLDivElement} PersonalData 
 * @returns {HTMLDivElement} PersonnalData
 */

function SidenavDivSys(PersonalData) {
    PersonalData.style.top = '0';
    PersonalData.style.position = 'fixed';
    PersonalData.style.right = '0';
    PersonalData.style.backgroundColor = 'rgb(127,127,127)';
    PersonalData.style.color = 'black';
    PersonalData.style.padding = '10px';
    PersonalData.overflowx = 'hidden';
    PersonalData.style.zIndex = '4';
    PersonalData.style.transition = '0.5s';

    setTimeout(() => {
        PersonalData.style.width = '30%';
    }, 500)

    return PersonalData
}

// Function who create the close button
/**
 * 
 * @param {HTMLButtonElement} CloseButton 
 * @param {HTMLDivElement} div 
 */
function SidenavDivSysClose(CloseButton, div) {
    CloseButton.innerText = 'X';
    CloseButton.style.position = 'absolute';
    CloseButton.style.right = '10px';
    CloseButton.style.top = '10px';
    CloseButton.style.backgroundColor = 'rgba(0,0,0,0.5)';
    CloseButton.style.color = 'white';
    CloseButton.style.border = 'none';
    CloseButton.style.borderRadius = '50%';
    CloseButton.style.padding = '5px';
    CloseButton.style.cursor = 'pointer';
    CloseButton.style.zIndex = '5';
    CloseButton.style.transition = '0.5s';
    CloseButton.style.animation = 'fadeIn 1s';
    div.appendChild(CloseButton);
}

// Function who display the personal data
/**
 * 
 * @param {Data} Result 
 * @param {HTMLDivElement} div 
 */
function InfInsideDiv(Result, div) {
    let infosTable = document.createElement('ul');
    infosTable.style.display = 'flex';
    infosTable.style.flexDirection = 'column';
    infosTable.style.justifyContent = 'flex-start';
    infosTable.style.color = 'white';
    infosTable.style.gap = '50px';
    infosTable.style.padding = '10px';
    infosTable.style.top = '10px';

    let Name = document.createElement('h1');
    Name.innerText = Result.data.user[0].attrs.firstName + ' ' + Result.data.user[0].attrs.lastName;
    Name.style.fontSize = '50px';
    Name.style.fontWeight = 'bold';
    Name.style.textAlign = 'center';
    infosTable.appendChild(Name);

    let Email = document.createElement('h2');
    Email.style.fontSize = '12px';
    Email.innerText = 'Email :' + Result.data.user[0].attrs.email;
    infosTable.appendChild(Email);

    let Adress = document.createElement('h2');
    Adress.innerText = 'Adress :' + Result.data.user[0].attrs.addressStreet + ',' + Result.data.user[0].attrs.addressPostalCode + ',' + Result.data.user[0].attrs.addressCity;
    Adress.style.fontSize = '12px';
    infosTable.appendChild(Adress)

    let Phone = document.createElement('h2');
    Phone.innerText = 'Phone :' + Result.data.user[0].attrs.Phone;
    Phone.style.fontSize = '12px';
    infosTable.appendChild(Phone);

    let Birth = document.createElement('h2');
    var NewDate = new Date(Result.data.user[0].attrs.dateOfBirth);
    Birth.innerText = 'Birthday :   Born ' + NewDate.toDateString() + ' in ' + Result.data.user[0].attrs.placeOfBirth;
    Birth.style.fontSize = '12px';
    infosTable.appendChild(Birth);

    div.appendChild(infosTable);

}