import { Link } from 'react-router-dom';

const Navigation = () => {
    return (
        <nav className="bg-gray-100 p-4 shadow-md">
            <ul className="space-y-4">
                <li><Link to="/" className="text-blue-500 hover:text-blue-700">Home</Link></li>
                <li><Link to="/posts/new" className="text-blue-500 hover:text-blue-700">新規投稿</Link></li>
                <li><Link to="/join-club" className="text-blue-500 hover:text-blue-700">ファンクラブ入会</Link></li>
            </ul>
        </nav>
    );
};

export default Navigation;
