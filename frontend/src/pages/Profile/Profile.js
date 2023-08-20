import Select from "react-select";
import Pokeball from "../../components/Pokeball/Pokeball";
import TypeIcon from "../../components/TypeIcon/TypeIcon";
import style from "./style.module.scss";
import { useEffect, useState } from "react";
import PublicInfos from "../../components/PublicInfos/PublicInfos";
import PrivateInfos from "../../components/PrivateInfos/PrivateInfos";
import PhotoGallery from "../../components/PhotoGallery/PhotoGallery";

const Profile = () => {
  // const [user, setUser] = useState(null);
  const [tagsOptions, setTagsOptions] = useState(null);
  //TODO
  //GET api/v1/tags
  //GET api/v1/me jwt

  const user = {
    images: [
      "https://cdn.pixabay.com/photo/2015/04/23/22/00/tree-736885__480.jpg",
      "https://cdn.pixabay.com/photo/2015/04/23/22/00/tree-736885__480.jpg",
      "",
      "https://cdn.pixabay.com/photo/2015/04/23/22/00/tree-736885__480.jpg",
      "https://cdn.pixabay.com/photo/2015/04/23/22/00/tree-736885__480.jpg",
      "",
    ],
    pokeball: "masterball",
    type: "plante",
    firstname: "Jean",
    lastname: "Dupont",
    bio: "Je suis un grand fan de Pokémon et j'adore passer des heures à explorer les régions et attraper de nouveaux Pokémon. Si vous partagez cette passion ou que vous êtes curieux d'en savoir plus, n'hésitez pas à me contacter !",
    tags: [0],
    email: "jean.dupont@nasa.fr",
    gender: "male",
    attract: { female: true, male: false, nb: false },
  };
  //TODO:
  useEffect(() => {
    fetch(":8080/api/v1/me")
      .then((res) => res.json())
      .then((data) =>
        //setUser(data)
        console.log(data)
      );

    fetch(":8080/api/v1/tags")
      .then((res) => res.json())
      .then((data) =>
        //setTagsOptions(data)
        console.log(data)
      );
  }, []);

  //TODO: protect when no user
  return (
    <div className={style.Profile}>
      <PublicInfos user={user} tagsOptions={tagsOptions} />
      <PhotoGallery user={user} />
      <PrivateInfos user={user} />
    </div>
  );
};

export default Profile;
