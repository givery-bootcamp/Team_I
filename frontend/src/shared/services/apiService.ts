import {Post} from '../models/Post';
import API_BASE_URL from '../../config';
import {IFormInput} from '../models/Post.ts';


type APIRequestOptions<T = undefined> = T extends undefined ? undefined : {data: T};


const apiRequest = async <T = undefined>(url: string, method = 'GET', options?:  APIRequestOptions<T>): Promise<any> => {
    const requestOptions: RequestInit = {
        method,
        headers: {
            'Content-Type': 'application/json',
        },
        credentials: 'include',
    };

    const requestData = options && 'data' in options ? options.data : undefined;

    if (requestData) {
        requestOptions.body = JSON.stringify(requestData);
    }

    console.log('url', url);
    console.log('requestOptions', requestOptions);
    const response = await fetch(url, requestOptions);
    if (!response.ok) {
        console.error(`Failed to ${method}`, response);

        // Extract error message
        let errorText = 'Unknown error occurred';
        try {
            const errorData = await response.json();
            errorText = errorData.error || JSON.stringify(errorData);
        } catch(e) {
            // Failed to parse response
            errorText = await response.text();
        }
        throw new Error(errorText);
    }

    if (response.ok && response.headers.get('content-type')?.includes('application/json')) {
        return response.json();
    } else {
        throw new Error(`HTTP error! status: ${response.status}`);
    }
};

export const fetchPosts = (): Promise<Post[]> => apiRequest(`${API_BASE_URL}/posts`);

export const fetchPostById = (id: number): Promise<Post> => apiRequest(`${API_BASE_URL}/posts/${id}`);

export const deletePost = (id: number): Promise<void> => apiRequest(`${API_BASE_URL}/posts/${id}`, 'DELETE');

export const updatePost = (data: IFormInput, id: string | undefined): Promise<any> => apiRequest(`${API_BASE_URL}/posts/${id}`, 'PUT', {data});

export const createPost = (data: IFormInput): Promise<any> => apiRequest(`${API_BASE_URL}/posts`, 'POST', {data});

export const signIn = (data: { name: string; password: string }): Promise<any> =>
    apiRequest(`${API_BASE_URL}/signin`, 'POST', { data: { name: data.name, password: data.password } });

export const signOut = (): Promise<any> => apiRequest(`${API_BASE_URL}/signout`, 'POST');
