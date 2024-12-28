import React, { useEffect } from "react";
import { Button, FormControl, TextField } from "@mui/material";
import { useForm } from "react-hook-form";
import style from "./EditContactPage.module.css";
import { Backdrop } from "../../component/UI/backdrop";
import { useNavigate, useParams } from "react-router-dom";
import { editContact, getContactByID } from "../../store/contactSlice";
import { useDispatch, useSelector } from "react-redux";
import { AppDispatch, RootState } from "../../store/store";
import { contact } from "../../constants/types/typeContact";

export const EditContactPage: React.FC = () => {
    const { register, handleSubmit } = useForm();
    const { id } = useParams<{ id: string }>();
    const navigate = useNavigate();
    const dispatch = useDispatch<AppDispatch>();
    const { selectedContact } = useSelector(
        (state: RootState) => state.contactSlice
    );
    useEffect(() => {
        if (id) {
            dispatch(getContactByID(id))
        }
    }, [id]);

    const clickSubmit = async (data: any) => {
        const changeContact: contact = data;
        if (id) {
            changeContact.id = Number(id)
        }
        try {
            await dispatch(editContact(changeContact));
        } catch (error) {
            console.log(error);
            navigate("/login")
        } finally {
            navigate("/");
        }
    };

    return (
        <div className={style.login_page}>
            <Backdrop>
                <h2>Edit Contact</h2>
                <form className={style.login_form} noValidate onSubmit={handleSubmit((data: any) => clickSubmit(data))}>
                    <FormControl sx={{ m: 1, width: '40ch' }} variant="outlined">
                        <TextField id="outlined-basic" defaultValue={selectedContact?.Data.name ?? ""} label="Name" variant="outlined" {...register("name", {
                            required: "Required",
                        })} />
                    </FormControl>
                    <FormControl sx={{ m: 1, width: '40ch' }} variant="outlined">
                        <TextField id="outlined-basic" defaultValue={selectedContact?.Data.phone ?? ""} label="Phone" variant="outlined" {...register("phone", {
                            required: "Required",
                        })} />
                    </FormControl>
                    <FormControl sx={{ m: 1, width: '40ch' }} variant="outlined">
                        <TextField id="outlined-basic" defaultValue={selectedContact?.Data.email ?? ""} label="Email" variant="outlined" {...register("email", {
                            required: "Required",
                        })} />
                    </FormControl>
                    <Button type="submit">
                        Submit
                    </Button>
                </form>
            </Backdrop>
        </div >
    )
}