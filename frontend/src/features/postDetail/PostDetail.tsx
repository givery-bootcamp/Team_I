import React, { useState, useEffect } from 'react';
import { Post } from '../../shared/models/Post';
import { useParams, Link } from 'react-router-dom';
import { fetchPosts } from '../../shared/services/mockApiService'; // モックAPIサービスをインポート

const PostDetail: React.FC = () => {
    const [posts, setPosts] = useState<Post[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const { postId } = useParams<{ postId: string }>();

    // ページが読み込まれた時に実行
    useEffect(() => {
        const getPosts = async () => {
            try {
                // モックAPIサービスを呼び出し
                const data = await fetchPosts();
                setPosts(data);
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

        getPosts();
    }, []);

    // ローディング中の場合
    if (loading) {
        return <div className="flex justify-center items-center h-full"><span className="text-lg font-semibold">Loading...</span></div>;
    }

    // エラーが発生した場合
    if (error) {
        return <div className="flex justify-center items-center h-full"><span className="text-lg font-semibold text-red-500">Error: {error}</span></div>;
    }

    // パラメータで指定されたIDの投稿を取得
    const post = postId? posts.find(post => post.id === parseInt(postId, 10)):undefined;

    // 投稿が見つからない場合
    if (!post) {
        return <div className="flex justify-center items-center h-full"><span className="text-lg font-semibold text-red-500">Post not found</span></div>;
    }

    return(
        // 投稿詳細を表示
        <div className="p-6 bg-white shadow-lg rounded-lg">
            <h2 className="text-2xl font-bold text-gray-800 mb-2">{post.title}</h2>
            <p className="text-gray-600">ユーザー名: <span className="font-semibold">{post.username}</span></p>
            <p className="text-gray-500">更新日: {post.updatedAt}</p>
            <p className="text-gray-500 mt-4">{post.content}</p>
            <Link to="/" className="text-blue-500 mt-4 block">Back to home</Link>
        </div>
    );
}

export default PostDetail;