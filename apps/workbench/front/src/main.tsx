import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.less'

//route
import { RouterProvider } from 'react-router-dom'
import { router } from './router/router.tsx'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
