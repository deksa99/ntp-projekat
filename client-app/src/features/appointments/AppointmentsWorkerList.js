import React, { useEffect } from "react";
import List from "@mui/material/List";
import { useSelector, useDispatch } from "react-redux";
import {
  selectAppointments,
  getAppointmentsForWorker,
} from "../appointments/AppointmentsSlice";
import AppointmentListItem from "./AppointmentListItem";

const AppointmentsWorkerList = () => {
  const appointments = useSelector(selectAppointments);

  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(getAppointmentsForWorker());
  }, [dispatch]);

  return (
    <List sx={{ width: "100%", minWidth: 360, bgcolor: "background.paper" }}>
      {appointments.map((a) => (
        <AppointmentListItem app={a} />
      ))}
    </List>
  );
};

export default AppointmentsWorkerList;
