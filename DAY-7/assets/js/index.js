let barsopen = true


function buttonbar() {
    let floatbar = document.getElementById('floatbar')
    let openbutt = document.getElementById('openbutt')
    let closebutt = document.getElementById('closebutt')

    if (barsopen) {
        floatbar.style.display = 'none'
        openbutt.style.display = 'block'
        closebutt.style.display = 'none'
        barsopen = false
    } else {
        floatbar.style.display = 'block'
        closebutt.style.display = 'block'
        openbutt.style.display = 'none'
        barsopen = true

    }
 

}