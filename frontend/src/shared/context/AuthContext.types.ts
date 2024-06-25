import {User} from "../models/User.ts";
import {ReactNode} from "react";

export interface IAuthContext {
    isLoggedIn: boolean;
    isCheckingAuth: boolean;
    user: User | null;
    signIn: (user: User) => void;
    signOut: () => void;
    setIsCheckingAuth: (isCheckingAuth: boolean) => void;
}

export interface AuthProviderProps {
    children: ReactNode;
}
