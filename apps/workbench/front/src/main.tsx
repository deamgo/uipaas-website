import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.less'
import '@assets/globle.less'

//route
import { RouterProvider } from 'react-router-dom'
import { router } from './router/router.tsx'
//
import { Provider } from 'mobx-react'
import { appStore, tokenStore } from '@store/store'
import { currentWorkspaceStore, wsStore } from './store/wsStore.ts'
import { applicationStore } from './store/application.ts'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <Provider
      appStore={appStore}
      tokenStore={tokenStore}
      wsStore={wsStore}
      currentWorkspaceStore={currentWorkspaceStore}
      applicationStore={applicationStore}>
      <RouterProvider router={router} />
    </Provider>
  </React.StrictMode>
)
