import { Container } from "@mui/system";
import React, { useState } from "react";
import Button from "@mui/material/Button";
import AddEditVehicleModal from "./AddEditVehicleModal";

const AddVehicle = () => {
  const [open, setOpen] = useState(false);

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <Container>
      <Button
        fullWidth
        variant="outlined"
        color="primary"
        onClick={() => setOpen(true)}
      >
        Dodaj vozilo
      </Button>
      <AddEditVehicleModal open={open} handleClose={handleClose} v={null} />
    </Container>
  );
};

export default AddVehicle;
