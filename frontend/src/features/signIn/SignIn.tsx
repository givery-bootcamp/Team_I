import React, {useState} from 'react';
import {useNavigate} from "react-router-dom";
import {SubmitHandler, useForm} from "react-hook-form"
import {useAuth} from "../../shared/components/AuthContext.tsx";

interface IFormInput{
    userName: string
    password: string
}

// テストユーザー情報
const TEST_USER = {
    userName: 'testuser',
    password: 'testpassword',
};


// const API_URL = 'https://team-9.member0005.track-bootcamp.run/signin'
const API_URL = 'http://localhost:9000/signin'

const SignIn: React.FC = () => {
    // ユーザ名・パスワードのステートを
    const [signinError, setSignInError] = useState<string | null>(null);
    const navigate = useNavigate();
    const {signIn, setIsCheckingAuth} = useAuth();

    // const {register, handleSubmit} = useForm<IFormInput>()
    const { register, handleSubmit, formState: { errors }, } = useForm<IFormInput>();
    
    const onSubmit: SubmitHandler<IFormInput> = async (data) => {
        // バックエンドを呼ばずにテストユーザーでのサインインを検証
        if (data.userName === TEST_USER.userName && data.password === TEST_USER.password) {
            signIn(data.userName);
            navigate('/');
        }

        //TODO: サブミットした時の処理を書くぞ

        // userName, passwordをバックエンドにpost
        // data.name = data.userName;
        const sendData = {
            name: data.userName,
            password: data.password,
        }
        try {
            // ユーザ認証開始
            setIsCheckingAuth(true);
            const response = await fetch(API_URL, {
                method: 'POST',
                headers: {
                    // 'Accept': 'application/json',
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(sendData),
                credentials: 'include',
            })
            if (!response.ok) {
                // 失敗したらサインインエラー
                setSignInError(await response.json())
                return;
            }
            
            // jwtを保存
            const jwt = response.headers.get('Authorization');
            if (jwt) {
                localStorage.setItem('jwt', jwt);
            }

            // ユーザを保存
            const signInUser = await response.json();
            console.log(signInUser)

            signIn(signInUser.userName);

            // ログインが成功したらホーム
            navigate("/");     
          } catch (e) {
            console.error(e);
          } finally {
            // ユーザ認証終了
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
