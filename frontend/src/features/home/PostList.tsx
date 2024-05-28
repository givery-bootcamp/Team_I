import React, { useState, useEffect } from 'react';
import { Post } from '../../shared/models/Post';
import { fetchPosts } from '../../shared/services/mockApiService'; // モックAPIサービスをインポート

const PostList: React.FC = () => {
    const [posts, setPosts] = useState<Post[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const getPosts = async () => {
            try {
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

    if (loading) {
        return <div className="flex justify-center items-center h-full"><span className="text-lg font-semibold">Loading...</span></div>;
    }

    if (error) {
        return <div className="flex justify-center items-center h-full"><span className="text-lg font-semibold text-red-500">Error: {error}</span></div>;
    }

    return (
        <div className="p-6 bg-white shadow-lg rounded-lg">
            {posts.map(post => (
                <div key={post.id} className="border-b mb-6 pb-4 last:border-b-0 last:mb-0 last:pb-0">
                    <h2 className="text-2xl font-bold text-gray-800 mb-2">{post.title}</h2>
                    <p className="text-gray-600">ユーザー名: <span className="font-semibold">{post.username}</span></p>
                    <p className="text-gray-500">更新日: {post.updatedAt}</p>
                </div>
            ))}
        </div>
    );
};

export default PostList;
