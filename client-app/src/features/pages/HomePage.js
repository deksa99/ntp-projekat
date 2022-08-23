import React from "react";
import { Container, CssBaseline, Box, Typography } from "@mui/material";

const HomePage = () => {
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
        <Typography component="h1" variant="h2">
          kreativna početna strana
        </Typography>
      </Box>
    </Container>
  );
};

export default HomePage;
