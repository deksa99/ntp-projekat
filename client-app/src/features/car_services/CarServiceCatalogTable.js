import React, { useEffect, useState } from "react";
import api from "../../api/api";
import {
  TableCell,
  TableContainer,
  Paper,
  TableHead,
  TableRow,
  Table,
  TableBody,
} from "@mui/material";

const CarServiceCatalogTable = ({ cs }) => {
  const [catalog, setCatalog] = useState([]);

  useEffect(() => {
    api
      .get(`/car-services/${cs.Id}/catalog`)
      .then((res) => {
        setCatalog(res.data);
      })
      .catch(() => {
        setCatalog([]);
      });
  }, []);

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
            <TableCell />
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
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default CarServiceCatalogTable;
