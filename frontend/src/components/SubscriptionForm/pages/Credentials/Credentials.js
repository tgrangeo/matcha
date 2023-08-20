import Input from "../../../Input/Input";
import style from "./style.module.scss";

const Credentials = ({ password, repassword, email, onChange }) => {
  return (
    <div className={style.Credentials}>
      {/* <p style={{ width: "calc(50% - 10px)" }}>
        <label htmlFor="email">Email</label>
        <input
          type="email"
          id="email"
          required
          value={email}
          placeholder="Entrer votre email..."
          onChange={(e) => onChange("email", e.target.value)}
        ></input>
      </p> */}
      <Input
        type="email"
        name="email"
        required
        classNames={style.input0}
        value={email}
        placeholder="Entrer votre email..."
        onChange={(e) => onChange("email", e.target.value)}
      >
        Email
      </Input>
      <Input
        type="password"
        name="password"
        required
        classNames={style.input1}
        value={password}
        placeholder="Entrer votre mot de passe..."
        onChange={(e) => onChange("password", e.target.value)}
      >
        Mot de passe
      </Input>
      <Input
        type="password"
        name="repassword"
        classNames={style.input2}
        required
        value={repassword}
        placeholder="Confirmer votre mot de passe..."
        onChange={(e) => onChange("repassword", e.target.value)}
      >
        Confirmer le mot de passe
      </Input>
    </div>
  );
};

export default Credentials;
