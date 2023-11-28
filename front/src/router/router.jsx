import { createBrowserRouter } from 'react-router-dom'
import App from '../App.jsx'
import PriPolicy from '../layout/Privacy-policy.jsx'
import ErrorPage from '../layout/errorPage.jsx'

export const router = createBrowserRouter([
    {
        path: '/',
        Component: App,
        errorElement: <ErrorPage />
    },
    {
        path: '/privacy',
        Component: PriPolicy,
    },
])