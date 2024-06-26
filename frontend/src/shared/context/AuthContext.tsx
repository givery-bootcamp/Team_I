import React, {createContext, useEffect, useState} from "react";
import { IAuthContext, AuthProviderProps } from "./AuthContext.types";
import {User} from "../models/User.ts";
import {getUser} from "../services/apiService.ts";


export const AuthContext = createContext<IAuthContext | undefined>(undefined);

const AuthProvider: React.FC<AuthProviderProps> = ({children}) => {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [isCheckingAuth, setIsCheckingAuth] = useState(false);
    const [user, setUser] = useState<User | null>(null);

    const signIn = (user: User) => {
        setIsLoggedIn(true);
        setUser(user);
    };

    const signOut = () => {
        setIsLoggedIn(false);
        setUser(null);
    };

    useEffect(() => {
        getUser()
            .then(user => signIn(user))
            .catch(error => {
                console.error(error);
                signOut()
            })
            .finally(() => setIsCheckingAuth(false))
    }, []);

    return (
        <AuthContext.Provider value={{isLoggedIn,isCheckingAuth, user, signIn, signOut, setIsCheckingAuth}}>
            {children}
        </AuthContext.Provider>
    );
};

export default AuthProvider;
