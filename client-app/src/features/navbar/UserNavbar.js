import React from "react";
import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import Link from "@mui/material/Link";
import { HomeOutlined } from "@mui/icons-material";
import { Tooltip } from "@mui/material";
import { useDispatch } from "react-redux";

import { logout } from "../users/UsersSlice";

const UserNavbar = () => {
  const dispatch = useDispatch();

  return (
    <AppBar
      position="static"
      color="default"
      elevation={0}
      sx={{ borderBottom: (theme) => `1px solid ${theme.palette.divider}` }}
    >
      <Toolbar sx={{ flexWrap: "wrap" }}>
        <Link href="/">
          <Tooltip title="Početna strana">
            <Avatar
              sx={{
                m: 1,
                bgcolor: "secondary.main",
                marginLeft: "0px",
              }}
              href="/"
            >
              <HomeOutlined />
            </Avatar>
          </Tooltip>
        </Link>
        <Typography
          variant="h6"
          color="inherit"
          noWrap
          sx={{ flexGrow: 1, margin: "0px 0px 0px 10px" }}
        >
          NTP projekat
        </Typography>
        <Button href="/services" variant="outlined" sx={{ my: 1, mx: 1.5 }}>
          Auto Servisi
        </Button>
        <Button href="/vehicles" variant="outlined" sx={{ my: 1, mx: 1.5 }}>
          Vozila
        </Button>
        <Button href="/vehicles" variant="outlined" sx={{ my: 1, mx: 1.5 }}>
          Termini
        </Button>
        <Button href="/profile" variant="outlined" sx={{ my: 1, mx: 1.5 }}>
          Profil
        </Button>
        <Button
          href="/"
          onClick={() => dispatch(logout())}
          variant="contained"
          sx={{ my: 1, mx: 1.5 }}
        >
          Odjava
        </Button>
      </Toolbar>
    </AppBar>
  );
};

export default UserNavbar;
