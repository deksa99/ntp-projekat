import {
  Box,
  Button,
  FormControl,
  Grid,
  InputLabel,
  MenuItem,
  Select,
} from "@mui/material";
import React, { useEffect, useState } from "react";
import { useSelector, useDispatch } from "react-redux";
import api from "../../api/api";
import { selectCarServices } from "../car_services/CarServiceSlice";
import { selectVehicles } from "../vehicles/VehicleSlice";
import { getUserIdFromToken } from "../../util/JwtToken";
import { addRequest } from "./RequestsSlice";

const AddRequestForm = () => {
  const vehicles = useSelector(selectVehicles);
  const carServices = useSelector(selectCarServices);
  const [services, serServices] = useState([]);

  const [carServiceId, setCarServiceId] = useState(0);
  const [serviceId, setServiceId] = useState(0);
  const [vehicleId, setVehicleId] = useState(0);

  const dispatch = useDispatch();

  useEffect(() => {
    api
      .get(`/car-services/${carServiceId}/catalog`)
      .then((res) => {
        serServices(res.data);
      })
      .catch(() => {
        serServices([]);
      });
  }, [carServiceId]);

  const handleSubmit = (event) => {
    event.preventDefault();
    const post = {
      userId: getUserIdFromToken(),
      vehicleId,
      serviceId,
      carServiceId,
    };
    dispatch(addRequest({ post }));
  };

  const handleCSChange = (event) => {
    setCarServiceId(event.target.value);
  };

  const handleSChange = (event) => {
    setServiceId(event.target.value);
  };

  const handleVChange = (event) => {
    setVehicleId(event.target.value);
  };

  return (
    <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 3 }}>
      <Grid container spacing={2}>
        <Grid item xs={12}>
          <FormControl fullWidth>
            <InputLabel id="cs-select-label">Auto Servis</InputLabel>
            <Select
              labelId="cs-select-label"
              id="cs-select"
              label="Auto Servis"
              value={carServiceId}
              onChange={handleCSChange}
            >
              {carServices.map((cs) => (
                <MenuItem value={cs.Id}>{`${cs.Name} - ${cs.Street}`}</MenuItem>
              ))}
            </Select>
          </FormControl>
        </Grid>
        <Grid item xs={12}>
          <FormControl fullWidth>
            <InputLabel id="s-select-label">Usluga</InputLabel>
            <Select
              labelId="s-select-label"
              id="s-select"
              label="Usluga"
              value={serviceId}
              onChange={handleSChange}
            >
              {services.map((s) => (
                <MenuItem value={s.Id}>{`${s.Name} - ${s.Price}`}</MenuItem>
              ))}
            </Select>
          </FormControl>
        </Grid>
        <Grid item xs={12}>
          <FormControl fullWidth>
            <InputLabel id="v-select-label">Vozilo</InputLabel>
            <Select
              labelId="v-select-label"
              id="v-select"
              label="Vozilo"
              value={vehicleId}
              onChange={handleVChange}
            >
              {vehicles.map((v) => (
                <MenuItem
                  value={v.Id}
                >{`${v.Manufacturer} ${v.CarModel} - ${v.LicencePlate}`}</MenuItem>
              ))}
            </Select>
          </FormControl>
        </Grid>
      </Grid>
      <Button type="submit" fullWidth variant="contained" sx={{ mt: 3, mb: 2 }}>
        Podnesi zahtev
      </Button>
    </Box>
  );
};

export default AddRequestForm;
