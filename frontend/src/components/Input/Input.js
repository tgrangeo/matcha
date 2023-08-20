import clsx from "clsx";
import Button from "../Button/Button";
import style from "./style.module.scss";

const Input = ({ children, name, type, classNames, ...props }) => {
  return (
    <p
      className={clsx(
        style.inputContainer,
        classNames,
        type === "file" && style.file
      )}
    >
      <label htmlFor={name}>{children}</label>
      <input {...props} type={type} id={name} name={name}></input>
    </p>
  );
};

export default Input;
