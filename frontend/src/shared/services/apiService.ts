import { Post } from '../models/Post';

const API_BASE_URL = 'https://api.example.com';

export const fetchPosts = async (): Promise<Post[]> => {
    const response = await fetch(`${API_BASE_URL}/posts`);
    if (!response.ok) {
        throw new Error('Failed to fetch posts');
    }
    return response.json();
};
