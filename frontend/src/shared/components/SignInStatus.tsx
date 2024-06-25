import React from 'react';
import {useAuth} from './AuthContext.tsx';
import {signOut as apiSignOut} from "../services/apiService.ts";


const SignInStatus: React.FC = () => {
    const {isLoggedIn, userName, signOut} = useAuth();

    const handleSignOut = async () => {
        await apiSignOut();

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
