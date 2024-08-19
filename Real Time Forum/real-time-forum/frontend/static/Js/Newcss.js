/**
 * Remove all the css files and set the new css files link gives in arguments
 * 
 * @param {Array.String} arr contain all the link to the css files to set 
 */

export const NewCSS = (arr) => {
    Array.from(document.getElementsByName('css')).forEach(elem => {
        elem.remove()
    })

    arr.forEach(v => {
        const newCssLink = document.createElement('link')
        newCssLink.setAttribute('rel', 'stylesheet')
        newCssLink.setAttribute('type', 'text/css')
        newCssLink.setAttribute('name', 'css')
        newCssLink.setAttribute('href', v)

        document.head.append(newCssLink)
    })
}