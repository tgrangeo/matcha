import Select from "react-select";
import Pokeball from "../Pokeball/Pokeball";
import TypeIcon from "../TypeIcon/TypeIcon";
import style from "./style.module.scss";
import { useEffect, useState } from "react";
import Gender from "../SubscriptionForm/pages/Gender/Gender";
import Attract from "../SubscriptionForm/pages/Attract/Attract";
import ElementGallery from "../PhotoGallery/ElementGallery/ElementGallery";
import Button from "../Button/Button";

const PublicInfos = ({ user, tagsOptions }) => {
  const [modifiedUser, setModifiedUser] = useState(user);
  const [isEditingBio, setIsEditingBio] = useState(false);
  const options = [
    { value: "cinema", label: 0 },
    { value: "la drogue", label: 1 },
  ];

  const setUser = (key, value) => {
    const tempUser = { ...user };
    tempUser[key] = value;
    setModifiedUser(tempUser);
  };

  return (
    <>
      {/* <h1>Profile</h1> */}
      <div className={style.profile}>
        <div className={style.left}>
          <ElementGallery
            className={style.avatar}
            imageInit={modifiedUser.images[0]}
            i={0}
          />
          <div className={style.infos}>
            <div className={style.pokeball}>
              <h4>Pokeball</h4>
              <Pokeball pokeball={user.pokeball} />
              <div>
                <p>{user.pokeball}</p>
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
            <h2>{user.firstname}</h2>
            <h2>{user.lastname}</h2>
          </div>

          {isEditingBio ? (
            <textarea
              className={style.bio}
              value={modifiedUser.bio}
              maxLength={300}
              on={(e) => {
                e.target.style.height = "1px";
                e.target.style.height = e.target.scrollHeight + "px";
              }}
              onChange={(e) => {
                //update height
                e.target.style.height = "1px";
                e.target.style.height = e.target.scrollHeight + "px";
                setUser("bio", e.target.value);
              }}
            ></textarea>
          ) : (
            <p className={style.bio}>
              Je suis un grand fan de Pokémon et j'adore passer des heures à
              explorer les régions et attraper de nouveaux Pokémon. Si vous
              partagez cette passion ou que vous êtes curieux d'en savoir plus,
              n'hésitez pas à me contacter !{" "}
              <i
                class="fi fi-rr-edit"
                onClick={() => setIsEditingBio(true)}
              ></i>
            </p>
          )}

          <Select
            isMulti
            name="colors"
            options={options}
            value={options.filter((obj) =>
              modifiedUser.tags.includes(obj.value)
            )}
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
      {user !== modifiedUser && (
        <Button
          Wrapper="button"
          align="center"
          styl="filled"
          className={style.button}
        >
          Sauvegarder
        </Button>
      )}
    </>
  );
};

export default PublicInfos;
