import React, { useEffect } from "react";
import { useDispatch } from "react-redux";
import { getCarServices } from "../car_services/CarServiceSlice";
import CarServiceTable from "../car_services/CarServiceTable";
import { Container, CssBaseline, Box } from "@mui/material";

const ServiceInfoPage = () => {
  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(getCarServices());
  }, [dispatch]);

  return (
    <Container type="main" maxWidth="md">
      <CssBaseline />
      <Box
        xs={{
          marginTop: 50,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <CarServiceTable />
      </Box>
    </Container>
  );
};

export default ServiceInfoPage;
