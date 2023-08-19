import Button from "../Button/Button";
import Popup from "../Popup/Popup";

const PopupButton = ({
  button,
  children,
  open = false,
  onChange,
  style,
  ...props
}) => {
  return (
    <>
      {button && (
        <Button
          children={button}
          {...props}
          style={style}
          onClick={(e) => {
            onChange(true);
          }}
        />
      )}
      <Popup open={open} onClose={() => onChange(false)}>
        {children}
      </Popup>
    </>
  );
};

export default PopupButton;
