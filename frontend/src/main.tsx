// src/main.tsx
import React from "react";
import ReactDOM from "react-dom/client";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import App from "./App.tsx";
import LandingPage from "./pages/LandingPage.tsx";
import "./index.css";
import SignUpPage from "./pages/SignUpPage.tsx";
import InstructorDashboard from "./pages/InstructorDashboard.tsx";
import InstructorAssignment from "./pages/INS_Assignment.tsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
  },
  {
    path: "/landing",
    element: <LandingPage />,
  },
  {
    path: "/signup",
    element: <SignUpPage />,
  },
  {
    path: "/dashboard",
    element: <InstructorDashboard />,
  },
  {
    path: "/assignment",
    element: <InstructorAssignment />,
  },
]);

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
