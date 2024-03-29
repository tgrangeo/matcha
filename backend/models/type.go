package models

var Tags []string = []string{
	"Amour",
	"Amitié",
	"Romance",
	"Relation sérieuse",
	"Relation occasionnelle",
	"Sorties",
	"Voyages",
	"Sport",
	"Cinéma",
	"Musique",
	"Littérature",
	"Art",
	"Danse",
	"Théâtre",
	"Cuisine",
	"Nature",
	"Animaux",
	"Photographie",
	"Mode",
	"Shopping",
	"Technologie",
	"Science",
	"Éducation",
	"Spiritualité",
	"Voyance",
	"Astrologie",
	"Méditation",
	"Yoga",
	"Écologie",
	"Politique",
	"Actualités",
	"Histoire",
	"Gastronomie",
	"Végétarien",
	"Voyage",
	"Aventure",
	"Plage",
	"Montagne",
	"Randonnée",
	"Camping",
	"Pêche",
	"Chasse",
	"Plongée",
	"Ski",
	"Snowboard",
	"Surf",
	"Escalade",
	"Cyclisme",
	"Natation",
}

// import "fmt"

// type Type struct{
// 	id int `json:"id"`
//     Name  string  `json:"name"`
// 	Desc  string  `json:"desc"`
// }

// type TypeStruct struct {
// 	Feu Type;
// 	Normal Type;
// 	Eau Type;
// 	Plante Type;
// 	Electric Type;
// 	Glace Type;
// 	Combat Type;
// 	Poison Type;
// 	Sol Type;
// 	Vol Type;
// 	Psy Type;
// 	Insecte Type;
// 	Roche Type;
// 	Spectre Type;
// 	Dragon Type;
// 	Tenebre Type;
// 	Acier Type;
// 	Fée Type;
// }

// var TypeList = TypeStruct{
// 	Feu:Type{"Feu", "Les individus de type Feu ont tendance à être passionnés et impulsifs. Ils peuvent être très énergiques et ont souvent un fort désir de réussite. Cependant, ils peuvent aussi être enclins à la colère et à l'agressivité."},
// 	Normal :Type{"Normal", "Les individus de type Normal sont souvent des personnes stables et équilibrées. Ils sont généralement amicaux et sociables, et peuvent être considérés comme des personnes \"normales\" ou \"ordinaires\"."},
// 	Eau :Type{"Eau", "Les individus de type Eau sont souvent calmes et paisibles. Ils ont un fort sens de l'empathie et sont souvent très émotionnels. Ils peuvent être considérés comme des personnes sensibles et intuitives."},
// 	Plante :Type{"Plante", " Les individus de type Plante ont tendance à être calmes et réfléchis. Ils sont souvent très concentrés sur leur environnement et leur santé physique et mentale. Ils peuvent être considérés comme des personnes calmes et terre-à-terre."},
// 	Electric :Type{"Glace", " Les individus de type Électrique sont souvent très énergiques et pleins de vie. Ils peuvent être considérés comme des personnes dynamiques et audacieuses, mais peuvent aussi être enclins à l'impatience et à l'impulsivité."},
// 	Glace :Type{"Glace", "Les individus de type Glace sont souvent très réservés et introvertis. Ils peuvent être considérés comme des personnes froides ou insensibles, mais en réalité, ils sont souvent très prévenants et réfléchis. Ils ont souvent un fort sens de la discipline et de l'autodiscipline."},
// 	Combat :Type{"Combat", "Les individus de type Combat sont souvent très déterminés et compétitifs. Ils peuvent être considérés comme des personnes agressives et combatives, mais en réalité, ils ont souvent un fort sens de l'honneur et de la justice. Ils peuvent être très passionnés dans leur quête pour atteindre leurs objectifs."},
// 	Poison :Type{"Poison", "Les individus de type Poison peuvent être considérés comme des personnes sournoises et malveillantes. Cependant, ils ont souvent un esprit vif et sont très habiles à naviguer dans des situations difficiles. Ils ont souvent un sens aigu de la méfiance et de la prudence."},
// 	Sol :Type{"Sol", "Les individus de type Sol ont tendance à être très stables et résistants. Ils peuvent être considérés comme des personnes pratiques et résolues, avec un fort sens de la responsabilité et de l'engagement. Ils ont souvent une grande confiance en eux-mêmes et peuvent être très persévérants."},
// 	Vol :Type{"Vol", "Les individus de type Vol ont tendance à être très libres d'esprit et à avoir un sens de l'aventure. Ils peuvent être considérés comme des personnes imprévisibles et souvent changeantes, mais en réalité, ils ont souvent un fort désir de liberté et d'indépendance."},
// 	Psy :Type{"Psy", "Comportement humain mystérieux et profond"},
// 	Insecte :Type{"Insecte", "Les individus de type Insecte ont souvent un esprit très travailleur et sont très persévérants. Ils peuvent être considérés comme des personnes minutieuses et méticuleuses, qui travaillent dur pour atteindre leurs objectifs. Ils ont souvent un fort sens de la communauté et peuvent être très sociables."},
// 	Roche :Type{"Roche", " Les individus de type Roche ont tendance à être très stables et résistants. Ils peuvent être considérés comme des personnes pratiques et résolues, avec un fort sens de la responsabilité et de l'engagement. Ils ont souvent une grande confiance en eux-mêmes et peuvent être très persévérants."},
// 	Spectre :Type{"Spectre", "Les individus de type Spectre peuvent être considérés comme des personnes mystérieuses et énigmatiques. Ils ont souvent un esprit très créatif et sont très imaginatifs. Ils peuvent être très réservés et introvertis, mais en réalité, ils ont souvent un grand sens de l'empathie et sont très prévenants envers les autres."},
// 	Dragon :Type{"Dragon", "Les individus de type Dragon ont souvent un fort désir de succès et de puissance. Ils peuvent être considérés comme des personnes ambitieuses et audacieuses, avec un fort désir de réussite. Ils ont souvent un esprit très compétitif et peuvent être très déterminés dans leur quête pour atteindre leurs objectifs."},
// 	Tenebre :Type{"Ténèbres", "Les individus de type Ténèbres peuvent être considérés comme des personnes sombres et mystérieuses. Ils ont souvent un fort sens de l'individualité et de l'indépendance, mais peuvent aussi être enclins à la méfiance et à la méchanceté. Ils peuvent être très réservés et introvertis, mais en réalité, ils ont souvent un grand sens de l'empathie et sont très prévenants envers les autres."},
// 	Acier :Type{"Acier", "Les individus de type Acier peuvent être considérés comme des personnes très résistantes et durables. Ils ont souvent un fort sens de la loyauté et de l'honneur, et sont très fiables. Ils peuvent être très persévérants et sont souvent très habiles à naviguer dans des situations difficiles."},
// 	Fée :Type{"Fée", "Le type Fée est connu pour être extrêmement bienveillant et empathique envers les autres. Les Pokémon de type Fée ont également une nature protectrice et sont souvent décrits comme ayant des personnalités douces et aimantes. Enfin, ils sont également associés à la magie et à la fantaisie, leur donnant un côté mystérieux et fascinant."},
// }

