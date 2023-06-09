import { useState } from "react";
import style from "./style.module.scss";
import NameAndBirthdate from "./pages/NameAndBirthdate/NameAndBirthdate";
import ProfChenMsg from "../ProfChenMsg/ProfChenMsg";
import PokeType from "./pages/PokeType/PokeType";
import PokeballInput from "./pages/PokeballInput/PokeballInput";
import Gender from "./pages/Gender/Gender";
import Attract from "./pages/Attract/Attract";
import Tags from "./pages/Tags/Tags";
import Credentials from "./pages/Credentials/Credentials";

const SubscriptionForm = () => {
	const now = new Date(); // get the current date and time
	const eighteenYearsAgo = new Date(now.getFullYear() - 18, now.getMonth(), now.getDate());

	const [formValues, setFormValues] = useState({
		birthdate: eighteenYearsAgo,
		type: 0,
		pokeball: 0,
		gender: "female",
	});
	const [pageIndex, setPageIndex] = useState(0);
	const monologues = [
		"Bienvenue sur Pokémeet ! Je suis le Professeur Chen, et je suis heureux de t'accueillir sur notre toute nouvelle plateforme de rencontre pour les fans de Pokémon. Avant de commencer, permet-moi de te poser quelques questions pour mieux te connaître.",
		"Enchanté <name>. Est-ce que tu te reconnais dans un des types feu, eau, plante, électrique ou peut-être quelque chose de plus rare ? Cette question nous permettra de mieux comprendre tes préférences en matière de temperement et de t'aider à trouver des PokeMeets en fonction de tes intérêts.",
		"Ensuite, quelle Pokeball correspond le mieux à tes attentes ? Cherches-tu une relation à long terme avec une Pokeball, un coup d'un soir avec une Rapidball, un plan cul régulier avec une Loveball ou simplement prendre un verre avec une Luxeball ?.",
		"Une fois que nous aurons compris tes préférences en matière de Pokémon et de Pokeball, peux tu me dire quelle sont les choses que tu aimes ou qui t'interesse?",
		"Nous pouvons passer à des questions plus personnelles. Par exemple, peux-tu nous dire qui tu es ? Es-tu un homme, une femme ou peut-être un dresseur non-binaire ? Cette question nous permettra de mieux comprendre ton identité de genre et de te trouver des matchs compatibles.",
		"Ensuite, quel genre de personne cherches-tu sur Pokémeet ? Préfères-tu les hommes, les femmes ou peut-être les deux ? Cette question nous aidera à comprendre tes préférences en matière de rencontres et de te trouver des correspondances qui répondent à tes attentes.",
		"Enfin, pour t'inscrire sur Pokémeet, tu devras nous donner ta date de naissance, ton adresse e-mail et un mot de passe contenant des chiffres et des lettres en majuscules et en minuscules. Il est important que tu sois majeur pour pouvoir utiliser notre plateforme, et nous voulons nous assurer que ton compte reste en sécurité avec un mot de passe solide.",
		"Merci de prendre le temps de répondre à ces questions, dresseur. Nous sommes impatients de t'aider à trouver ton poke-match parfait !",
	];
	const handlePrevious = () => {
		if (pageIndex !== 0) setPageIndex(pageIndex - 1);
	};
	const NB_PAGE = 6;

	const HandleChange = (key, value) => {
		console.log("ici");
		// let temp = formValues;
		// temp[key] = value;
		let temp = { ...formValues, [key]: value };
		setFormValues(temp);
		console.log(temp);
	};
	const HandleSubmit = (e) => {
		// e.stopPropagation();
		e.preventDefault();
		if (pageIndex < NB_PAGE - 1) {
			setPageIndex(pageIndex + 1);
		} else {
			if (formValues.password !== formValues.repassword) {
				const event = new CustomEvent("ProfChenMessage", {
					detail: { message: "Les mots de passe ne correspondent pas !" },
				});
				document.dispatchEvent(event);
				return false;
			}
			console.log("request");
		}
		return false;
	};
	const getFieldValue = (key, def) => {
		return formValues[key] ?? def;
	};
	return (
		<div className={style.SubscriptionForm}>
			<ProfChenMsg message={monologues[pageIndex].replace("<name>", formValues.firstname ?? "")} />
			<div className={style.line} />
			<form className={style.formPages} onSubmit={(event) => HandleSubmit(event)}>
				{pageIndex === 0 && (
					<NameAndBirthdate {...formValues} onChange={(key, value) => HandleChange(key, value)} />
				)}
				{pageIndex === 1 && <PokeType onChange={HandleChange} activeIndex={formValues.type} />}
				{pageIndex === 2 && <PokeballInput onChange={HandleChange} activeIndex={formValues.pokeball} />}
				{pageIndex === 3 && <Tags />}
				{pageIndex === 4 && <Gender onChange={HandleChange} value={formValues.gender} />}
				{pageIndex === 5 && <Attract onChange={HandleChange} value={formValues.attract} />}
				{pageIndex === 6 && (
					<Credentials
						onChange={HandleChange}
						password={formValues.password}
						repassword={formValues.repassword}
						email={formValues.email}
					/>
				)}
				<div className={style.pagination}>
					{pageIndex !== 0 && (
						<div className={style.previous} onClick={handlePrevious}>
							Precedent
						</div>
					)}
					<input
						type="submit"
						className={style.next}
						value={pageIndex !== NB_PAGE - 1 ? "Suivant" : "S'inscrire"}
					></input>
				</div>
			</form>
		</div>
	);
};

