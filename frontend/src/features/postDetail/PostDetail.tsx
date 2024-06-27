import React, {useEffect, useState} from 'react';
import {Post} from '../../shared/models/Post';
import {Link, useNavigate, useParams} from 'react-router-dom';
import {useAuth} from '../../shared/context/useAuth';
import {deletePost, fetchPostById, postIntention, fetchIntentionState} from '../../shared/services/apiService';
import {useAlert} from '../../shared/components/AlertContext';
import ConfirmModal from '../../shared/components/Modal';
import { useConfirmModal } from '../../shared/hooks/useConfirmModal';


export type IntentionState = 'attend' | 'skip'


const PostDetail: React.FC = () => {
    const [post, setPost] = useState<Post | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const {postId} = useParams<{ postId: string }>();
    const navigate = useNavigate();

    const { user } = useAuth();
    const userId = user?.id;

    const {showAlert} = useAlert();

    const [intention, setIntention] = useState<IntentionState | null>(null);
    const [attendees, setAttendees] = useState<string[]>([]);
    const [nonAttendees, setNonAttendees] = useState<string[]>([]);

    const [isHoveringAttendees, setIsHoveringAttendees] = useState(false);
    const [isHoveringNonAttendees, setIsHoveringNonAttendees] = useState(false);

    // Modalを表示するためのカスタムフック
    const {modalRef, confirmMessage, onConfirm, onCancel, customConfirm} = useConfirmModal();

    // ページが読み込まれた時に実行
    useEffect(() => {
        const getPost = async () => {
            try {
                // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
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

        // Load attendees
        // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
        fetchIntentionState(parseInt(postId!, 10), 'attend')
            .then(data => setAttendees(data.usernames)); // Adjust as per your API response

        // Load non-attendees
        // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
        fetchIntentionState(parseInt(postId!, 10), 'skip')
            .then(data => setNonAttendees(data.usernames)); // Adjust as per your API response
    }, [postId]);

    const handleIntention = (intention: IntentionState) => {
        // Make sure user id exists.
        if (!userId || !post) {
            return;
        }

        postIntention(post.id, intention, userId)
            .then(() => {
                setIntention(intention);
                intention === 'attend' ?
                    setAttendees(prev => [...prev, user?.name]) :
                    setNonAttendees(prev => [...prev, user?.name]);
            })
            .catch((error) => {
                console.error(error);
                setError('Failed to post response');
            });
    };

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

    // 投稿が見つからない場合
    if (!post) {
        return <div className="flex justify-center items-center h-full"><span
            className="text-lg font-semibold text-red-500">Post not found</span></div>;
    }

    // 投稿を削除する関数
    const handleDelete = async () => {
        // モーダルを表示
        const result = await customConfirm('投稿を削除しますか？');
        if (result) {
            try {
                await deletePost(post.id);
                showAlert('投稿が削除されました');
                setTimeout(() => {
                    navigate('/');
                }, 1000);
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
        <div className="p-6 bg-white shadow-lg rounded-lg relative">
            <button onClick={() => navigate("/new-post")}
                    className="text-white bg-blue-500 hover:bg-blue-600 px-4 py-2 rounded absolute top-0 right-0 m-4">新しい投稿を作成
            </button>
            <h2 className="text-2xl font-bold text-gray-800 mb-2">{post.title}</h2>
            <p className="text-gray-600">ユーザー名: <span className="font-semibold">{post.username}</span></p>
            <p className="text-gray-500">更新日: {post.updated_at}</p>
            <p className="text-gray-500 mt-4">{post.body}</p>

            {userId === post.user_id && (
                <div className="mt-4">
                    <Link to={`/posts/${post.id}/edit`} className="text-blue-500 mr-4">編集</Link>
                    <button onClick={handleDelete} className="text-red-500">削除</button>
                </div>
            )}

            {error && <div className="text-red-500 mt-4">{error}</div>}

            {post.tag === 'FanMeeting' && (
                <div className="mt-4">
                    <button
                        onClick={() => handleIntention('attend')}
                        disabled={loading}
                        className={intention === 'attend' ? 'button-attend' : ''}>
                        Attend
                    </button>
                    <button
                        onClick={() => handleIntention('skip')}
                        disabled={loading}
                        className={intention === 'skip' ? 'button-attend' : ''}>
                        Not Attend
                    </button>
                </div>
            )}

            <div
                onMouseEnter={() => setIsHoveringAttendees(true)}
                onMouseLeave={() => setIsHoveringAttendees(false)}
            >
                <h3>Attendees</h3>
                {isHoveringAttendees && attendees.map(name => (
                    <p key={name}>{name}</p>
                ))}
            </div>
            <div
                onMouseEnter={() => setIsHoveringNonAttendees(true)}
                onMouseLeave={() => setIsHoveringNonAttendees(false)}
            >
                <h3>Non Attendees</h3>
                {isHoveringNonAttendees && nonAttendees.map(name => (
                    <p key={name}>{name}</p>
                ))}
            </div>

            <Link to="/" className="text-blue-500 mt-4 block">ホームに戻る</Link>

            <ConfirmModal message={confirmMessage}
                          modalRef={modalRef}
                          onConfirm={onConfirm}
                          onCancel={onCancel}></ConfirmModal>
        </div>
    );
}

export default PostDetail;
