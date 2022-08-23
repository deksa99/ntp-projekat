import React, { useEffect } from "react";
import UsersTable from "../users/UsersTable";
import { useDispatch } from "react-redux";
import { getAllUsers } from "../users/UsersSlice";
import { Box, Container } from "@mui/system";
import { CssBaseline } from "@mui/material";

const UsersPage = () => {
  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(getAllUsers());
  }, [dispatch]);

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
        <UsersTable />;
      </Box>
    </Container>
  );
};

export default UsersPage;
