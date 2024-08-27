import { query, query2, query3 } from "./Fetch.js";


// Function who fetch the data from the api
/**
 * 
 * @param {Jwt} token 
 * @returns {Promise} response
 */
export async function Data(token) {
    const response = await fetch('https://zone01normandie.org/api/graphql-engine/v1/graphql', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + token
        },
        body: JSON.stringify({
            query: query
        })
    })
    return response.json();
}

// Function who fetch the data from the api for the last 3 projects
/**
 * 
 * @param {Jwt} token 
 * @returns {Promise} response
 */
export async function ThreeLastProjects(token) {
    const response = await fetch('https://zone01normandie.org/api/graphql-engine/v1/graphql', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + token
        },
        body: JSON.stringify({
            query: query2
        })
    })
    return response.json();
}

// Function who fetch the data from the api for the xps
/**
 * 
 * @param {Jwt} token 
 * @returns {Promise} response
 */
export async function XpsGraph(token) {
    const response = await fetch('https://zone01normandie.org/api/graphql-engine/v1/graphql', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + token
        },
        body: JSON.stringify({
            query: query3
        })
    })
    return response.json();
}

// Function who calculate the total of xps
/**
 * 
 * @param {Array} data 
 * @returns {Number} amount
 */
export async function Xps(data) {
    let amount = 0;
    data.forEach(element => {
        var Path = element.path

        if (!(Path.includes('piscine-go/')) && !(Path.includes('piscine-js/')) && !(Path.includes('checkpoint'))) {
            amount += +element.amount;
        }
    });
    return amount;
}



