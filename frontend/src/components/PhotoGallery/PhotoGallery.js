import style from "./style.module.scss";
import ElementGallery from "./ElementGallery/ElementGallery";

const PhotoGallery = ({ user }) => {
  const imagesWithoutFirst = user?.images?.slice(1);
  return (
    <>
      <h2>Gallerie photo</h2>
      <div className={style.images}>
        {(imagesWithoutFirst ?? []).map((image, i) => (
          <ElementGallery key={i} imageInit={image} i={i + 1} />
        ))}
      </div>
    </>
  );
};

export default PhotoGallery;
