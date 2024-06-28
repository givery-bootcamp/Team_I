import React, {useEffect, useState} from 'react';
import {Post} from '../../shared/models/Post';
import {Link} from 'react-router-dom';
import {fetchPosts} from '../../shared/services/apiService';


type FetchPostsOptions = {
    page?: number;
    limit?: number;
    type?: 'official' | 'yamada';
}


type PostListProps = Partial<FetchPostsOptions>;

const PostList: React.FC<PostListProps> = ({ type }) => {
    const [posts, setPosts] = useState<Post[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    const POSTS_PER_PAGE = 7;
    const [page, setPage] = useState<number>(1);

    useEffect(() => {
        const getPosts = async () => {
            try {
                setLoading(true);
                const data = await fetchPosts({ type, page, limit: POSTS_PER_PAGE });
                setPosts(data);
            } catch (err) {
                setError('An unexpected error occurred');
            } finally {
                setLoading(false);
            }
        };

        getPosts();
    }, [type, page]);

    if (loading) {
        return <div className="flex justify-center items-center h-full"><span className="text-lg font-semibold">Loading...</span></div>;
    }

    if (error) {
        return <div className="flex justify-center items-center h-full"><span className="text-lg font-semibold text-red-500">Error: {error}</span></div>;
    }

    return (
        <div className="p-6 bg-white shadow-lg rounded-lg">
            {posts.map(post => (
                <Link to={`/posts/${post.id}`} key={post.id}
                      className='block border-b mb-6 last:border-b-0 last:mb-0 last:pb-0 hover:bg-gray-100 transition-colors duration-200'>
                    <div key={post.id} className="border-b last:border-b-0 last:mb-0 last:pb-0">
                        <h2 className="text-2xl font-bold text-gray-800 mb-2">{post.title}</h2>
                        <p className="text-gray-600">ユーザー名: <span className="font-semibold">{post.username}</span></p>
                        <p className="text-gray-500">更新日: {post.updated_at}</p>
                    </div>
                </Link>
            ))}
            <div className="mt-4 d-flex justify-content-center">
                <button className="mx-2" onClick={() => setPage(prevPage => Math.max(prevPage - 1, 1))} disabled={page === 1}>
                    Previous Page: {page - 1}
                </button>
                <span>Current Page: {page}</span>
                <button className="mx-2" onClick={() => setPage(prevPage => prevPage + 1)}>
                    Next Page: {page + 1}
                </button>
            </div>
        </div>
    );
};

export default PostList;
