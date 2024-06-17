import React from 'react';
import {useAuth} from './AuthContext.tsx';

const SignInStatus: React.FC = () => {
    const {isLoggedIn, userName, signOut} = useAuth();

    return (
        <div className="flex items-center space-x-4">
            {isLoggedIn ? (
                <>
                    <span>{userName}</span>
                    <button onClick={signOut} className="btn">
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
