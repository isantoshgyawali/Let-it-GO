
const submitForm = () => {
    const form = document.querySelector("#mainForm")
    form.submit()

    const formData = getData();
    console.log(formData)
}

const getData = () => {
        const name = document.getElementById("name").value
        const address= document.getElementById("address").value
        const email= document.getElementById("email").value

        return {name,address,email}
}

const putData = () => {
    
}