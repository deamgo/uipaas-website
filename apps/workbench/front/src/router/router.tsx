import { RouteObject, createBrowserRouter } from 'react-router-dom'
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

export const routes: RouteObject[] = [
  {
    path: '/',
    Component: Layout,
    children: [
      {
        path: '/apps',
        Component: ContentApp,
      },
      {
        path: '/u',
        Component: ContentUsr,
        children: [
          {
            path: '/u/profile',
            element: <h1>Profile</h1>
          },
          {
            path: '/u/invite',
            element: <h1>Invite</h1>
          },
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
        Component: SignUp
      },
      {
        path: '/s/ev',
        Component: EmailVerif
      },
      {
        path: '/s/in',
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