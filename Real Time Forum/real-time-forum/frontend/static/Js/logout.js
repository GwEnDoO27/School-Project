import { deleteSession, deleteCookie } from "./Cookie.js";

//delete coockie and sensd to back infos for totally disconect
export async function Logout() {
    await deleteSession()
    deleteCookie('session-token')
    window.location.reload()
}