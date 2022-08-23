import React from "react";
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
} from "react-router-dom";

import { selectToken } from "./features/users/UsersSlice";
import { getRoleFromToken } from "./util/JwtToken";

import ProtectedRoute from "./util/ProtectedRoute";
import Navbar from "./features/navbar/Navbar";

import HomePage from "./features/pages/HomePage";
import LoginPage from "./features/pages/LoginPage";
import RegistrationPage from "./features/pages/RegistrationPage";
import ServiceInfoPage from "./features/pages/ServiceInfoPage";
import AppointmentsPage from "./features/pages/AppointmentsPage";
import ProfilePage from "./features/pages/ProfilePage";

import { useEffect, useState } from "react";
import { useSelector } from "react-redux";

function App() {
  const token = useSelector(selectToken);
  const [role, setRole] = useState(getRoleFromToken());
  useEffect(() => {
    setRole(getRoleFromToken());
  }, [token]);

  return (
    <Router>
      <Navbar />
      <Routes>
        <Route index element={<HomePage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/sign-up" element={<RegistrationPage />} />
        <Route path="/services" element={<ServiceInfoPage />} />
        <Route
          element={
            <ProtectedRoute isAllowed={!!role && role.includes("user")} />
          }
        >
          <Route path="/appointments" element={<AppointmentsPage />} />
          <Route path="/profile" element={<ProfilePage />} />
        </Route>
        <Route path="*" element={<Navigate to="/" />} />
      </Routes>
    </Router>
  );
}

export default App;
