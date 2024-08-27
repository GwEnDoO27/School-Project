
import { ThreeLastProjects } from '../../Api/Data.js';
import { CreateBackDiv } from './BackDiv.js';

// Function who create the last projects cart
/**
 * 
 * @param {Jwt} token 
 */
export async function LastProjectsCart(token) {
    let Result2 = await ThreeLastProjects(token)
    CreateBackDiv('LastProjectsCart', '100%', '100%', 'MiddleCards');
    let div = document.getElementById('LastProjectsCart');

    var LastThreeProjects = Result2.data.user[0].xps;
    let Names = []

    LastThreeProjects.forEach(element => {
        let Path = element.path
        let Parties = Path.split('/')
        let LastPart = Parties[Parties.length - 1]
        let AddSpace = LastPart.replaceAll('-', ' ')
        let UpperCase = AddSpace.charAt(0).toUpperCase() + AddSpace.slice(1)
        Names.push(UpperCase)
    });

    CreateTitle('Last Projects', div)

    let Projects = CreateTable()

    Names.forEach(element => {
        CreateTr(element, Projects)
    });

    div.appendChild(Projects);
}

// Function who create the title
/**
 * 
 * @param {string} Name 
 * @param {HtmlElement} ParentDiv 
 */
function CreateTitle(Name, ParentDiv) {
    let Title = document.createElement('h1');
    Title.style.display = 'flex';
    Title.style.paddingLeft = '20px';
    Title.style.alignItems = 'flex-start';
    Title.innerText = Name;
    ParentDiv.appendChild(Title);
}

// Function who create the table
/**
 * 
 * @returns {HtmlElement} Table
 */
function CreateTable() {
    let Table = document.createElement('table');
    Table.style.display = 'flex';
    Table.style.justifyContent = 'center';
    Table.style.alignItems = 'flex-start';
    Table.style.flexDirection = 'column';
    Table.style.width = '100%';
    Table.style.height = '100%';
    Table.style.borderRadius = '10px';
    Table.style.paddingLeft = '20px';
    Table.style.gap = '20px';
    return Table;
}

// Function who create the tr
/**
 * 
 * @param {string} element 
 * @param {HtmlElement} Projects 
 */
function CreateTr(element, Projects) {
    let tr = document.createElement('tr');
    tr.style.display = 'flex';
    tr.style.justifyContent = 'flex-start';
    tr.style.alignItems = 'center';
    tr.style.fontSize = '20px';
    tr.style.fontSize = 'italic';
    tr.innerText = '➡ ' + element;
    Projects.appendChild(tr);
}