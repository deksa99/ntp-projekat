import React, { useEffect } from "react";
import { Container, CssBaseline, Box } from "@mui/material";
import RequestsTable from "../requests/RequestsTable";
import { useDispatch } from "react-redux";
import { getRequestsForService } from "../requests/RequestsSlice";
import { getServiceFromToken } from "../../util/JwtToken";

const RequestsWorkerPage = () => {
  const dispatch = useDispatch();

  useEffect(() => {
    const sid = getServiceFromToken();
    console.log(sid);
    dispatch(getRequestsForService(sid));
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
      </Box>
    </Container>
  );
};

export default RequestsWorkerPage;
