import React from "react";
import { Outlet, Navigate } from "react-router-dom";
import "react-toastify/dist/ReactToastify.css";

const ProtectedRoute = ({ isAllowed, redirectPath = "/landing", children }) => {
  if (!isAllowed) {
    return <Navigate to={redirectPath} replace />;
  }

  return children ? children : <Outlet />;
};

export default ProtectedRoute;
