import React from "react";
import { Container, CssBaseline, Box } from "@mui/material";
import AppointmentsWorkerList from "../appointments/AppointmentsWorkerList";

const AppointmentsWorkerPage = () => {
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
        <AppointmentsWorkerList />
      </Box>
    </Container>
  );
};

export default AppointmentsWorkerPage;
