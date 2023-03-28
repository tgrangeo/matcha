import style from "./style.module.scss";
import { ReactComponent as Bug } from "./icons/bug.svg";
import { ReactComponent as Dark } from "./icons/dark.svg";
import { ReactComponent as Dragon } from "./icons/dragon.svg";
import { ReactComponent as Electric } from "./icons/electric.svg";
import { ReactComponent as Fairy } from "./icons/fairy.svg";
import { ReactComponent as Fighting } from "./icons/fighting.svg";
import { ReactComponent as Fire } from "./icons/fire.svg";
import { ReactComponent as Flying } from "./icons/flying.svg";
import { ReactComponent as Ghost } from "./icons/ghost.svg";
import { ReactComponent as Grass } from "./icons/grass.svg";
import { ReactComponent as Ground } from "./icons/ground.svg";
import { ReactComponent as Ice } from "./icons/ice.svg";
import { ReactComponent as Normal } from "./icons/normal.svg";
import { ReactComponent as Poison } from "./icons/poison.svg";
import { ReactComponent as Psychic } from "./icons/psychic.svg";
import { ReactComponent as Rock } from "./icons/rock.svg";
import { ReactComponent as Steel } from "./icons/steel.svg";
import { ReactComponent as Water } from "./icons/water.svg";
import clsx from "clsx";
const TypeIcon = ({ type }) => {
	return (
		<div className={clsx(style.TypeIcon, style[type])}>
			{type === "insecte" && <Bug />}
			{type === "sombre" && <Dark />}
			{type === "dragon" && <Dragon />}
			{type === "electrik" && <Electric />}
			{type === "fee" && <Fairy />}
			{type === "combat" && <Fighting />}
			{type === "feu" && <Fire />}
			{type === "vol" && <Flying />}
			{type === "spectre" && <Ghost />}
			{type === "plante" && <Grass />}
			{type === "sol" && <Ground />}
			{type === "glace" && <Ice />}
			{type === "normal" && <Normal />}
			{type === "poison" && <Poison />}
			{type === "psy" && <Psychic />}
			{type === "roche" && <Rock />}
			{type === "acier" && <Steel />}
			{type === "eau" && <Water />}
		</div>
	);
};

export default TypeIcon;
