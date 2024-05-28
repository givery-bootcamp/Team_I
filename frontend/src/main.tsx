import React from 'react';
import { createRoot } from 'react-dom/client';
import './main.scss'; // Tailwind CSSをインポート
import App from './app/App';

const container = document.getElementById('root');

if (container !== null) {
    const root = createRoot(container);

    root.render(
        <React.StrictMode>
            <App />
        </React.StrictMode>
    );
} else {
    console.error('Root container missing in index.html');
}
