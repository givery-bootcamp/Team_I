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

    const handleSignUp = async () => {
        navigate('/signup');
    }

    return (
        <div className="flex items-center space-x-4">
            {isLoggedIn ? (
                <>
                    <span className="mr-4 text-2xl">{user?.name}</span>
                    <button onClick={handleSignOut}
                            className="px-4 py-2 bg-gray-500 text-white rounded hover:bg-gray-600 shadow-lg transition duration-200 flex items-center">
                        <img src="/okazaki.png" className="h-6 mr-2" alt="Logo"/>
                        サインアウト
                    </button>
                </>
            ) : (
                <>
                    <button onClick={handleSignUp}
                            className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 shadow-lg transition duration-200 flex items-center">
                        <img src="/yamada.png" className="h-6 mr-2" alt="Logo"/>
                        サインアップ
                    </button>
                    <button onClick={handleSignIn}
                            className="ml-2 px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600 shadow-lg transition duration-200 flex items-center">
                        <img src="/nagatani.png" className="h-6 mr-2" alt="Logo"/>
                        サインイン
                    </button>
                </>
            )}
        </div>
    );
};

export default SignInStatus;
