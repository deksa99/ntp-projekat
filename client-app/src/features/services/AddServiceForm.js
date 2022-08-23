import { Box, Button, Grid, TextField } from "@mui/material";
import React, { useState } from "react";
import { getServiceFromToken } from "../../util/JwtToken";
import api from "../../api/api";

const AddServiceForm = () => {
  const [Name, setName] = useState("");
  const [Description, setDescription] = useState("");
  const [Price, setPrice] = useState(0);

  const handleSubmit = (e) => {
    e.preventDefault();
    const body = {
      CarServiceId: getServiceFromToken(),
      Name,
      Description,
      Price: parseInt(Price),
      ExpectedTime: null,
    };
    console.log(body);
    api
      .post("/car-services/service", body)
      .then((res) => {
        alert("Uspesno");
      })
      .catch((err) => {
        alert("Greska");
      });
  };

  return (
    <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 3 }}>
      <Grid container spacing={2}>
        <Grid item xs={12}>
          <TextField
            required
            value={Name}
            fullWidth
            label="Naziv"
            onChange={(e) => setName(e.target.value)}
            autoFocus
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            required
            value={Description}
            fullWidth
            label="Opis"
            onChange={(e) => setDescription(e.target.value)}
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            required
            value={Price}
            fullWidth
            type={"number"}
            label="Cena"
            onChange={(e) => setPrice(e.target.value)}
          />
        </Grid>
      </Grid>
      <Button
        type="submit"
        onSubmit={handleSubmit}
        fullWidth
        variant="contained"
        sx={{ mt: 3, mb: 2 }}
      >
        Dodaj
      </Button>
    </Box>
  );
};

export default AddServiceForm;
