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
import { resize } from '@/utils/adapt'
import { getDevelopers } from '@/api/workspace_settings'
import { IUserInfo } from '@/api/account'

const tokenLoader = async () => {
  resize()
  console.log('tokenLoading');
  const token = tokenStore.getToken()
  if (!token) {
    return redirect('/s')
  }
  await WorkspaceListLoader()
  await getUserInfo().then(res => {
    console.log('enter get info');

    if (res.value?.code === 0) {
      console.log('enter 0');
      // sessionStorage.setItem('userInfo', JSON.stringify(res.value.data))
      appStore.setUserInfo(res.value.data as IUserInfo)
    } else if (res.code === 2005) {
      console.log('enter 2005');

      return redirect('/s')
    } else if (res.code === 2006) {
      console.log('enter 2006');
      console.log('updating token...');

      tokenStore.setToken(res.data.token)
    }
  }).catch(err => {
    return redirect('/s')
    console.log(err);
  })
  return null
}

// const UserProfileLoader = async () => {
//   try {
//     const { value } = await getUserInfo()
//     return value.data
//   } catch (err) {
//     console.log(err);
//     return redirect('/s')
//   }
// }

const WorkspaceListLoader = async () => {
  try {
    const { value } = await workspaceList()
    if (value.data) {
      wsStore.setWsList(value.data)
      currentWorkspaceStore.setCurrentWorkspace(value.data[0])
    }
    return value.data ? value.data : []
  } catch (err) {
    return []
  }

}

const DeveloperListLoader = async () => {
  try {
    const { value } = await getDevelopers(1)
    return value.data ? value.data : []
  } catch (err) {
    return []
  }
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
            // loader: UserProfileLoader,
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
            loader: DeveloperListLoader,
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
