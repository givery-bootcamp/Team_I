import { BrowserRouter as Router } from 'react-router-dom';
import Navigation from '../shared/components/Navigation';
import AppRouter from './AppRouter.tsx';

function App() {
    return (
        <Router>
            <div className="flex">
                <aside className="w-1/4">
                    <Navigation />
                </aside>
                <main className="w-3/4 p-4">
                    <AppRouter />
                </main>
            </div>
        </Router>
    );
}

export default App;
