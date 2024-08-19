import { NewCSS } from "./Newcss.js"
import { Home } from "./Home.js"

//Managin the posts 
export async function MoreInfos() {
    NewCSS(['/static/Css/informations.css'])

    const regdiv = document.getElementsByClassName('container')[0]
    regdiv.classList.add('container')
    //html
    regdiv.innerHTML = `
        <div class="cont" id="Name">
            <div class="form_area" id="Name">
                <p class=" title" id="Name">Name</p>
                <!-- <label class=" sub_title" id="Title" for="name">Name</label> -->
                <input placeholder="Your Name ?" type="text" id="InputName" name="Name" class="form_style"required><br>
            </div>
        </div>

        <div class="cont" id="LastName">
            <div class="form_area" id="LastName">
                <p class=" title" id="LastName">LastName</p>
                <!-- <label class=" sub_title" id="Title" for="name">Name</label> -->
                <input placeholder="Your Last Name ?" type="text" id="InputLastName" name="LastName" class="form_style"required><br>
            </div>
        </div>

        <div class="cont" id="Birthday">
            <div class="form_area" id="Birthday">
                <p class=" title" id="Birthday">Birthday</p>
                <!-- <label class=" sub_title" id="Title" for="name">Name</label> -->
                <input placeholder="Birthday date ?" type="date" id="InputBirthday" name="Birthday" class="form_style" required><br>
            </div>
        </div>

        <div class="cont" id="Gender">
            <div class="form_area" id="Gender">
                <p class="title" id="Gender">Gender</p>
                <select name="Gender" id="SelectGender" class="form_style" id="Gender">
                    <option name="Gender">Man</option>
                    <option name="Gender">Woman</option>
                    <option name="Gender">Helicoptere de Combat</option>
                </select>
            </div>
        </div>


        <div class="cont" id="Town">
            <div class="form_area" id="Town">
                <p class="title" id="Town">Town</p>
                <!-- <label class=" sub_title" id="Title" for="name">Name</label> -->
                <input placeholder="Your Town ?" type="text" id="InputTown" name="Town" class="form_style" required><br>
            </div>
        </div>

        <div class="cont" id="Country">
            <div class="form_area" id="Country">
                <p class="title" id="Country">Country</p>
                <!-- <label class=" sub_title" id="Title" for="name">Name</label> -->
                <input placeholder="Country ?" type="text" id="InputCountry" name="Country" class="form_style" required><br>
            </div>
        </div>
        <div class="cont" id="post-button">
            <div class="form_area" id="post-button">
                <p class="title" id="post-button">Finish the register</p>
                <button class="form_style" id="post-button">Finish the register</button>
            </div>
        </div>
    `

    //Sending  the data
    document.getElementById('post-button').addEventListener('click', async () => {
        let name = document.getElementById('InputName').value
        let Lastname = document.getElementById('InputLastName').value
        let Birthday = document.getElementById('InputBirthday').value
        let Gender = document.getElementById('SelectGender').value
        let Town = document.getElementById('InputTown').value
        let Country = document.getElementById('InputCountry').value

        let data = {
            name: name,
            Lastname: Lastname,
            Birthday: Birthday,
            Gender: Gender,
            Town: Town,
            Country: Country
        }

        //console.log(data)
        let GetfromApi = await fetch('http://localhost:8080/infos', {
            method: 'POST',
            body: JSON.stringify(data),
            headers: {
                'Content-Type': 'application/json'
            }
        }).then(val => val.json())
        Home()
    })

}