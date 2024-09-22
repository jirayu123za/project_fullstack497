// src/main.tsx
import React from "react";
import ReactDOM from "react-dom/client";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import App from "./App.tsx";
import LandingPage from "./pages/LandingPage.tsx";
import "./index.css";
import SignUpPage from "./pages/SignUpPage.tsx";
import InstructorDashboard from "./pages/InstructorDashboard.tsx";
import INS_Course from "./pages/INS_Course.tsx";
import INS_Create from "./pages/INS_Create.tsx";
import INS_Assignment from "./pages/INS_Assignment.tsx";

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
    path: "/course",
    element: <INS_Course />,
  },
  {
    path: "/create",
    element: <INS_Create />,
  },
  {
    path: "/course/:courseId",
    element: <INS_Course />,
  },
  {
    path: "/assignment",
    element: <INS_Assignment />,
  },
]);

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
