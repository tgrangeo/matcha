import { useState } from "react";
import style from "./style.module.scss";
import ProfChenMsg from "../ProfChenMsg/ProfChenMsg";

const SubscriptionForm = () => {
	const [open, setOpen] = useState(0);
	const [pageIndex, setPageIndex] = useState(0);
	const monologues = [
		"Bienvenue sur Pokémeet ! Je suis le Professeur Chen, et je suis heureux de t'accueillir sur notre toute nouvelle plateforme de rencontre pour les fans de Pokémon. Avant de commencer, permet-moi de te poser quelques questions pour mieux te connaître.",
		"J'aurais besoin de quelques information ...",
		"Tres bien nous pouvons commencer.",
		"Tout d'abord, quel est ton type de Pokémon préféré ? Est-ce que tu aimes les Pokémon de type feu, eau, plante, électrique ou peut-être quelque chose de plus rare ? Cette question nous permettra de mieux comprendre tes préférences en matière de Pokémon et de t'aider à trouver des PokeMeets en fonction de tes intérêts.",
		"Hmmmmm... Tres interessant :)",
		"Ensuite, quelle Pokeball correspond le mieux à tes attentes ? Cherches-tu une relation à long terme avec une Pokeball, un coup d'un soir avec une Rapidball, un plan cul régulier avec une Loveball ou simplement prendre un verre avec une Luxeball ?.",
		"Une fois que nous aurons compris tes préférences en matière de Pokémon et de Pokeball, nous pourrons passer à des questions plus personnelles. Par exemple, peux-tu nous dire qui tu es ? Es-tu un homme, une femme ou peut-être un dresseur non-binaire ? Cette question nous permettra de mieux comprendre ton identité de genre et de te trouver des matchs compatibles.",
		"Ensuite, quel genre de personne cherches-tu sur Pokémeet ? Préfères-tu les hommes, les femmes ou peut-être les deux ? Cette question nous aidera à comprendre tes préférences en matière de rencontres et de te trouver des correspondances qui répondent à tes attentes.",
		"Enfin, pour t'inscrire sur Pokémeet, tu devras nous donner ta date de naissance, ton adresse e-mail et un mot de passe contenant des chiffres et des lettres en majuscules et en minuscules. Il est important que tu sois majeur pour pouvoir utiliser notre plateforme, et nous voulons nous assurer que ton compte reste en sécurité avec un mot de passe solide.",
		"Merci de prendre le temps de répondre à ces questions, dresseur. Nous sommes impatients de t'aider à trouver ton poke-match parfait !",
	];

	const pages = [
		<>
			<div>
				<p>
					<label htmlFor="nom">Nom</label>
					<input type="text" id="nom" required></input>
				</p>
				<p>
					<label htmlFor="prenom">Prenom</label>
					<input type="text" id="prenom" required></input>
				</p>
			</div>
			{/* <p>
				<label htmlFor="Nom">Description</label>
				<textarea type="textarea" id="nom" required></textarea>
			</p> */}
		</>,
	];
	return (
		<form className={style.SubscriptionForm}>
			<ProfChenMsg
				message={monologues[open]}
				onClick={() => setOpen(open + 1)}
				onAnimEnd={() => setPageIndex(pageIndex + 1)}
			/>
			<div className={style.line} />
			<div className={style.formPages}>{pages[pageIndex] ?? null}</div>
		</form>
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
