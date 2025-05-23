import LoginForm from "../../components/auth/login.tsx";
import "./login-registration.css"

export default function Login() {
    return (
        <div id={"main-container"}>
            <div id={"logo-slogan"}>
                <h1 id={"slogan"}>Your Home, Your Way</h1>
            </div>
            <LoginForm/>
        </div>
    )
}