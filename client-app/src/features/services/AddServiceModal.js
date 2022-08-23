import React from "react";
import { Modal, Box, Typography } from "@mui/material";
import AddServiceForm from "./AddServiceForm";

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

const AddServiceModal = ({ open, handleClose }) => {
  return (
    <Modal open={open} onClose={handleClose}>
      <Box sx={style}>
        <Box>
          <Typography variant="h5" component="h5" sx={{ m: 0 }}>
            {"Dodavanje stavke"}
          </Typography>
        </Box>
        <Box sx={{ mt: 2 }}>
          <AddServiceForm />
        </Box>
      </Box>
    </Modal>
  );
};

export default AddServiceModal;
