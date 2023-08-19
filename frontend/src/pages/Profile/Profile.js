import Select from "react-select";
import Pokeball from "../../components/Pokeball/Pokeball";
import TypeIcon from "../../components/TypeIcon/TypeIcon";
import style from "./style.module.scss";
import { useEffect, useState } from "react";
import PublicInfos from "../../components/PublicInfos/PublicInfos";
import PrivateInfos from "../../components/PrivateInfos/PrivateInfos";
import PhotoGallery from "../../components/PhotoGallery/PhotoGallery";

const Profile = () => {
  const [user, setUser] = useState(null);
  const [tagsOptions, setTagsOptions] = useState(null);
  //TODO
  //GET api/v1/tags
  //GET api/v1/me jwt

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
