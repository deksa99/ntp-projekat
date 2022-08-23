import React from "react";
import {
  TableCell,
  TableContainer,
  Paper,
  TableHead,
  TableRow,
  Table,
  TableBody,
} from "@mui/material";
import { useSelector } from "react-redux";
import { selectRequests } from "./RequestsSlice";
import RequestOptions from "./RequestOptions";

const RequestsTable = () => {
  const requests = useSelector(selectRequests);

  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>
              <b>Auto Servis</b>
            </TableCell>
            <TableCell>
              <b>Usluga</b>
            </TableCell>
            <TableCell>
              <b>Vozilo</b>
            </TableCell>
            <TableCell>
              <b>Cena</b>
            </TableCell>
            <TableCell>
              <b>Poslato</b>
            </TableCell>
            <TableCell>
              <b>Status</b>
            </TableCell>
            <TableCell />
          </TableRow>
        </TableHead>
        <TableBody>
          {requests.map((c) => (
            <TableRow
              key={c.Id}
              sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
            >
              <TableCell align="left">{`${c.CarServiceName} - ${c.Street}`}</TableCell>
              <TableCell align="left">{`${c.ServiceName}`}</TableCell>
              <TableCell align="left">{`${c.Manufacturer} ${c.CarModel}`}</TableCell>
              <TableCell align="left">{c.ServicePrice}</TableCell>
              <TableCell align="left">{c.SubmittedAt}</TableCell>
              <TableCell align="left">{c.Status}</TableCell>
              <TableCell align="right">
                <RequestOptions r={c} />
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default RequestsTable;
