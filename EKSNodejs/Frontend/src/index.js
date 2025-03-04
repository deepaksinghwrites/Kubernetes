import { createRoot } from "react-dom/client";
import React from "react";
import "./index.css";
import App from "./App";
import { Paper, TextField, Checkbox, Button } from "@mui/material";

const root = createRoot(document.getElementById("root"));
root.render(
    <React.StrictMode>
        <div className="app-wrapper">
            <App />
        </div>
    </React.StrictMode>
);
