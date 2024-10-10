import React from "react";
import ReactDOM from "react-dom/client";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import App from "./App.tsx";
import LandingPage from "./pages/LandingPage.tsx";
import "./index.css";
import SignUpPage from "./pages/SignUpPage.tsx";
import INS_Dashboard from "./pages/INS_Dashboard.tsx";
import INS_Course from "./pages/INS_Course.tsx";
import INS_Create from "./pages/INS_Create.tsx";
import INS_Assignment from "./pages/INS_Assignment.tsx";
import STD_Assignment from "./pages/STD_Assignment.tsx";
import STD_Dashboard from "./pages/STD_Dashboard.tsx";
import STD_Course from "./pages/STD_Course.tsx";
import STD_Join from "./pages/STD_Join.tsx";
import AuthRoute from "./middleware/AuthRoute.tsx";

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
    path: "/insdash",
    element: (
      <AuthRoute>
        <INS_Dashboard />
      </AuthRoute>
    ),
  },
  {
    path: "/create",
    element: (
      <AuthRoute>
        <INS_Create />
      </AuthRoute>
    ),
  },
  {
    path: "/course/:course_id",
    element: (
      <AuthRoute>
        <INS_Course />
      </AuthRoute>
    ),
  },
  {
    path: "/course/:course_id/assignment/:assignment_id",
    element: (
      <AuthRoute>
        <INS_Assignment />
      </AuthRoute>
    ),
  },
  {
    path: "/stddash",
    element: <STD_Dashboard />,
  },
  {
    path: "/stdcourse/:courseId",
    element: <STD_Course />,
  },
  {
    path: "/stdassignment/:assignmentId",
    element: <STD_Assignment />,
  },
  {
    path: "/stdcreate",
    element: <STD_Join />,
  },
]);

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
