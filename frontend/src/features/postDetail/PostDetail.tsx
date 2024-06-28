import React, {useEffect, useRef, useState} from 'react';
import {Post} from '../../shared/models/Post';
import {Link, useNavigate, useParams} from 'react-router-dom';
import {useAuth} from '../../shared/context/useAuth';
import {deletePost, fetchPostById, postIntention, fetchIntentionState, createComment, deleteComment, updateComment} from '../../shared/services/apiService';
import {useAlert} from '../../shared/components/AlertContext';
import ConfirmModal from '../../shared/components/Modal';
import { useConfirmModal } from '../../shared/hooks/useConfirmModal';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Comment, CommentIFormInput } from '../../shared/models/Comment';

interface CommentCardProps {
    comment: Comment;
    userId: number | undefined;
    setError: (message: string) => void;
    customConfirm: (message: string) => Promise<boolean>;
}

const CommentCard: React.FC<CommentCardProps> = ({comment, userId, setError, customConfirm}) => {
    const [isEditComment, setIsEditComment] = useState(false);
    const { register, handleSubmit, setValue, formState: { errors }, } = useForm<CommentIFormInput>();
    const [newCommentError, setNewCommentError] = React.useState<string | null>(null);
    const {showAlert} = useAlert();
    // ボタン連打防止用のフラグ
    const isProcessing = useRef(false);

    useEffect(() => {
        setValue('body', comment.body);
    }   , [comment.body, setValue]);

    const handleEditComment = () => {
        setIsEditComment(!isEditComment);
    }

    const onSubmit: SubmitHandler<CommentIFormInput> = async (data) => {
        // 処理中なら何もしない
        if (isProcessing.current) {
            return;
        }

        try {
            // 処理開始
            isProcessing.current = true;
            await updateComment(data, comment.id);
            showAlert('投稿が更新されました。');
            // ページをリロード
            window.location.reload();
        } catch (err) {
            if (err instanceof Error) {
                setNewCommentError(err.message);
            } else {
                setNewCommentError('An unexpected error occurred');
            }
        } finally {
            // 処理終了
            isProcessing.current = false;
        }
    };

    // コメントを削除する関数
    const handleDeleteComment = async (comment_id: number) => {
        // モーダルを表示
        const result = await customConfirm('コメントを削除しますか？');
        if (result) {
            try {
                // コメントを削除
                await deleteComment(comment_id);
                showAlert('コメントが削除されました');
                // ページをリロード
                window.location.reload();
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
        <div className="p-4 bg-white shadow-lg rounded-lg">
            {(isEditComment
            ?
            <div className="border-t border-gray-200 mt-4  pt-4">
                <form onSubmit={handleSubmit(onSubmit)}>
                    <div className="mb-4">
                        <label className="block text-gray-600 mb-2">コメント</label>
                        <textarea
                            className={`w-full p-2 border ${errors.body ? 'border-red-500' : 'border-gray-300'} rounded-md`}
                            {...register('body', {
                                required: 'コメントは必須です。',
                            })}
                        />
                        {/* コメントエラーメッセージ */}
                        {errors.body && <div className="text-red-500 mt-2">{errors.body.message}</div>}
                    </div>
                    <button onClick={handleEditComment}
                            className="px-4 py-2 bg-gray-500 text-white rounded-md hover:bg-gray-600 mx-4">キャンセル
                    </button>
                    <button type="submit"
                            className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 mx-4">更新する
                    </button>
                    {newCommentError && <div className='text-red-500'>{newCommentError}</div>}
                </form>
            </div>
            :<div className="border-t border-gray-200">
                <div key={comment.id} className='m-4'>
                    <div className='my-4'>
                        <p>{comment.body}</p>
                    </div>
                    <div>
                        <p className="text-gray-600 text-xs">ユーザー名: <span className="font-semibold">{comment.user_name}</span></p>
                        <p className="text-gray-500 text-xs">更新日: {comment.updated_at}</p>
                    </div>

                </div>
                { comment.user_id === userId &&
                    <div className="t-4">
                        <button onClick={handleEditComment} className="text-blue-500 mr-4">編集</button>
                        <button onClick={() => handleDeleteComment(comment.id)} className="text-red-500">削除</button>
                    </div>
                }
            </div>

            )}
        </div>
    );
}

export type IntentionState = 'attend' | 'skip'


const PostDetail: React.FC = () => {
    const [post, setPost] = useState<Post | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const {postId} = useParams<{ postId: string }>();
    const navigate = useNavigate();
    const [isAddComment, setIsAddComment] = useState(false);
    const [newCommentError, setNewCommentError] = React.useState<string | null>(null);
    const { register, handleSubmit, formState: { errors }, } = useForm<CommentIFormInput>();


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
    // ボタン連打防止用のフラグ
    const isProcessing = useRef(false);

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

    const handleAddComment = () => {
        setIsAddComment(!isAddComment);
    }


    const onSubmit: SubmitHandler<CommentIFormInput> = async (data) => {
        // 処理中なら何もしない
        if (isProcessing.current) {
            return;
        }

        try {
            // 処理開始
            isProcessing.current = true;
            data.post_id = post.id;
            await createComment(data);

            // 成功したらアラート
            showAlert('コメントしました。');

            // 成功したらぺージをリロード
            window.location.reload();
        } catch (err) {
            if (err instanceof Error) {
                setNewCommentError(err.message);
            } else {
                setNewCommentError('An unexpected error occurred');
            }
        } finally {
            // 処理終了
            isProcessing.current = false;
        }
    }

    return (
        <div className="p-6 bg-white shadow-lg rounded-lg relative">
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

                {(post.type === 'official' || post.type === 'yamada') && (
                    <>
                        <div className="mt-4">
                            <button
                                onClick={() => handleIntention('attend')}
                                disabled={loading}
                                className={`flex-1 mr-2 py-2 px-4 rounded bg-blue-500 hover:bg-blue-600 text-white font-semibold
                            ${intention === 'attend' ? 'opacity-100' : 'opacity-50'}`}
                            >
                                参加
                            </button>
                            <button
                                onClick={() => handleIntention('skip')}
                                disabled={loading}
                                className={`flex-1 ml-2 py-2 px-4 rounded bg-red-500 hover:bg-red-600 text-white font-semibold
                            ${intention === 'skip' ? 'opacity-100' : 'opacity-50'}`}
                            >
                                不参加
                            </button>
                        </div>

                        <div className="flex gap-8">
                            <div
                                onMouseEnter={() => setIsHoveringAttendees(true)}
                                onMouseLeave={() => setIsHoveringAttendees(false)}
                            >
                                <h3 className="inline-block py-1 px-2 text-sm rounded-full bg-blue-500 text-white my-2">Attendees</h3>
                                {isHoveringAttendees && attendees.map(name => (
                                    <p key={name}>{name}</p>
                                ))}
                            </div>
                            <div
                                onMouseEnter={() => setIsHoveringNonAttendees(true)}
                                onMouseLeave={() => setIsHoveringNonAttendees(false)}
                            >
                                <h3 className="inline-block py-1 px-2 text-sm rounded-full bg-blue-500 text-white my-2">Non
                                    Attendees</h3>
                                {isHoveringNonAttendees && nonAttendees.map(name => (
                                    <p key={name}>{name}</p>
                                ))}
                            </div>
                        </div>
                    </>
                )}


                <Link to="/" className="text-blue-500 mt-4 block">ホームに戻る</Link>

                <ConfirmModal message={confirmMessage}
                              modalRef={modalRef}
                              onConfirm={onConfirm}
                              onCancel={onCancel}></ConfirmModal>

            </div>
            {/* コメントのidがuserのidと一致するときに編集ボタンを表示 */}
            {post.comments?.length > 0 && post.comments.map(comment => {
                return (
                    <CommentCard key={comment.id} comment={comment} userId={userId} setError={setError}
                                 customConfirm={customConfirm}/>
                );
            })
            }


            {isAddComment ?
                <div className="p-4 mt-2 bg-white shadow-lg rounded-lg">
                <form onSubmit={handleSubmit(onSubmit)}>
                    <div className="mb-4">
                        <label className="block text-gray-600 mb-2">コメント</label>
                        <textarea
                            className={`w-full p-2 border ${errors.body ? 'border-red-500' : 'border-gray-300'} rounded-md`}
                            {...register('body', {
                                required: 'コメントは必須です。',
                            })}
                        />
                        {/* コメントエラーメッセージ */}
                        {errors.body && <div className="text-red-500 mt-2">{errors.body.message}</div>}
                    </div>
                    <button onClick={handleAddComment}
                            className="px-4 py-2 bg-gray-500 text-white rounded-md hover:bg-gray-600 mx-4">キャンセル
                    </button>
                    <button type="submit"
                            className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 mx-4">投稿する
                    </button>
                    {newCommentError && <div className='text-red-500'>{newCommentError}</div>}
                </form>
            </div>
                :
                <button onClick={handleAddComment} className="text-white bg-blue-500 hover:bg-blue-600 px-4 py-2 rounded m-4">
                    コメントを追加する
                </button>
            }
            <ConfirmModal message={confirmMessage} modalRef={modalRef} onConfirm={onConfirm} onCancel={onCancel}></ConfirmModal>
        </div>
    );
}

export default PostDetail;
