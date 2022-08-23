import React from "react";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

import api from "../../api/api";

const ChangePassword = () => {
  const changePassword = (data) => {
    api
      .patch(`/users/change-password`, data)
      .then(() => {
        alert("Uspesno promenjeno");
      })
      .catch(() => {
        alert("Greska prilikom menjanja");
      });
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    const data = new FormData(event.currentTarget);
    const postData = {
      Password: data.get("old_password"),
      NewPassword: data.get("new_password"),
    };
    changePassword(postData);
  };

  return (
    <Box
      sx={{
        marginTop: 8,
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
      }}
    >
      <Typography component="h5" variant="subtitle1">
        Promena lozinke
      </Typography>
      <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 3 }}>
        <Grid container spacing={2}>
          <Grid item xs={6}>
            <TextField
              required
              fullWidth
              id="old_password"
              name="old_password"
              label="Stara lozinka"
              type="password"
            />
          </Grid>
          <Grid item xs={6}>
            <TextField
              required
              fullWidth
              id="new_password"
              name="new_password"
              label="Nova lozinka"
              type="password"
            />
          </Grid>
        </Grid>
        <Button
          type="submit"
          fullWidth
          variant="contained"
          sx={{ mt: 3, mb: 2 }}
        >
          Promeni lozinku
        </Button>
      </Box>
    </Box>
  );
};

export default ChangePassword;
