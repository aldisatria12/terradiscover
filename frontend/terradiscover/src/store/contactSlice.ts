import { createSlice, Dispatch } from "@reduxjs/toolkit";
import { ResError } from "../constants/response/resAuth";
import { ResContact } from "../constants/response/resContact";
import { inputContact } from "../constants/types/typeContact";

export interface ContactState {
    contactList: ResContact | null;
    isGetContactError: boolean;
    isGetContactSuccess: boolean;
    errorGetContactMsg: string | null;
    isInsertContactError: boolean;
    isInsertContactSuccess: boolean;
    errorInsertContactMsg: string | null;
}

const initialState: ContactState = {
    contactList: null,
    isGetContactError: false,
    isGetContactSuccess: false,
    errorGetContactMsg: null,
    isInsertContactError: false,
    isInsertContactSuccess: false,
    errorInsertContactMsg: null,
};

export const appSlice = createSlice({
    name: "contact",
    initialState: initialState,
    reducers: {
        setContactList: (state, action) => {
            state.contactList = action.payload;
        },
        setIsGetContactError: (state, action) => {
            state.isGetContactError = action.payload;
        },
        setErrorGetContactMsg: (state, action) => {
            state.isGetContactSuccess = action.payload;
        },
        setGetContactSuccess: (state, action) => {
            state.errorGetContactMsg = action.payload;
        },
        setIsInsertContactError: (state, action) => {
            state.isInsertContactError = action.payload;
        },
        setErrorInsertContactMsg: (state, action) => {
            state.isInsertContactSuccess = action.payload;
        },
        setInsertContactSuccess: (state, action) => {
            state.errorInsertContactMsg = action.payload;
        },
    },
});

export const {
    setContactList,
    setIsGetContactError,
    setErrorGetContactMsg,
    setGetContactSuccess,
    setIsInsertContactError,
    setErrorInsertContactMsg,
    setInsertContactSuccess,
} = appSlice.actions;

export const getContactList = () => {
    return async (dispatch: Dispatch): Promise<void> => {
        try {
            dispatch(setIsGetContactError(false));
            const link = "http://localhost:8081/contact";
            const response = await fetch(link, {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + localStorage.getItem("token")
                },
            });
            if (!response.ok) {
                const errorMsg: ResError = await response.json();
                if (!errorMsg.errors) {
                    throw new Error(errorMsg.message);
                }
                throw new Error(errorMsg.errors[0].message);
            }
            const data: ResContact = await response.json();
            dispatch(setContactList(data));
        } catch (error) {
            dispatch(setIsGetContactError(true));
            if (error instanceof Error) {
                dispatch(setErrorGetContactMsg(error.message));
            } else {
                dispatch(setErrorGetContactMsg("Something wrong with the server"));
            }
            throw error;
        }
    };
};

export const insertContact = (input: inputContact) => {
    return async (dispatch: Dispatch): Promise<void> => {
        try {
            dispatch(setIsInsertContactError(false));
            const link = "http://localhost:8081/contact/insert";
            const response = await fetch(link, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + localStorage.getItem("token")
                },
                body: JSON.stringify(input)
            });
            if (!response.ok) {
                const errorMsg: ResError = await response.json();
                if (!errorMsg.errors) {
                    throw new Error(errorMsg.message);
                }
                throw new Error(errorMsg.errors[0].message);
            }
        } catch (error) {
            dispatch(setIsInsertContactError(true));
            if (error instanceof Error) {
                dispatch(setErrorInsertContactMsg(error.message));
            } else {
                dispatch(setErrorInsertContactMsg("Something wrong with the server"));
            }
            throw error;
        }
    };
};

export default appSlice.reducer;