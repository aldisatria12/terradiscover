import { createSlice, Dispatch } from "@reduxjs/toolkit";
import { ResError } from "../constants/response/resAuth";
import { ResContact } from "../constants/response/resContact";

export interface ContactState {
    contactList: ResContact | null;
    isGetContactError: boolean;
    isGetContactSuccess: boolean;
    errorGetContactMsg: string | null;
}

const initialState: ContactState = {
    contactList: null,
    isGetContactError: false,
    isGetContactSuccess: false,
    errorGetContactMsg: null,
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
    },
});

export const {
    setContactList,
    setIsGetContactError,
    setErrorGetContactMsg,
    setGetContactSuccess,
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

export default appSlice.reducer;