import { Data } from '../../Api/Data.js';
import { BarGraphDown, BarGraphUp } from '../Graphs/BarGraph.js';
import { CreateBackDiv } from './BackDiv.js';

// Function who create the ratio cart
/**
 * 
 * @param {Jwt} token 
 */
export async function RatioCart(token) {
    let Result = await Data(token)
    let div = CreateBackDiv('Ratio', '100%', '100%', 'MiddleCards');
    BarGraphUp(Result.data.user[0].totalUp, 'blue', div,);
    BarGraphDown(Result.data.user[0].totalDown, 'orange', div);
}
