import React from "react";
import { Button, FormControl, TextField } from "@mui/material";
import { useForm } from "react-hook-form";
import style from "./EditContactPage.module.css";
import { Backdrop } from "../../component/UI/backdrop";

export const EditContactPage: React.FC = () => {
    const { register } = useForm();

    return (
        <div className={style.login_page}>
            <Backdrop>
                <h2>Edit Contact</h2>
                <form className={style.login_form} noValidate>
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