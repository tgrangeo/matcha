import style from "./style.module.scss";
import React, { useState } from 'react';

import Logo from "../../assets/pokemeet.png";
const Login = () => {
	const [email, setEmail] = useState('');
	const [password, setPassword] = useState('');

	const handleLogin = (event) => {
		event.preventDefault();

		// Préparez les données à envoyer au serveur
		const loginData = {
			email: email,
			password: password,
		};

		// Effectuez la requête fetch vers le serveur
		fetch("http://localhost:8080/api/v1/signin", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify(loginData),
		})
			.then((response) => response.json())
			.then((data) => {
				// Traitez la réponse du serveur ici (ex: redirigez l'utilisateur)
			})
			.catch((error) => {
				console.error("An error occurred:", error);
			});
	};

	return (
		<div className={style.Login}>
			<form onSubmit={handleLogin}>
				<img src={Logo} alt={"PokeMeet"} />
				<p className={style.row}>
					<label htmlFor="email">Email</label>
					<input
						type="email"
						name="email"
						id="email"
						placeholder="Enter email..."
						value={email}
						onChange={(e) => setEmail(e.target.value)}
						required
					/>
				</p>
				<p className={style.row}>
					<label htmlFor="password">Password</label>
					<input
						type="password"
						name="password"
						id="password"
						placeholder="Enter password..."
						value={password}
						onChange={(e) => setPassword(e.target.value)}
						required
					/>
				</p>
				<input type="submit" value="Log in" />
				<p className={style.noAccount}>
					No account? <a href="/subscription">Create account</a>
				</p>
				{/* //TODO: create link to /api/v1/resetpass */}
				<p className={style.noAccount}>
					Forgot password? <a href="">Recover password</a>
				</p>
			</form>
		</div>
	);
};

export default Login;
