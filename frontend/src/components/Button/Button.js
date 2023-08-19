import clsx from "clsx";
import styles from "./style.module.scss";

const Button = ({
  children,
  Wrapper = "button",
  align = "center",
  styl = "filled", //ghost , filled , outlined
  ...props
}) => (
  <div className={clsx(styles.container, props?.className)}>
    <Wrapper
      {...props}
      className={clsx(styles[styl], props.buttonClassName)}
      style={{
        margin:
          align === "center" ? "0 auto" : align === "left" ? "0" : "0 auto 0 0",
      }}
    >
      {children}
    </Wrapper>
  </div>
);

export default Button;
