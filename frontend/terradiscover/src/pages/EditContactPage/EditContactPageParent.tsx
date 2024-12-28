import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { getContactByID } from "../../store/contactSlice";
import { useDispatch, useSelector } from "react-redux";
import { AppDispatch, RootState } from "../../store/store";
import { EditContactPage } from "./EditContactPage";

export const EditContactParent: React.FC = () => {
    const { id } = useParams<{ id: string }>();
    const [isLoading, setLoading] = useState<boolean>(true);
    const dispatch = useDispatch<AppDispatch>();
    const { selectedContact } = useSelector(
        (state: RootState) => state.contactSlice
    );

    useEffect(() => {
        if (id) {
            dispatch(getContactByID(id))
        }
    }, [id]);

    useEffect(() => {
        if (selectedContact) {
            setLoading(false)
        }
    }, [selectedContact])

    return (
        <div>
            {
                isLoading ?
                    <h2>Loading...</h2> :
                    <EditContactPage />
            }
        </div>
    )
}