export default SubscriptionForm;

/*chatgpt question:
fais moi un dialog de jeux video du professeur chen qui poserait des question afin de guider un utilisateur sur le site de rencontre pokemeet, il voudrait savoir dans l'ordre, nom, son type(type pokemon), la pokeball qui , prenom,qui on est(homme, femme, bi), ce qu'on recherche (homme, femme, bi. choix multiple) , date de naissance(specifier qu'il doit etre majeur),
*/
/*
Bienvenue sur Pokémeet ! Je suis le Professeur Chen, et je suis heureux de t'accueillir sur notre toute nouvelle plateforme de rencontre pour les fans de Pokémon. Avant de commencer, permet-moi de te poser quelques questions pour mieux te connaître.

J'aurais besoin de quelques information ...

Tres bien nous pouvons commencer.

Tout d'abord, quel est ton type de Pokémon préféré ? Est-ce que tu aimes les Pokémon de type feu, eau, plante, électrique ou peut-être quelque chose de plus rare ? Cette question nous permettra de mieux comprendre tes préférences en matière de Pokémon et de t'aider à trouver des PokeMeets en fonction de tes intérêts.

Hmmmmm... Tres interessant :)

Ensuite, quelle Pokeball correspond le mieux à tes attentes ? Cherches-tu une relation à long terme avec une Pokeball, un coup d'un soir avec une Rapidball, un plan cul régulier avec une Loveball ou simplement prendre un verre avec une Luxeball ?.

TROUVER REP MARRANTE

Une fois que nous aurons compris tes préférences en matière de Pokémon et de Pokeball, nous pourrons passer à des questions plus personnelles. Par exemple, peux-tu nous dire qui tu es ? Es-tu un homme, une femme ou peut-être un dresseur non-binaire ? Cette question nous permettra de mieux comprendre ton identité de genre et de te trouver des matchs compatibles.

Ensuite, quel genre de personne cherches-tu sur Pokémeet ? Préfères-tu les hommes, les femmes ou peut-être les deux ? Cette question nous aidera à comprendre tes préférences en matière de rencontres et de te trouver des correspondances qui répondent à tes attentes.

Enfin, pour t'inscrire sur Pokémeet, tu devras nous donner ta date de naissance, ton adresse e-mail et un mot de passe contenant des chiffres et des lettres en majuscules et en minuscules. Il est important que tu sois majeur pour pouvoir utiliser notre plateforme, et nous voulons nous assurer que ton compte reste en sécurité avec un mot de passe solide.

Merci de prendre le temps de répondre à ces questions, dresseur. Nous sommes impatients de t'aider à trouver ton poke-match parfait !
*/
