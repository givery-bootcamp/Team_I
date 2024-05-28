import React from 'react';
import { Routes, Route } from 'react-router-dom';
import PostList from '../components/PostList';
import NewPost from '../components/NewPost';
import JoinClub from '../components/JoinClub';

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
