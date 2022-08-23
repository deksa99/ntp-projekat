import React from "react";
import { useSelector } from "react-redux";
import { selectUsers } from "./UsersSlice";

import BlockOutlined from "@mui/icons-material/BlockOutlined";
import {
  IconButton,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from "@mui/material";
import api from "../../api/api";

const UsersTable = () => {
  const users = useSelector(selectUsers);

  const handleBlockClick = (c) => {
    api
      .patch(`/users/${c.AccountId}/block`, {})
      .then(() => {
        alert("Blokiran na 3 dana");
      })
      .catch(() => {
        alert("Greskas");
      });
  };

  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>
              <b>Email</b>
            </TableCell>
            <TableCell>
              <b>Username</b>
            </TableCell>
            <TableCell>
              <b>FirstName</b>
            </TableCell>
            <TableCell>
              <b>LastName</b>
            </TableCell>
            <TableCell />
          </TableRow>
        </TableHead>
        <TableBody>
          {users.map((c) => (
            <TableRow
              key={c.Id}
              sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
            >
              <TableCell align="left">{c.Email}</TableCell>
              <TableCell align="left">{c.Username}</TableCell>
              <TableCell align="left">{c.FirstName}</TableCell>
              <TableCell align="left">{c.LastName}</TableCell>
              <TableCell align="right">
                <IconButton
                  aria-label="edit"
                  size="large"
                  color="secondary"
                  onClick={() => handleBlockClick(c)}
                >
                  <BlockOutlined fontSize="inherit" />
                </IconButton>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default UsersTable;
