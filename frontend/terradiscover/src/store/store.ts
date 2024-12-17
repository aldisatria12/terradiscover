import { configureStore } from "@reduxjs/toolkit";
import authSlice from "./authSlice";
import contactSlice from "./contactSlice"

const store = configureStore({
  reducer: {
    authSlice,
    contactSlice
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export default store;
