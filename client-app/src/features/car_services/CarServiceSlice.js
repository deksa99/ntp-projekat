import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import api from "../../api/api";

export const getCarServices = createAsyncThunk(
  "car_service/get_all",
  async (_) => {
    return await api.get(`/car-services`).then((res) => {
      return res.data;
    });
  }
);

const CarServicesSlice = createSlice({
  name: "carService",
  initialState: {
    carServices: [],
  },
  reducers: {},
  extraReducers: {
    [getCarServices.pending]: (state, _) => {
      state.carServices = [];
    },
    [getCarServices.fulfilled]: (state, action) => {
      state.carServices = action.payload;
    },
    [getCarServices.rejected]: (state, _) => {
      state.carServices = [];
    },
  },
});

export const selectCarServices = (state) => state.carService.carServices;
export default CarServicesSlice.reducer;
