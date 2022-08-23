import React from "react";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";

import { getUserIdFromToken } from "../../util/JwtToken";
import { useDispatch } from "react-redux";
import { updateVehicle, addVehicle } from "./VehicleSlice";

const AddEditVehicleForm = ({ v, close }) => {
  const dispatch = useDispatch();

  const handleSubmit = (event) => {
    event.preventDefault();
    const data = new FormData(event.currentTarget);
    if (!!v) {
      const patch = {
        id: v.Id,
        color: data.get("Color"),
      };
      dispatch(updateVehicle({ patch }));
      close();
    } else {
      const id = getUserIdFromToken();

      const vehicle = {
        UserId: id,
        Manufacturer: data.get("Manufacturer"),
        CarModel: data.get("CarModel"),
        Color: data.get("Color"),
        LicencePlate: data.get("LicencePlate"),
        ChassisNumber: data.get("ChassisNumber"),
      };
      dispatch(addVehicle({ vehicle }));
      close();
    }
  };

  return (
    <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 3 }}>
      <Grid container spacing={2}>
        <Grid item xs={12} sm={6}>
          <TextField
            name="Manufacturer"
            required
            defaultValue={!!v ? v.Manufacturer : ""}
            disabled={!!v}
            fullWidth
            id="Manufacturer"
            label="Proizvodjac"
            autoFocus
          />
        </Grid>
        <Grid item xs={12} sm={6}>
          <TextField
            required
            defaultValue={!!v ? v.CarModel : ""}
            disabled={!!v}
            fullWidth
            id="CarModel"
            label="Model"
            name="CarModel"
          />
        </Grid>
        <Grid item xs={12} sm={6}>
          <TextField
            name="Color"
            required
            defaultValue={!!v ? v.Color : ""}
            fullWidth
            id="Color"
            label="Boja"
            autoFocus
          />
        </Grid>
        <Grid item xs={12} sm={6}>
          <TextField
            required
            defaultValue={!!v ? v.LicencePlate : ""}
            disabled={!!v}
            fullWidth
            id="LicencePlate"
            label="Broj tablica"
            name="LicencePlate"
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            required
            fullWidth
            defaultValue={!!v ? v.ChassisNumber : ""}
            disabled={!!v}
            id="ChassisNumber"
            name="ChassisNumber"
            label="Broj sasije"
          />
        </Grid>
      </Grid>
      <Button type="submit" fullWidth variant="contained" sx={{ mt: 3, mb: 2 }}>
        Potvrdi
      </Button>
    </Box>
  );
};

export default AddEditVehicleForm;
