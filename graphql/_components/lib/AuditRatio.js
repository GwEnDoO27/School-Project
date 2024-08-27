import { Data } from '../../Api/Data.js';
import { CreateBackDiv } from './BackDiv.js';

// Function who display the AuditRatio  div
/**
 * 
 * @param {Jwt} token 
 */
export async function AuditRatio(token) {
    let Result = await Data(token)

    var Nb = Result.data.user[0].auditRatio

    CreateBackDiv('RatioNumber', '100%', '174px', 'MiddleCards');
    let div = document.getElementById('RatioNumber')

    var RoundedNb = parseFloat(Nb).toFixed(1)

    let infosTable = document.createElement('h1');
    infosTable.style.display = 'flex';
    infosTable.style.justifyContent = 'center';
    infosTable.style.alignItems = 'center';
    infosTable.style.fontSize = '80px';
    infosTable.innerText = RoundedNb;

    div.appendChild(infosTable);
}