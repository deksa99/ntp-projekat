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

const CarServiceReviewsTable = ({ cs }) => {
  const [reviews, setReviews] = useState([]);

  useEffect(() => {
    api
      .get(`/reviews/car-service/${cs.Id}`)
      .then((res) => {
        setReviews(res.data);
      })
      .catch(() => {
        setReviews([]);
      });
  }, []);

  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>
              <b>Usluga</b>
            </TableCell>
            <TableCell>
              <b>Ocena</b>
            </TableCell>
            <TableCell>
              <b>Komentar</b>
            </TableCell>
            <TableCell />
          </TableRow>
        </TableHead>
        <TableBody>
          {reviews.map((c) => (
            <TableRow
              key={c.Id}
              sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
            >
              <TableCell align="left">{c.ServiceName}</TableCell>
              <TableCell align="left">{c.Rating}</TableCell>
              <TableCell align="left">{c.Comment}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default CarServiceReviewsTable;
