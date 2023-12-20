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

const tokenLoader = async () => {
  const token = sessionStorage.getItem('token')
  if (!token) {
    return redirect('/s')
  }
  return null
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
            Component: UserProfile,
          }
        ]
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
