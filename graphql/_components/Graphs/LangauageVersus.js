
import { Data } from '../../Api/Data.js';
import { CreateRounded } from './RondedGraph.js';

// Function who create the RoundedGraphs div
/**
 * 
 * @param {HTMLDivElement} main 
 * @param {Jwt} token 
 */
export async function CreateBvsFt(main, token) {
  let Result = await Data(token)

  let GraphsDivs = CreateGraphDivs(main)

  let DivSvg =  CreateRounded('#6bdba7', '565', main, 'go', Result, GraphsDivs)
  let DivSvg2 =  CreateRounded('#de516a', '565', main, 'js', Result, GraphsDivs)
  let DivSvg3 =  CreateRounded('#864aff', '565', main, 'back-end', Result, GraphsDivs)
  let DivSvg4 =  CreateRounded('#dc5a3c', '565', main, 'front-end', Result, GraphsDivs)

}

//Function to create the div who contain the graphs
/**
 * 
 * @param {HTMLDivElement} main 
 * @returns {HTMLDivElement} GraphsDivs
 */
function CreateGraphDivs(main) {
  let GraphsDivs = document.createElement('div');
  GraphsDivs.id = 'GraphsDivs';
  GraphsDivs.style.display = 'flex';
  GraphsDivs.style.flexDirection = 'row';
  GraphsDivs.style.justifyContent = 'space-around';
  GraphsDivs.style.alignContent = 'flex-start';
  GraphsDivs.style.height = '300px';
  GraphsDivs.style.backgroundColor = 'rgba(127, 127, 127, 0.93)';
  GraphsDivs.style.borderRadius = '10px';
  GraphsDivs.style.padding = '20px'
  GraphsDivs.style.margin = '20px 100px'

  main.appendChild(GraphsDivs);

  return GraphsDivs
}



