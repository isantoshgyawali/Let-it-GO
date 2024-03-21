const goToUser = () => {
  window.location.href = "/user/"
}

const fetchUserData = (fetchUrl) => {
 return axios.get(fetchUrl)
    .then((response) => {
      const users = response.data;
      return users;
    })
    .catch((error) => {
      console.error("There was an error fetching user data", error);
    });
}

const listUserData = async () => {
  try {
    let userDataItems = await fetchUserData("http://localhost:6969/user/");
    console.log(userDataItems);

    const UserListUl = document.getElementsByClassName("user-data")[0];
    const messageValue = document.getElementById("hidden-message").value;
    const ListTitle = document.getElementsByClassName("user-list-title")[0];
    const deleteAll = document.getElementsByClassName("delete-all")[0];

    if (Array.isArray(userDataItems) && userDataItems.length > 0) {
      ListTitle.textContent = messageValue;

      userDataItems.forEach((user) => {
        const userlist = document.createElement("li");
        userlist.classList.add("user")
        
        //-- Anchor tag for user details
        const userLink = document.createElement("a");
        const userDetails = document.createElement("p");
        userDetails.textContent = `${user.id}. ${user.name}`;
        userLink.appendChild(userDetails);
        userlist.appendChild(userLink);
        
        //-- if there is the click on the any of the specific user 
        //-- then fire the function userRedirect
        userLink.addEventListener("click", () => {
          console.log(user)
          userRedirect(user.id)
        })

        //-- delete specific users on click
        const deleteButton = document.createElement("button");
        deleteButton.textContent = "Delete";
        deleteButton.classList.add("delete-btn");
        deleteButton.classList.add("delete-user");
        userlist.appendChild(deleteButton);

        //-- if the button is clicked then the user data from the database
        //-- is cleared and the data is then remove the ui too
        deleteButton.addEventListener("click", () => {
          deleteUser(user.id)
        })

        UserListUl.appendChild(userlist);
      });
    } else {
      ListTitle.textContent = "No data is available";
      deleteAll.style.display = "none"; 
      UserListUl.style.display= "none"
      UserListUl.style.display= "none"
    }

  } catch (err) {
    console.log("Error occurred listing the user data items", err);
  }
};

//-- just to ensure no need to fetch always the window load
if (window.location.pathname === "/user/") {
  listUserData();
}

const userRedirect= (id) => {
  window.location.href = `/user/${id}/`
}

const deleteUser = (id) => {
  console.log("delete")
}