import React, { useEffect } from "react";
import { Button, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, TextField } from "@mui/material";
import { useDispatch, useSelector } from "react-redux";
import { AppDispatch, RootState } from "../../store/store";
import { getContactList } from "../../store/contactSlice";
import AddIcon from '@mui/icons-material/Add';
import { FloatingButton } from "../../component/UI/floatingButton";
import { useNavigate } from "react-router-dom";
import EditIcon from '@mui/icons-material/Edit';

export const ContactPage: React.FC = () => {
    const { contactList } = useSelector(
        (state: RootState) => state.contactSlice
    );

    const dispatch = useDispatch<AppDispatch>();
    const navigate = useNavigate();

    useEffect(() => {
        dispatch(getContactList()).catch((_) => navigate("/login"));
    }, []);

    const handleClickToRegister = () => {
        navigate("/contact/create")
    }

    return (
        <div>
            <TableContainer component={Paper}>
                <Table sx={{ minWidth: 650 }} aria-label="simple table">
                    <TableHead>
                        <TableRow>
                            <TableCell align="right">No</TableCell>
                            <TableCell align="right">Name</TableCell>
                            <TableCell align="right">Phone</TableCell>
                            <TableCell align="right">Email</TableCell>
                            <TableCell align="right">Edit</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {contactList?.Data ? contactList.Data.map((row, n) => (
                            <TableRow
                                key={row.id}
                                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                            >
                                <TableCell component="th" scope="row">
                                    {n + 1}
                                </TableCell>
                                <TableCell align="right">{row.name}</TableCell>
                                <TableCell align="right">{row.phone}</TableCell>
                                <TableCell align="right">{row.email}</TableCell>
                                <TableCell align="right"><Button variant="contained" color="primary" onClick={() => {
                                    navigate("/contact/edit/" + row.id)
                                }}>Edit</Button></TableCell>
                            </TableRow>
                        )) : <></>}
                    </TableBody>
                </Table>
            </TableContainer>
            <FloatingButton color="primary" aria-label="add" onClick={handleClickToRegister}>
                <AddIcon />
            </FloatingButton>
        </div>
    )
}