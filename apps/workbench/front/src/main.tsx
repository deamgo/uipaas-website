import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.less'

//route
import { RouterProvider } from 'react-router-dom'
import { router } from './router/router.tsx'
//
import { Provider } from 'mobx-react'
import appStore from '@store/store'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <Provider appStore={appStore}>
      <RouterProvider router={router} />
    </Provider>
  </React.StrictMode>,
)
