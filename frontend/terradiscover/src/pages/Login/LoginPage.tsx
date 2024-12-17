import React from "react";
import style from './LoginPage.module.css';
import { Button, FormControl, IconButton, InputAdornment, OutlinedInput, TextField } from "@mui/material";
import InputLabel from '@mui/material/InputLabel';
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';
import { useForm } from "react-hook-form"
import { inputLogin } from "../../constants/types/typeAuth";
import { login } from "../../store/authSlice";
import { useDispatch } from "react-redux";
import { AppDispatch } from "../../store/store";
import { useNavigate } from "react-router-dom";

export const LoginPage: React.FC = () => {
    const { register, handleSubmit } = useForm()


    const dispatch = useDispatch<AppDispatch>();
    const navigate = useNavigate();

    const [showPassword, setShowPassword] = React.useState(false);

    const handleClickShowPassword = () => setShowPassword((show) => !show);

    const handleMouseDownPassword = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
    };

    const handleMouseUpPassword = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
    };

    const clickSubmit = async (data: any) => {
        const user: inputLogin = data;
        try {
            await dispatch(login(user));
        } catch (error) {
            console.log(error);
        } finally {
            navigate("/");
        }
    };

    return (
        <div className={style.login_page}>
            <form className={style.login_form} noValidate onSubmit={handleSubmit((data: any) => clickSubmit(data))}>
                <FormControl sx={{ m: 1, width: '40ch' }} variant="outlined">
                    <TextField id="outlined-basic" label="Email" variant="outlined" {...register("email", {
                        required: "Required",
                    })} />
                </FormControl>
                <FormControl sx={{ m: 1, width: '40ch' }} variant="outlined">
                    <InputLabel htmlFor="outlined-adornment-password">Password
                    </InputLabel>
                    <OutlinedInput id="outlined-basic" type={showPassword ? 'text' : 'password'}
                        endAdornment={
                            <InputAdornment position="end">
                                <IconButton
                                    aria-label={
                                        showPassword ? 'hide the password' : 'display the password'
                                    }
                                    onClick={handleClickShowPassword}
                                    onMouseDown={handleMouseDownPassword}
                                    onMouseUp={handleMouseUpPassword}
                                    edge="end"
                                >
                                    {showPassword ? <VisibilityOff /> : <Visibility />}
                                </IconButton>
                            </InputAdornment>
                        } label="Password"  {...register("password", {
                            required: "Required",
                        })} />
                </FormControl>
                <Button type="submit">
                    Submit
                </Button>
            </form>
        </div >
    )
}