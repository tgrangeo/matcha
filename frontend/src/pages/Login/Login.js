import style from "./style.module.scss";
import React, { useState } from 'react';

import Logo from "../../assets/pokemeet.png";
import Input from "../../components/Input/Input";
import Button from "../../components/Button/Button";
const Login = () => {

	const handleLogin = (event) => {
		event.preventDefault();

		// Préparez les données à envoyer au serveur
		const loginData = {
			email: event.target[0].value,
			password: event.target[1].value,
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
