import React from 'react';
import {Route, Routes} from 'react-router-dom';
import PostList from '../features/home/PostList';
import NewPost from '../features/newPost/NewPost';
import JoinClub from '../features/joinClub/JoinClub';
import SignIn from '../features/signIn/SignIn';
import SignUp from '../features/signUp/SignUp';
import PostDetail from '../features/postDetail/PostDetail';
import EditPost from '../features/editPost/EditPost';

const AppRouter: React.FC = () => {
    return (
        <Routes>
            <Route path="/" element={<PostList />} />
            <Route path="/posts/new" element={<NewPost />} />
            <Route path="/join-club" element={<JoinClub />} />
            <Route path="/signin" element={<SignIn />} />
            <Route path="/posts/:postId" element={<PostDetail />} />
            <Route path="/posts/:postId/edit" element={<EditPost/>}/>
            <Route path="/signup" element={<SignUp />} />
        </Routes>
    );
};

export default AppRouter;
