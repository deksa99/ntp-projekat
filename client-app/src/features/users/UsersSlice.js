import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import api from "../../api/api";

export const getAllUsers = createAsyncThunk("users/get_all", async () => {
  return await api.get(`/users`).then((res) => {
    return res.data;
  });
});

export const login = createAsyncThunk(
  "users/login",
  async ({ role, username, password }) => {
    return await api
      .post(`/users/login?role=${role}`, {
        Username: username,
        Password: password,
      })
      .then((res) => {
        return res.data;
      });
  }
);

const UsersSlice = createSlice({
  name: "user",
  initialState: {
    token: null,
    users: [],
  },
  reducers: {
    logout(state, _) {
      state.token = null;
      localStorage.removeItem("user");
    },
  },
  extraReducers: {
    [login.pending]: (state, _) => {
      state.token = null;
      localStorage.removeItem("user");
    },
    [login.fulfilled]: (state, action) => {
      state.token = action.payload;
      localStorage.setItem("user", JSON.stringify(action.payload));
    },
    [login.rejected]: (state, _) => {
      state.token = null;
      localStorage.removeItem("user");
    },
    [getAllUsers.pending]: (state, _) => {
      state.users = [];
    },
    [getAllUsers.fulfilled]: (state, action) => {
      state.users = action.payload;
    },
    [getAllUsers.rejected]: (state, _) => {
      state.users = [];
    },
  },
});

export const selectToken = (state) => state.user.token;
export const selectUsers = (state) => state.user.users;
export const { logout } = UsersSlice.actions;
export default UsersSlice.reducer;
