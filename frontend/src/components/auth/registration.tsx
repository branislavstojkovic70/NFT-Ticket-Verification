import {
    Button,
    MenuItem,
    TextField
} from "@mui/material";
import { useFormik } from "formik";
import * as yup from "yup";
import { useNavigate } from "react-router-dom";
import toast from "react-hot-toast";
import api from "../../config/axios-config.tsx";
import "./login-registration-form.css";

export default function RegistrationForm() {
    const navigate = useNavigate();

    const validationSchema = yup.object({
        email: yup.string().email("Enter a valid email").required("Email is required"),
        password: yup.string().min(8, "Minimum 8 characters").required("Password is required"),
        name: yup.string().required("Name is required"),
        surname: yup.string().required("Surname is required"),
        wallet: yup.string().required("Wallet address is required"),
        location: yup.string().required("Location is required"),
        age: yup.number().min(13, "You must be at least 13").required("Age is required"),
        gender: yup.string().oneOf(['Male', 'Female']).required("Gender is required"),
        interests: yup.string().required("Interests are required"),
    });

    const formik = useFormik({
        initialValues: {
            email: '',
            password: '',
            name: '',
            surname: '',
            wallet: '',
            location: '',
            age: '',
            gender: '',
            interests: '',
        },
        validationSchema: validationSchema,
        onSubmit: (values) => {
            api.post("auth/register", {
                ...values,
                interests: values.interests.split(',').map(s => s.trim()), // ako očekuje array
            }).then(res => {
                toast.success("Registration successful!");
                navigate("/login");
            }).catch(err => {
                toast.error(err?.response?.data?.message || "Registration failed");
            });
        }
    });

    return (
        <div id="form-container" style={{ width: "60%", padding: "5%" }}>
            <div className="headings">
                <h2>Create your account</h2>
                <h1>Register</h1>
            </div>

            <form onSubmit={formik.handleSubmit}>
                <TextField
                    fullWidth label="Email" id="email" name="email"
                    value={formik.values.email} onChange={formik.handleChange} onBlur={formik.handleBlur}
                    error={formik.touched.email && Boolean(formik.errors.email)}
                    helperText={formik.touched.email && formik.errors.email}
                />
                <TextField
                    fullWidth label="Password" id="password" name="password" type="password"
                    value={formik.values.password} onChange={formik.handleChange} onBlur={formik.handleBlur}
                    error={formik.touched.password && Boolean(formik.errors.password)}
                    helperText={formik.touched.password && formik.errors.password}
                />
                <TextField
                    fullWidth label="Name" id="name" name="name"
                    value={formik.values.name} onChange={formik.handleChange} onBlur={formik.handleBlur}
                    error={formik.touched.name && Boolean(formik.errors.name)}
                    helperText={formik.touched.name && formik.errors.name}
                />
                <TextField
                    fullWidth label="Surname" id="surname" name="surname"
                    value={formik.values.surname} onChange={formik.handleChange} onBlur={formik.handleBlur}
                    error={formik.touched.surname && Boolean(formik.errors.surname)}
                    helperText={formik.touched.surname && formik.errors.surname}
                />
                <TextField
                    fullWidth label="Wallet Address" id="wallet" name="wallet"
                    value={formik.values.wallet} onChange={formik.handleChange} onBlur={formik.handleBlur}
                    error={formik.touched.wallet && Boolean(formik.errors.wallet)}
                    helperText={formik.touched.wallet && formik.errors.wallet}
                />
                <TextField
                    fullWidth label="Location" id="location" name="location"
                    value={formik.values.location} onChange={formik.handleChange} onBlur={formik.handleBlur}
                    error={formik.touched.location && Boolean(formik.errors.location)}
                    helperText={formik.touched.location && formik.errors.location}
                />
                <TextField
                    fullWidth label="Age" id="age" name="age" type="number"
                    value={formik.values.age} onChange={formik.handleChange} onBlur={formik.handleBlur}
                    error={formik.touched.age && Boolean(formik.errors.age)}
                    helperText={formik.touched.age && formik.errors.age}
                />
                <TextField
                    fullWidth select label="Gender" id="gender" name="gender"
                    value={formik.values.gender} onChange={formik.handleChange} onBlur={formik.handleBlur}
                    error={formik.touched.gender && Boolean(formik.errors.gender)}
                    helperText={formik.touched.gender && formik.errors.gender}
                >
                    <MenuItem value="Male">Male</MenuItem>
                    <MenuItem value="Female">Female</MenuItem>
                    <MenuItem value="Other">Other</MenuItem>
                </TextField>
                <TextField
                    fullWidth label="Interests (comma-separated)" id="interests" name="interests"
                    value={formik.values.interests} onChange={formik.handleChange} onBlur={formik.handleBlur}
                    error={formik.touched.interests && Boolean(formik.errors.interests)}
                    helperText={formik.touched.interests && formik.errors.interests}
                />
                <Button
                    sx={{ fontSize: "20px", textTransform: "capitalize" }}
                    variant="contained" fullWidth type="submit">
                    Register
                </Button>
            </form>
            <div>
                <p>Already have an account? <span onClick={() => navigate('/login')}>Login</span></p>
            </div>
        </div>
    );
}
