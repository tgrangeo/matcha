import style from "./style.module.scss";

import Logo from "../../assets/pokemeet.png";
const Login = () => {
	return (
		<div className={style.Login}>
			<form>
				<img src={Logo} alt={"PokeMeet"} />
				<p className={style.row}>
					<label htmlFor="email">Email</label>
					<input type="email" name="email" id="email" placeholder="Enter email..." />
				</p>
				<p className={style.row}>
					<label htmlFor="password">Password</label>
					<input type="password" name="password" id="password" placeholder="Enter password..." />
				</p>
				<input type="submit" value="Log in" />
				<p className={style.noAccount}>
					No account? <a href="/subscription">Create account</a>
				</p>
			</form>
		</div>
	);
};

export default Login;
