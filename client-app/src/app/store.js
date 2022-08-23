import { configureStore } from "@reduxjs/toolkit";
import CarServiceReducer from "../features/car_services/CarServiceSlice";
import UserReducer from "../features/users/UsersSlice";
import VehicleReducer from "../features/vehicles/VehicleSlice";

export const store = configureStore({
  reducer: {
    user: UserReducer,
    carService: CarServiceReducer,
    vehicle: VehicleReducer,
  },
});
