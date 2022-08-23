import React from "react";
import { Modal, Box, Typography } from "@mui/material";
import AddEditVehicleForm from "./AddEditVehicleForm";

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

const AddEditVehicleModal = ({ open, handleClose, v }) => {
  return (
    <Modal open={open} onClose={handleClose}>
      <Box sx={style}>
        <Box>
          <Typography variant="h5" component="h5" sx={{ m: 0 }}>
            {!!v ? v.Manufacturer : "Dodavanje vozila"}
          </Typography>
          <Typography variant="subtitle2" component="h6" sx={{ m: 0 }}>
            {!!v ? v.CarModel : "Unesite podatke u tabelu ispod"}
          </Typography>
        </Box>
        <Box sx={{ mt: 2 }}>
          <AddEditVehicleForm v={v} close={handleClose} />
        </Box>
      </Box>
    </Modal>
  );
};

export default AddEditVehicleModal;
