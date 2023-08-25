import style from "./style.module.scss";
import Logo from "../../assets/pokemeet.png";
import HeaderProfile from "../HeaderProfile/HeaderProfile";

// import { ReactComponent as Logo } from '../../assets/meetic-logo-vector.svg'
const Header = ({ name }) => {
	return (
		<header className={style.header}>
			<div className={style.content}>
				<img src={Logo} alt={"PokeMeet"} onClick={() => window.location.href = "http://localhost:3000/"} />
				<HeaderProfile />
			</div>
		</header>
	);
};

export default Header;
