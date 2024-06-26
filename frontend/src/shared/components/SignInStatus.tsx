import React from 'react';
import { useAuth } from '../context/useAuth.ts'
import {signOut as apiSignOut} from "../services/apiService.ts";
import {useNavigate} from "react-router-dom";


const SignInStatus: React.FC = () => {
    const {isLoggedIn, user, signOut} = useAuth();
    const navigate = useNavigate();


    const handleSignOut = async () => {
        await apiSignOut();
        signOut();
    }

    const handleSignIn = async () => {
        navigate('/signin');
    }

    return (
        <div className="flex items-center space-x-4">
            {isLoggedIn ? (
                <>
                    <span>{user?.name}</span>
                    <button onClick={handleSignOut} className="btn">
                        サインアウト
                    </button>
                </>
            ) : (
                <>
                    <span></span>
                    <button onClick={handleSignIn} className="btn">
                        サインイン
                    </button>
                </>
            )}
        </div>
    );
};

export default SignInStatus;
