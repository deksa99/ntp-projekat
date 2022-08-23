import React, { useEffect } from "react";
import { Container, CssBaseline, Box } from "@mui/material";
import { getVehicles } from "../vehicles/VehicleSlice";
import VehiclesTable from "../vehicles/VehiclesTable";
import { useDispatch } from "react-redux";
import AddVehicle from "../vehicles/AddVehicle";

const VehiclesPage = () => {
  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(getVehicles());
  }, [dispatch]);

  return (
    <Container type="main" maxWidth="md">
      <CssBaseline />
      <Box
        sx={{
          marginTop: 8,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <VehiclesTable />
        <AddVehicle />
      </Box>
    </Container>
  );
};

export default VehiclesPage;
