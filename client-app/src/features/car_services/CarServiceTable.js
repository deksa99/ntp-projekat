import React from "react";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import IconButton from "@mui/material/IconButton";
import InfoOutlined from "@mui/icons-material/InfoOutlined";

import { useSelector } from "react-redux";
import { selectCarServices } from "../car_services/CarServiceSlice";

const CarServiceTable = () => {
  const carServices = useSelector(selectCarServices);

  const handleInfoClick = (id) => {
    console.log(id);
  };

  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>
              <b>Naziv servisa</b>
            </TableCell>
            <TableCell>
              <b>Adresa</b>
            </TableCell>
            <TableCell />
          </TableRow>
        </TableHead>
        <TableBody>
          {carServices.map((cs) => (
            <TableRow
              key={cs.Id}
              sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
            >
              <TableCell align="left">{cs.Name}</TableCell>
              <TableCell align="left">{cs.Street}</TableCell>
              <TableCell align="right">
                <IconButton
                  aria-label="delete"
                  size="large"
                  color="secondary"
                  onClick={() => handleInfoClick(cs.Id)}
                >
                  <InfoOutlined fontSize="inherit" />
                </IconButton>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default CarServiceTable;
