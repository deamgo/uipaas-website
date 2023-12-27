import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.less'

//route
import { RouterProvider } from 'react-router-dom'
import { router } from './router/router.tsx'
//
import { Provider } from 'mobx-react'
import { appStore, tokenStore } from '@store/store'
import { currentWorkspaceStore, wsStore } from './store/wsStore.ts'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <Provider appStore={appStore} tokenStore={tokenStore} wsStore={wsStore} currentWorkspaceStore={currentWorkspaceStore}>
      <RouterProvider router={router} />
    </Provider>
  </React.StrictMode>
)
