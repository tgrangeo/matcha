import { useState } from "react";
import style from "./style.module.scss";
import DatePicker from "react-datepicker";

import "react-datepicker/dist/react-datepicker.css";
import Input from "../../../Input/Input";

const NameAndBirthdate = ({ firstname, lastname, birthdate, onChange }) => {
  const handleDateChange = (date) => {
    const age = calculateAge(date);
    if (age < 18) {
      // alert("You must be of majority age to use this service.");
      const event = new CustomEvent("ProfChenMessage", {
        detail: { message: "Tu as besoin d'etre majeur pour t'inscrire ici!" },
      });
      document.dispatchEvent(event);
      return;
    }
    onChange("birthdate", date);
  };

  const calculateAge = (birthday) => {
    const today = new Date();
    let age = today.getFullYear() - birthday.getFullYear();
    const monthDifference = today.getMonth() - birthday.getMonth();
    if (
      monthDifference < 0 ||
      (monthDifference === 0 && today.getDate() < birthday.getDate())
    ) {
      age--;
    }
    return age;
  };
  // console.log(inputs["birthdate"]);
  return (
    <div className={style.NameAndBirthdate}>
      <div>
        <Input
          type="text"
          name="lastname"
          required
          value={lastname}
          placeholder="Entrer votre nom..."
          onChange={(e) => onChange("lastname", e.target.value)}
        >
          Nom
        </Input>

        <Input
          type="text"
          name="firstname"
          required
          value={firstname}
          placeholder="Entrer votre prenom..."
          onChange={(e) => onChange("firstname", e.target.value)}
        >
          Prenom
        </Input>
      </div>

      <p className={style.datePicker}>
        <label htmlFor="date">Date de naissance</label>
        <DatePicker
          required
          value={birthdate}
          selected={birthdate}
          onChange={handleDateChange}
          id="date"
        />
      </p>
    </div>
  );
};

export default NameAndBirthdate;
