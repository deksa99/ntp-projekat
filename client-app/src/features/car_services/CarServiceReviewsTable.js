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
  IconButton,
} from "@mui/material";
import { useSelector } from "react-redux";
import ErrorOutlineOutlined from "@mui/icons-material/ErrorOutlineOutlined";
import { getRoleFromToken } from "../../util/JwtToken";
import { selectToken } from "../users/UsersSlice";

const CarServiceReviewsTable = ({ cs }) => {
  const token = useSelector(selectToken);
  const [role, setRole] = useState(getRoleFromToken());
  useEffect(() => {
    setRole(getRoleFromToken());
  }, [token]);

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

  const handleReportClick = (c) => {
    console.log(c);
    api
      .patch(`/reviews/${c.ID}/report`, {})
      .then((res) => {
        alert("Prijavljeno");
      })
      .catch(() => {
        alert(!!token ? "Vec prijavljeno" : "Prijavite se");
      });
  };

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
              sx={{
                "&:last-child td, &:last-child th": { border: 0 },
                bgcolor: c.Inappropriate ? "red" : "default",
              }}
            >
              <TableCell align="left">{c.ServiceName}</TableCell>
              <TableCell align="left">{c.Rating}</TableCell>
              <TableCell align="left">{c.Comment}</TableCell>
              <TableCell align="right">
                <IconButton
                  aria-label="delete"
                  size="large"
                  color="secondary"
                  onClick={() => handleReportClick(c)}
                >
                  <ErrorOutlineOutlined fontSize="inherit" />
                </IconButton>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default CarServiceReviewsTable;
