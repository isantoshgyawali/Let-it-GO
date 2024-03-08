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
        
        // Anchor tag for user details
        const userLink = document.createElement("a");
        const userDetails = document.createElement("p");
        userDetails.textContent = `${user.id}. ${user.name}`;
        userLink.appendChild(userDetails);
        userlist.appendChild(userLink);

        // delete button for specific users
        const deleteButton = document.createElement("button");
        deleteButton.textContent = "Delete";
        deleteButton.classList.add("delete-btn");
        deleteButton.classList.add("delete-user");
        userlist.appendChild(deleteButton);

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

