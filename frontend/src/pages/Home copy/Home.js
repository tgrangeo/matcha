// import { useState } from 'react';
import CardGrid from "../../components/CardGrid/CardGrid";
import style from "./style.module.scss";
//listing
const Home = () => {
	// const [users, setUsers] = useState({[]})
	return (
		<div className={style.Home}>
			<CardGrid />
		</div>
	);
};

export default Home;
