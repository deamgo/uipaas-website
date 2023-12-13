import { createBrowserRouter } from 'react-router-dom'
//layout
import Layout from '@/views/layout'
import Sign from '@views/sign'
import SignUp from '@/views/sign/SignUp'
import EmailVerif from '@/views/sign/EmailVerif'
import SignIn from '@/views/sign/SignIn'
import PwdReset from '@/views/sign/PwdReset'
import Privacy from '@/views/privacy'


export const router = createBrowserRouter([
  {
    path: '/',
    Component: Layout
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
])