
//Function that creates the top bar of the graph
/**
 * 
 * @param {number} Result 
 * @param {string} Couleur 
 * @param {HTMLDivElement} div 
 */
export function BarGraphUp(Result, Couleur, div,) {
    let Updiv = document.createElement('div');

    Updiv.style.display = 'flex';
    Updiv.style.flexDirection = 'column';

    let Divtitle = document.createElement('h1');
    Divtitle.innerText = 'Audits ratio';
    Divtitle.style.marginBottom = '10px';
    Updiv.appendChild(Divtitle);

    let insideDiv = document.createElement('div');
    insideDiv.style.display = 'flex';
    insideDiv.style.justifyContent = 'space-between';
    insideDiv.style.flexDirection = 'row';

    let SvgBarUp = document.createElementNS('http://www.w3.org/2000/svg', 'svg');
    SvgBarUp.setAttribute('width', '70%');
    SvgBarUp.setAttribute('height', '50px');

    CreatesAxes(Couleur, Result, SvgBarUp)
    insideDiv.appendChild(SvgBarUp);
    Updiv.appendChild(insideDiv);


    let Title = document.createElement('h1');
    var CommaUp = numberWithCommas(Result)
    Title.innerText = CommaUp + ' MB';

    insideDiv.appendChild(Title);
    div.appendChild(Updiv);

}

//Function that creates the down bar of the graph
/**
 * 
 * @param {number} Result 
 * @param {string} Couleur 
 * @param {HTMLDivElement} div 
 */
export function BarGraphDown(Result, Couleur, div,) {

    let DownDiv = document.createElement('div');
    DownDiv.style.display = 'flex';
    DownDiv.style.justifyContent = 'space-between';
    DownDiv.style.flexDirection = 'row';

    let SvgBarDown = document.createElementNS('http://www.w3.org/2000/svg', 'svg');
    SvgBarDown.setAttribute('width', '70%');
    SvgBarDown.setAttribute('height', '50px');

    CreatesAxes(Couleur, Result, SvgBarDown)

    DownDiv.appendChild(SvgBarDown);


    let Title = document.createElement('h1');
    var CommaUp = numberWithCommas(Result)
    Title.innerText = CommaUp + ' MB';
    DownDiv.appendChild(Title);

    div.appendChild(DownDiv);

}

// Function who format the number with commas
/**
 * 
 * @param {number} x 
 * @returns {string} formatted
 */
function numberWithCommas(x) {
    let str = x.toString();
    let firstThreeDigits = str.slice(0, 3);
    let formatted = firstThreeDigits.replace(/^(\d)(\d{2})/, '$1,$2');
    return formatted;
}

// Function who format the number with one comma
/**
 * 
 * @param {number} x 
 * @returns {string} result
 */
function numberWithOneCommas(x) {
    let numStr = x.toString();
    let firstThreeDigits = numStr.slice(0, 3);
    let result = firstThreeDigits + "." + numStr.slice(3);
    return result;
}

// Function who creates the axes
/**
 * 
 * @param {string} Couleur 
 * @param {number} Result 
 * @param {HTMLDivElement} ParentDiv 
 */
function CreatesAxes(Couleur, Result, ParentDiv) {
    let SvgRect = document.createElementNS('http://www.w3.org/2000/svg', 'rect');
    let X2Down = numberWithOneCommas(Result);

    SvgRect.setAttribute('fill', Couleur);
    SvgRect.setAttribute('width', (X2Down * 1.7));
    SvgRect.setAttribute('height', '30');
    SvgRect.setAttribute('rx', '10');

    ParentDiv.appendChild(SvgRect);
}
