import React, { useEffect } from "react";
import { Modal, Box, Typography } from "@mui/material";
import AddRequestForm from "./AddRequestForm";
import { useDispatch } from "react-redux";
import { getVehicles } from "../vehicles/VehicleSlice";
import { getCarServices } from "../car_services/CarServiceSlice";

const style = {
  position: "absolute",
  left: "50%",
  transform: "translate(-50%, 0%)",
  width: 860,
  bgcolor: "background.paper",
  boxShadow: 12,
  p: 4,
  m: 0,
};

const AddRequestModal = ({ open, handleClose }) => {
  const dispatch = useDispatch();

  useEffect(() => {
    dispatch(getVehicles());
    dispatch(getCarServices());
  }, [dispatch]);

  return (
    <Modal open={open} onClose={handleClose}>
      <Box sx={style}>
        <Box>
          <Typography variant="h5" component="h5" sx={{ m: 0 }}>
            {"Podnosenje zahteva za servis"}
          </Typography>
          <Typography variant="subtitle2" component="h6" sx={{ m: 0 }}>
            {"Unesite podatke u tabelu ispod"}
          </Typography>
        </Box>
        <Box sx={{ mt: 2 }}>
          <AddRequestForm />
        </Box>
      </Box>
    </Modal>
  );
};

export default AddRequestModal;