// type PokeballStruct struct{
// 	Loveball Type;
// 	Masterball Type;
// 	Rapideball Type;
// 	Pokeball Type;
// }

// var Pokeballs = PokeballStruct{
// 	Loveball: Type{"Loveball", "cherche un plan q regulier"},
// 	Masterball: Type{"Masterball", "cherche une relation serieuse"},
// 	Rapideball: Type{"Rapideball", "cherche un coup d'un soir"},
// 	Pokeball: Type{"Pokeball", "cherche une personne pour boire un verre"},
// }

// var Tags = [100]string{
// 	"Voyage",
// 	"Photographie",
// 	"Musique",
// 	"Danse",
// 	"Art",
// 	"Cinéma",
// 	"Lecture",
// 	"Écriture",
// 	"Théâtre",
// 	"Mode",
// 	"Fitness",
// 	"Yoga",
// 	"Nutrition",
// 	"Cuisine",
// 	"Vin",
// 	"Bière artisanale",
// 	"Cocktails",
// 	"Spiritueux",
// 	"Café",
// 	"Thé",
// 	"Jardinage",
// 	"Randonnée",
// 	"Camping",
// 	"Pêche",
// 	"Chasse",
// 	"Sports nautiques",
// 	"Sports de montagne",
// 	"Sports de raquette",
// 	"Sports d'équipe",
// 	"Sports de combat",
// 	"Équitation",
// 	"Animaux de compagnie",
// 	"Volontariat",
// 	"Activisme",
// 	"Politique",
// 	"Environnement",
// 	"Science",
// 	"Technologie",
// 	"Entrepreneuriat",
// 	"Éducation",
// 	"Langues étrangères",
// 	"Histoire",
// 	"Mythologie",
// 	"Philosophie",
// 	"Psychologie",
// 	"Spiritualité",
// 	"Voyance",
// 	"Astrologie",
// 	"Tarot",
// 	"Méditation",
// 	"Religion",
// 	"Musées",
// 	"Galeries d'art",
// 	"Festivals",
// 	"Concerts",
// 	"Spectacles comiques",
// 	"Opéra",
// 	"Ballet",
// 	"Cirque",
// 	"Parcs d'attractions",
// 	"Plages",
// 	"Montagnes",
// 	"Villes historiques",
// 	"Châteaux",
// 	"Palais",
// 	"Édifices religieux",
// 	"Marchés aux puces",
// 	"Antiquités",
// 	"Artisanat",
// 	"Design d'intérieur",
// 	"Architecture",
// 	"Photographie de paysage",
// 	"Photographie de rue",
// 	"Photographie de portrait",
// 	"Écriture de fiction",
// 	"Écriture de non-fiction",
// 	"Poésie",
// 	"Nouvelles technologies",
// 	"Programmation informatique",
// 	"Jeux vidéo",
// 	"Cinéma d'horreur",
// 	"Cinéma indépendant",
// 	"Comédies romantiques",
// 	"Films d'action",
// 	"Films d'animation",
// 	"Séries télévisées",
// 	"Documentaires",
// 	"Fantasy",
// 	"Science-fiction",
// 	"Mystères",
// 	"Suspense",
// 	"Crime",
// 	"Romance",
// 	"Jeux de société",
// 	"Jeux de rôle",
// 	"Jeux de cartes",
// 	"Jeux de stratégie",
// 	"Jeux de hasard",
// 	"Décoration d'intérieur",
// 	"DIY (bricolage)",
// }
