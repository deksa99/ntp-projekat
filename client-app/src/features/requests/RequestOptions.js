import React, { useEffect, useState } from "react";
import { IconButton } from "@mui/material";
import { useSelector, useDispatch } from "react-redux";
import { selectToken } from "../users/UsersSlice";
import { cancellRequest, rejectRequest } from "./RequestsSlice";
import { getRoleFromToken, getUserIdFromToken } from "../../util/JwtToken";
import CancelOutlined from "@mui/icons-material/CancelOutlined";
import CheckOutlined from "@mui/icons-material/CheckOutlined";
import StartDatePicker from "./StartDatePicker";

const RequestOptions = ({ r }) => {
  const [open, setOpen] = useState(false);

  const token = useSelector(selectToken);
  const [role, setRole] = useState(getRoleFromToken());
  useEffect(() => {
    setRole(getRoleFromToken());
  }, [token]);

  const dispatch = useDispatch();

  const handleCancelClick = () => {
    const patch = {
      userId: getUserIdFromToken(),
      appointmentRequestId: r.Id,
    };
    dispatch(cancellRequest({ patch }));
  };

  const handleAcceptClick = () => {
    setOpen(true);
  };

  const handleCloseDP = () => {
    setOpen(false);
  };

  const handleRejectClick = () => {
    dispatch(rejectRequest(r.Id));
  };

  return (
    <>
      {r.Status === "Submitted" && role === "user" && (
        <IconButton
          aria-label="edit"
          size="large"
          color="secondary"
          onClick={() => handleCancelClick()}
        >
          <CancelOutlined fontSize="inherit" />
        </IconButton>
      )}
      {r.Status === "Submitted" && role === "worker" && (
        <>
          <IconButton
            aria-label="edit"
            size="large"
            color="secondary"
            onClick={() => handleAcceptClick()}
          >
            <CheckOutlined fontSize="inherit" />
          </IconButton>
          <IconButton
            aria-label="edit"
            size="large"
            color="secondary"
            onClick={() => handleRejectClick()}
          >
            <CancelOutlined fontSize="inherit" />
          </IconButton>
          <StartDatePicker id={r.Id} open={open} handleClose={handleCloseDP} />
        </>
      )}
    </>
  );
};

export default RequestOptions;
