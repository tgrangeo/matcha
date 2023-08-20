import { useRef, useState } from "react";
import style from "./style.module.scss";
import Popup from "../../Popup/Popup";
import Button from "../../Button/Button";
import PopupButton from "../../PopupButton/PopupButton";
import AvatarEditor from "react-avatar-editor";
import Input from "../../Input/Input";
import clsx from "clsx";

const ElementGallery = ({ imageInit, i, ...props }) => {
  const [popup, setPopup] = useState(false);
  const [popupDelete, setPopupDelete] = useState(false);
  const [addPopup, setAddPopup] = useState(false);
  const [imageFile, setImageFile] = useState(null);
  const [scale, setScale] = useState(0);
  const [image, setImage] = useState(imageInit);
  const editor = useRef(null);

  const handleChangeImage = (imageFile, index) => {
    const formData = new FormData();
    formData.append("image", imageFile ?? null);
    formData.append("index", index);

    // fetch(":8080/api/v1/setImage", {
    //   method: "POST",
    //   headers: {
    //   ContentType: "multipart/form-data",
    //     Authorization: "Bearer " + localStorage.getItem("jwt"),
    //   },
    //   body: formData,
    // })
    //   .then((response) => response.json())
    //   .then((data) => {
    //     console.log(data);
    //   })
    //   .catch((error) => {
    //     console.error("Erreur:", error);
    //   });

    console.log(imageFile, index);
  };
  const handleAddImage = (file) => {
    console.log(file);
    setImageFile(file);
    console.log(addPopup);
    setAddPopup(true);
  };
  return (
    <>
      {image === "" ? (
        <div className={clsx(style.imageAdd, props?.className)}>
          <Input
            type="file"
            onChange={(e) => handleAddImage(e.target.files[0])}
            name={"image_" + i}
            classNames={style.input}
            accept="image/*"
          >
            <i class="fi fi-rr-add"></i>
          </Input>
        </div>
      ) : (
        <>
          <div
            className={clsx(style.image, props?.className)}
            style={{ backgroundImage: `url(${image})` }}
            onClick={() => setPopup(true)}
          >
            <div
              className={style.buttons}
              onClick={(e) => {
                e.stopPropagation();
              }}
            >
              <Input
                type="file"
                onChange={(e) => handleAddImage(e.target.files[0])}
                name={"image_" + i}
                className={style.input}
                accept="image/*"
              >
                <i class="fi fi-rr-replace"></i>
              </Input>
              <Popup
                button={<i class="fi fi-rr-cross-small"></i>}
                open={popup}
                onClose={setPopup}
              >
                test
              </Popup>

              <PopupButton
                button={<i class="fi fi-rr-trash"></i>}
                styl="ghost"
                open={popupDelete}
                onChange={setPopupDelete}
              >
                <p>Etes vous sur de vouloir supprimer cette image ?</p>
                <Button
                  Wrapper="button"
                  align="center"
                  styl="filled"
                  className={style.button}
                  onClick={() => {
                    setImage("");
                    handleChangeImage("", i);
                    setPopupDelete(false);
                  }}
                >
                  Confirmer
                </Button>
              </PopupButton>
            </div>
          </div>
        </>
      )}
      <Popup
        button={<i class="fi fi-rr-cross-small"></i>}
        open={addPopup}
        onClose={setAddPopup}
      >
        <AvatarEditor
          ref={editor}
          image={imageFile}
          width={500}
          height={500}
          backgroundColor={style.blue}
          border={50}
          color={[255, 255, 255, 0.6]} // RGBA
          scale={scale / 100 + 1}
          rotate={0}
        />
        {/* slider for scale */}
        <input
          type="range"
          max={100}
          value={scale}
          onChange={(e) => {
            setScale(e.target.value);
          }}
        />
        <Button
          onClick={() => {
            if (editor) {
              // This returns a HTMLCanvasElement, it can be made into a data URL or a blob,
              // drawn on another canvas, or added to the DOM.

              // If you want the image resized to the canvas size (also a HTMLCanvasElement)
              const canvasScaled = editor.current.getImageScaledToCanvas();
              console.log(canvasScaled);
              //to file
              const test = canvasScaled.toBlob(
                (blob) => {
                  console.log(blob);
                  // handleChangeImage(blob, i);
                },
                "image/jpeg",
                0.95
              );
              const file = new File([test], "test.jpg", {
                type: "image/jpeg",
              });
              // console.log(test);
              setImage(canvasScaled.toDataURL("image/jpeg", 0.95));
              console.log(file);
              setAddPopup(false);
            }
          }}
        >
          Valider
        </Button>
      </Popup>
    </>
  );
};

export default ElementGallery;
