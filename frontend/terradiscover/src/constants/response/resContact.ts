import { contact } from "../types/typeContact";

export type ResContact = {
    Data: contact[];
    Msg: string;
};

export type ResSelectedContact = {
    Data: contact;
    Msg: string;
};
