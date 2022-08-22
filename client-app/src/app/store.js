import { configureStore } from '@reduxjs/toolkit';
import UserReducer from '../features/users/UsersSlice';

export const store = configureStore({
  reducer: {
    user: UserReducer,
  },
});
