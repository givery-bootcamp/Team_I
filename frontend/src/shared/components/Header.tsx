import React from 'react';
import SignInStatus from './SignInStatus';

const Header: React.FC = () => {

    return (
        <header className="bg-gray-100 p-4 shadow-md flex justify-between items-center">
            <h1 className="text-xl font-bold">あおいファンクラブ</h1>
            <div className="flex items-center">
                <SignInStatus/>
            </div>
        </header>
    );
};

export default Header;
