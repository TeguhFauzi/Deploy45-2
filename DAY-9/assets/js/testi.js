// class Testimonial {

//     constructor(quotes, images) {
//         this._quotes = quotes
//         this._images = images
//     }

//     get quote() {
//         return this._quotes
//     }

//     get image() {
//         return this._images
//     }

//     get html() {
//         return `<div class="w5-card-4" id="testicard">
//         <img class="pict" src="${this.image}" alt="kend">
//         <div class="texting">
//             <p class="quotes">"${this.quote}"</p>
//             <p class="authors">- ${this.author}</p>
//         </div>
//     </div>`
//     }

// }

// class AuthorTestimonial extends Testimonial {

//     constructor(authors, quotes, images) {
//         super(quotes, images)
//         this._authors = authors
//     }

//     get author() {
//         return this._authors
//     }

// }

// class CompanyTestimonial extends Testimonial {

//     constructor(company, quotes, images) {
//         super(quotes, images)
//         this._company = company
//     }

//     get author() {
//         return this._company + " Company"
//     }

// }


// const testimonial1 = new AuthorTestimonial("Jamesy Ed", "Mantap Slurr", "https://i.pinimg.com/originals/13/da/27/13da27e146f78880980458486586ea0c.jpg")

// const testimonial2 = new AuthorTestimonial("Michael Ipin", "Onichan!!!", "https://yt3.googleusercontent.com/ifZJFKoXaasijKlvMxM43d1lXCsvTjJgujLvEWCvv68e8rHKSg2OmixwdRBp_L_50uQBviQOHA=s900-c-k-c0x00ffffff-no-rj")

// const testimonial3 = new CompanyTestimonial("Lato-lato", "Terima Kasih bang", "https://doktersehat.com/wp-content/uploads/2023/01/gambar-manfaat-lato-lato-doktersehat-800x534.jpg")

// let testimonialData = [testimonial1, testimonial2, testimonial3]

// let testimonialHTML = ""


// // for (let i = 0; i < testimonialData.length; i++) {
// //     document.getElementById("testimonials").innerHTML += testimonialData[i].html 
// // }

// for ( let i = 0; i < testimonialData.length; i++){
//     testimonialHTML += testimonialData[i].html
// }
// document.getElementById("testimonials").innerHTML= testimonialHTML


const testimonialData = [

    {
        author : "Jamesy Ed",
        quote : "Mantap Slurr",
        image : "https://i.pinimg.com/originals/13/da/27/13da27e146f78880980458486586ea0c.jpg",
        rating : 4
    },


    {
        author : "Michael Ipin",
        quote : "Onichan!!!",
        image : "https://yt3.googleusercontent.com/ifZJFKoXaasijKlvMxM43d1lXCsvTjJgujLvEWCvv68e8rHKSg2OmixwdRBp_L_50uQBviQOHA=s900-c-k-c0x00ffffff-no-rj",
        rating : 2
    },


    {
        author : "Lato-lato",
        quote : "Terima Kasih bang",
        image : "https://doktersehat.com/wp-content/uploads/2023/01/gambar-manfaat-lato-lato-doktersehat-800x534.jpg",
        rating : 5
    },


    {
        author : "Jordi Baik",
        quote : "Mantul bang saya suka",
        image : "https://m.media-amazon.com/images/M/MV5BYTU5ZWJhZDctNGZlOS00NjRjLWFiNGMtNDU0NmIwNTM1Y2ViXkEyXkFqcGdeQXVyNDUzOTQ5MjY@._V1_.jpg",
        rating : 5
    },

    
    {
        author : "Bernard beard",
        quote : "Gaje ah bang, Gabut yaa",
        image : "https://i.imgflip.com/52l2rx.png",
        rating : 1
    }

]

function allTestimonials() {
    let testimonialHTML = " "

    testimonialData.forEach((item) => {
        testimonialHTML += `<div class="w5-card-4" id="testicard">
                 <img class="pict" src="${item.image}" alt="kend">
                 <div class="texting">
                     <p class="quotes">"${item.quote}"</p>
                     <p class="authors">- ${item.author}</p>
                     <p class="authors">${item.rating} <i class="fa-solid fa-star"></i></p>
                 </div>
             </div>`
    })

    document.getElementById("testimonials").innerHTML = testimonialHTML

}

allTestimonials()


function filterTestimonials(rating) {
    let testimonialHTML = "";

    // rating : 1

    const testimonialFiltered = testimonialData.filter((item) => {
        return item.rating === rating
    })

    // [

    //     {
    //         author : "User vpn",
    //         quote : "501",
    //         image : "https://i.pinimg.com/originals/97/2f/1b/972f1b8aca65479e3c401b800a4bd76a.jpg",
    //         rating : 2
    //     }

    // ]

    if (testimonialFiltered.length === 0){
        testimonialHTML = `<h2> Data not found! </h2>`
    } else{
        testimonialFiltered.forEach((item) => {
            testimonialHTML += `<div class="w5-card-4" id="testicard">
            <img class="pict" src="${item.image}" alt="kend">
            <div class="texting">
                <p class="quotes">"${item.quote}"</p>
                <p class="authors">- ${item.author}</p>
                <p class="authors">${item.rating} <i class="fa-solid fa-star"></i></p>
            </div>
        </div>`
        })

    }


    document.getElementById("testimonials").innerHTML = testimonialHTML
}


