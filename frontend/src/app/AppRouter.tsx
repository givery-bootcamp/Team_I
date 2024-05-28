import React from 'react';
import { Routes, Route } from 'react-router-dom';
import PostList from '../features/home/PostList';
import NewPost from '../features/newPost/NewPost';
import JoinClub from '../features/joinClub/JoinClub';

const AppRouter: React.FC = () => {
    return (
        <Routes>
            <Route path="/" element={<PostList />} />
            <Route path="/new-post" element={<NewPost />} />
            <Route path="/join-club" element={<JoinClub />} />
        </Routes>
    );
};

export default AppRouter;