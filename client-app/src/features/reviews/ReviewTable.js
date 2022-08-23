import React, { useEffect, useState } from "react";
import api from "../../api/api";
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
import DangerousOutlined from "@mui/icons-material/DangerousOutlined";

const ReviewTable = () => {
  const [reviews, setReviews] = useState([]);

  useEffect(() => {
    api
      .get("/reviews/reported")
      .then((res) => {
        setReviews(res.data);
      })
      .catch(() => {
        setReviews([]);
      });
  }, []);

  const handleInappropriate = (c) => {
    const body = {
      Inappropriate: true,
    };
    api
      .patch(`/reviews/report/${c.ID}/process`, body)
      .then((res) => {
        alert("Neprikladan");
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
              <b>Korisnik</b>
            </TableCell>
            <TableCell>
              <b>Auto servis</b>
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
              key={c.ID}
              sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
            >
              <TableCell align="left">
                {c.FirstName} {c.LastName}
              </TableCell>
              <TableCell align="left">{c.CarServiceName}</TableCell>
              <TableCell align="left">{c.Rating}</TableCell>
              <TableCell align="left">{c.Comment}</TableCell>
              {!c.Processed && (
                <TableCell align="right">
                  <IconButton
                    aria-label="delete"
                    size="large"
                    color="secondary"
                    onClick={() => handleInappropriate(c)}
                  >
                    <DangerousOutlined fontSize="inherit" />
                  </IconButton>
                </TableCell>
              )}
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default ReviewTable;
