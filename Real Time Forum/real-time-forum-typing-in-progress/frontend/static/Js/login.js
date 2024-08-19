import { NewCSS } from "./Newcss.js"
import { register } from "./index.js"
import { Home } from "./Home.js"
import { ErroDis } from "./error.js"


let regdiv = document.getElementsByClassName('container')[0]
//Function For the login page
export async function login() {

    NewCSS(['/static/Css/login.css'])

    //Html of the page
    try {
        regdiv.innerHTML = `
            <div class="login">
                <div class="form_area">
                    <p class="title">SIGN IN</p>
                    <div class="form_group">
                        <label class="sub_title" for="name">Username/Mail</label>
                        <input placeholder="Username/Mail" class="form_style" type="text" id="username" name="username">
                    </div>
                    <div class="form_group">
                        <label class="sub_title" for="password">Password</label>
                        <input placeholder="Enter your password" id="password" class="form_style" type="text"
                            name="password">
                    </div>
                    <div>
                        <button class="btn" id="logbtn" type="submit">LOGIN</button>
                        <p> Don't Have an Account? <a class="link" id="registerbtn" >Sign Up Here!</a></p>
                    </div>
                </div>
            </div>

        `
        //When login button is pressed get value of fileds
        document.getElementById('logbtn').addEventListener('click', async () => {
            let username = document.getElementById('username').value
            let password = document.getElementById('password').value


            let data = {
                username: username,
                password: password
            }
            // Send to back
            let GetfromApi = await fetch('http://localhost:8080/login', {
                method: 'POST',
                body: JSON.stringify(data),
                headers: {
                    'Content-Type': 'application/json'
                }
            }).then(val => val.json())
            //Managing the back response
            if (GetfromApi.PostInf.Logged === true && GetfromApi.Auth.DisErr === false) {
                Home()
            } else if (GetfromApi.PostInf.Logged === false && GetfromApi.Auth.DisErr === true) {
                ErroDis(GetfromApi.Auth.Err)
            }
        })
        //Send to register page
        document.getElementById('registerbtn').addEventListener('click', () => {
            register()
        })

    }
    catch (error) {
        console.error(error)
    }
}

