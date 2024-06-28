import React from 'react';
import {Route, Routes} from 'react-router-dom';
import PostList from '../features/posView/PostList';
import OfficialMeeting from '../features/posView/OfficialMeeting';
import YamadaMeeting from '../features/posView/YamadaMeeting';
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
            <Route path="/meeting/official" element={<OfficialMeeting />}/>
            <Route path="/meeting/yamada" element={<YamadaMeeting />}/>
        </Routes>
    );
};

export default AppRouter;
