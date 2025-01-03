import { createSlice, Dispatch } from "@reduxjs/toolkit";
import { ResError } from "../constants/response/resAuth";
import { ResContact, ResSelectedContact } from "../constants/response/resContact";
import { contact, inputContact } from "../constants/types/typeContact";
import { backEndURL } from "../constants/constants";

export interface ContactState {
    contactList: ResContact | null;
    selectedContact: ResSelectedContact | null;
    isGetContactError: boolean;
    isGetContactSuccess: boolean;
    errorGetContactMsg: string | null;
    isInsertContactError: boolean;
    isInsertContactSuccess: boolean;
    errorInsertContactMsg: string | null;
}

const initialState: ContactState = {
    contactList: null,
    selectedContact: null,
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
        setSelectedContact: (state, action) => {
            state.selectedContact = action.payload;
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
    setSelectedContact,
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
            const link = backEndURL + "/contact";
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

export const getContactByID = (id: string) => {
    return async (dispatch: Dispatch): Promise<void> => {
        try {
            dispatch(setIsGetContactError(false));
            const link = backEndURL + "/contact/" + id;
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
            const data: ResSelectedContact = await response.json();
            dispatch(setSelectedContact(data));
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
            const link = backEndURL + "/contact/insert";
            const response = await fetch(link, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + localStorage.getItem("token")
                },
                body: JSON.stringify(input)
            });
            console.log(JSON.stringify(input))
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

export const editContact = (input: contact) => {
    return async (dispatch: Dispatch): Promise<void> => {
        try {
            dispatch(setIsInsertContactError(false));
            const link = backEndURL + "/contact/edit";
            const response = await fetch(link, {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + localStorage.getItem("token")
                },
                body: JSON.stringify(input)
            });
            console.log(JSON.stringify(input))
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