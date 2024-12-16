import React from "react";
import { Navigate } from "react-router-dom";
import { useSelector } from "react-redux";
import { RootState } from "../store/store";
import { setUserToken } from "../store/authSlice";

type ProtectedRouteProps = {
  children: JSX.Element;
};

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ children }) => {
  const { userToken } = useSelector((state: RootState) => state.authSlice);

  if (!userToken) {
    if (!localStorage.getItem("token")) {
      return <Navigate to="/login" />;
    }
    setUserToken(localStorage.getItem("token"));
  }

  return children;
};

export default ProtectedRoute;