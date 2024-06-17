import React from 'react';
import { useRef } from 'react';
import { useNavigate } from "react-router-dom";
import { useForm, SubmitHandler } from "react-hook-form";

// const API_URL = 'https://team-9.member0005.track-bootcamp.run/posts';
const API_URL = 'http://localhost:3000/posts';

interface IFormInput {
    title: string;
    content: string;
}

const NewPost: React.FC = () => {
    const navigate = useNavigate();
    const [newPostError, setNewPostError] = React.useState<string | null>(null);
    const { register, handleSubmit, formState: { errors }, } = useForm<IFormInput>();

    // ボタン連打防止用のフラグ
    const isProcessing = useRef(false);
    const onSubmit: SubmitHandler<IFormInput> = async (data) => {
        // 処理中なら何もしない
        if (isProcessing.current) {
            return;
        }

        try {
            // 処理開始
            isProcessing.current = true;
            const response = await fetch(API_URL, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data)
            });

            if (!response.ok) {
                // 失敗したらエラー
                setNewPostError(await response.json());
                return;
            }

            
            // 成功したらアラート
            alert('投稿しました。');

            // 成功したら投稿一覧画面に戻る
            navigate("/");
        } catch (e) {
            console.error(e);
        } finally {
            // 処理終了
            isProcessing.current = false;
        }
    }

    return (
        <div className="p-6 bg-white shadow-lg rounded-lg">
            <h1 className="text-2xl font-bold text-gray-800 mb-4">新規投稿</h1>
            <form onSubmit={handleSubmit(onSubmit)}>
                <div className="mb-4">
                    <label className="block text-gray-600 mb-2">タイトル</label>
                    <input
                        type="text"
                        className={`w-full p-2 border ${errors.title ? 'border-red-500' : 'border-gray-300'} rounded-md`}
                        {...register('title', {
                            required: 'タイトルは必須です。',
                            pattern: {
                            value: /^.{0,100}$/,
                            message: 'タイトルは100文字以内です。',
                            },
                        })}
                    />
                    {/* タイトルエラーメッセージ */}
                    {errors.title && <div className="text-red-500 mt-2">{errors.title.message}</div>}
                </div>
                <div className="mb-4">
                    <label className="block text-gray-600 mb-2">内容</label>
                    <textarea
                        className={`w-full p-2 border ${errors.content ? 'border-red-500' : 'border-gray-300'} rounded-md`}
                        {...register('content', {
                            required: '内容は必須です。',
                        })}
                    />
                    {/* 内容エラーメッセージ */}
                    {errors.content && <div className="text-red-500 mt-2">{errors.content.message}</div>}
                </div>
                <button type="submit" className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600">投稿する</button>
                {newPostError && <div className='text-red-500'>{newPostError}</div>}
            </form>
        </div>
    );
};

export default NewPost;
