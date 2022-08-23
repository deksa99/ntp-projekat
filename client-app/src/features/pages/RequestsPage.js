import React, { useEffect, useState } from "react";
import { Container, CssBaseline, Box } from "@mui/material";
import RequestsTable from "../requests/RequestsTable";
import { useDispatch, useSelector } from "react-redux";
import { getRequestsForUser } from "../requests/RequestsSlice";
import AddRequest from "../requests/AddRequest";
import { getRoleFromToken } from "../../util/JwtToken";
import { selectToken } from "../users/UsersSlice";

const RequestsPage = () => {
  const token = useSelector(selectToken);
  const [role, setRole] = useState(getRoleFromToken());
  useEffect(() => {
    setRole(getRoleFromToken());
  }, [token]);
  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(getRequestsForUser());
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
        <RequestsTable />
        {role === "user" && <AddRequest />}
      </Box>
    </Container>
  );
};

export default RequestsPage;
