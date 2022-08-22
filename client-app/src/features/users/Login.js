import {
  Box,
  Container,
  CssBaseline,
  Avatar,
  Typography,
  TextField,
  Select,
  MenuItem,
  Button,
  Link,
} from "@mui/material";
import { LockOutlined } from "@mui/icons-material";
import React from "react";

const Login = () => {
  const handleSubmit = (event) => {
    event.preventDefault();
    const data = new FormData(event.currentTarget);
    console.log({
      role: data.get("role"),
      username: data.get("username"),
      password: data.get("password"),
    });
    // TODO request
  };

  return (
    <Container component="main" maxWidth="xs">
      <CssBaseline />
      <Box
        sx={{
          marginTop: 8,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Avatar sx={{ m: 1, bgcolor: "secondary.main" }}>
          <LockOutlined />
        </Avatar>
        <Typography component="h1" variant="h5">
          Prijava
        </Typography>
        <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}>
          <Select
            fullWidth
            required
            defaultValue={"user"}
            id="role"
            name="role"
          >
            <MenuItem value={"user"}>Korisnik</MenuItem>
            <MenuItem value={"worker"}>Radnik</MenuItem>
            <MenuItem value={"admin"}>Admin</MenuItem>
          </Select>
          <TextField
            margin="normal"
            required
            fullWidth
            id="username"
            label="Korisničko ime"
            name="username"
            autoFocus
          />
          <TextField
            margin="normal"
            required
            fullWidth
            name="password"
            label="Lozinka"
            type="password"
            id="password"
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
          >
            Prijavi se
          </Button>
          <Link href="/sign-up" variant="body2">
            {"Nemaš nalog? Registruj se"}
          </Link>
        </Box>
      </Box>
    </Container>
  );
};

export default Login;
