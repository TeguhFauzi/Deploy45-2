
// FUNCTION GETDATA
function getData() {
    let name = document.getElementById("name").value
    let email = document.getElementById("email").value
    let phone = document.getElementById("phone").value
    let subject = document.getElementById("subject").value
    let message = document.getElementById("message").value
    let form = document.getElementById("form")




    // CONDITIONAL
    if (name == "") {
        return alert("name cannot be empty!")
    } else if (email == "") {
        return alert("email cannot be empty!")
    } else if (phone == "") {
        return alert("phone cannot be empty!")
    } else if (subject == "") {
        return alert("subject cannot be empty!")
    } else if (message == "") {
        return alert("message cannot be empty!")
    }

    const emailP = "teguhfauzi55@gmail.com"

    let a = document.createElement('a')
    a.href = `mailto:${emailP}?subject=${subject}&body=Hi! My name is ${name}, ${message} , this is my phone number ${phone}, please call me later.`
    a.click()

    let data = {
        name,
        email,
        phone,
        subject,
        address
    }

    console.log(data)
    form.reset()


}






// window.addEventListener("beforeunload", event =>
// {
//     event.preventDefault();
//     event.returnValue = "";
// });
