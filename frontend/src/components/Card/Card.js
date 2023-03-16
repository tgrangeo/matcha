import style from "./style.module.scss";
import CardBg from "../../assets/CardBg.png";
import PP from "../../assets/pp.png";
import { getRandomType, getTypeGradient } from "../../utils";
import TypeIcon from "../TypeIcon/TypeIcon";
const Card = () => {
	const type = getRandomType();
	const gradient = getTypeGradient(type);

	return (
		<div className={style.cardBg} style={{ backgroundImage: gradient }}>
			<div className={style.card} style={{ backgroundImage: `url(${CardBg})` }}>
				<div className={style.main}>
					<div className={style.top}>
						<div className={style.right}>
							<p className={style.name}>Eliott Depauw</p>
							<p className={style.fame}>
								FAME <span className={style.fameValue}>60</span>
							</p>
						</div>
						<TypeIcon type={type} />
					</div>

					<div className={style.image}>
						<img src={PP} alt="pp" />
					</div>

					<div className={style.bottom}>
						<p className={style.title}>Description</p>
						<p className={style.subtitle}>
							Je suis un grand fan de Pokémon et j'adore passer des heures à explorer les régions
							et attraper de nouveaux Pokémon. Si vous partagez cette passion ou que vous êtes
							curieux d'en savoir plus, n'hésitez pas à me contacter !
						</p>

						<p className={style.title}>Type</p>
						<p className={style.subtitle}>Lorem ipsum dolor sit samere.</p>
					</div>
				</div>
			</div>
		</div>
	);
};

export default Card;
