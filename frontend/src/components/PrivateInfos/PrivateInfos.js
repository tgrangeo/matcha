import Button from "../Button/Button";
import style from "./style.module.scss";

const PrivateInfos = ({ user, onChange }) => {
  const handleEmailChange = (e) => {
    //fetch api/v1/changeemail (jwt + email)
    //fetch api/v1/changepassword (jwt + email)
    fetch(":8080/api/v1/changeemail", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${localStorage.getItem("jwt")}`,
      },
      body: JSON.stringify({
        email: e.target.email.value,
        password: e.target.password.value,
      }),
    }).then((res) => {
      if (res.status === 200) {
        console.log("you received an email to confirm your new email");
      }
    });
  };

  const handlePasswordChange = (e) => {
    //fetch api/v1/changeemail (jwt + email)
    //fetch api/v1/changepassword (jwt + email)
    fetch(":8080/api/v1/changepassword", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${localStorage.getItem("jwt")}`,
      },
      body: JSON.stringify({
        password: e.target.password.value,
        newpassword: e.target.newpassword.value,
        repassword: e.target.repassword.value,
      }),
    }).then((res) => {
      if (res.status === 200) {
        console.log("you received an email to confirm your new password");
      }
    });
  };

  return (
    <>
      <h2>information personnelle</h2>
      <div className={style.NameAndBirthdate} onSubmit={null}>
        <form className={style.formEmail} onSubmit={handleEmailChange}>
          <div>
            <p>
              <label htmlFor="email">Email</label>
              <input
                type="email"
                id="email"
                name="email"
                required
                defaultValue={user?.email}
                placeholder="Entrer le nouvel email..."
              ></input>
            </p>
            <p>
              <label htmlFor="password">Mot de passe</label>
              <input
                type="password"
                id="password"
                name="password"
                required
                placeholder="Entrer votre mot de passe..."
              ></input>
            </p>

            <Button style="filled" align="left">
              Changer d'email
            </Button>
          </div>
        </form>
        <form className={style.formPassword} onSubmit={handlePasswordChange}>
          <div>
            <p>
              <label htmlFor="password">Nouveau mot de passe</label>
              <input
                type="password"
                id="password"
                name="password"
                required
                placeholder="Entrer votre nouveau mot de passe..."
              ></input>
            </p>
            <p>
              <label htmlFor="repassword">Confirmer mot de passe</label>
              <input
                type="password"
                id="repassword"
                name="repassword"
                required
                placeholder="Confirmer votre nouveau mot de passe..."
              ></input>
            </p>
            <p>
              <label htmlFor="currentpassword">Mot de passe actuel</label>
              <input
                type="password"
                id="currentpassword"
                name="currentpassword"
                required
                placeholder="Entrer votre mot de passe actuel..."
              ></input>
            </p>
            <Button style="filled" align="left">
              Changer de mot de passe
            </Button>
          </div>
        </form>
      </div>
    </>
  );
};

export default PrivateInfos;
