import style from "./style.module.scss";
import Card from "../Card/Card";

const CardGrid = () => {
	return (
		<div className={style.CardGrid}>
			<div className={style.content}>
				<Card />
				<Card />
				<Card />
				<Card />
				<Card />
				<Card />
				<Card />
				<Card />
				<Card />
			</div>
			<div className={style.searchBar}></div>
		</div>
	);
};

export default CardGrid;
