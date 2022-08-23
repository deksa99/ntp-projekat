import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import api from "../../api/api";

export const getAppointmentsForUser = createAsyncThunk(
  "appointments/get_all",
  async () => {
    return await api.get(`/appointments/user`).then((res) => {
      return res.data;
    });
  }
);

export const getAppointmentsForWorker = createAsyncThunk(
  "appointments/get_all_worker",
  async () => {
    return await api.get(`/appointments/worker`).then((res) => {
      return res.data;
    });
  }
);

export const getAppointmentsForService = createAsyncThunk(
  "appointments/get_all_service",
  async ({ id }) => {
    return await api.get(`/appointments/car-service/${id}`).then((res) => {
      return res.data;
    });
  }
);

const AppointmentsSlice = createSlice({
  name: "appointments",
  initialState: {
    appointments: [],
  },
  reducers: {},
  extraReducers: {
    [getAppointmentsForUser.pending]: (state, _) => {
      state.appointments = [];
    },
    [getAppointmentsForUser.fulfilled]: (state, action) => {
      state.appointments = action.payload;
    },
    [getAppointmentsForUser.rejected]: (state, _) => {
      state.appointments = [];
    },
    [getAppointmentsForWorker.pending]: (state, _) => {
      state.appointments = [];
    },
    [getAppointmentsForWorker.fulfilled]: (state, action) => {
      state.appointments = action.payload;
    },
    [getAppointmentsForWorker.rejected]: (state, _) => {
      state.appointments = [];
    },
  },
});

export const selectAppointments = (state) => state.appointment.appointments;
export default AppointmentsSlice.reducer;
