class Testimonial {

    constructor(quotes, images) {
        this._quotes = quotes
        this._images = images
    }

    get quote() {
        return this._quotes
    }

    get image() {
        return this._images
    }

    get html() {
        return `<div class="w5-card-4" id="testicard">
        <img class="pict" src="${this.image}" alt="kend">
        <div class="texting">
            <p class="quotes">"${this.quote}"</p>
            <p class="authors">- ${this.author}</p>
        </div>
    </div>`
    }

}

class AuthorTestimonial extends Testimonial {

    constructor(authors, quotes, images) {
        super(quotes, images)
        this._authors = authors
    }

    get author() {
        return this._authors
    }

}

class CompanyTestimonial extends Testimonial {

    constructor(company, quotes, images) {
        super(quotes, images)
        this._company = company
    }

    get author() {
        return this._company + " Company"
    }

}


const testimonial1 = new AuthorTestimonial("Jamesy Ed", "Mantap Slurr", "https://i.pinimg.com/originals/13/da/27/13da27e146f78880980458486586ea0c.jpg")

const testimonial2 = new AuthorTestimonial("Michael Ipin", "Onichan!!!", "https://yt3.googleusercontent.com/ifZJFKoXaasijKlvMxM43d1lXCsvTjJgujLvEWCvv68e8rHKSg2OmixwdRBp_L_50uQBviQOHA=s900-c-k-c0x00ffffff-no-rj")

const testimonial3 = new CompanyTestimonial("Lato-lato", "Terima Kasih bang", "https://doktersehat.com/wp-content/uploads/2023/01/gambar-manfaat-lato-lato-doktersehat-800x534.jpg")

let testimonialData = [testimonial1, testimonial2, testimonial3]

let testimonialHTML = ""


// for (let i = 0; i < testimonialData.length; i++) {
//     document.getElementById("testimonials").innerHTML += testimonialData[i].html 
// }

for ( let i = 0; i < testimonialData.length; i++){
    testimonialHTML += testimonialData[i].html
}
document.getElementById("testimonials").innerHTML= testimonialHTML
