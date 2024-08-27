import { Body } from "../Body.js";


//Function that create the login field
/**
 * @param {string} username 
 * @param {string} password 
 * @returns {HtmlElement} loginField
 */
export async function Auth(username, password) {
    let data = {
        username: username,
        password: password
    }
    //Fetch the data from the api
    const response = await fetch('https://zone01normandie.org/api/auth/signin', {
        method: 'POST',
        body: JSON.stringify(data),
        headers: {
            'Authorization': "Basic " + btoa(username + ":" + password)
        }
    }).then(response => response.json())
    //If the response has an error display it
    if (response.error) {
        await ErrorDiv(response.error)
    } else {
        document.getElementById('LeftContainer').remove()
        document.getElementById('RightContainer').remove()
        //Go to the landing page
        Body(response)
    }

}

//Function that create the error div
/**
 * 
 * @param {object} response 
 */
async function ErrorDiv(response) {
    let div = document.createElement('div');
    let container = document.getElementById('LeftContainer');
    div.id = 'ErrorDiv';
    div.style.height = '50%%';
    div.style.width = '90%';
    div.style.alignItems = 'center';
    div.style.justifyContent = 'center';
    div.style.fontFamily = 'IBM Plex Mono, monospace';
    div.style.display = 'flex';
    div.style.fontSize = '30px';
    div.style.fontWeight = "bold";
    div.style.color = 'rgb(133, 74, 250)';
    div.style.borderRadius = '10px';
    div.style.padding = '20px';
    div.style.backgroundColor = 'rgba(255, 255, 255,0.9)';

    div.innerText = response + ' Please try again';
    container.appendChild(div)
}