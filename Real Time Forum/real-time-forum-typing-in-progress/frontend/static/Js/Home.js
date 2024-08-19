import { NewCSS } from "./Newcss.js"
import { MoreInfos } from "./Post.js"
import { CreatPost } from "./CreateNewPost.js"
import { Mess } from "./Mess.js"
import { Logout } from "./logout.js"
import { getCookie, DecodeUsername, username } from "./Cookie.js";

let cont = document.getElementById('container')
let regdiv = document.getElementsByClassName('container')[0]


async function emptyconst() {
    return await fetch('http://localhost:8080/Home').then(val => val.json())
}

//Display all the post on homepage
export async function Home() {
    NewCSS(['/static/Css/index.css'])
    document.getElementById('nav').style.display = 'block'

    const sessionCookie = getCookie('session-token');
    const username = DecodeUsername(sessionCookie)
    const response = await emptyconst()
    regdiv.innerHTML = ``
    try {
        regdiv.classList.add('container')
        response.forEach(element => {
            let postdiv = document.createElement('div')
            postdiv.classList.add('cont')
            postdiv.id = 'post'
            postdiv.innerHTML = `
                <div class="form_area" id="post">
                    <div class="area" id="post">
                        <div class="form_style" id="post">
                            <h1 class="title" id="post">${element.Title}</h1>
                            <h2 class="Categories" id="post">${element.Categories}</h2>
                            <p class="Description" id="post">${element.Description}</p>
                            <p class="Post" id="post">${element.Post}</p>
                            <p class="Author" id="post">${element.Author}</p>
                            <img src="${element.Img}">
                            <button value="${element.ID}" class="Comments" id="postbyID${element.ID}"">View Comments</button>
                        </div>
                    </div
                </div>
                `
            cont.appendChild(postdiv)

            //get hte id of the post and allow to see the coms
            document.getElementById(`postbyID${element.ID}`).addEventListener('click', () => {
                let id = document.getElementById(`postbyID${element.ID}`).value
                MoreInfos(id)
            })
        });

        //Manage the different buttons from the nav bar
        document.getElementById('home').addEventListener('click', () => {
            Home()
        })
        document.getElementById('CreatPost').addEventListener('click', () => {
            CreatPost()
        })
        document.getElementById('Messages').addEventListener('click', () => {
            Mess()
        })
        document.getElementById('logout').addEventListener('click', () => {
            Logout()
        })

    }
    catch (error) {
        console.error(error)
    }
}
