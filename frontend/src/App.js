import "./App.scss";
import React, { useEffect, useState } from "react";
import Footer from "./components/Footer/Footer";
import Header from "./components/Header/Header";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import NotFound from "./pages/404/404";
import Home from "./pages/Home/Home";
import Login from "./pages/Login/Login";
import Profile from "./pages/Profile/Profile";
import Subscription from "./pages/Subscription/Subscription";

const Test = () => {
    useEffect(() => {
        const token = window.location.href.split("=")[1];
        fetch("/api/v1/validate", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ token }),
        })
            .then((res) => res.json())
            .then((data) => console.log(data));
    }, []);
    return <div>validating...</div>;
};

const ResetPass = () => {
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const [token, setToken] = useState("");

    useEffect(() => {
        const token = window.location.href.split("=")[1];
        setToken(token);
    }, []);

    const handleSubmit = () => {
        // Vérifiez si les mots de passe correspondent
        if (password !== confirmPassword) {
            console.log("Les mots de passe ne correspondent pas");
            return;
        }

        // Effectuez la requête de réinitialisation de mot de passe
        fetch("/api/v1/newpass", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ token, password }),
        })
            .then((res) => res.json())
            .then((data) => console.log(data));
    };

    return (
        <div>
            <h1>Réinitialisation de mot de passe</h1>
            <input
                type="password"
                placeholder="Nouveau mot de passe"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
            />
            <br />
            <input
                type="password"
                placeholder="Confirmer le nouveau mot de passe"
                value={confirmPassword}
                onChange={(e) => setConfirmPassword(e.target.value)}
            />
            <br />
            <button onClick={handleSubmit}>Réinitialiser le mot de passe</button>
        </div>
    );
};



const App = () => {
    const router = createBrowserRouter([
        {
            path: "*",
            element: <NotFound />,
        },
        {
            path: "/validate",
            element: <Test />
        },
        {
            path: "/resetpass",
            element: <ResetPass />
        },
        {
            path: "/",
            element: <Home />,
        },
        {
            path: "/login",
            element: <Login />,
        },
        {
            path: "/profile",
            element: <Profile />,
        },
        {
            path: "/subscription",
            element: <Subscription />,
        },
    ]);

    return (
        <React.StrictMode>
            <div className="App">
                <Header />
                <section className="content">
                    <RouterProvider router={router} />
                </section>
                <Footer />
            </div>
        </React.StrictMode>
    );
};

export default App;
