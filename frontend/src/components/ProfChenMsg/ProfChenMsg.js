import style from "./style.module.scss";
import ProfChen from "../../assets/prof_Chen (1).png";
import { useEffect, useState } from "react";

const ProfChenMsg = ({ message, children, onClick }) => {
	const [length, setLength] = useState(1);
	const [ended, setEnded] = useState(false);

	useEffect(() => {
		// let len = 1;
		// console.log(length, message);
		if (ended) return;
		if (length === message.length) {
			setEnded(true);
			return;
		}
		if (length < message.length) {
			let timeout = setTimeout(() => {
				setLength(length + 1);
			}, 10);
			return () => clearTimeout(timeout);
		}
	}, [length, message, ended]);

	useEffect(() => {
		setLength(1);
		setEnded(false);
	}, [message]);

	return (
		<div className={style.ProfChenMsg}>
			<img src={ProfChen} alt="profChen" />
			<div className={style.Bubble} onClick={() => onClick()}>
				<p className={style.message}>{message.substring(0, length) ?? "..."}</p>
				{children}
			</div>
		</div>
	);
};

export default ProfChenMsg;
