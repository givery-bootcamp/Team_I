import React, {useEffect, useRef, useState} from 'react';
import {Post} from '../../shared/models/Post';
import {Link, useNavigate, useParams} from 'react-router-dom';
import {useAuth} from '../../shared/context/useAuth';
import {createComment, deletePost, fetchPostById} from '../../shared/services/apiService';
import {useAlert} from '../../shared/components/AlertContext';
import ConfirmModal from '../../shared/components/Modal';
import { useConfirmModal } from '../../shared/hooks/useConfirmModal';
import { SubmitHandler, useForm } from 'react-hook-form';
import { CommentIFormInput } from '../../shared/models/Comment';

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
    }, [postId]);

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
        // 投稿詳細を表示
        <div className="p-6 bg-white shadow-lg rounded-lg relative">
            <div>
            <button onClick={() => navigate("/new-post")}
                    className="text-white bg-blue-500 hover:bg-blue-600 px-4 py-2 rounded absolute top-0 right-0 m-4">新しい投稿を作成
            </button>
            <h2 className="text-2xl font-bold text-gray-800 mb-2">{post.title}</h2>
            <p className="text-gray-600">ユーザー名: <span className="font-semibold">{post.username}</span></p>
            <p className="text-gray-500">更新日: {post.updated_at}</p>
            <p className="text-gray-500 mt-4">{post.body}</p>
            <Link to="/" className="text-blue-500 mt-4 block">ホームに戻る</Link>

            {userId === post.user_id && (
                <div className="mt-4">
                    <Link to={`/posts/${post.id}/edit`} className="text-blue-500 mr-4">編集</Link>
                    <button onClick={handleDelete} className="text-red-500">削除</button>
                </div>
            )}

            {error && <div className="text-red-500 mt-4">{error}</div>}
            </div>
            {post.comments?.length > 0 && post.comments.map(comment => {
                return (
                    <div key={comment.id} className="border-t border-gray-200 mt-4 pt-4">
                        <p className="text-gray-600">ユーザー名: <span className="font-semibold">{comment.user_name}</span></p>
                        <p className="text-gray-500">コメント: {comment.body}</p>
                        <p className="text-gray-500">更新日: {comment.updated_at}</p>
                    </div>
                );
                })}

            {isAddComment ?
                <form onSubmit={handleSubmit(onSubmit)}>
                    <div className="mb-4">
                        <label className="block text-gray-600 mb-2">内容</label>
                        <textarea
                            className={`w-full p-2 border ${errors.body ? 'border-red-500' : 'border-gray-300'} rounded-md`}
                            {...register('body', {
                                required: '内容は必須です。',
                            })}
                        />
                        {/* 内容エラーメッセージ */}
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
