DO $$
DECLARE
  prenom_list TEXT[] := ARRAY[
    'Alice', 'Bob', 'Charlie', 'David', 'Emma', 'Frank', 'Grace', 'Hannah', 'Isaac', 'Julia',
    'Kevin', 'Laura', 'Michael', 'Nora', 'Oliver', 'Penny', 'Quinn', 'Rachel', 'Samuel', 'Tara',
    'Ulysses', 'Vera', 'William', 'Xander', 'Yasmine', 'Zachary'
  ];
  nom_list TEXT[] := ARRAY[
    'Smith', 'Johnson', 'Williams', 'Jones', 'Brown', 'Davis', 'Miller', 'Wilson', 'Moore', 'Taylor',
    'Anderson', 'Thomas', 'Jackson', 'White', 'Harris', 'Martin', 'Thompson', 'Garcia', 'Martinez', 'Lopez',
    'Lee', 'Walker', 'Hall', 'Young', 'Allen', 'King', 'Wright', 'Scott', 'Green', 'Adams'
  ];
BEGIN
INSERT INTO users (fname, lname, email, birthdate, pass, bio, imageurl, age, gender, desiredgender, fame, tags, pokeball, type, userliked, likedfrom, seenfrom, blocked, convlist, coord, notifs, isactive, temp_token, login)
SELECT 
  COALESCE(prenom_list[(random() * array_length(prenom_list, 1))::integer], 'DefaultFirstName'),
  COALESCE(nom_list[(random() * array_length(nom_list, 1))::integer], 'DefaultFirstName'),
  'utilisateur' || generate_series(1, 100) || '@exemple.com',
   (CURRENT_DATE - INTERVAL '1 year' * ((random() * (66 - 18 + 1) + 18)::integer))::date,
  'motdepasse' || generate_series(1, 100),
  'Bio utilisateur' || generate_series(1, 100),
  ARRAY['url_image1', 'url_image2', 'url_image3'],
  (random() * (66 - 18 + 1) + 18)::integer, -- L'âge varie entre 25 et 35
  (random() * 2)::integer, -- Le genre varie entre 0 et 1
  ARRAY[(random() * 2)::integer], -- Le genre désiré varie entre 0 et 1
  0, -- La renommée initiale est définie à 0
  ARRAY[(random() * 10)::integer, (random() * 10)::integer], -- Les tags varient entre 0 et 9
  ARRAY[(random() * 10)::integer], -- Les pokeballs varient entre 0 et 9
  ARRAY[(random() * 10)::integer], -- Les types varient entre 0 et 9
  ARRAY[]::integer[], -- Aucun utilisateur n'a encore été aimé
  ARRAY[]::integer[], -- Aucun utilisateur n'a encore été aimé par d'autres
  ARRAY[]::integer[], -- Aucun utilisateur n'a encore été vu
  ARRAY[]::integer[], -- Aucun utilisateur n'est bloqué
  ARRAY[]::integer[], -- Aucune conversation pour le moment
  '{"lat": 0, "long": 0}', -- Coordonnées de base
  '{}', -- Aucune notification
  true, -- L'utilisateur est actif
  '', -- Aucun jeton temporaire
  '' || generate_series(1, 100) -- Nom d'utilisateur
FROM generate_series(1, 100);
END $$;