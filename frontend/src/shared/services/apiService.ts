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
        let errorText = 'Unknown error occurred';
        if (response.headers.get('content-type')?.includes('application/json')) {
            try {
                const errorData = await response.json();
                errorText = errorData.error || JSON.stringify(errorData);
            } catch(e) {
                errorText = 'Error occurred while parsing the error message';
            }
        } else {
            errorText = `HTTP error! status: ${response.status}`;
        }
        throw new Error(errorText);
    }

    if (response.status === 204) {
        return;
    }
    
    if (response.headers.get('content-type')?.includes('application/json')) {
        return response.json();
    } else {
        throw new Error(`Could not get a valid JSON response. status: ${response.status}`);
    }
};

type FetchPostsOptions = {
    page?: number;
    limit?: number;
}

export const fetchPosts = (options: FetchPostsOptions = {}): Promise<Post[]> => {
    let url = `${API_BASE_URL}/posts`;

    const params = new URLSearchParams();
    if (options.page != null) {
        params.append('page', String(options.page));
    }
    if (options.limit != null) {
        params.append('limit', String(options.limit));
    }
    if (params.toString()) {
        url += `?${params.toString()}`;
    }

    return apiRequest(url);
};

export const fetchPostById = (id: number): Promise<Post> => apiRequest(`${API_BASE_URL}/posts/${id}`);

export const deletePost = (id: number): Promise<void> => apiRequest(`${API_BASE_URL}/posts/${id}`, 'DELETE');

export const updatePost = (data: IFormInput, id: string | undefined): Promise<any> => apiRequest(`${API_BASE_URL}/posts/${id}`, 'PUT', {data});

export const createPost = (data: IFormInput): Promise<any> => apiRequest(`${API_BASE_URL}/posts`, 'POST', {data});

export const signIn = (data: { name: string; password: string }): Promise<any> =>
    apiRequest(`${API_BASE_URL}/signin`, 'POST', { data: { name: data.name, password: data.password } });

export const signOut = (): Promise<any> => apiRequest(`${API_BASE_URL}/signout`, 'POST');

export const signUp = (data: { name: string; password: string }): Promise<any> =>
    apiRequest(`${API_BASE_URL}/signup`, 'POST', { data: { name: data.name, password: data.password } });

export const getUser = (): Promise<any> => apiRequest(`${API_BASE_URL}/user`);
