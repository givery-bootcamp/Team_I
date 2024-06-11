import React from 'react';
import { useState } from 'react';
import { useNavigate } from "react-router-dom";
import { useForm, SubmitHandler } from "react-hook-form"

interface IFormInput{
    userName: string
    password: string
}


const SignIn: React.FC = () => {
    // ユーザ名・パスワードのステートを
    const [signinError, setSignInError] = useState<string | null>(null);
    const navigate = useNavigate();

    // const {register, handleSubmit} = useForm<IFormInput>()
    const { register, handleSubmit, formState: { errors }, } = useForm<IFormInput>();
    const onSubmit: SubmitHandler<IFormInput> = (data) => {
        //TODO: サブミットした時の処理を書くぞ  
        
        // サインイン
        // 失敗したらサインインエラー
        // ログインが成功したらホーム
        const result = true;
        if (result) {
            navigate("/");
        } else {
            setSignInError("サインインが失敗したお")
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
