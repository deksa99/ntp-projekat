import React, { useEffect, useState } from "react";
import ListItem from "@mui/material/ListItem";
import Divider from "@mui/material/Divider";
import ListItemText from "@mui/material/ListItemText";
import Typography from "@mui/material/Typography";
import { getRoleFromToken } from "../../util/JwtToken";
import { useSelector } from "react-redux";
import { selectToken } from "../users/UsersSlice";
import { ListItemButton } from "@mui/material";
import { useDispatch } from "react-redux";
import { cancellAppointment, finishAppointment } from "./AppointmentsSlice";

const AppointmentListItem = ({ app }) => {
  const dispatch = useDispatch();

  const token = useSelector(selectToken);
  const [role, setRole] = useState(getRoleFromToken());
  useEffect(() => {
    setRole(getRoleFromToken());
  }, [token]);

  const cancellAppointmentH = () => {
    dispatch(cancellAppointment(app.Id));
  };

  const finishAppointmentH = () => {
    dispatch(finishAppointment(app.Id));
  };

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
        {role === "worker" && app.Status === "Scheduled" && (
          <ListItemButton onClick={cancellAppointmentH}>Otkazi</ListItemButton>
        )}
        {role === "worker" && app.Status === "Scheduled" && (
          <ListItemButton onClick={finishAppointmentH}>Zavrsi</ListItemButton>
        )}
      </ListItem>
      <Divider variant="inset" component="li" />
    </>
  );
};

export default AppointmentListItem;
