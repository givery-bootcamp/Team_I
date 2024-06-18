import React, {useEffect, useState} from 'react';
import {Post} from '../../shared/models/Post';
import {Link, useParams} from 'react-router-dom';
import {fetchPostById} from '../../shared/services/apiService';

const PostDetail: React.FC = () => {
    const [post, setPost] = useState<Post | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const { postId } = useParams<{ postId: string }>();


    // ページが読み込まれた時に実行
    useEffect(() => {
        const getPost = async () => {
            try {
                const data = await fetchPostById(parseInt(postId!, 10));
                console.log(data);
                setPost(data);
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

        if (postId) {
            getPost();
        }
    }, [postId]);

    // ローディング中の場合
    if (loading) {
        return <div className="flex justify-center items-center h-full"><span className="text-lg font-semibold">Loading...</span></div>;
    }

    // エラーが発生した場合
    if (error) {
        return <div className="flex justify-center items-center h-full"><span className="text-lg font-semibold text-red-500">Error: {error}</span></div>;
    }

    // 投稿が見つからない場合
    if (!post) {
        return <div className="flex justify-center items-center h-full"><span className="text-lg font-semibold text-red-500">Post not found</span></div>;
    }

    return (
        <div className="p-6 bg-white shadow-lg rounded-lg">
            <h2 className="text-2xl font-bold text-gray-800 mb-2">{post.title}</h2>
            <p className="text-gray-600">ユーザー名: <span className="font-semibold">{post.username}</span></p>
            <p className="text-gray-500">更新日: {post.updated_at}</p>
            <p className="text-gray-500 mt-4">{post.body}</p>
            <Link to="/" className="text-blue-500 mt-4 block">Back to home</Link>

            {error && <div className="text-red-500 mt-4">{error}</div>}
        </div>
    );
}

export default PostDetail;