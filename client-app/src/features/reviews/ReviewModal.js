import { Box, Button, Grid, Modal, TextField, Typography } from "@mui/material";
import api from "../../api/api";
import React, { useState } from "react";

const style = {
  position: "absolute",
  left: "50%",
  transform: "translate(-50%, 0%)",
  width: 860,
  bgcolor: "background.paper",
  boxShadow: 12,
  p: 4,
  m: 0,
};

const ReviewModal = ({ id, open, handleClose }) => {
  const [rating, setRating] = useState(0);
  const [comment, setComment] = useState("");

  const handleSubmit = (event) => {
    event.preventDefault();
    const body = {
      AppointmentID: id,
      Rating: parseInt(rating),
      Comment: comment,
    };
    api
      .post("/reviews/add", body)
      .then(() => {
        alert("Uspesno dodato");
      })
      .catch(() => {
        alert("Vec postoji komentar");
      });
  };

  return (
    <Modal open={open} onClose={handleClose}>
      <Box sx={style}>
        <Box>
          <Typography variant="h5" component="h5" sx={{ m: 0 }}>
            {"Ostavi komentar"}
          </Typography>
        </Box>
        <Box sx={{ mt: 2 }}>
          <Box
            component="form"
            noValidate
            onSubmit={handleSubmit}
            sx={{ mt: 3 }}
          >
            <Grid container spacing={2}>
              <Grid item xs={12} sm={6}>
                <TextField
                  required
                  value={rating}
                  onChange={(e) => setRating(e.target.value)}
                  fullWidth
                  type={"number"}
                  label="Ocena"
                  autoFocus
                />
              </Grid>
              <Grid item xs={12} sm={6}>
                <TextField
                  required
                  value={comment}
                  onChange={(e) => setComment(e.target.value)}
                  fullWidth
                  label="Komentar"
                />
              </Grid>
            </Grid>
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              Potvrdi
            </Button>
          </Box>
        </Box>
      </Box>
    </Modal>
  );
};

export default ReviewModal;
