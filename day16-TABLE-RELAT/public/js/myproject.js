
// RESPONSIVE NAV
let barsopen = false


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






let datas = []

function getData(event) {
    event.preventDefault()

    let title = document.getElementById("title").value
    let start = new Date(document.getElementById("start").value)
    let end = new Date(document.getElementById("end").value)
    let description = document.getElementById("description").value
    let form = document.getElementById("form")


    // GET IMAGE IN FAKEPATH 
    let image = document.getElementById("image").files


    // GET CHECKBOX VALUE
    let java = document.getElementById("java").checked ? document.getElementById("java").value : false;
    let javascript = document.getElementById("javascript").checked ? document.getElementById("javascript").value : false;
    let react = document.getElementById("react").checked ? document.getElementById("react").value : false;
    let node = document.getElementById("node").checked ? document.getElementById("node").value : false;





    // VALIDATION
    if (title == "") {
        alert("Title cannot be empty!")
    } else if (start == "") {
        alert("Date cannot be empty!")
    } else if (end == "") {
        alert("Date cannot be empty!")
    } else if (description == "") {
        alert("Description cannot be empty!")
    } else if (image == "") {
        alert("Image cannot be empty!")
    } else if (java == false && javascript == false && react == false && node == false) {
        alert("Technologies cannot be empty!")
    } else {



        // GET IMAGE BY CONVERT URL FROM FAKEPATH 
        image = URL.createObjectURL(image[0])



        // HOW TO COMPARE TECHNOLOGIES BY VALUE
        java = java != false ? `<i class="fa-brands fa-java"></i>` : "";
        javascript = javascript != false ? `<i class="fa-brands fa-js"></i>` : "";
        react = react != false ? `<i class="fa-brands fa-react"></i>` : "";
        node = node != false ? `<i class="fa-brands fa-node"></i>` : "";



        // GET DURATION
        let duration = end.getTime() - start.getTime()
        if (duration < 0) {
            alert("Invalid time input!")
        } else {
            let yeardur = Math.floor(duration / (12 * 30 * 24 * 60 * 60 * 1000))
            let monthdur = Math.floor(duration / (30 * 24 * 60 * 60 * 1000))
            let weekdur = Math.floor(duration / (7 * 24 * 60 * 60 * 1000))
            let daydur = Math.floor(duration / (24 * 60 * 60 * 1000))

            if (yeardur > 0) {
                duration = yeardur + " year(s)"
            } else if (monthdur > 0) {
                duration = monthdur + " month(s)"
            } else if (weekdur > 0) {
                duration = weekdur + " week(s)"
            } else if (daydur > 0) {
                duration = daydur + " day(s)"
            }
        }





        // CREATE BLOG OBJECT
        let data = {
            title,
            description,
            image,
            java,
            javascript,
            react,
            node,
            duration,
            posted: new Date(),
        }



        form.reset()
        datas.push(data)
        console.log(datas)
        showData()

    }



}






// function manipulation time.get
function createTime(time) {
    // declaration variable
    let years = time.getFullYear()
    let monthIndex = time.getMonth()
    let date = time.getDate()
    let hour = time.getHours()
    let minutes = time.getMinutes()

    const month = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"]

    return `${date} ${month[monthIndex]} ${years} ${hour}:${minutes} WIB`
}





const showData = () => {
    document.getElementById("content").innerHTML =

        `<div class="card-content">
        <div class="contentblog">
        <img class="blog-img" src="./assets/images/blog.png" alt="blog-image">
        <div class="content-text">
            <a href="project-detail.html">"Pasar Coding di Indonesia dinilai..."</a>
            <h6>28 November 2022 02:30 WIB || Author : Teguh Fauzi</h6>
            <h5>Duration : 3 month(s).</h5>
            <p>Lorem ipsum dolor sit amet consectetur, adipisicing elit. Quas impedit provident quia
            voluptatum sit adipisci, modi voluptatem. Minima, quide.</p>
            <h5>3 week ago(s).</h5>
            <div class="icon">
                <i class="fa-brands fa-java"></i>
                <i class="fa-brands fa-js"></i>
                <i class="fa-brands fa-react"></i>
                <i class="fa-brands fa-node"></i>
            </div>
            <div class="button">
                <div class="edit">
                    <button>edit</button>
                </div>
                <div class="del">
                    <button>delete</button>
                </div>
            </div>
        </div>

    </div>
</div>`




    for (let i = 0; i <= datas.length; i++) {
        document.getElementById("content").innerHTML +=
            `<div class="card-content">
        <div class="contentblog">
            <img class="blog-img" src="${datas[i].image}" alt="blog-image">
            <div class="content-text">
                <a href="project-detail.html">"${datas[i].title}..."</a>
                <h6>${createTime(datas[i].posted)} || Author : Teguh Fauzi</h6>
                <h5>Duration :${datas[i].duration} .</h5>
                <p>${datas[i].description}.</p>
                <h5>${between(datas[i].posted)}</h5>
                <div class="icon">
                    ${datas != false ? datas[i].java : ""}
                    ${datas != false ? datas[i].javascript : ""}
                    ${datas != false ? datas[i].react : ""}
                    ${datas != false ? datas[i].node : ""}
                </div>
                <div class="button">
                    <div class="edit">
                        <button>edit</button>
                    </div>
                    <div class="del">
                        <button>delete</button>
                    </div>
                </div>
            </div>

        </div>
    </div>`


    }


}






// MANIPULATION TIME POSTED AT 
const between = (timePost) => {
    let timePosting = timePost
    let timeNow = new Date()
    let agos = timeNow - timePosting
    let yearBetween = Math.floor(agos / (12 * 30 * 24 * 60 * 60 * 1000))
    let monthBetween = Math.floor(agos / (30 * 24 * 60 * 60 * 1000))
    let weekBeetween = Math.floor(agos / (7 * 24 * 60 * 60 * 1000))
    let dayBetween = Math.floor(agos / (24 * 60 * 60 * 1000))
    let hourBetween = Math.floor(agos / (60 * 60 * 1000))
    let minuteBetween = Math.floor(agos / (60 * 1000))
    let secondBetween = Math.floor(agos / 1000)


    if (yearBetween > 0) {
        return yearBetween + " year ago(s)."
    } else if (monthBetween > 0) {
        return monthBetween + " month ago(s)."
    } else if (weekBeetween > 0) {
        return weekBeetween + " week ago(s)."
    } else if (dayBetween > 0) {
        return dayBetween + " day ago(s)."
    } else if (hourBetween > 0) {
        return hourBetween + " hour ago(s)."
    } else if (minuteBetween > 0) {
        return minuteBetween + " minute ago(s)."
    } else if (secondBetween > 0) {
        return secondBetween + " second ago(s)."
    }


}






setInterval(showData, 1000)






// DELETE
const blog = document.querySelector('#content')
const button = document.createElement('button')
button.textContent = 'delete'


blog.addEventListener('click', (event) => {
    if (event.target.tagName === 'BUTTON') {
        const button = event.target
        const contentblog = button.parentNode.parentNode.parentNode
        const cardcontent = contentblog.parentNode
        const blog = cardcontent.parentNode


        if (button.textContent === 'delete') {
            blog.removeChild(cardcontent)
        }

    }

})








