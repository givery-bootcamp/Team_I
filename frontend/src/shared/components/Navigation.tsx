import { Link } from 'react-router-dom';
import { useAuth } from "../context/useAuth.ts";


const Navigation = () => {
    const { isLoggedIn } = useAuth();
    return (
        <nav className="bg-gray-100 p-4 shadow-md">
            <ul className="space-y-4">
                <li><Link to="/" className="text-blue-500 hover:text-blue-700">Home</Link></li>
                <li><Link to="/meeting/official"
                          className="text-yellow-600 hover:text-yellow-900 flex">
                    (公式)ファンミ
                    <img src="/nagatani.png" className="h-6 mr-2" alt="Logo"/>
                </Link></li>
                <li><Link to="/meeting/yamada"
                          className="text-gray-500 hover:text-gray-700 flex">
                    (山田)ファンミ
                    <img src="/yamada.png" className="h-6 mr-2" alt="Logo"/>
                </Link></li>
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
