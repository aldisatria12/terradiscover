import React from "react";
import { TextField } from "@mui/material";

export const RegisterPage: React.FC = () => {
    return (
        <div>
            <TextField id="outlined-basic" label="Username" variant="outlined" />
            <TextField id="outlined-basic" label="Email" variant="outlined" />
            <TextField id="outlined-basic" label="Password" variant="outlined" />
        </div>
    )
}