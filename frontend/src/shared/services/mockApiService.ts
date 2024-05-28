import { Post } from '../models/Post';

export const fetchPosts = async (): Promise<Post[]> => {
    // モックデータ
    const mockPosts: Post[] = [
        { id: 1, title: 'バナナは持っていっていいですか？', username: 'hoge', date: '2022/4/25 18:38:00' },
        { id: 2, title: 'test2', username: 'hoge', date: '2022/4/25 18:38:00' },
    ];

    return new Promise(resolve => {
        setTimeout(() => {
            resolve(mockPosts);
        }, 1000); // 1秒の遅延をシミュレート
    });
};
