import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import App from './App.tsx';
import './index.css';
import ConfirmationLanding from './pages/Social/ConfirmationLanding.tsx';

const router = createBrowserRouter([
	{
		path: '/',
		element: <App />
	},
	{
		path: '/confirm/:token',
		element: <ConfirmationLanding />
	}
]);

createRoot(document.getElementById('root')!).render(
	<StrictMode>
		<RouterProvider router={router} />
	</StrictMode>
);
