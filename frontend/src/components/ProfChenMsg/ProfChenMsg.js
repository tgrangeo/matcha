import style from "./style.module.scss";
import ProfChen from "../../assets/prof_Chen (1).png";
import { useEffect, useState } from "react";

const ProfChenMsg = ({ message, children, onClick }) => {
	const [messageCopy, setMessageCopy] = useState(message);
	const [length, setLength] = useState(1);
	const [ended, setEnded] = useState(false);

	useEffect(() => {
		const handleClick = (event) => {
			// Display the message with the string passed as an argument
			console.log(`Message: ${event.detail.message}`);
			setMessageCopy(event.detail.message);
		};

		document.addEventListener("ProfChenMessage", handleClick);

		return () => {
			document.removeEventListener("ProfChenMessage", handleClick);
		};
	}, []);

	useEffect(() => {
		// let len = 1;
		// console.log(length, message);
		if (ended) return;
		if (length === messageCopy.length) {
			setEnded(true);
			return;
		}
		if (length < messageCopy.length) {
			let timeout = setTimeout(() => {
				setLength(length + 1);
			}, 10);
			return () => clearTimeout(timeout);
		}
	}, [length, messageCopy, ended]);

	useEffect(() => {
		setLength(1);
		setEnded(false);
	}, [messageCopy]);

	useEffect(() => {
		setMessageCopy(message);
	}, [message]);

	return (
		<div className={style.ProfChenMsg}>
			<img src={ProfChen} alt="profChen" />
			<div className={style.Bubble} onClick={() => onClick()}>
				<p className={style.message}>{messageCopy.substring(0, length) ?? "..."}</p>
				{children}
			</div>
		</div>
	);
};

export default ProfChenMsg;
