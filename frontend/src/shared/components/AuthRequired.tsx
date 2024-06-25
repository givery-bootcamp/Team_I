// src/components/AuthRequired.tsx
import React from 'react';
import { Navigate } from 'react-router-dom';
import { useAuth } from '../context/useAuth';

interface AuthRequiredProps {
    children: React.ReactNode;
}

const AuthRequired: React.FC<AuthRequiredProps> = ({ children }) => {
    const { isLoggedIn, isCheckingAuth } = useAuth();
    
    if (isCheckingAuth) {
        return <div>ユーザ認証中...</div>;
    }
    
    if (!isLoggedIn) {
        return <Navigate replace to="/signin"></Navigate>;
    }

    return <>{children}</>;
};

export default AuthRequired;
