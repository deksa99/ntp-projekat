import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import api from "../../api/api";

export const getVehicles = createAsyncThunk("vehicles/get_all", async () => {
  return await api.get(`/vehicles/all`).then((res) => {
    return res.data;
  });
});

export const addVehicle = createAsyncThunk(
  "vehicles/add",
  async ({ vehicle }) => {
    return await api.post(`/vehicles/add`, vehicle).then((res) => {
      return res.data;
    });
  }
);

export const updateVehicle = createAsyncThunk(
  "vehicles/update",
  async ({ patch }) => {
    return await api
      .patch(`/vehicles/${patch.id}`, { Color: patch.color })
      .then((res) => {
        return res.data;
      });
  }
);

const VehiclesSlice = createSlice({
  name: "vehicles",
  initialState: {
    vehicles: [],
  },
  reducers: {},
  extraReducers: {
    [getVehicles.pending]: (state, _) => {
      state.vehicles = [];
    },
    [getVehicles.fulfilled]: (state, action) => {
      state.vehicles = action.payload;
    },
    [getVehicles.rejected]: (state, _) => {
      state.vehicles = [];
    },
    [addVehicle.pending]: () => {},
    [addVehicle.fulfilled]: (state, action) => {
      state.vehicles.push(action.payload);
    },
    [addVehicle.rejected]: () => {},
    [updateVehicle.pending]: () => {},
    [updateVehicle.fulfilled]: (state, action) => {
      state.vehicles.map((e) => {
        if (e.Id === action.payload.Id) {
          return { ...e, Color: action.payload.Color };
        }
        return e;
      });
    },
    [updateVehicle.rejected]: () => {},
  },
});

export const selectVehicles = (state) => state.vehicle.vehicles;
export default VehiclesSlice.reducer;
