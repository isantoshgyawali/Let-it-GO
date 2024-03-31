const baseUrl = "http://localhost:6969";

const fetchOrgData = (fetchUrl) => {
  return axios
    .get(fetchUrl)
    .then((response) => {
      const orgs = response.data;
      return orgs;
    })
    .catch((error) => {
      console.error("There was an error fetching org data", error);
    });
};

const deleteOrgData = (id) => {
  axios
    .delete(`${baseUrl}/org/${id}/`)
    .then(() => {
      alert("Data was successfully deleted");
    })
    .catch((error) => {
      console.log(error);
      alert("Error occurred deleting the org data");
    });
};

const deleteAllOrg = () => {
  axios
    .delete(`${baseUrl}/org/`)
    .then(() => {
      alert("Data was successfully deleted");
    })
    .catch((error) => {
      console.log(error);
      alert("Error occurred deleting the all org data");
    });
};

const listOrgData = async () => {
  try {
    let orgDataItems = await fetchOrgData(`${baseUrl}/org/`);
    console.log(orgDataItems);

    const OrgListUl = document.getElementsByClassName("org-data")[0];
    const messageValue = document.getElementById("hidden-message").value;
    const ListTitle = document.getElementsByClassName("org-list-title")[0];
    const deleteAll = document.getElementsByClassName("delete-all")[0];

    if (Array.isArray(orgDataItems) && orgDataItems.length > 0) {
      ListTitle.textContent = messageValue;

      orgDataItems.forEach((org, index) => {
        const orgList = document.createElement("li");
        orgList.classList.add("org");
        orgList.id = `org-${index+1}`; // Assigning generated ID to list items

        //-- Anchor tag for orgs details
        const orglink = document.createElement("a");
        const orgDetails = document.createElement("p");
        orgDetails.textContent = `${index+1}. ${org.name}`;
        orglink.appendChild(orgDetails);
        orgList.appendChild(orglink);

        //-- if there is the click on the any of the specific org
        //-- then fire the function orgRedirect
        orglink.addEventListener("click", () => {
          orgRedirect(org.id);
        });

        //-- delete specific org on click
        const deleteButton = document.createElement("button");
        deleteButton.textContent = "Delete";
        deleteButton.classList.add("delete-btn");
        deleteButton.classList.add("delete-org");
        orgList.appendChild(deleteButton);

        //-- if the button is clicked then the org data from the database
        //-- is cleared and the data is then remove the ui too
        deleteButton.addEventListener("click", () => {
          deleteOrgData(org.id);
        });

        deleteAll.addEventListener("click", () => {
          deleteAllOrg();
        })

        OrgListUl.appendChild(orgList);
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

//-- just to ensure no need to fetch always the window load
if (window.location.pathname === "/org/") {
  listOrgData();
}

const orgRedirect = (id) => {
  window.location.href = `/org/${id}/`;
};