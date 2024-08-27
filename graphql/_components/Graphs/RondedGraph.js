
//Function to create a rounded graph
/**
 * 
 * @param {string} stroke 
 * @param {string} Circum 
 * @param {string} main 
 * @param {string} name 
 * @param {string} Result 
 * @param {HTMLDivElement} div 
 * @returns {HTMLDivElement}  DivSvg
 */
export async function CreateRounded(stroke, Circum, main, name, Result, div) {
    let DivSvg = document.createElement('div');
    let data = TransctionExtractSkill(Result.data.transaction, name)
    let strokeDashoffset = Math.round(Circum * ((100 - data[0].amount) / 100)) + 'px'
    let svg = CreateRoundedSvg(main)
    let Circle1 = CreateCircle1(stroke, Circum, '0',)
    let Circle2 = CreateCircle2(stroke, Circum, strokeDashoffset,)
    let text = Title(data[0].amount, stroke)
    let skill = NameOfSkill(name, stroke)
    svg.appendChild(Circle1)
    svg.appendChild(Circle2)
    svg.appendChild(text)
    DivSvg.appendChild(skill)
    DivSvg.appendChild(svg)
    div.appendChild(DivSvg)
    return DivSvg
}

//Function to create the svg
/**
 * 
 * @param {HTMLDivElement} main 
 * @returns 
 */
export function CreateRoundedSvg(main) {
    let div = document.createElementNS('http://www.w3.org/2000/svg', 'svg');
    div.id = 'RoundedSvg';
    div.setAttribute('width', 200);
    div.setAttribute('height', 200);
    div.setAttribute('viewBox', '-25 -25 25à 250');
    div.setAttribute('version', "1.1")
    div.setAttribute('style', 'transform:rotate(-90deg)')
    main.appendChild(div);
    return div
}

//Function to create the first circle
/**
 * 
 * @param {string} Circum 
 * @param {string} strokeDashoffset 
 * @returns {HTMLDivElement} Circle1
 */
export function CreateCircle1(Circum, strokeDashoffset) {
    let Circle1 = document.createElementNS('http://www.w3.org/2000/svg', 'circle');
    Circle1.setAttribute('cx', 100);
    Circle1.setAttribute('cy', 100);
    Circle1.setAttribute('r', 90);
    Circle1.style.fill = 'transparent';
    Circle1.style.stroke = '#e0e0e0';
    Circle1.style.strokeWidth = '16px';
    Circle1.style.strokeDasharray = Circum;
    Circle1.setAttribute('stroke-dashoffset', strokeDashoffset);

    return Circle1
}


//Function to create the second circle
/**
 * 
 * @param {string} stroke 
 * @param {string} Circum 
 * @param {string} strokeDashoffset 
 * @returns {HTMLDivElement} Circle2
 */

export function CreateCircle2(stroke, Circum, strokeDashoffset) {
    let Circle2 = document.createElementNS('http://www.w3.org/2000/svg', 'circle');
    Circle2.setAttribute('cx', 100);
    Circle2.setAttribute('cy', 100);
    Circle2.setAttribute('r', 90);
    Circle2.style.fill = 'transparent';
    Circle2.style.stroke = stroke;
    Circle2.style.strokeWidth = '16px';
    Circle2.style.strokeLinecap = 'round';
    Circle2.style.strokeDasharray = Circum;
    Circle2.setAttribute('stroke-dashoffset', strokeDashoffset);

    return Circle2
}

//Function to create the Percent text
/**
 * 
 * @param {string} Percent 
 * @param {string} color 
 * @returns {HTMLDivElement} text
 */
export function Title(Percent, color) {
    let text = document.createElementNS('http://www.w3.org/2000/svg', 'text');
    text.setAttribute('x', 69);
    text.setAttribute('y', 155);
    text.setAttribute('fill', color);
    text.setAttribute('font-size', '52px');
    text.setAttribute('font-weight', 'bold');
    text.setAttribute('style', 'transform:rotate(90deg) translate(-12px, -235px)');
    text.textContent = Percent + '%';

    return text
}

//Function to create the name of the skill
/**
 * 
 * @param {string} name 
 * @param {string} color 
 * @returns {HTMLDivElement} text
 */
export function NameOfSkill(name, color) {
    let text = document.createElement('h1')
    text.style.color = color;
    text.style.fontSize = '24px';
    text.style.fontWeight = 'bold';
    text.style.textAlign = 'center';
    text.style.justifyContent = 'center';
    text.style.alignContent = 'center';
    text.style.marginBottom = '20px';
    text.textContent = name;

    return text
}

//Function to extract the skill for the svg from the data
/**
 * 
 * @param {Array} data 
 * @param {string} name 
 * @returns {Array} Skill
 */
export function TransctionExtractSkill(data, name) {
    let Skill = {};
    let skillName = 'skill_' + name;
    data.forEach(element => {
        var SkillContain = element.type;
        if (SkillContain.includes(skillName)) {
            Skill[element.type] = element;
        }
    });
    Skill = SortSkills(Object.values(Skill));
    return Skill;
}

//Function to sort the skills
/**
 * 
 * @param {Array} Skills 
 * @returns {Array} sorted
 */
function SortSkills(Skills) {
    let sorted = Skills.sort((a, b) => {
        return a.amount - b.amount;
    })
    return sorted
}

