import React from "react";
import { Container, CssBaseline, Box } from "@mui/material";
import Profile from "../users/Profile";
import ChangePassword from "../users/ChangePassword";

const ProfilePage = () => {
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
        <Profile />
        <ChangePassword />
      </Box>
    </Container>
  );
};

export default ProfilePage;
