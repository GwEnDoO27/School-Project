
import { Graphxp } from "./_components/Graphs/Graphxp.js";
import { CreateBvsFt } from "./_components/Graphs/LangauageVersus.js";
import { AuditRatio } from "./_components/lib/AuditRatio.js";
import { Header } from "./_components/lib/Header.js";
import { LastProjectsCart } from "./_components/lib/LastProjectsCart.js";
import { RatioCart } from "./_components/lib/RatioCart.js";


//Function that create the body of the page
/**
 * 
 * @param {string} token 
 */
export async function Body(token) {
    let cont = BodyAttrs(token);
    let main = CreateMaincont(cont);

    await CreateLittleCard(main, token);
    await CreateXpGraph(main, token);
    await CreateBvsFt(main, token)
}

//Function that create the body of the page
/**
 * 
 * @param {string} token 
 * @returns {HTMLDivElement} cont
 */
function BodyAttrs(token) {
    let cont = document.querySelector('body');
    cont.id = 'Body';
    cont.style.width = '100%';
    cont.style.minHeight = '100vh';
    cont.style.zIndex = '0';
    cont.style.position = 'relative';
    cont.style.display = 'flex';
    cont.style.flexDirection = 'column';
    cont.style.fontFamily = 'IBM Plex Mono, monospace';
    cont.style.gap = '10px';
    Header(cont, token)

    return cont
}

//Function that create the little cards div
/**
 * 
 * @param {string} token 
 * @param {HTMLDivElement} cont
 */
async function CreateLittleCard(cont, token) {

    let MiddleCards = document.createElement('div');
    MiddleCards.setAttribute('id', 'MiddleCards');
    MiddleCards.style.padding = '0 100px'
    MiddleCards.id = 'MiddleCards';
    MiddleCards.style.display = 'flex';
    MiddleCards.style.flexDirection = 'row';
    MiddleCards.style.justifyContent = 'space-around';
    MiddleCards.style.alignContent = 'flex-start';
    MiddleCards.style.height = '200px';
    MiddleCards.style.marginTop = '10px';
    MiddleCards.style.marginBottom = '10px';
    MiddleCards.style.zIndex = '1';
    cont.appendChild(MiddleCards);

    // Create the cards
    RatioCart(token)
    AuditRatio(token)
    LastProjectsCart(token)
}

//Function that create the xp graph
/**
 * 
 * @param {HTMLDivElement} cont 
 * @param {string} token 
 * @returns {HTMLDivElement} xp
 */
async function CreateXpGraph(cont, token) {
    let xp = document.createElement('div');
    xp.id = 'xp';
    xp.style.width = '100%';
    xp.style.height = '300px';
    xp.style.display = 'flex';
    xp.style.flexDirection = 'row';
    xp.style.padding = '0 100px'
    xp.style.justifyContent = 'space-around';
    xp.style.gap = '10px';
    cont.appendChild(xp);
    CreateTitle(xp)
    Graphxp(token)
    return xp
}


//Function that create the title of the page
/**
 * 
 * @param {HTMLDivElement} cont 
 * @returns {HTMLDivElement} title
 */
function CreateTitle(cont) {
    let title = document.createElement('div');
    title.id = 'Title';
    title.style.width = '50%';
    title.style.height = '300px';
    title.style.display = 'flex';
    title.style.fontfamily = 'IBM Plex Mono, monospace';
    title.style.justifyContent = 'center';
    title.style.alignItems = 'center';
    title.style.fontSize = '100px';
    title.style.fontWeight = 'bold';
    title.style.textAlign = 'center';
    title.style.color = 'rgb(76, 0, 250)';
    title.textContent = 'Experience Graph';
    cont.appendChild(title);
    return title
}

//Function that create the main container
/**
 * 
 * @param {HTMLDivElement} cont 
 * @returns {HTMLDivElement} main
 */
function CreateMaincont(cont) {
    let main = document.createElement('div');
    main.id = 'Main';
    main.style.width = '100%';
    main.style.height = '100%';
    main.style.display = 'flex';
    main.style.flexDirection = 'column';
    main.style.padding = '0 100px'
    main.style.gap = '10px';

    cont.appendChild(main);
    return main

} 
