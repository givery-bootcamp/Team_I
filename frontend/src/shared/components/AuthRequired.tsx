// src/components/AuthRequired.tsx
import React from 'react';
import { Navigate } from 'react-router-dom';
import { useAuth } from './AuthContext';

interface AuthRequiredProps {
    children: React.ReactNode;
}

const AuthRequired: React.FC<AuthRequiredProps> = ({ children }) => {
    const { isLoggedIn } = useAuth();
    
    if (!isLoggedIn) {
        return <Navigate replace to="/signin"></Navigate>;
    }

    return <>{children}</>;
};

export default AuthRequired;