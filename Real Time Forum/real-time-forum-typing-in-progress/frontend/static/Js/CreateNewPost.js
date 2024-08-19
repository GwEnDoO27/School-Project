import { NewCSS } from "./Newcss.js"
import { fetchCat } from "./Fetching.js"
import { Home } from "./Home.js"

//Managing creation post page
export async function CreatPost() {

    NewCSS(['/static/Css/CreatePost.css'])

    let regdiv = document.getElementsByClassName('container')[0]
    regdiv.classList.add('container')
    //fetch the categories from backend
    let categories = await fetchCat()
    //Html
    regdiv.innerHTML = `
            <div class="cont" id="Title">
                <div class="form_area" id="Title">
                    <p class=" title" id="Title">Title</p>
                    <input placeholder="What is the Title of your post ?" type="text" id="InputTitle" name="Title" class="form_style" required><br>
                </div>
            </div>

            <div class="cont" id="Description">
                <div class="form_area" id="Description">
                    <p class="title" id="Description">Description</p>
                    <input placeholder="Short Description of your post" type="text" id="InputDescription" name="Description" maxlength="250" class="form_style" required><br>
                </div>
            </div>

            <div class="cont" id="Categories">
                <div class="form_area" id="CategoriesForm">
                    <p class="title" id="Categories">Categories</p>
                    <div class="area" id="CategoriesDiv">
                        <input type="text" id="CategoriesInput" name="Categories" placeholder="Enter new category" class="form_style"><br>
                        <select name="category" id="InputCategories" class="form_style"></select required><br>
                    </div>
                </div>
            </div>

            <div class="cont" id="Post">
                <div class="form_area" id="Post">
                    <p class="title" id="Post">Post</p>
                    <textarea id="InputPost" name="Post" maxlength="1000" rows="20" cols="100" class="form_style">
                    </textarea><br>
                    <input type="file" name="Img">
                </div>
            </div>

            <div class="cont" id="post-button">
                <div class="form_area" id="post-button">
                    <p class="title" id="post-button">Post</p>
                    <!-- <label class="sub_title" id="post-button" for="name">Name</label> -->
                    <button type="submit" class="form_style" id="post-button">Post</button>
                </div>
            </div>
    `

    //Create the select with all the categories
    let Catdiv = document.getElementById('InputCategories')
    categories.forEach(element => {
        console.log(element.categories)
        let selectValue = document.createElement('option')
        selectValue.innerHTML = `
            <option value=${element.Categories}>${element.Categories}<option>
        `
        Catdiv.appendChild(selectValue)
    });

    //Sending the values
    document.getElementById('post-button').addEventListener('click', async () => {
        let Title = document.getElementById('InputTitle').value
        let Description = document.getElementById('InputDescription').value
        let Categories = document.getElementById('CategoriesInput').value
        let Post = document.getElementById('InputPost').value
        let select = document.getElementById('InputCategories').value

        let data = {
            Title: Title,
            Description: Description,
            Categories: Categories,
            Post: Post,
            SelectValue: select
        }

        let GetfromApi = await fetch('http://localhost:8080/newpost', {
            method: 'POST',
            body: JSON.stringify(data),
            headers: {
                'Content-Type': 'application/json'
            }
        }).then(val => val.json())
        //if(GetfromApi.CreatePost.Err != ""){
        //} 
        Home()
    })


}