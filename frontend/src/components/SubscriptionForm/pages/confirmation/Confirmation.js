import React, { useState } from "react";
import Input from "../../../Input/Input";
import style from "./style.module.scss";

const Credentials = ({ onConfirmation }) => {
	const [confirmationCode, setConfirmationCode] = useState("");

	const handleConfirmationCodeChange = (e) => {
		const code = e.target.value;
		if (/^\d{0,6}$/.test(code)) {
			setConfirmationCode(code);
		}
	};

	const handleConfirmButtonClick = () => {
		if (confirmationCode.length === 6) {
			const requestBody = JSON.stringify({ Token: confirmationCode });

			fetch("/api/v1/validate", {
				method: "POST",
				// mode: "no-cors", //TODO: remove
				headers: {
					"Content-Type": "application/json",
				},
				body: requestBody,
			})
				.then((data) => {
					console.log(data.status)
					if (data.status === 200) {
						window.location.href = "http://localhost:3000/";
						//TODO: connecter le gars
					} else {
						//TODO: Gérer d'autres cas de réponse si nécessaire
					}
				})
				.catch((error) => {
					console.error("Erreur lors de l'envoi de la requête :", error);
				});
		}
	};

	return (
		<div className={style.Conf}>
			<Input
				type="text"
				name="confirmationCode"
				required
				classNames={style.input0}
				value={confirmationCode}
				placeholder="Entrer votre code de confirmation..."
				onChange={handleConfirmationCodeChange}
			>
				Code de Confirmation
			</Input>
			<button className="but" onClick={handleConfirmButtonClick}>Confirmer</button>
		</div>
	);
};

export default Credentials;