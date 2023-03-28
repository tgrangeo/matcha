import style from "./style.module.scss";

const Credentials = ({ password, repassword, email, onChange }) => {
	return (
		<div className={style.NameAndBirthdate}>
			<p style={{ width: "calc(50% - 10px)" }}>
				<label htmlFor="email">Email</label>
				<input
					type="email"
					id="email"
					required
					value={email}
					placeholder="Entrer votre email..."
					onChange={(e) => onChange("email", e.target.value)}
				></input>
			</p>
			<div>
				<p>
					<label htmlFor="password">Mot de passe</label>
					<input
						type="password"
						id="password"
						required
						value={password}
						placeholder="Entrer votre mot de passe..."
						onChange={(e) => onChange("password", e.target.value)}
					></input>
				</p>

				<p>
					<label htmlFor="repassword">Confimation du mot de passe</label>
					<input
						type="password"
						id="repassword"
						required
						value={repassword}
						placeholder="Confirmer votre mot de passe..."
						onChange={(e) => onChange("repassword", e.target.value)}
					></input>
				</p>
			</div>
		</div>
	);
};

export default Credentials;
