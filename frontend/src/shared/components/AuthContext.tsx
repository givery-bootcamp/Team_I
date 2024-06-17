// src/shared/context/AuthContext.tsx
import React, {createContext, ReactNode, useContext, useState} from 'react';

interface AuthContextType {
    isLoggedIn: boolean;
    userName: string | null;
    signIn: (userName: string) => void;
    signOut: () => void;
}

interface AuthProviderProps {
    children: ReactNode;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider: React.FC<AuthProviderProps> = ({children}) => {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [userName, setUserName] = useState<string | null>(null);

    const signIn = (userName: string) => {
        setIsLoggedIn(true);
        setUserName(userName);
    };

    const signOut = () => {
        setIsLoggedIn(false);
        setUserName(null);
    };

    return (
        <AuthContext.Provider value={{isLoggedIn, userName, signIn, signOut}}>
            {children}
        </AuthContext.Provider>
    );
};

export const useAuth = () => {
    const context = useContext(AuthContext);
    if (!context) {
        throw new Error('useAuth must be used within an AuthProvider');
    }
    return context;
};
