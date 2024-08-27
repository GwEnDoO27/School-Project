import { XpsGraph } from '../../Api/Data.js';
import { CreateBackDiv } from '../lib/BackDiv.js';

//Somes variables
let Points = []
let HeightPoint = []
var y = 0

// Function who create the graph of the xps
/**
 * 
 * @param {Jwt} token 
 */
export async function Graphxp(token) {
    await SorteValues(token, HeightPoint, Points)
    CreateBackDiv('graphXp', '300px', '50%', 'xp');
    let div = document.getElementById('graphXp');

    let maxY = Math.max(...Points.map(p => p.y));
    let minX = Math.min(...Points.map(p => p.x));
    let maxX = Math.max(...Points.map(p => p.x));
    let minY = Math.min(...Points.map(p => p.y));

    let CoordHeight = maxY - minY;

    let graphSvg = Createsvg(CoordHeight, minX)

    GraphPoint(Points, HeightPoint, graphSvg, maxY)

    div.appendChild(graphSvg);

    let LineXxorig = ((minX * 10) - 25) - 200
    let LineXyorig = (maxX * 4) + 70
    let LineXy2orig = (maxY / 1000) - 200

    CreateAxes(LineXxorig, LineXyorig, LineXy2orig, LineXyorig, 'LineX')

    let LineYxorig = ((minX * 10) - 225)
    let LineYyorig = (maxX * 4) + 70
    CreateAxes(LineYxorig, LineYyorig, LineYxorig, '0', 'LineY')

}

// Function who create the axes
/**
 * 
 * @param {number} x1 
 * @param {number} y1 
 * @param {number} x2 
 * @param {number} y2 
 * @param {string} name 
 */
function CreateAxes(x1, y1, x2, y2, name) {
    let ParentDiv = document.getElementById('xpSvg');
    let Axes = document.createElementNS('http://www.w3.org/2000/svg', 'line');
    Axes.setAttribute('id', name);
    Axes.setAttribute('x1', x1);
    Axes.setAttribute('y1', y1);
    Axes.setAttribute('x2', x2);
    Axes.setAttribute('y2', y2);
    Axes.setAttribute('stroke', 'white');
    Axes.setAttribute('stroke-width', '2');
    ParentDiv.appendChild(Axes);
}

// Function who create the svg
/**
 * 
 * @param {number} CoordHeight 
 * @param {number} minX 
 * @returns {HTMLDivElement} graphSvg
 */
function Createsvg(CoordHeight, minX) {
    let graphSvg = document.createElementNS('http://www.w3.org/2000/svg', 'svg');
    graphSvg.setAttribute('id', 'xpSvg');
    graphSvg.setAttribute('viewBox', `${(minX * 10) - 100} 0 400 ${CoordHeight / 1800} `);
    graphSvg.style.margin = 'auto';
    return graphSvg
}

// Function who create the points of the graph
/**
 * 
 * @param {Array} Points 
 * @param {Array} HeightPoint 
 * @param {HTMLDivElement} graphSvg 
 * @param {number} maxY 
 */
function GraphPoint(Points, HeightPoint, graphSvg, maxY) {
    Points.forEach(element => {
        let circle = document.createElementNS('http://www.w3.org/2000/svg', 'circle');
        // Scale x from 0 to 400 (for the SVG width minus margins)
        let cx = (element.x * 10) - 200;
        // Scale y from 0 to 400 (for the SVG height minus margins), but we subtract to invert the Y axis
        let cy = 450 - ((element.y - Math.min(...HeightPoint)) * 800 / (maxY - Math.min(...HeightPoint)));

        circle.setAttribute('cx', cx);
        circle.setAttribute('cy', cy);
        circle.setAttribute('r', '5');
        circle.setAttribute('fill', 'blue');
        graphSvg.appendChild(circle);
    });
}

// Function who sort the values
/**
 * 
 * @param {Jwt} token 
 * @param {Array} HeightPoint 
 * @param {Array} Points 
 */
async function SorteValues(token, HeightPoint, Points) {
    let Result = await XpsGraph(token)
    let xps = Result.data.transaction

    let XpsSorted = xps.sort((a, b) => new Date(a.createdAt) - new Date(b.createdAt))

    let DistDate = Date.parse(XpsSorted[XpsSorted.length - 1].createdAt) - Date.parse(XpsSorted[0].createdAt)

    XpsSorted.forEach(element => {
        if (!(element.path.includes('piscine-go') || element.path.includes('piscine-js') || element.path.includes('checkpoint'))) {
            HeightPoint.push(element.amount)
            var NewDate = Date.parse(element.createdAt) - Date.parse(XpsSorted[0].createdAt);

            var x = ((NewDate * 100) / DistDate);
            y += element.amount
            Points.push({ x, y });
        }
    });
}