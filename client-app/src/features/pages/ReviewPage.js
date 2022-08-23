import { CssBaseline } from "@mui/material";
import { Box, Container } from "@mui/system";
import React from "react";

import ReviewTable from "../reviews/ReviewTable";

const ReviewPage = () => {
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
        <ReviewTable />
      </Box>
    </Container>
  );
};

export default ReviewPage;
