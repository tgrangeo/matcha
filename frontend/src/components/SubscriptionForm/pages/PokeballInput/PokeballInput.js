import style from "./style.module.scss";
import clsx from "clsx";
import MasterBall from "../../../../assets/MasterBall.png";
import RapidBall from "../../../../assets/rapidBall.png";
import LoveBall from "../../../../assets/loveball.png";
import Pokeball from "../../../../assets/pokeball.png";

const PokeballInput = ({ onChange, activeIndex }) => {
	const balls = [
		{
			name: "masterball",
			subtitle: "Relation serieuse",
			desc: "Je suis prêt(e) à me laisser capturer par une Masterball pour trouver la personne avec qui je pourrai construire une relation sérieuse et engagée.",
			imageSrc: MasterBall,
		},
		{
			name: "rapideball",
			subtitle: "Sans lendemain",
			desc: "Je suis ouvert(e) à me faire capturer par une Rapidball pour un coup d'un soir, avec quelqu'un prêt à partager un moment de plaisir intense et éphémère.",
			imageSrc: RapidBall,
		},
		{
			name: "loveball",
			subtitle: "Plan cul regulier",
			desc: "Je suis à la recherche d'une Loveball pour un plan cul régulier, avec une personne prête à me capturer et à partager des moments de plaisir intense.",
			imageSrc: LoveBall,
		},
		{
			name: "pokeball",
			subtitle: "Boire un verre",
			desc: "Je cherche quelqu'un qui serait prêt(e) à me capturer le temps d'une soirée, avec une Pokeball pleine de bons moments et de rires autour d'un verre",
			imageSrc: Pokeball,
		},
	];
	return (
		<div className={style.Pokeballinput}>
			{balls.map((el, i) => (
				<div
					key={i}
					onClick={() => onChange("pokeball", i)}
					className={clsx(style.base, activeIndex === i && style.active)}
				>
					<div className={style.type}>
						<img src={el.imageSrc} alt={el.name} />
						<div className={style.title}>
							{el.name}
							<p className={style.subtitle}>{el.subtitle}</p>
						</div>
					</div>
					<p>{el.desc}</p>
				</div>
			))}
		</div>
	);
};

export default PokeballInput;
