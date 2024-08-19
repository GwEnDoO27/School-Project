import { NewCSS } from "./Newcss.js"
import { login } from "./login.js"
import { MoreInfos } from "./account.js"
import { ErroDis } from "./error.js"


document.getElementById('nav').style.display = 'none'

//Manage the register page
export async function register() {
    NewCSS(['/static/Css/register.css'])
    try {
        const regdiv = document.getElementsByClassName('container')[0]
        regdiv.classList.add('container')
        regdiv.innerHTML = `
        <div class="login">
                <div class="form_area">
                    <p class="title">SIGN UP</p>
                    
                        <div class="form_group">
                            <label class="sub_title" for="name">Username</label>
                            <input placeholder="Enter your Username" class="form_style" type="text" id="username"
                                name="username">
                        </div>
                        <div class="form_group">
                            <label class="sub_title" for="email">Email</label>
                            <input placeholder="Enter your email" id="Email" class="form_style" type="email" name="email">
                        </div>
                        <div class="form_group">
                            <label class="sub_title" for="password">Password</label>
                            <input placeholder="Enter your password" id="password" class="form_style" type="password"
                                name="password">
                        </div>
                        <div>
                            <button id="regbtn" class="btn">SIGN UP</button>
                            <p>Have an Account? <a class="link" id="loginf">Login Here!</a></p>
                        </div>
                </div>
            </div>
        </div>
        `
        //Send the data
        document.getElementById('regbtn').addEventListener('click', async () => {
            let username = document.getElementById('username').value
            let email = document.getElementById('Email').value
            let password = document.getElementById('password').value


            let data = {
                username: username,
                email: email,
                password: password
            }
            let GetfromApi = await fetch('http://localhost:8080/register', {
                method: 'POST',
                body: JSON.stringify(data),
                headers: {
                    'Content-Type': 'application/json'
                }
            }).then(val => val.json())
            //Error managing 
            if (GetfromApi.PostInf.Logged === true && GetfromApi.Auth.DisErr === false) {
                MoreInfos()
            } else if (GetfromApi.PostInf.Logged === false && GetfromApi.Auth.DisErr === true) {

                ErroDis(GetfromApi.Auth.Err)
            }
        })
        //send to login page
        document.getElementById('loginf').addEventListener('click', () => {
            login()
        })

    }
    catch (error) {
        console.error(error)
    }
}


register()