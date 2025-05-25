import {Button, IconButton, InputAdornment, TextField} from "@mui/material";
import * as yup from 'yup';
import {useFormik} from "formik";
import {useNavigate} from "react-router-dom";
import "./login-registration-form.css";
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';
import {useState} from "react";
import api from "../../config/axios-config.tsx";
import toast from "react-hot-toast";
import type { UserTokenState } from "../../model/user-model.tsx";

export default function LoginForm() {

    const navigate = useNavigate()
    const [showPassword, setShowPassword] = useState(false);
    const handleClickShowPassword = () => setShowPassword(!showPassword);

    const handleRegister = () => {
        navigate('/registration')
    }

    // @ts-ignore
    const validationSchema = yup.object({
        email: yup
            .string()
            .test('is-email', 'Enter a valid email', (value) => {
                if (value === 'admin') {
                    return true;
                }

                try {
                    yup.string().email().validateSync(value, { abortEarly: true });
                    return true;
                } catch (error) {
                    return false;
                }
            })
            .required('Email is required'),
        password: yup
            .string()
            .min(3, 'Password should be of minimum 8 characters length')
            .required('Password is required'),
    });

    const formik = useFormik({
        initialValues: {
            email: '',
            password: '',
        },
        validationSchema: validationSchema,
        onSubmit: (values) => {
            api.post<UserTokenState>('auth/login', {
                email: values.email,
                password: values.password,
                role: "user"
            }).then(res => {
                if (res.status === 200) {
                    localStorage.setItem('user', JSON.stringify(res.data));

                    navigate('/home');
                }
            }).catch((error) => {
                toast.error(error.response.data.message);
            });
        },
    });

    return (
        <div id={"form-container"} style={{width: "50%", padding: "10%"}}>
        <div className={"headings"}>
                <h2>Welcome back to NFT-Ticket-Shop</h2>
                <h1>Login</h1>
        </div>

            <form onSubmit={formik.handleSubmit}>
                <TextField
                    fullWidth
                    id="email"
                    name="email"
                    label="Email"
                    value={formik.values.email}
                    onChange={formik.handleChange}
                    onBlur={formik.handleBlur}
                    error={formik.touched.email && Boolean(formik.errors.email)}
                    helperText={formik.touched.email && formik.errors.email}
                />
                <TextField
                    fullWidth
                    id="password"
                    name="password"
                    label="Password"
                    type={showPassword ? 'text' : 'password'}
                    value={formik.values.password}
                    onChange={formik.handleChange}
                    onBlur={formik.handleBlur}
                    error={formik.touched.password && Boolean(formik.errors.password)}
                    helperText={formik.touched.password && formik.errors.password}
                    InputProps={{
                        endAdornment: (
                            <InputAdornment position="end">
                                <IconButton
                                    aria-label="toggle password visibility"
                                    onClick={handleClickShowPassword}
                                >
                                    {showPassword ? <Visibility /> : <VisibilityOff />}
                                </IconButton>
                            </InputAdornment>
                        )
                    }}
                />
                <Button
                    sx={{
                        fontSize: "20px",
                        textTransform: "capitalize"
                    }}
                    variant="contained" fullWidth type="submit">
                    Login
                </Button>
            </form>
            <div>
                <p>Don't have an account? <span onClick={handleRegister}>Sign up</span> </p>
            </div>
        </div>
    );
}