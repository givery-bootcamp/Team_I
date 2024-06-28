import { Link } from 'react-router-dom';
import { useAuth } from "../context/useAuth.ts";


const Navigation = () => {
    const { isLoggedIn } = useAuth();
    return (
        <nav className="bg-gray-100 p-4 shadow-md">
            <ul className="space-y-4">
                <li><Link to="/" className="text-blue-500 hover:text-blue-700">Home</Link></li>
                {isLoggedIn ?
                    <li><Link to="/posts/new" className="text-blue-500 hover:text-blue-700">新規投稿</Link></li> : null}
                <li>
                    <a
                        href="https://dena.enterprise.slack.com/archives/C070S5JS9FX"
                        className="text-blue-500 hover:text-blue-700"
                        target="_blank"
                        rel="noopener noreferrer"
                    >
                        ファンクラブ入会
                    </a>
                </li>
            </ul>
        </nav>
    );
};

export default Navigation;
