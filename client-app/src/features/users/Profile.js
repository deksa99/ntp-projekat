import React, { useEffect, useState } from "react";
import {
  TableCell,
  TableContainer,
  Paper,
  TableHead,
  TableRow,
  Table,
  TableBody,
} from "@mui/material";
import { getUserIdFromToken } from "../../util/JwtToken";
import api from "../../api/api";

const Profile = () => {
  const [user, setUser] = useState(null);

  useEffect(() => {
    const id = getUserIdFromToken();

    api
      .get(`/users/${id}`)
      .then((res) => {
        setUser(res.data);
      })
      .catch(() => {
        setUser([]);
      });
  }, []);

  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>
              <b>Ime</b>
            </TableCell>
            <TableCell>
              <b>Prezime</b>
            </TableCell>
            <TableCell>
              <b>Korisničko ime</b>
            </TableCell>
            <TableCell>
              <b>Email</b>
            </TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {!!user && (
            <TableRow key={user.Id}>
              <TableCell align="left">{user.FirstName}</TableCell>
              <TableCell align="left">{user.LastName}</TableCell>
              <TableCell align="left">{user.Username}</TableCell>
              <TableCell align="left">{user.Email}</TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default Profile;
