import { NewCSS } from "./Newcss.js"
import { Home } from "./Home.js"

// Fetching post by ids
async function getPost(id) {
    try {
        const response = await fetch('http://localhost:8080/post', {
            method: 'POST',
            body: JSON.stringify({ id: id, getPost: 'true' }),
            headers: {
                'Content-Type': 'application/json'
            }
        });
        return await response.json();
    } catch (error) {
        console.error("Failed to fetch post:", error);
        throw error;
    }
}

let cont = document.getElementById('container')
let regdiv = document.getElementsByClassName('container')[0]

//Displyaing post by ids
export async function MoreInfos(id) {
    console.log(id)
    NewCSS(['/static/Css/post.css'])

    try {
        const response = await getPost(id)
        const page = document.createElement('div')
        page.classList.add('Page')

        regdiv.classList.add('container')
        regdiv.innerHTML = ``
        regdiv.innerHTML = `
            <div class="cont" id="post">
                <div class="form_area" id="post">
                    <div class="form_style" id="post">
                        <h1 class="title" id="post">${response.Post.Title}</h1>
                        <h2 class="Categories" id="post">${response.Post.Categories}</h2>
                        <p class="Description" id="post">${response.Post.Description}</p>
                        <p class="Post" id="post">${response.Post.Post}</p>
                        <p class="Author" id="post">${response.Post.Author}</p>
                    </div>
                </div>
            </div>

            <div class="PostForm">
                <div class="cont" id="reply">
                    <div class="form_area" id="reply">
                        <div class="area" id="reply">
                            <input placeholder="Enter your message" class="form_style" id="replyinput" name="reply">
                            <button class="form_style" id="replyButton">Send</button>
                        </div>
                    </div>
                </div>
            </div>
        `;

        //Display all coms
        if (response.Comment != null) {
            response.Comment.forEach(element => {
                console.log(element);
                let coms = document.createElement('div')
                coms.classList.add('cont')
                coms.innerHTML = `
                <div class="form_area" id="commentaries">
                    <div class="form_style" id="commentaries">
                        <p class="Comments" id="commentaries">${element.Coms}</p>
                        <p class="username" id="commentaries">${element.Author}</p>
                    </div>
                </div>
            `;
                page.appendChild(coms)
            });
        }

        cont.appendChild(page)

        // Manage the sending of post
        document.getElementById('replyButton').addEventListener('click', async () => {
            let com = document.getElementById('replyinput').value

            let data = {
                Coms: com,
                Posting: 'true',
                id: id,
            }

            console.log(data)

            try {
                let GetfromApi = await fetch('http://localhost:8080/post', {
                    method: 'POST',
                    body: JSON.stringify(data),
                    headers: {
                        'Content-Type': 'application/json'
                    }
                });
                await GetfromApi.json();
                Home();
            } catch (error) {
                console.error("Failed to post comment:", error);
            }
        })
    }
    catch (error) {
        console.error(error)
    }
}
