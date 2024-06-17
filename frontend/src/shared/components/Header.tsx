import React from 'react';
import {useNavigate} from 'react-router-dom';
import SignInStatus from './SignInStatus';
import {useAuth} from "./AuthContext.tsx";

const Header: React.FC = () => {
    const navigate = useNavigate();
    const {isLoggedIn} = useAuth();

    const handleSignInClick = () => {
        navigate('/signIn');
    };

    return (
        <header className="bg-gray-100 p-4 shadow-md flex justify-between items-center">
            <h1 className="text-xl font-bold">あおいファンクラブ</h1>
            <div className="flex items-center">
                <SignInStatus/>
                {!isLoggedIn && (
                    <button onClick={handleSignInClick} className="ml-4 btn">
                        サインイン
                    </button>
                )}
            </div>
        </header>
    );
};

export default Header;
