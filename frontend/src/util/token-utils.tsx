import {jwtDecode} from "jwt-decode";

export function getUserId(): number {
    if (isLoggedIn()) {
        const userString: any = localStorage.getItem('user');
        const user = JSON.parse(userString);
        // @ts-ignore
        return jwtDecode(user.accessToken).id;
    }
    return -1;
}

export function isLoggedIn(): boolean {
    return localStorage.getItem('user') != null;
}

export function getUserMail(): string {
    if (isLoggedIn()) {
        const userString: any = localStorage.getItem('user');
        const user = JSON.parse(userString);
        // @ts-ignore
        return jwtDecode(user.accessToken).sub;
    }
    return "";
}

export function getRole(): string | null {
    if (isLoggedIn()) {
        const userString = localStorage.getItem('user');
        if (!userString) return null;

        try {
            const user = JSON.parse(userString);
            const decoded: any = jwtDecode(user.accessToken || user.token); // zavisi kako si nazvao polje
            return decoded.role;
        } catch (e) {
            console.error("Token decode error:", e);
            return null;
        }
    }
    return null;
}

export function getUserWallet(): string {
    if (isLoggedIn()) {
        const userString: any = localStorage.getItem('user');
        const user = JSON.parse(userString);
        // @ts-ignore
        return jwtDecode(user.accessToken).wallet;
    }
    return "";
}

export function getUserPasswordResetDate(): number {
    if (isLoggedIn()) {
        const userString: any = localStorage.getItem('user');
        const user = JSON.parse(userString);
        // @ts-ignore
        return jwtDecode(user.accessToken).prdate;
    }
    return -1;
}

export function logout() {
    localStorage.removeItem('user');
}