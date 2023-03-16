function getTypeGradient(type) {
	switch (type) {
		case "normal":
			return "linear-gradient(to bottom left, #A8A77A, #CDD1C4)";
		case "fire":
			return "linear-gradient(to bottom left, #F95643, #FFC6A5)";
		case "water":
			return "linear-gradient(to bottom left, #4D90D5, #A2CFFE)";
		case "electric":
			return "linear-gradient(to bottom left, #F8D030, #FFF1A8)";
		case "grass":
			return "linear-gradient(to bottom left, #7AC74C, #B5D99C)";
		case "ice":
			return "linear-gradient(to bottom left, #96D9D6, #BCE5E5)";
		case "fighting":
			return "linear-gradient(to bottom left, #C22E28, #FF9C9D)";
		case "poison":
			return "linear-gradient(to bottom left, #A33EA1, #DCBDFB)";
		case "ground":
			return "linear-gradient(to bottom left, #E2BF65, #F5D58F)";
		case "flying":
			return "linear-gradient(to bottom left, #A98FF3, #D4B1FF)";
		case "psychic":
			return "linear-gradient(to bottom left, #F95587, #FFA5D2)";
		case "bug":
			return "linear-gradient(to bottom left, #A6B91A, #C9DE55)";
		case "rock":
			return "linear-gradient(to bottom left, #B6A136, #D5C88C)";
		case "ghost":
			return "linear-gradient(to bottom left, #735797, #A291BC)";
		case "dragon":
			return "linear-gradient(to bottom left, #6F35FC, #B29DFD)";
		case "dark":
			return "linear-gradient(to bottom left, #705746, #A29288)";
		case "steel":
			return "linear-gradient(to bottom left, #B7B7CE, #D1D1E0)";
		case "fairy":
			return "linear-gradient(to bottom left, #D685AD, #FFB7DD)";
		default:
			return "linear-gradient(to bottom left, #68A090, #A4D4AE)";
	}
}
function getTypeColor(type) {
	switch (type) {
		case "normal":
			return "#A8A77A";
		case "fire":
			return "#F95643";
		case "water":
			return "#4D90D5";
		case "electric":
			return "#F8D030";
		case "grass":
			return "#7AC74C";
		case "ice":
			return "#96D9D6";
		case "fighting":
			return "#C22E28";
		case "poison":
			return "#A33EA1";
		case "ground":
			return "#E2BF65";
		case "flying":
			return "#A98FF3";
		case "psychic":
			return "#F95587";
		case "bug":
			return "#A6B91A";
		case "rock":
			return "#B6A136";
		case "ghost":
			return "#735797";
		case "dragon":
			return "#6F35FC";
		case "dark":
			return "#705746";
		case "steel":
			return "#B7B7CE";
		case "fairy":
			return "#D685AD";
		default:
			return "#68A090";
	}
}
function getRandomType() {
	const types = [
		"normal",
		"fire",
		"water",
		"electric",
		"grass",
		"ice",
		"fighting",
		"poison",
		"ground",
		"flying",
		"psychic",
		"bug",
		"rock",
		"ghost",
		"dragon",
		"dark",
		"steel",
		"fairy",
	];
	const randomIndex = Math.floor(Math.random() * types.length);
	return types[randomIndex];
}
export { getTypeColor, getTypeGradient, getRandomType };
