import { useState, useEffect } from "react";
import Cookies from 'js-cookie';
import style from "./style.module.scss";
import KeyboardArrowDownSharpIcon from "@mui/icons-material/KeyboardArrowDownSharp";
const HeaderProfile = ({ onLogin }) => {
    const [logged, setLogged] = useState(false);
    const [open, setOpen] = useState(false);
    const user = {
        firstname: "Eliott",
        lastname: "Depauw",
        fame: 75,
        profilePictureUrl:
            "https://upload.wikimedia.org/wikipedia/commons/5/57/Chicken_-_melbourne_show_2005.jpg?uselang=fr",
    };

    useEffect(() => {
        const granola = Cookies.get("matcha")
        console.log(granola)
        if (granola)
            setLogged(true)
    }, []);

    const login = () => {
        const granola = Cookies.get("matcha")
        console.log(granola)
        if (granola)
            setLogged(true)
        if (logged === false) {
            window.location.href = "http://localhost:3000/login";
        }
    };


    return logged ? (
        <div className={style.headerProfile} onClick={() => setOpen(!open)}>
            <div className={style.left}>
                <p className={style.name}>
                    {user.firstname + " "}
                    {user.lastname}
                </p>
                <div className={style.fameContainer}>
                    <p className={style.fame}>fame</p>
                    <p className={style.fameValue}>{user.fame}</p>
                </div>
            </div>
            <img className={style.logo} src={user.profilePictureUrl} alt="pp" />
            <KeyboardArrowDownSharpIcon />
        </div>
    ) : (
        <div className={style.cta} onClick={login}>
            Se connecter
        </div>
    );
};

export default HeaderProfile;
