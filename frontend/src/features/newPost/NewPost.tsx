import React from 'react';

const NewPost: React.FC = () => {
    return (
        <div className="p-6 bg-white shadow-lg rounded-lg">
            <h1 className="text-2xl font-bold text-gray-800 mb-4">新規投稿</h1>
            <form>
                <div className="mb-4">
                    <label className="block text-gray-600 mb-2">タイトル</label>
                    <input type="text" name="title" className="w-full p-2 border border-gray-300 rounded-md" />
                </div>
                <div className="mb-4">
                    <label className="block text-gray-600 mb-2">内容</label>
                    <textarea name="content" className="w-full p-2 border border-gray-300 rounded-md" />
                </div>
                <button type="submit" className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600">投稿する</button>
            </form>
        </div>
    );
};

export default NewPost;
