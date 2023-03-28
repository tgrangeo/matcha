import { getTypeDescription } from "../../../../utils";
import TypeIcon from "../../../TypeIcon/TypeIcon";
import style from "./style.module.scss";
import clsx from "clsx";

const PokeType = ({ onChange, activeIndex }) => {
	const types = [
		"normal",
		"feu",
		"eau",
		"electrik",
		"plante",
		"glace",
		"combat",
		"poison",
		"sol",
		"vol",
		"psy",
		"insecte",
		"roche",
		"spectre",
		"dragon",
		"sombre",
		"acier",
		"fee",
	];
	return (
		<div className={style.PokeType}>
			{types.map((el, i) => (
				<div
					key={i}
					onClick={() => onChange("type", i)}
					className={clsx(style.base, activeIndex === i && style.active)}
				>
					<div className={style.type}>
						<TypeIcon type={el} />
						<p className={style.title}>{el}</p>
					</div>
					<p>{getTypeDescription(el)}</p>
				</div>
			))}
		</div>
	);
};

export default PokeType;
