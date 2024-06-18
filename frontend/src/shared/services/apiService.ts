import {Post} from '../models/Post';
import API_BASE_URL from '../../config';
import {IFormInput} from '../models/Post.ts';


export const fetchPosts = async (): Promise<Post[]> => {
    const response = await fetch(`${API_BASE_URL}/posts`, {
        headers: {
            'Content-Type': 'application/json',
        }
    });
    console.log(response);
    if (!response.ok) {
        throw new Error('Failed to fetch posts');
    }
    return response.json();
};

export const fetchPostById = async (id: number): Promise<Post> => {
    const response = await fetch(`${API_BASE_URL}/posts/${id}`, {
        headers: {
            'Content-Type': 'application/json',
        }
    });
    if (!response.ok) {
        throw new Error('Failed to fetch post');
    }
    return response.json();
}

export const deletePost = async (id: number): Promise<void> => {
    const response = await fetch(`${API_BASE_URL}/posts/${id}`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
        }
    });
    if (!response.ok) {
        throw new Error('Failed to delete post');
    }
};

export const updatePost = async (id: string | undefined, data: IFormInput): Promise<void> => {
    const response = await fetch(`${API_BASE_URL}/posts/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data)
    });
    if (!response.ok) {
        throw new Error('Failed to update post');
    }
};
