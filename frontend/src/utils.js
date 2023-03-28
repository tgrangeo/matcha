function getTypeGradient(type) {
	switch (type) {
		case "normal":
			return "linear-gradient(to bottom left, #A8A77A, #CDD1C4)";
		case "feu":
			return "linear-gradient(to bottom left, #F95643, #FFC6A5)";
		case "eau":
			return "linear-gradient(to bottom left, #4D90D5, #A2CFFE)";
		case "electrik":
			return "linear-gradient(to bottom left, #F8D030, #FFF1A8)";
		case "plante":
			return "linear-gradient(to bottom left, #7AC74C, #B5D99C)";
		case "glace":
			return "linear-gradient(to bottom left, #96D9D6, #BCE5E5)";
		case "combat":
			return "linear-gradient(to bottom left, #C22E28, #FF9C9D)";
		case "poison":
			return "linear-gradient(to bottom left, #A33EA1, #DCBDFB)";
		case "sol":
			return "linear-gradient(to bottom left, #E2BF65, #F5D58F)";
		case "vol":
			return "linear-gradient(to bottom left, #A98FF3, #D4B1FF)";
		case "psy":
			return "linear-gradient(to bottom left, #F95587, #FFA5D2)";
		case "insecte":
			return "linear-gradient(to bottom left, #A6B91A, #C9DE55)";
		case "roche":
			return "linear-gradient(to bottom left, #B6A136, #D5C88C)";
		case "spectre":
			return "linear-gradient(to bottom left, #735797, #A291BC)";
		case "dragon":
			return "linear-gradient(to bottom left, #6F35FC, #B29DFD)";
		case "sombre":
			return "linear-gradient(to bottom left, #705746, #A29288)";
		case "acier":
			return "linear-gradient(to bottom left, #B7B7CE, #D1D1E0)";
		case "fee":
			return "linear-gradient(to bottom left, #D685AD, #FFB7DD)";
		default:
			return "linear-gradient(to bottom left, #68A090, #A4D4AE)";
	}
}
function getTypeColor(type) {
	switch (type) {
		case "normal":
			return "#A8A77A";
		case "feu":
			return "#F95643";
		case "eau":
			return "#4D90D5";
		case "electrik":
			return "#F8D030";
		case "plante":
			return "#7AC74C";
		case "glace":
			return "#96D9D6";
		case "combat":
			return "#C22E28";
		case "poison":
			return "#A33EA1";
		case "sol":
			return "#E2BF65";
		case "vol":
			return "#A98FF3";
		case "psy":
			return "#F95587";
		case "insecte":
			return "#A6B91A";
		case "roche":
			return "#B6A136";
		case "spectre":
			return "#735797";
		case "dragon":
			return "#6F35FC";
		case "sombre":
			return "#705746";
		case "acier":
			return "#B7B7CE";
		case "fee":
			return "#D685AD";
		default:
			return "#68A090";
	}
}

function getTypeDescription(type) {
	switch (type) {
		case "normal":
			return "Le type Normal se caractérise par son équilibre et son absence de forces ou de faiblesses particulières. Les personnes de ce type sont souvent des compagnons fiables et stables.";
		case "feu":
			return "Le type Feu se distingue par sa passion et son impulsivité, mais aussi par sa grande protection envers ceux qu'il aime. Les personnes de ce type peuvent être des partenaires ardents et loyaux.";
		case "eau":
			return "Le type Eau est souvent calme et apaisant, mais peut également être imprévisible et changeant. Les personnes de ce type ont tendance à s'adapter facilement aux situations et à être très empathiques.";
		case "electrik":
			return "Le type Électrique est énergique et stimulant, mais peut également être imprévisible et chaotique. Les personnes de ce type ont tendance à être très expressives et communicatives.";
		case "plante":
			return "Le type Plante se caractérise par sa patience et son attention, mais peut aussi être possessif et jaloux. Les personnes de ce type ont tendance à être très sensibles à leur environnement et à leurs émotions.";
		case "glace":
			return "Le type Glace est souvent réservé et introverti, mais peut également être passionné et romantique. Les personnes de ce type ont tendance à être très loyales envers leurs proches.";
		case "combat":
			return "Le type Combat est déterminé et persévérant, mais peut aussi être impulsif et agressif. Les personnes de ce type ont tendance à être très compétitives et à aimer relever des défis.";
		case "poison":
			return "Le type Poison est souvent mystérieux et séduisant, mais peut aussi être manipulateur et dangereux. Les personnes de ce type ont tendance à être très habiles dans la manipulation des émotions et des relations.";
		case "sol":
			return "Le type Sol est souvent stable et fiable, mais peut aussi être têtu et obstiné. Les personnes de ce type ont tendance à être très terre-à-terre et pragmatiques.";
		case "vol":
			return "Le type Vol est souvent libre et aventureux, mais peut aussi être imprévisible et inconstant. Les personnes de ce type ont tendance à être très indépendantes et à avoir besoin de leur liberté.";
		case "psy":
			return "Le type Psy est souvent intuitif et sage, mais peut aussi être émotionnellement instable et difficile à comprendre. Les personnes de ce type ont tendance à être très profondes et à réfléchir beaucoup.";
		case "insecte":
			return "Le type Insecte est souvent travailleur et déterminé, mais peut aussi être obsessionnel et un peu bizarre. Les personnes de ce type ont tendance à être très efficaces dans l'accomplissement de tâches précises.";
		case "roche":
			return "Le type Roche est souvent solide et résistant, mais peut aussi être rigide et peu flexible. Les personnes de ce type ont tendance à être très résistantes face aux difficultés.";
		case "spectre":
			return "Le type Spectre est souvent mystérieux et introverti, mais peut aussi être très affectueux et attachant. Les personnes de ce type ont tendance à avoir une grande profondeur émotionnelle et à être très empathiques.";
		case "dragon":
			return "Le type Dragon est souvent puissant et majestueux, mais peut aussi être solitaire et inaccessible. Les personnes de ce type ont tendance à avoir une grande confiance en eux et à être très ambitieuses.";
		case "sombre":
			return "Le type Ténèbres est souvent mystérieux et intrigant, mais peut aussi être très protecteur et loyal envers ses proches. Les personnes de ce type ont tendance à avoir une grande capacité d'adaptation et à être très déterminées.";
		case "acier":
			return "Le type Acier est souvent solide et résistant, mais peut aussi être rigide et peu flexible. Les personnes de ce type ont tendance à être très résistantes face aux difficultés et à avoir une grande détermination.";
		case "fee":
			return "Le type Fée est souvent doux et affectueux, mais peut aussi être très fort et protecteur envers ses proches. Les personnes de ce type ont tendance à avoir une grande sensibilité et à être très empathiques envers les autres.";
		default:
			return "Fait pas le fou";
	}
}
function getRandomType() {
	const types = [
		"normal",
		"feu",
		"eau",
		"electrik",
		"plante",
		"glace",
		"combat",
		"poison",
		"sol",
		"vol",
		"psy",
		"insecte",
		"roche",
		"spectre",
		"dragon",
		"sombre",
		"acier",
		"fee",
	];
	const randomIndex = Math.floor(Math.random() * types.length);
	return types[randomIndex];
}
export { getTypeColor, getTypeGradient, getRandomType, getTypeDescription };
