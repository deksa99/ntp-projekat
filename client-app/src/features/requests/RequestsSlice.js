import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import api from "../../api/api";

export const getRequestsForUser = createAsyncThunk(
  "requests/get_all",
  async () => {
    return await api.get(`/appointments/requests`).then((res) => {
      return res.data;
    });
  }
);

export const cancellRequest = createAsyncThunk(
  "requests/cancel",
  async ({ patch }) => {
    return await api
      .patch(`/appointments/requests/cancel`, patch)
      .then((res) => {
        return res.data;
      });
  }
);

export const addRequest = createAsyncThunk("requests/add", async ({ post }) => {
  return await api.post(`/appointments/requests`, post).then((res) => {
    return res.data;
  });
});

export const getRequestsForService = createAsyncThunk(
  "requests/get_all_service",
  async (id) => {
    return await api.get(`/appointments/requests/service/${id}`).then((res) => {
      return res.data;
    });
  }
);

export const rejectRequest = createAsyncThunk("requests/reject", async (id) => {
  return await api
    .patch(`/appointments/requests/${id}/reject`, {})
    .then((res) => {
      return res.data;
    });
});

export const acceptRequest = createAsyncThunk(
  "requests/accept",
  async ({ patch }) => {
    return await api
      .patch(`/appointments/requests/${patch.id}/accept`, {
        Start: patch.start,
      })
      .then((res) => {
        return res.data;
      });
  }
);

const RequestsSlice = createSlice({
  name: "request",
  initialState: {
    requests: [],
  },
  reducers: {},
  extraReducers: {
    [getRequestsForUser.pending]: (state, _) => {
      state.requests = [];
    },
    [getRequestsForUser.fulfilled]: (state, action) => {
      state.requests = action.payload;
    },
    [getRequestsForUser.rejected]: (state, _) => {
      state.requests = [];
    },
    [getRequestsForService.pending]: (state, _) => {
      state.requests = [];
    },
    [getRequestsForService.fulfilled]: (state, action) => {
      state.requests = action.payload;
    },
    [getRequestsForService.rejected]: (state, _) => {
      state.requests = [];
    },
  },
});

export const selectRequests = (state) => state.request.requests;
export default RequestsSlice.reducer;
