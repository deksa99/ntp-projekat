import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import api from "../../api/api"

export const login = createAsyncThunk(
    "users/login",
    async ({role, username, password}) => {
        return await api.post(`/users/login?role=${role}`, {"Username": username, "Password": password})
        .then((res) => {
            return res.data
         })
    }
);

const UsersSlice = createSlice({
    name: "user",
    initialState: {
        token: JSON.parse(localStorage.getItem("user")),
    },
    reducers: {
        logout(state, _) {
            state.token = null;
            console.log("LOGOUT")
            localStorage.removeItem("user");
        }
    },
    extraReducers: {
        [login.pending]: (state, _) => {
            state.token = null;
            console.log("PENDING")
        },
        [login.fulfilled]: (state, action) => {
            state.token = action.payload;
            console.log("FULFILLED");
            localStorage.setItem("user", JSON.stringify(action.payload));
        },
        [login.rejected]: (state, _) => {
            state.token = null;
            console.log("REJECTED");
            localStorage.removeItem("user");
        }
    }
})

export const selectToken = (state) => state.token;
export const { logout } = UsersSlice.actions;
export default UsersSlice.reducer;