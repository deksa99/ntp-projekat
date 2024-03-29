import { configureStore } from "@reduxjs/toolkit";
import CarServiceReducer from "../features/car_services/CarServiceSlice";
import UserReducer from "../features/users/UsersSlice";
import VehicleReducer from "../features/vehicles/VehicleSlice";
import AppointmentReducer from "../features/appointments/AppointmentsSlice";
import RequestReducer from "../features/requests/RequestsSlice";

export const store = configureStore({
  reducer: {
    user: UserReducer,
    carService: CarServiceReducer,
    vehicle: VehicleReducer,
    appointment: AppointmentReducer,
    request: RequestReducer,
  },
});
