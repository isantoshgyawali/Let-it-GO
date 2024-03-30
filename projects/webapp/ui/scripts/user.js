const baseUrl = "http://localhost:6969";

const goToUser = () => {
  window.location.href = "/user/";
};

const fetchUserData = (fetchUrl) => {
  return axios
    .get(fetchUrl)
    .then((response) => {
      const users = response.data;
      return users;
    })
    .catch((error) => {
      console.error("There was an error fetching user data", error);
    });
};

const deleteUserData = (id) => {
  axios
    .delete(`${baseUrl}/user/${id}/`)
    .then(() => {
      alert("Data was successfully deleted");
      // remove the deleted item from the ui
    })
    .catch((error) => {
      console.log(error);
      alert("Error occurred deleting the user data");
    });
};

const deleteAllUser = () => {
  axios
    .delete(`${baseUrl}/user/`)
    .then(() => {
      alert("Data was successfully deleted");
    })
    .catch((error) => {
      console.log(error);
      alert("Error occurred deleting the all user's data");
    });
};

const listUserData = async () => {
  try {
    let userDataItems = await fetchUserData(`${baseUrl}/user/`);
    console.log(userDataItems);

    const UserListUl = document.getElementsByClassName("user-data")[0];
    const messageValue = document.getElementById("hidden-message").value;
    const ListTitle = document.getElementsByClassName("user-list-title")[0];
    const deleteAll = document.getElementsByClassName("delete-all")[0];

    if (Array.isArray(userDataItems) && userDataItems.length > 0) {
      ListTitle.textContent = messageValue;

      userDataItems.forEach((user, index) => {
        const userlist = document.createElement("li");
        userlist.classList.add("user");
        userlist.id = `user-${index+1}`; // Assigning generated ID to list items

        //-- Anchor tag for user details
        const userLink = document.createElement("a");
        const userDetails = document.createElement("p");
        userDetails.textContent = `${index+1}. ${user.name}`;
        userLink.appendChild(userDetails);
        userlist.appendChild(userLink);

        //-- if there is the click on the any of the specific user
        //-- then fire the function userRedirect
        userLink.addEventListener("click", () => {
          console.log(user);
          userRedirect(user.id);
        });

        //-- delete specific users on click
        const deleteButton = document.createElement("button");
        deleteButton.textContent = "Delete";
        deleteButton.classList.add("delete-btn");
        deleteButton.classList.add("delete-user");
        userlist.appendChild(deleteButton);

        //-- if the button is clicked then the user data from the database
        //-- is cleared and the data is then remove the ui too
        deleteButton.addEventListener("click", () => {
          deleteUserData(user.id);
        });

        deleteAll.addEventListener("click", () => {
          deleteAllUser();
        })

        UserListUl.appendChild(userlist);
      });
    } else {
      ListTitle.textContent = "No data is available";
      deleteAll.style.display = "none";
      UserListUl.style.display = "none";
      UserListUl.style.display = "none";
    }
  } catch (err) {
    console.log("Error occurred listing the user data items", err);
  }
};

//-- just to ensure no need to fetch always the window load
if (window.location.pathname === "/user/") {
  listUserData();
}

const userRedirect = (id) => {
  window.location.href = `/user/${id}/`;
};
