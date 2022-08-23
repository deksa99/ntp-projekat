import React, { useEffect, useState } from "react";
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
import { useSelector } from "react-redux";
import { selectVehicles } from "./VehicleSlice";
import EditOutlined from "@mui/icons-material/EditOutlined";
import AddEditVehicleModal from "./AddEditVehicleModal";

const VehiclesTable = () => {
  const vehicles = useSelector(selectVehicles);
  const [open, setOpen] = useState(false);
  const [v, setV] = useState(null);

  const handleClose = () => {
    setOpen(false);
  };

  const handleEditClick = (v) => {
    setV(v);
  };

  useEffect(() => {
    if (v !== null) {
      setOpen(true);
    }
  }, [v]);

  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>
              <b>Proizvodjac</b>
            </TableCell>
            <TableCell>
              <b>Model</b>
            </TableCell>
            <TableCell>
              <b>Boja</b>
            </TableCell>
            <TableCell>
              <b>Registarski broj</b>
            </TableCell>
            <TableCell>
              <b>Broj sasije</b>
            </TableCell>
            <TableCell />
          </TableRow>
        </TableHead>
        <TableBody>
          {vehicles.map((c) => (
            <TableRow
              key={c.Id}
              sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
            >
              <TableCell align="left">{c.Manufacturer}</TableCell>
              <TableCell align="left">{c.CarModel}</TableCell>
              <TableCell align="left">{c.Color}</TableCell>
              <TableCell align="left">{c.LicencePlate}</TableCell>
              <TableCell align="left">{c.ChassisNumber}</TableCell>
              <TableCell align="right">
                <IconButton
                  aria-label="edit"
                  size="large"
                  color="secondary"
                  onClick={() => handleEditClick(c)}
                >
                  <EditOutlined fontSize="inherit" />
                </IconButton>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
      <AddEditVehicleModal open={open} handleClose={handleClose} v={v} />
    </TableContainer>
  );
};

export default VehiclesTable;
