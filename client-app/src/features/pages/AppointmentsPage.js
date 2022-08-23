import React from "react";
import { Container, CssBaseline, Box } from "@mui/material";
import AppointmentsList from "../appointments/AppointmentsList";

const AppointmentsPage = () => {
  return (
    <Container type="main" maxWidth="lg">
      <CssBaseline />
      <Box
        sx={{
          marginTop: 8,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <AppointmentsList />
      </Box>
    </Container>
  );
};

export default AppointmentsPage;
