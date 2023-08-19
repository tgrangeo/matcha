import style from "./style.module.scss";
import ElementGallery from "./ElementGallery/ElementGallery";

const PhotoGallery = () => {
  const user = {
    images: [
      "https://cdn.pixabay.com/photo/2015/04/23/22/00/tree-736885__480.jpg",
      "",
      "https://cdn.pixabay.com/photo/2015/04/23/22/00/tree-736885__480.jpg",
      "https://cdn.pixabay.com/photo/2015/04/23/22/00/tree-736885__480.jpg",
      "",
    ],
  };
  return (
    <>
      <h2>Gallerie photo</h2>
      <div className={style.images}>
        {(user?.images ?? []).map((image, i) => (
          <ElementGallery key={i} imageInit={image} i={i} />
        ))}
      </div>
    </>
  );
};

export default PhotoGallery;
