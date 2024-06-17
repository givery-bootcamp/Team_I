import {Post} from '../models/Post';

// モックデータ
let posts: Post[] = [
    {
        id: 1,
        title: 'バナナは持っていっていいですか？',
        username: 'testuser',
        createdAt: '2022/4/25 18:38:00',
        updatedAt: '2022/4/25 18:38:00',
        content: '２本目も持って行っていいですかね？'
    },
    {
        id: 2,
        title: 'test2',
        username: 'hoge',
        createdAt: '2022/4/25 18:38:00',
        updatedAt: '2022/4/25 18:38:00',
        content: 'これは二つ目のテストです。'
    },
];

export const fetchPosts = async (): Promise<Post[]> => {
    return new Promise(resolve => {
        setTimeout(() => {
            resolve(posts);
        }, 1000); // 1秒の遅延をシミュレート
    });
};

export const deletePost = async (postId: number): Promise<void> => {
    posts = posts.filter(post => post.id !== postId);
}

export const updatePost = async (postId: number, data: { title: string, content: string }): Promise<void> => {
    const postIndex = posts.findIndex(post => post.id === postId);
    if (postIndex > -1) {
        posts[postIndex] = {...posts[postIndex], ...data, updatedAt: new Date().toISOString()};
    }
};
