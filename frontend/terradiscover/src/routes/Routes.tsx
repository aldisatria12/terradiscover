import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { LoginPage } from "../pages/Login/LoginPage";
import { RegisterPage } from "../pages/Register/RegisterPage";
import ProtectedRoute from "./ProtectedRoute";
import { ContactPage } from "../pages/ContactPage";

const routes = createBrowserRouter([
  {
    path: "/",
    element: (
      <ProtectedRoute>
        <ContactPage />
      </ProtectedRoute>
    ),
  },
  {
    path: "/login",
    element: <LoginPage />,
  },
  {
    path: "/register",
    element: <RegisterPage />,
  },
  {
    path: "/contact",
    element: (
      <ProtectedRoute>
        <ContactPage />
      </ProtectedRoute>
    ),
  }
]);

const AppRouter: React.FC = () => <RouterProvider router={routes} />;

export default AppRouter;
