const testimonialPromise = new Promise ((resolve, reject) => {
    const xhr = new XMLHttpRequest()

    xhr.open("GET", "https://api.npoint.io/f08ee6b958392a47f476")

    xhr.onload = function () {

        if (xhr.status === 200) {
            resolve(JSON.parse(xhr.response))
        } else {
            reject("Error loading data!")
        }
    }
    xhr.onerror = function () {
        reject ("Network Error!")
    }
    xhr.send()
})

//=============================================================================================================================================

async function getAllTestimonials () {
    const response = await testimonialPromise
    console.log(response)

    let testimonialHTML = ""
    response.forEach((item) => {
        testimonialHTML += `<div class="w5-card-4" id="testicard">
        <img class="pict" src="${item.image}" alt="kend">
        <div class="texting">
            <p class="quotes">"${item.quote}"</p>
            <p class="authors">- ${item.author} -</p>
            <p class="authors">${item.rating} <i class="fa-solid fa-star"></i></p>
        </div>
    </div>`
        
    })

    document.getElementById("testimonials").innerHTML = testimonialHTML
}

getAllTestimonials()


async function getFilteredTestimonials(rating){
    const responseFilter = await testimonialPromise

    let testimonialHTML = ""

    const testimonialFiltered = responseFilter.filter((item) => {
        return item.rating === rating
    })

    if (testimonialFiltered.length === 0) {
        testimonialHTML = `<h2> Data not found! </h2>`
    } else {
        
        testimonialFiltered.forEach((item) => {
            testimonialHTML = `<div class="w5-card-4" id="testicard">
            <img class="pict" src="${item.image}" alt="kend">
            <div class="texting">
                <p class="quotes">"${item.quote}"</p>
                <p class="authors">- ${item.author} -</p>
                <p class="authors">${item.rating} <i class="fa-solid fa-star"></i></p>
            </div>
        </div>`
        })
    } 
    document.getElementById("testimonials").innerHTML = testimonialHTML
}
