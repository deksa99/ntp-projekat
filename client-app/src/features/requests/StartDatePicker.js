import React from "react";
import { Modal, Box, Typography, Button } from "@mui/material";
import TextField from "@mui/material/TextField";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DateTimePicker } from "@mui/x-date-pickers/DateTimePicker";
import { acceptRequest } from "./RequestsSlice";
import { useDispatch } from "react-redux";

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

const StartDatePicker = ({ id, open, handleClose }) => {
  const dispatch = useDispatch();
  const [value, setValue] = React.useState(new Date("2022-08-25T10:00:00"));

  const handleSubmit = () => {
    const patch = {
      id: id,
      start: value,
    };
    dispatch(acceptRequest({ patch }));
  };

  const handleChange = (newValue) => {
    setValue(newValue);
  };

  return (
    <Modal open={open} onClose={handleClose}>
      <Box sx={style}>
        <Box>
          <Typography variant="h5" component="h5" sx={{ m: 0 }}>
            {"Unesite vreme pocetka servisa"}
          </Typography>
        </Box>
        <Box sx={{ mt: 2 }}>
          <LocalizationProvider dateAdapter={AdapterDateFns}>
            <DateTimePicker
              label="Pocetak servisa"
              value={value}
              onChange={handleChange}
              renderInput={(params) => <TextField {...params} />}
            />
          </LocalizationProvider>
          <Button variant="contained" color="primary" onClick={handleSubmit}>
            Potvrdi
          </Button>
        </Box>
      </Box>
    </Modal>
  );
};

export default StartDatePicker;
