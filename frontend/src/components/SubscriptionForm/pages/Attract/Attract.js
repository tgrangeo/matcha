import { useState } from "react";
import style from "./style.module.scss";
import clsx from "clsx";

const Attract = ({ onChange, value = { male: false, female: false, nb: false } }) => {
	console.log(value.male || value.female || value.nb);
	const valid = value.male || value.female || value.nb;
	const handleChange = (newValue) => {
		let temp = value;
		temp[newValue] = !temp[newValue];
		onChange("attract", temp);
	};
	return (
		<div className={style.Attract}>
			<div className={style.checkbox} onClick={() => handleChange("female")}>
				<div className={clsx(style.check, value.female && style.active)} />
				<p>Femme</p>
			</div>
			<div className={style.checkbox} onClick={() => handleChange("male")}>
				<div className={clsx(style.check, value.male && style.active)} />
				<p>Homme</p>
			</div>
			<div className={style.checkbox} onClick={() => handleChange("nb")}>
				<div className={clsx(style.check, value.nb && style.active)} />
				<p>Non binaire</p>
			</div>
			<input type="checkbox" onChange={() => {}} checked={valid} required />
		</div>
	);
};

export default Attract;
