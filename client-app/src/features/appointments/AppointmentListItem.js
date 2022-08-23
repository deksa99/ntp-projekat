import React from "react";
import ListItem from "@mui/material/ListItem";
import Divider from "@mui/material/Divider";
import ListItemText from "@mui/material/ListItemText";
import Typography from "@mui/material/Typography";

const AppointmentListItem = ({ app }) => {
  return (
    <>
      <ListItem alignItems="flex-start">
        <ListItemText
          primary={`${app.ServiceName} | ${app.CarServiceName} - ${app.Street} ... Cena: ${app.ServicePrice}`}
          secondary={
            <React.Fragment>
              <Typography
                sx={{ display: "inline" }}
                component="span"
                variant="body2"
                color="text.primary"
              >
                {`${app.StartTime} - ${app.Status}`}
              </Typography>
            </React.Fragment>
          }
        />
      </ListItem>
      <Divider variant="inset" component="li" />
    </>
  );
};

export default AppointmentListItem;
