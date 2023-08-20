import style from "./style.module.scss";
import Card from "../Card/Card";
import MasterBall from "../..//assets/MasterBall.png";
import RapidBall from "../../assets/rapidBall.png";
import LoveBall from "../../assets/loveball.png";
import Pokeballl from "../../assets/pokeball.png";

const Pokeball = ({ pokeball }) => {
  switch (pokeball) {
    case "masterball":
      return <img src={MasterBall} alt="masterball" />;
    case "rapidball":
      return <img src={RapidBall} alt="rapidball" />;
    case "loveball":
      return <img src={LoveBall} alt="loveball" />;
    case "pokeball":
      return <img src={Pokeballl} alt="pokeball" />;
      return null;
  }
};

export default Pokeball;
