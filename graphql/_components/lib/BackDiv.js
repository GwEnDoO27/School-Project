
// This function creates the div back for each data
/**
 * 
 * @param {String} Name 
 * @param {String} Height 
 * @param {String} width 
 * @param {String} container 
 * @returns {HTMLDivElement} div
 */
export function CreateBackDiv(Name, Height, width, container) {
    let div = document.createElement('div');
    div.id = Name;
    div.style.height = Height;
    div.style.width = width;
    div.style.borderRadius = '10px';
    div.style.backgroundColor = 'rgba(127, 127, 127, 0.93)';
    div.style.color = 'white';
    div.style.padding = '10px';
    div.style.display = 'flex';
    div.style.flexDirection = 'column';
    div.style.marginTop = '10px';
    div.style.marginRight = '10px';
    div.style.alignContent = 'center';
    div.style.justifyContent = 'center';

    let Appending = document.getElementById(container);
    Appending.appendChild(div);

    return div
}