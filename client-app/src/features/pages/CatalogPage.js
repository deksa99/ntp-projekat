import React from "react";
import { Container, CssBaseline, Box } from "@mui/material";
import Catalog from "../services/Catalog";
import AddService from "../services/AddService";

const CatalogPage = () => {
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
        <Catalog />
        <AddService />
      </Box>
    </Container>
  );
};

export default CatalogPage;
