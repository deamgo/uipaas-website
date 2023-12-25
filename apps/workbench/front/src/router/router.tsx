import { RouteObject, createBrowserRouter, redirect } from 'react-router-dom'
//layout
import Layout from '@/views/layout'
import Sign from '@views/sign'
import SignUp from '@/views/sign/SignUp'
import EmailVerif from '@/views/sign/EmailVerif'
import SignIn from '@/views/sign/SignIn'
import PwdReset from '@/views/sign/PwdReset'
import Privacy from '@/views/privacy'
import ContentApp from '@/views/layout/content-app'
import ContentUsr from '@/views/layout/content-usr'
import UserProfile from '@/views/layout/content-usr/content/profile'
import ContentWorkSpace from '@/views/layout/content-ws-sys'
import WSDevelopers from '@/views/layout/content-ws-sys/content/developers'
import WSSettings from '@/views/layout/content-ws-sys/content/settings'
import { tokenStore, appStore } from '@/store/store'
import { getUserInfo } from '@/api/developer_profile'
import { workspaceList } from '@/api/workspace'
import { currentWorkspaceStore, wsStore } from '@/store/wsStore'
import _Blank from '@/views/layout/_blank'

const tokenLoader = async () => {
  console.log('tokenLoading');

  const token = tokenStore.getToken()
  if (!token) {
    return redirect('/s')
  } else {
    await getUserInfo().then(res => {
      if (res.value?.code === 0) {
        sessionStorage.setItem('userId', res.value.data.id)
        sessionStorage.setItem('userName', res.value.data.username)
        sessionStorage.setItem('userEmail', res.value.data.email)
        sessionStorage.setItem('userInfo', JSON.stringify(res.value.data))
        appStore.setUserInfo(res.value.data)
      } else if (res.code === 2005) {
        return redirect('/s')
      }
    }).catch(err => {
      console.log(err);
      return redirect('/s')
    })

    WorkspaceListLoader()

  }
  return null
}

const UserProfileLoader = async () => {
  await getUserInfo().then(res => {
    if (res.value.code === 0) {
      return res.value.data
    } else {
      return {}
    }
  }).catch(err => {
    return {}
  })
  return {}
}

const WorkspaceListLoader = async () => {
  await workspaceList().then(res => {
    if (res.value.code === 0) {
      wsStore.setWsList(res.value.data)
      if (!currentWorkspaceStore.currentWorkspace.name) {
        console.log(currentWorkspaceStore.currentWorkspace);

        currentWorkspaceStore.setCurrentWorkspace(res.value.data[0])
      }
      return res.value.data
    } else {
    }
  }).catch(err => {
    console.log(err);
  })
  return []
}

export const routes: RouteObject[] = [
  {
    path: '/',
    Component: Layout,
    loader: tokenLoader,
    children: [
      {
        index: true,
        Component: ContentApp,
      },
      {
        path: '/u',
        Component: ContentUsr,
        children: [
          {
            index: true,
            loader: UserProfileLoader,
            Component: UserProfile,
          }
        ]
      },
      {
        path: '/_blank',
        Component: _Blank,
      },
      {
        path: '/workspace',
        Component: ContentWorkSpace,
        children: [
          {
            index: true,
            Component: WSDevelopers,
          },
          {
            path: '/workspace/settings',
            Component: WSSettings,
          }
        ]
      },
    ]
  },
  {
    path: '/s',
    Component: Sign,
    children: [
      {
        path: '/s/up',
        Component: SignUp,
      },
      {
        path: '/s/ev',
        Component: EmailVerif
      },
      {
        index: true,
        Component: SignIn
      },
      {
        path: '/s/ryp',
        Component: PwdReset
      }
    ]
  },
  {
    path: '/privacy',
    Component: Privacy
  }
]


export const router = createBrowserRouter(routes)
