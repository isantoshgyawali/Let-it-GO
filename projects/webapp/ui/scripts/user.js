const goToUser = () => {
  window.location.href = "/user/"
}

//-- just to ensure no need to fetch always the window load
if (window.location.pathname === "/user/") {
    axios.get("http://localhost:8081/user/")
    .then(response => {
        const users = response.data
        console.log(users)
    })
    .catch(error => {
      console.error("There was an error fetching user data", error)
    })
}

