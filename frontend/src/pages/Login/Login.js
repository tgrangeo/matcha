import style from "./style.module.scss";

import Logo from "../../assets/pokemeet.png";
import Input from "../../components/Input/Input";
import Button from "../../components/Button/Button";
const Login = () => {
  return (
    <div className={style.Login}>
      <form>
        <img src={Logo} alt={"PokeMeet"} />
        <Input
          type="email"
          name="email"
          required
          placeholder="Entrer votre email..."
        >
          Email
        </Input>
        <Input
          type="password"
          name="password"
          required
          placeholder="Entrer votre mot de passe..."
        >
          Mot de passe
        </Input>

        <Button className={style.Button} style="filled" align="center">
          Log in
        </Button>
        <p className={style.noAccount}>
          No account? <a href="/subscription">Create account</a>
        </p>
      </form>
    </div>
  );
};

export default Login;
