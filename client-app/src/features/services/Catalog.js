import React, { useState, useEffect } from "react";
import {
  TableCell,
  TableContainer,
  Paper,
  TableHead,
  TableRow,
  Table,
  TableBody,
  IconButton,
} from "@mui/material";
import api from "../../api/api";
import { getServiceFromToken } from "../../util/JwtToken";
import ChangeCircleOutlined from "@mui/icons-material/ChangeCircleOutlined";

const Catalog = () => {
  const [catalog, setCatalog] = useState([]);

  useEffect(() => {
    const id = getServiceFromToken();
    api
      .get(`/car-services/${id}/catalog`)
      .then((res) => {
        setCatalog(res.data);
      })
      .catch(() => {
        setCatalog([]);
      });
  }, []);

  const handleToggleClick = (c) => {
    api
      .patch(`/car-services/service/${c.Id}/change-availability`)
      .then((res) => {
        alert("Uspesno");
        window.location.reload();
      })
      .catch(() => {
        alert("Greska");
      });
  };

  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>
              <b>Naziv</b>
            </TableCell>
            <TableCell>
              <b>Opis</b>
            </TableCell>
            <TableCell>
              <b>Cena</b>
            </TableCell>
            <TableCell>
              <b>Dostupan</b>
            </TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {catalog.map((c) => (
            <TableRow
              key={c.Id}
              sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
            >
              <TableCell align="left">{c.Name}</TableCell>
              <TableCell align="left">{c.Description}</TableCell>
              <TableCell align="left">{c.Price}</TableCell>
              <TableCell align="left">{c.Available ? "da" : "ne"}</TableCell>
              <TableCell align="right">
                <IconButton
                  aria-label="delete"
                  size="large"
                  color="secondary"
                  onClick={() => handleToggleClick(c)}
                >
                  <ChangeCircleOutlined fontSize="inherit" />
                </IconButton>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default Catalog;
