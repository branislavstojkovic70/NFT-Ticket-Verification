import {AppBar, Button, Dialog, DialogContent, Toolbar} from "@mui/material";
import {Outlet, useNavigate} from "react-router-dom";
import {useState} from "react";
import {getRole, logout} from "../util/token-utils.tsx";
import {Login, Logout, PersonAdd, Share, ElectricBolt} from "@mui/icons-material";
import "../index.css";

export default function Navbar () {
    const navigate = useNavigate()
    const [open, setOpen] = useState(false);
    const role = getRole();


    const handleLogin = () => {
        navigate('/login')
    }

    const handleClickRegistrationAdminOpen = () => {
        setOpen(true);
    };

    const handleClickRegistrationAdminClose = () => {
        setOpen(false);
    };

    const handleLogout = () => {
        logout();
        navigate("/")
    }

    const handleClickElectricityUsage = () => {
        navigate('/electricityUsage')
    };
    
    const handleClickShared = () => {
        navigate("/shared")
    }

    return (
        <>
        <div style={{height:"100%", width:"100%", display: "flex", flexDirection: "column"}}>
            <AppBar position={'relative'} style={{flex:'0 1 auto'}}>
                <Toolbar>
                    <div style={{
                        width: "200px",
                        height: "100px"
                    }}>
                        <img id={"logo"} src="/img/logo-light.svg" alt="Logo"
                             onClick={() => navigate(`home`)}
                             style={{
                                 width: "100%",
                                 height: "100%",
                                 objectFit: "cover"
                             }} />
                    </div>

                    {!role &&
                        <Button onClick={handleLogin}
                                startIcon={<Login sx={{color: "white"}} />}
                                sx={{
                                    textTransform: "capitalize",
                                    color: "#F5F5F5",
                                    marginLeft: "auto"
                                }}>Login</Button>
                    }
                    {role != null && true && role.includes("ROLE_SUPERADMIN") &&
                        <Button onClick={handleClickRegistrationAdminOpen}
                                startIcon={<PersonAdd sx={{color: "white"}} />}
                                sx={{
                                    textTransform: "capitalize",
                                    color: "#F5F5F5",
                                }}>Register admin</Button>
                    }
                    {role != null && true && role.includes("ROLE_ADMIN") &&
                        <Button onClick={handleClickElectricityUsage}
                                startIcon={<ElectricBolt sx={{color: "white"}} />}
                                sx={{
                                    textTransform: "capitalize",
                                    color: "#F5F5F5",
                                }}>Electricity usage</Button>
                    }
                    {role != null && true && role.includes("ROLE_USER") &&
                        <Button onClick={handleClickShared}
                                startIcon={<Share sx={{color: "white"}} />}
                                sx={{
                                    textTransform: "capitalize",
                                    color: "#F5F5F5",
                                }}>Shared</Button>
                    }
                    {role &&
                        <Button onClick={handleLogout}
                                startIcon={<Logout sx={{color: "white"}} />}
                                sx={{
                                    textTransform: "capitalize",
                                    color: "#F5F5F5",
                                    marginLeft: "auto"
                                }}>Logout</Button>
                    }
                </Toolbar>

            </AppBar>
            {/* <Dialog open={open} onClose={handleClickRegistrationAdminClose} fullWidth>
                <DialogContent>
                    <RegistrationForm/>
                </DialogContent>
            </Dialog> */}
            <div id="detail" style={{flex:'1 1 auto', width:"100%"}}>
                <Outlet />
            </div>
        </div>
        </>
    )
}