import { createBrowserRouter, Navigate } from 'react-router-dom'
import App from '../App.jsx'
import Signin from '../layout/Signin'

export const router = createBrowserRouter([
    {
        path: '/',
        Component: App,
        errorElement: <Navigate to="/" />,
    },
    {
        path: '/logo',
        Component: Signin,
    },
])