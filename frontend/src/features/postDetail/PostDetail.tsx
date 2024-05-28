import React, { useState, useEffect } from 'react';
import { Post } from '../../shared/models/Post';
import { useParams, Link } from 'react-router-dom';

const PostDetail: React.FC = () => {
    const [posts, setPosts] = useState<Post[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const { postId } = useParams<{ postId: string }>();


    return(
        <div>
            <h1>Post Detail</h1>
            <p>Post detail will be displayed here.</p>
        </div>
    )
}

export default PostDetail;