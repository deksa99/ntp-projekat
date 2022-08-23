import React from "react";
import { Modal, Box, Typography } from "@mui/material";
import CarServiceDetailsTabs from "./CarServiceDetailsTabs";

const style = {
  position: "absolute",
  top: "30%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  width: 860,
  bgcolor: "background.paper",
  boxShadow: 12,
  p: 4,
  m: 0,
};

const CarServiceDetailsModal = ({ open, handleClose, cs }) => {
  return (
    <Modal open={open} onClose={handleClose}>
      <Box sx={style}>
        <Box>
          <Typography variant="h5" component="h5" sx={{ m: 0 }}>
            {!!cs ? cs.Name : ""}
          </Typography>
          <Typography variant="subtitle2" component="h6" sx={{ m: 0 }}>
            {!!cs ? cs.Street : ""}
          </Typography>
        </Box>
        <Box sx={{ mt: 2 }}>
          <CarServiceDetailsTabs cs={cs} />
        </Box>
      </Box>
    </Modal>
  );
};

export default CarServiceDetailsModal;
