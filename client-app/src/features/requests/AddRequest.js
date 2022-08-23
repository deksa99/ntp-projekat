import { Button, Container } from "@mui/material";
import React, { useState } from "react";
import AddRequestModal from "./AddRequestModal";

const AddRequest = () => {
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
        Podnesi zahtev za termin
      </Button>
      <AddRequestModal open={open} handleClose={handleClose} />
    </Container>
  );
};

export default AddRequest;
