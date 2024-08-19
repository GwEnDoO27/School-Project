
//Display error
export function ErroDis(err) {
    const regdiv = document.getElementsByClassName('container')[0]
    let errorDiv = document.createElement('div')

    errorDiv.id = 'erroDiv'
    errorDiv.classList.add('erroDiv')
    errorDiv.innerHTML = `
    <div class="form_area" id="error">
    <p class="title">${err}</p>
    </div>
    `
    regdiv.appendChild(errorDiv)
}