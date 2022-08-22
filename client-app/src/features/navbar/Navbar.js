import React, { useEffect, useState } from "react";
import { useSelector } from "react-redux"

import { selectToken } from "../users/UsersSlice"
import { getRoleFromToken } from "../../util/JwtToken"

import UserNavbar from "./UserNavbar";
import AdminNavbar from "./AdminNavbar";
import WorkerNavbar from "./WorkerNavbar";
import DefaultNavbar from "./DefaultNavbar";

const Navbar = () => {
  const token = useSelector(selectToken);
  const [role, setRole] = useState(getRoleFromToken());
  useEffect(() => {
    setRole(getRoleFromToken())
  }, [token]);

  let navbar;
  switch (role) {
    case "user":
      navbar = <UserNavbar/>;
      break;
    case "admin":
      navbar = <AdminNavbar/>;
      break;
    case "worker":
      navbar = <WorkerNavbar/>;
      break;
    default:
      navbar = <DefaultNavbar/>;
  }
  return (
    <div>{navbar}</div>
  )
}

export default Navbar