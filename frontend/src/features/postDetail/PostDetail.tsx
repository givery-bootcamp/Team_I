import React, {useEffect, useState} from 'react';
import {Post} from '../../shared/models/Post';
import {Link, useNavigate, useParams} from 'react-router-dom';
import {deletePost, fetchPosts} from '../../shared/services/mockApiService'; // モックAPIサービスをインポート
import {useAuth} from '../../shared/components/AuthContext';

const PostDetail: React.FC = () => {
    const [posts, setPosts] = useState<Post[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const [message, setMessage] = useState<string | null>(null);
    const {postId} = useParams<{ postId: string }>();
    const navigate = useNavigate();
    const {userName} = useAuth();

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
        return <div className="flex justify-center items-center h-full"><span
            className="text-lg font-semibold">Loading...</span></div>;
    }

    // エラーが発生した場合
    if (error) {
        return <div className="flex justify-center items-center h-full"><span
            className="text-lg font-semibold text-red-500">Error: {error}</span></div>;
    }

    // パラメータで指定されたIDの投稿を取得
    const post = postId ? posts.find(post => post.id === parseInt(postId, 10)) : undefined;

    // 投稿が見つからない場合
    if (!post) {
        return <div className="flex justify-center items-center h-full"><span
            className="text-lg font-semibold text-red-500">Post not found</span></div>;
    }

    // 投稿を削除する関数
    const handleDelete = async () => {
        if (confirm('この投稿を削除しますか？')) {
            try {
                await deletePost(post.id);
                setMessage('投稿が削除されました');
                navigate('/');
            } catch (err) {
                if (err instanceof Error) {
                    setError(err.message);
                } else {
                    setError('An unexpected error occurred');
                }
            }
        }
    }

    return (
        // 投稿詳細を表示
        <div className="p-6 bg-white shadow-lg rounded-lg relative">
            <button onClick={() => navigate("/new-post")}
                    className="text-white bg-blue-500 hover:bg-blue-600 px-4 py-2 rounded absolute top-0 right-0 m-4">新しい投稿を作成
            </button>
            <h2 className="text-2xl font-bold text-gray-800 mb-2">{post.title}</h2>
            <p className="text-gray-600">ユーザー名: <span className="font-semibold">{post.username}</span></p>
            <p className="text-gray-500">更新日: {post.updatedAt}</p>
            <p className="text-gray-500 mt-4">{post.content}</p>
            <Link to="/" className="text-blue-500 mt-4 block">ホームに戻る</Link>

            {userName === post.username && (
                <div className="mt-4">
                    <Link to={`/posts/${post.id}/edit`} className="text-blue-500 mr-4">編集</Link>
                    <button onClick={handleDelete} className="text-red-500">削除</button>
                </div>
            )}

            {message && <div className="text-green-500 mt-4">{message}</div>}
            {error && <div className="text-red-500 mt-4">{error}</div>}
        </div>
    );
}

export default PostDetail;