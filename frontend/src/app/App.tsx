import {BrowserRouter as Router} from 'react-router-dom';
import Navigation from '../shared/components/Navigation';
import Header from '../shared/components/Header';
import AppRouter from './AppRouter';
import {AuthProvider} from '../shared/components/AuthContext';

function App() {
    return (
        <AuthProvider>
            <Router>
                <div className="flex flex-col min-h-screen">
                    <Header/>
                    <div className="flex flex-1">
                        <aside className="w-1/4">
                            <Navigation/>
                        </aside>
                        <main className="w-3/4 p-4">
                            <AppRouter/>
                        </main>
                    </div>
                </div>
            </Router>
        </AuthProvider>
    );
}

export default App;
