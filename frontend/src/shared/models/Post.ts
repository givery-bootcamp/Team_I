import { Comment } from "./Comment";

export interface Post {
    id: number;
    title: string;
    body: string;
    user_id: number;
    username: string;
    comments: Comment[];
    created_at: string;
    updated_at: string;
}

export interface IFormInput {
    title: string;
    content: string; // 'body' から 'content' に変更
}
