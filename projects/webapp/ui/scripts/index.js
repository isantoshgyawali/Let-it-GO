document.querySelector("#mainForm").addEventListener("submit", (e) => {
  e.preventDefault();
  const {type, name, address, email} = getFormData();

  let typeLC = type.toLowerCase();

  axios
    .post(`/${typeLC}/`, {
      type: type,
      name: name,
      address: address,
      email: email,
    })
    .then((response) => {
      console.log(response);
    });
  document.querySelector("#mainForm").reset();
});

const getFormData = () => {
  const type = document.getElementById("chooseOptions").value;
  const name = document.getElementById("name").value;
  const address = document.getElementById("address").value;
  const email = document.getElementById("email").value;
  return { type, name, address, email };
};

const goToUser = () => {
  window.location.href = "/user/";
};

const goToOrg = () => {
  window.location.href = "/org/";
};