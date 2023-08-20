// import { useState } from 'react';
import CardGrid from "../../components/CardGrid/CardGrid";
import style from "./style.module.scss";
//listing
const Home = () => {
  // const [users, setUsers] = useState({[]})
  return (
    <div className={style.Home}>
      <h1>Parcourir</h1>
      <div className={style.sorters}>
        <div className={style.left}></div>
        <div className={style.right}>
          <div className={style.sorters}></div>
          <CardGrid />
        </div>
      </div>
    </div>
  );
};

export default Home;
