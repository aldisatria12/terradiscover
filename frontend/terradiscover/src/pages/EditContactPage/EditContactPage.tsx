import React, { useEffect } from "react";
import { Button, FormControl, TextField } from "@mui/material";
import { useForm } from "react-hook-form";
import style from "./EditContactPage.module.css";
import { Backdrop } from "../../component/UI/backdrop";
import { useParams } from "react-router-dom";
import { getContactByID } from "../../store/contactSlice";
import { useDispatch, useSelector } from "react-redux";
import { AppDispatch, RootState } from "../../store/store";

export const EditContactPage: React.FC = () => {
    const { register } = useForm();
    const params = useParams();
    const dispatch = useDispatch<AppDispatch>();
    const { selectedContact } = useSelector(
        (state: RootState) => state.contactSlice
    );

    useEffect(() => {
        if (params.id) {
            const id = params.id
            dispatch(getContactByID(id))
        }
    }, []);

    return (
        <div className={style.login_page}>
            <Backdrop>
                <h2>Edit Contact</h2>
                <form className={style.login_form} noValidate>
                    <FormControl sx={{ m: 1, width: '40ch' }} variant="outlined">
                        <TextField id="outlined-basic" defaultValue={selectedContact?.Data.name} label="Name" variant="outlined" {...register("name", {
                            required: "Required",
                        })} />
                    </FormControl>
                    <FormControl sx={{ m: 1, width: '40ch' }} variant="outlined">
                        <TextField id="outlined-basic" defaultValue={selectedContact?.Data.phone} label="Phone" variant="outlined" {...register("phone", {
                            required: "Required",
                        })} />
                    </FormControl>
                    <FormControl sx={{ m: 1, width: '40ch' }} variant="outlined">
                        <TextField id="outlined-basic" defaultValue={selectedContact?.Data.email} label="Email" variant="outlined" {...register("email", {
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