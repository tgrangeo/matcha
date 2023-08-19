import Select from "react-select";
import Pokeball from "../Pokeball/Pokeball";
import TypeIcon from "../TypeIcon/TypeIcon";
import style from "./style.module.scss";
import { useEffect, useState } from "react";
import Gender from "../SubscriptionForm/pages/Gender/Gender";
import Attract from "../SubscriptionForm/pages/Attract/Attract";

const PublicInfos = ({ user, tagsOptions }) => {
  return (
    <>
      {/* <h1>Profile</h1> */}
      <div className={style.profile}>
        <div className={style.left}>
          <div className={style.avatar}>
            <img src="" alt="avatar" />
            <i class="fi fi-rr-edit"></i>
          </div>
          <div className={style.infos}>
            <div className={style.pokeball}>
              <h4>Pokeball</h4>
              <Pokeball pokeball={"masterball"} />
              <div>
                <p>Masterball</p>
                <i class="fi fi-rr-edit"></i>
              </div>
            </div>
            <div className={style.type}>
              <h4>Type</h4>
              <TypeIcon type="plante" />
              <div>
                <p>Plante</p>
                <i class="fi fi-rr-edit"></i>
              </div>
            </div>
          </div>
        </div>
        <div className={style.right}>
          <div className={style.name}>
            <h2>Prénom</h2>
            <h2>Nom</h2>
          </div>

          <p className={style.bio}>
            Je suis un grand fan de Pokémon et j'adore passer des heures à
            explorer les régions et attraper de nouveaux Pokémon. Si vous
            partagez cette passion ou que vous êtes curieux d'en savoir plus,
            n'hésitez pas à me contacter ! <i class="fi fi-rr-edit"></i>
          </p>

          <Select
            isMulti
            name="colors"
            options={[
              { label: "cinema", value: 0 },
              { label: "la drogue", value: 1 },
            ]}
            value={1}
            className="basic-multi-select"
            classNamePrefix="select"
            required
          />
        </div>
      </div>
      <div className={style.radios}>
        <div className={style.radio}>
          <h4>Je suis un(e): </h4>
          <Gender value={"female"} />
        </div>
        <div className={style.radio}>
          <h4>Je recherche: </h4>
          <Attract value={{ male: true, female: false, nb: true }} />
        </div>
      </div>
      <button className={style.button}>Modifier</button>
    </>
  );
};

export default PublicInfos;
