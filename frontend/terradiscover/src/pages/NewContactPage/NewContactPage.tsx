import React from "react";
import { Button, FormControl, TextField } from "@mui/material";
import { useDispatch } from "react-redux";
import { useNavigate } from "react-router-dom";
import { AppDispatch } from "../../store/store";
import { useForm } from "react-hook-form";
import style from "./NewContactPage.module.css";
import { Backdrop } from "../../component/UI/backdrop";
import { inputContact } from "../../constants/types/typeContact";
import { insertContact } from "../../store/contactSlice";

export const NewContactPage: React.FC = () => {
    const { register, handleSubmit } = useForm();

    const dispatch = useDispatch<AppDispatch>();
    const navigate = useNavigate();

    const clickSubmit = async (data: any) => {
        const newContact: inputContact = data;
        try {
            await dispatch(insertContact(newContact));
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
                <h2>Create Contact</h2>
                <form className={style.login_form} noValidate onSubmit={handleSubmit((data: any) => clickSubmit(data))}>
                    <FormControl sx={{ m: 1, width: '40ch' }} variant="outlined">
                        <TextField id="outlined-basic" label="Name" variant="outlined" {...register("name", {
                            required: "Required",
                        })} />
                    </FormControl>
                    <FormControl sx={{ m: 1, width: '40ch' }} variant="outlined">
                        <TextField id="outlined-basic" label="Phone" variant="outlined" {...register("phone", {
                            required: "Required",
                        })} />
                    </FormControl>
                    <FormControl sx={{ m: 1, width: '40ch' }} variant="outlined">
                        <TextField id="outlined-basic" label="Email" variant="outlined" {...register("email", {
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