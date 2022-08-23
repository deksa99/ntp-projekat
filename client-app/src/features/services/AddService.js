import React, { useState } from "react";
import { Container } from "@mui/system";
import Button from "@mui/material/Button";
import AddServiceModal from "./AddServiceModal";

const AddService = () => {
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
        Dodaj stavku
      </Button>
      <AddServiceModal open={open} handleClose={handleClose} v={null} />
    </Container>
  );
};

export default AddService;
