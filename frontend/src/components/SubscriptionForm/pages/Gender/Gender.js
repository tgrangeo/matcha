import style from "./style.module.scss";
import clsx from "clsx";

const Gender = ({ onChange, value }) => {
	console.log(value);
	return (
		<div className={style.Gender}>
			<div className={style.radio} onClick={() => onChange("gender", "female")}>
				<div className={clsx(style.check, value === "female" && style.active)} />
				<p>Femme</p>
			</div>
			<div className={style.radio} onClick={() => onChange("gender", "male")}>
				<div className={clsx(style.check, value === "male" && style.active)} />
				<p>Homme</p>
			</div>
			<div className={style.radio} onClick={() => onChange("gender", "nb")}>
				<div className={clsx(style.check, value === "nb" && style.active)} />
				<p>Non binaire</p>
			</div>
		</div>
	);
};

export default Gender;
