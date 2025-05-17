import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import App from "./App.tsx";
import UserList from "./pages/user.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <App />
    <UserList />
  </StrictMode>
);
