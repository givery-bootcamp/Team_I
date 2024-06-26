import React, {useState} from 'react';
import {useNavigate} from "react-router-dom";
import {SubmitHandler, useForm} from "react-hook-form"
import { useAuth } from "../../shared/context/useAuth.ts";
import { IAuthContext } from "../../shared/context/AuthContext.types.ts";
import {signIn as apiSignIn} from "../../shared/services/apiService.ts";
import {User} from "../../shared/models/User.ts";


interface IFormInput {
    userName: string,
    password: string,
}


const SignIn: React.FC = () => {
    const [signinError, setSigninError] = useState<string | undefined>(undefined);
    const navigate = useNavigate();
    const { signIn, setIsCheckingAuth } = useAuth() as IAuthContext;

    const { register, handleSubmit, formState: { errors }, } = useForm<IFormInput>();
    
    const onSubmit: SubmitHandler<IFormInput> = async (data) => {
        const sendData = {
            name: data.userName,
            password: data.password,
        }
        try {
            setIsCheckingAuth(true);
            const response = await apiSignIn(sendData);
            const user: User = {
                id: response.id,
                name: response.name,
            }
            console.log(user)

            signIn(user);

            // ログインが成功したらホーム
            navigate("/");     
          } catch (e) {
            console.error(e);
            setSigninError('ユーザー名またはパスワードが間違っています。');
          } finally {
            setIsCheckingAuth(false);
          }
    }

    // 提出ボタンを押さないとバリデーションがかかりません。が、面倒なので、やりません。
    return (
        <div className="p-6 bg-white shadow-lg rounded-lg">
            {/* <form onSubmit={handleSubmit}> */}
            <form onSubmit={handleSubmit(onSubmit)}> 
                <div className="mb-4">
                <label className="block text-gray-600 mb-2">ユーザー名</label>
                <input
                    type="text"
                    className={`w-full p-2 border ${errors.userName ? 'border-red-500' : 'border-gray-300'} rounded-md`}
                    {...register('userName', {
                        required: 'ユーザー名は必須です。',
                        pattern: {
                        value: /^[a-zA-Z0-9]+$/,
                        message: 'ユーザー名は英数字のみです。記号は使用できません。',
                        },
                    })}
                />
                    {/* ユーザー名エラーメッセージ */}
                    {errors.userName && <div className="text-red-500 mt-2">{errors.userName.message}</div>}
                </div>
                <div className="mb-4">
                <label className="block text-gray-600 mb-2">パスワード</label>
                <input
                    type="password"
                    className={`w-full p-2 border ${errors.password ? 'border-red-500' : 'border-gray-300'} rounded-md`}
                    {...register('password', {
                        required: 'パスワードは必須です。',
                        pattern: {
                        value: /^[\x20-\x7E]+$/,
                        message: 'パスワードはASCII範囲の英数記号のみだお',
                        },
                    })}
                />
                    {/* パスワードエラーメッセージ */}
                    {errors.password && <div className="text-red-500 mt-2">{errors.password.message}</div>}
                </div>
                <button type="submit" className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600">サインイン</button>
                {signinError && <div className='text-red-500'>{signinError}</div>}
            </form>
        </div>
    );
};

export default SignIn;
