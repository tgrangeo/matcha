import style from "./style.module.scss";
import clsx from "clsx";

const Popup = ({ children, open = false, onClose }) => {
  return (
    <div
      className={clsx(style.popupfull, open && style.open)}
      onClick={(e) => {
        e.stopPropagation();
        onClose();
      }}
    >
      {/* <div className={style.blur} /> */}
      <div
        className={style.popupContainer}
        onClick={(e) => {
          e.stopPropagation();
        }}
      >
        <i
          class="fi fi-rr-cross-small"
          onClick={(e) => {
            e.stopPropagation();
            onClose();
          }}
        ></i>
        <div className={style.popupContent}>{children}</div>
      </div>
    </div>
  );
};

export default Popup;
