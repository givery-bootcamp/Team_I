import React from 'react';
import {useAuth} from './AuthContext.tsx';
import API_BASE_URL from "../../config.ts";

const SignInStatus: React.FC = () => {
    const {isLoggedIn, userName, signOut} = useAuth();

    const handleSignOut = async () => {
        try {
            const response = await fetch(`${API_BASE_URL}/signout`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                credentials: 'include',
            });
            if (!response.ok) {
                throw new Error('Sign out failed');
            }
        } catch (error) {
            console.error(error);
        }

        signOut();
    }

    return (
        <div className="flex items-center space-x-4">
            {isLoggedIn ? (
                <>
                    <span>{userName}</span>
                    <button onClick={handleSignOut} className="btn">
                        サインアウト
                    </button>
                </>
            ) : (
                <span></span>
            )}
        </div>
    );
};

export default SignInStatus;
