import { createSlice, Dispatch } from "@reduxjs/toolkit";
import { ResError, ResLogin } from "../constants/response/resAuth";
import { inputLogin, inputRegister } from "../constants/types/typeAuth";

export interface AuthState {
  userToken: ResLogin | null;
  isRegisterError: boolean;
  isRegisterSuccess: boolean;
  errorRegisterMsg: string | null;
  isLoginError: boolean;
  errorLoginMsg: string | null;
}

const initialState: AuthState = {
  userToken: null,
  isRegisterError: false,
  isRegisterSuccess: false,
  errorRegisterMsg: null,
  isLoginError: false,
  errorLoginMsg: null,
};

export const appSlice = createSlice({
  name: "auth",
  initialState: initialState,
  reducers: {
    setUserToken: (state, action) => {
      state.userToken = action.payload;
    },
    setIsRegisterError: (state, action) => {
      state.isRegisterError = action.payload;
    },
    setErrorRegisterMsg: (state, action) => {
      state.errorRegisterMsg = action.payload;
    },
    setRegisterSuccess: (state, action) => {
      state.isRegisterSuccess = action.payload;
    },
    setIsLoginError: (state, action) => {
      state.isLoginError = action.payload;
    },
    setErrorLoginMsg: (state, action) => {
      state.errorLoginMsg = action.payload;
    },
  },
});

export const {
  setUserToken,
  setIsRegisterError,
  setErrorRegisterMsg,
  setRegisterSuccess,
  setIsLoginError,
  setErrorLoginMsg,
} = appSlice.actions;

export const login = (input: inputLogin) => {
  return async (dispatch: Dispatch): Promise<void> => {
    try {
      dispatch(setIsLoginError(false));
      const link = "http://localhost:8081/auth/login";
      const response = await fetch(link, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(input),
      });
      if (!response.ok) {
        const errorMsg: ResError = await response.json();
        if (!errorMsg.errors) {
          throw new Error(errorMsg.message);
        }
        throw new Error(errorMsg.errors[0].message);
      }
      const data: ResLogin = await response.json();
      dispatch(setUserToken(data.Data.token));
      localStorage.setItem("token", data.Data.token);
    } catch (error) {
      dispatch(setIsLoginError(true));
      if (error instanceof Error) {
        dispatch(setErrorLoginMsg(error.message));
      } else {
        dispatch(setErrorLoginMsg("Something wrong with the server"));
      }
      throw error;
    }
  };
};

export const register = (input: inputRegister) => {
  return async (dispatch: Dispatch): Promise<void> => {
    try {
      dispatch(setRegisterSuccess(false));
      dispatch(setIsRegisterError(false));
      const link = "http://localhost:8081/auth/register";
      const response = await fetch(link, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(input),
      });
      if (!response.ok) {
        const errorMsg: ResError = await response.json();
        if (!errorMsg.errors) {
          throw new Error(errorMsg.message);
        }
        throw new Error(errorMsg.errors[0].message);
      }
      dispatch(setRegisterSuccess(true));
      const newUser: inputLogin = {
        email: input.email,
        password: input.password,
      };
      login(newUser);
    } catch (error) {
      dispatch(setIsRegisterError(true));
      if (error instanceof Error) {
        dispatch(setErrorRegisterMsg(error.message));
      } else {
        dispatch(setErrorRegisterMsg("Something wrong with the server"));
      }
      throw error;
    }
  };
};

export default appSlice.reducer;
