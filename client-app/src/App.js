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
import AppointmentsWorkerPage from "./features/pages/AppointmentsWorkerPage";
import ProfilePage from "./features/pages/ProfilePage";
import RequestsPage from "./features/pages/RequestsPage";
import RequestsWorkerPage from "./features/pages/RequestsWorkerPage";

import { useEffect, useState } from "react";
import { useSelector } from "react-redux";
import VehiclesPage from "./features/pages/VehiclesPage";
import CatalogPage from "./features/pages/CatalogPage";
import UsersPage from "./features/pages/UsersPage";
import ReviewPage from "./features/pages/ReviewPage";

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
          <Route path="/profile" element={<ProfilePage />} />
          <Route path="/vehicles" element={<VehiclesPage />} />
          <Route path="/appointments" element={<AppointmentsPage />} />
          <Route path="/requests" element={<RequestsPage />} />
        </Route>
        <Route
          element={
            <ProtectedRoute isAllowed={!!role && role.includes("worker")} />
          }
        >
          <Route
            path="/appointments-worker"
            element={<AppointmentsWorkerPage />}
          />
          <Route path="/requests-worker" element={<RequestsWorkerPage />} />
          <Route path="/catalog" element={<CatalogPage />} />
        </Route>
        <Route
          element={
            <ProtectedRoute isAllowed={!!role && role.includes("admin")} />
          }
        >
          <Route path="/users" element={<UsersPage />} />
          <Route path="/reviews" element={<ReviewPage />} />
        </Route>
        <Route path="*" element={<Navigate to="/" />} />
      </Routes>
    </Router>
  );
}

export default App;
