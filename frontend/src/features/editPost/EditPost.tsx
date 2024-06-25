import React, {useEffect, useState} from 'react';
import {useNavigate, useParams} from 'react-router-dom';
import {SubmitHandler, useForm} from 'react-hook-form';
import {fetchPostById, updatePost} from '../../shared/services/apiService'; // APIサービスの関数をインポート
import {useAuth} from '../../shared/context/useAuth';
import {IFormInput} from '../../shared/models/Post';
import { useAlert } from '../../shared/components/AlertContext';


const EditPost: React.FC = () => {
    const {postId} = useParams<{ postId: string }>();
    const {user} = useAuth();
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const {register, handleSubmit, setValue, formState: {errors}} = useForm<IFormInput>();
    const navigate = useNavigate();
    const {showAlert} = useAlert();

    useEffect(() => {
        const getPost = async () => {
            try {
                const postIdNumber = postId && parseInt(postId, 10);
                if (!postIdNumber) {
                    setError('Post not found');
                    return;
                }

                const post = await fetchPostById(postIdNumber); // 単一の投稿を取得するAPI関数を使用
                if (!post) {
                    setError('Post not found');
                    return;
                }
                if (post.username !== user?.name) {
                    navigate('/');
                    return;
                }
                setValue('title', post.title);
                setValue('content', post.body);
            } catch (err) {
                if (err instanceof Error) {
                    setError(err.message);
                } else {
                    setError('An unexpected error occurred');
                }
            } finally {
                setLoading(false);
            }
        };

        getPost();
    }, [postId, user, navigate, setValue]);

    const onSubmit: SubmitHandler<IFormInput> = async (data) => {
        try {
            await updatePost(data, postId);
            showAlert('投稿が更新されました。');
            setTimeout(() => {
                navigate(`/posts/${postId}`);
            }, 1000);
        } catch (err) {
            if (err instanceof Error) {
                setError(err.message);
            } else {
                setError('An unexpected error occurred');
            }
        }
    };

    const handleCancel = () => {
        navigate(`/posts/${postId}`);
    };

    if (loading) {
        return <div className="flex justify-center items-center h-full"><span
            className="text-lg font-semibold">Loading...</span></div>;
    }

    if (error) {
        return <div className="flex justify-center items-center h-full"><span
            className="text-lg font-semibold text-red-500">Error: {error}</span></div>;
    }

    return (
        <div className="p-6 bg-white shadow-lg rounded-lg">
            <h2 className="text-2xl font-bold text-gray-800 mb-2">投稿の編集</h2>
            <form onSubmit={handleSubmit(onSubmit)}>
                <div className="mb-4">
                    <label className="block text-gray-600 mb-2">タイトル</label>
                    <input
                        type="text"
                        className={`w-full p-2 border ${errors.title ? 'border-red-500' : 'border-gray-300'} rounded-md`}
                        {...register('title', {required: 'タイトルは必須です。'})}
                    />
                    {errors.title && <div className="text-red-500 mt-2">{errors.title.message}</div>}
                </div>
                <div className="mb-4">
                    <label className="block text-gray-600 mb-2">内容</label>
                    <textarea
                        className={`w-full p-2 border ${errors.content ? 'border-red-500' : 'border-gray-300'} rounded-md`}
                        {...register('content', {required: '内容は必須です。'})}
                    />
                    {errors.content && <div className="text-red-500 mt-2">{errors.content.message}</div>}
                </div>
                <div className="flex justify-between">
                    <button type="button" onClick={handleCancel}
                            className="px-4 py-2 bg-gray-500 text-white rounded-md hover:bg-gray-600">キャンセル
                    </button>
                    <button type="submit"
                            className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600">変更を保存する
                    </button>
                </div>
                {error && <div className="text-red-500 mt-4">{error}</div>}
            </form>
        </div>
    );
};

export default EditPost;
