import Button from "../Button/Button";
import Input from "../Input/Input";
import style from "./style.module.scss";

const PrivateInfos = ({ user, onChange }) => {
  const handleEmailChange = (e) => {
    //fetch api/v1/changeemail (jwt + email)
    //fetch api/v1/changepassword (jwt + email)
    fetch("/api/v1/changeemail", {
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
    fetch("/api/v1/changepassword", {
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
            <Input
              type="email"
              name="email"
              required
              classNames={style.input}
              defaultValue={user?.email}
              placeholder="Entrer le nouvel email"
            >
              Email
            </Input>
            <Input
              type="password"
              name="password"
              classNames={style.input}
              required
              placeholder="Entrer le mot de passe actuel"
            >
              Mot de passe
            </Input>

            <Button style="filled" align="left">
              Changer d'email
            </Button>
          </div>
        </form>
        <form className={style.formPassword} onSubmit={handlePasswordChange}>
          <div>
            <Input
              type="password"
              name="password"
              classNames={style.input}
              required
              placeholder="Entrer votre nouveau mot de passe."
            >
              Mot de passe
            </Input>
            <Input
              type="password"
              name="repassword"
              classNames={style.input}
              required
              placeholder="Confirmer votre nouveau mot de passe."
            >
              Mot de passe
            </Input>
            <Input
              type="password"
              name="currentpassword"
              classNames={style.input}
              required
              placeholder="Entrer votre mot de passe actuel."
            >
              Mot de passe
            </Input>
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
