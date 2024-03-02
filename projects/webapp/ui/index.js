
const submitForm = () => {
    const form = document.querySelector("#mainForm")
    form.submit()

    const formData = getFormData();
    console.log(formData)

}

const getFormData = () => {
    const name = document.getElementById("name").value
    const address = document.getElementById("address").value
    const email = document.getElementById("email").value

    return { name, address, email }
}

const putData = () => {

}

const goToUser = () => {
    window.location.href = "users-data.html"

    const userElement = document.getElementsByClassName("user-data")
    fetch('http://localhost:8081/user/')
        .then(response => {
            if (!response.ok) {
                throw new Error("Network Response was not ok")
            } else {
                return response.json()
            }
        })
        .then(data => {
            userElement.innerHTML = ""; // Clear existing content
            data.forEach(user => {
              const listItem = document.createElement("li");
              listItem.textContent = user.name; 
              userElement.appendChild(listItem);
            });
        })
        .catch(error => {
            console.error('There was some problem while fetching : ', error)
        });

        window.onload = goToUser
}

