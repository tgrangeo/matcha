import { useEffect, useState } from "react";
import style from "./style.module.scss";
import Select from "react-select";

const Tags = () => {
	const [tagsOptions, setTagsOptions] = useState([
		{ label: "cinema", value: 0 },
		{ label: "la drogue", value: 1 },
	]);

	useEffect(() => {
		// fetch("")
		// 	.catch((res) => res.json())
		// 	.then((data) => setTagsOptions(data));
		//TODO: get tags as options
	}, []);
	return (
		<div className={style.Tags}>
			<Select
				isMulti
				name="colors"
				options={tagsOptions}
				className="basic-multi-select"
				classNamePrefix="select"
				required
			/>
		</div>
	);
};

export default Tags;
