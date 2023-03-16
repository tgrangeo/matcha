import "./App.scss";
import React from "react";
import Footer from "./components/Footer/Footer";
import Header from "./components/Header/Header";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import NotFound from "./pages/404/404";
import Home from "./pages/Home/Home";
import Login from "./pages/Login/Login";
import Profile from "./pages/Profile/Profile";
import Subscription from "./pages/Subscription/Subscription";
const App = () => {
    const router = createBrowserRouter([
        {
            path: "*",
            element: <NotFound />,
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
