import React, { useEffect } from "react";
import List from "@mui/material/List";
import { useSelector, useDispatch } from "react-redux";
import {
  selectAppointments,
  getAppointmentsForUser,
} from "../appointments/AppointmentsSlice";
import AppointmentListItem from "./AppointmentListItem";

const AppointmentsList = () => {
  const appointments = useSelector(selectAppointments);

  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(getAppointmentsForUser());
  }, [dispatch]);

  return (
    <List sx={{ width: "100%", minWidth: 360, bgcolor: "background.paper" }}>
      {appointments.map((a) => (
        <AppointmentListItem app={a} />
      ))}
    </List>
  );
};

export default AppointmentsList;
