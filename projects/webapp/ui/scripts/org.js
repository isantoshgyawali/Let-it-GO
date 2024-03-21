const goToOrg = () => {
  window.location.href = "/org/";
};

const fetchOrgData = (fetchUrl) => {
  return axios
    .get(fetchUrl)
    .then((response) => {
      const org = response.data;
      return org;
    })
    .catch((error) => {
      console.error("There was an error fetching org data", error);
    });
};

const listOrgData = async () => {
  try {
    let orgDataItems = await fetchOrgData("http://localhost:6969/org/");
    console.log(orgDataItems);

    const OrgListUl = document.getElementsByClassName("org-data")[0];
    const messageValue = document.getElementById("hidden-message").value;
    const ListTitle = document.getElementsByClassName("org-list-title")[0];
    const deleteAll = document.getElementsByClassName("delete-all")[0];

    if (Array.isArray(orgDataItems) && orgDataItems.length > 0) {
      ListTitle.textContent = messageValue;

      orgDataItems.forEach((org) => {
        const orglist = document.createElement("li");
        orglist.classList.add("org");

        // Anchor tag for org details
        const orglink = document.createElement("a");
        const orgDetails = document.createElement("p");
        orgDetails.textContent = `${org.id}. ${org.name}`;
        orglink.appendChild(orgDetails);
        orglist.appendChild(orglink);

        //-- if there is the click on the any of the specific org
        //-- then fire the function orgRedirect
        orglink.addEventListener("click", () => {
          console.log(org);
          orgRedirect(org.id);
        });

        // delete button for specific org
        const deleteButton = document.createElement("button");
        deleteButton.textContent = "Delete";
        deleteButton.classList.add("delete-btn");
        deleteButton.classList.add("delete-org");
        orglist.appendChild(deleteButton);

        //-- if the button is clicked then the user data from the database
        //-- is cleared and the data is then remove the ui too
        deleteButton.addEventListener("click", () => {
          deleteOrg(org.id);
        });

        OrgListUl.appendChild(orglist);
      });
    } else {
      ListTitle.textContent = "No data is available";
      deleteAll.style.display = "none";
      OrgListUl.style.display = "none";
      OrgListUl.style.display = "none";
    }
  } catch (err) {
    console.log("Error occurred listing the org data items", err);
  }
};

if (window.location.pathname === "/org/") {
  listOrgData();
}

const orgRedirect = (id) => {
  window.location.href = `/org/${id}/`;
};

const deleteOrg = (id) => {
  console.log("delete");
};
