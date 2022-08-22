import React from "react";
import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import Link from "@mui/material/Link";
import { HomeOutlined } from "@mui/icons-material";
import { Tooltip } from "@mui/material";

const DefaultNavbar = () => {
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
          Servisi
        </Button>
        <span>
          <Button href="/login" variant="outlined" sx={{ my: 1, mx: 1.5 }}>
            Prijava
          </Button>
          <Button href="/sign-up" variant="contained" sx={{ my: 1, mx: 1.5 }}>
            Registracija
          </Button>
        </span>
      </Toolbar>
    </AppBar>
  );
};

export default DefaultNavbar;
