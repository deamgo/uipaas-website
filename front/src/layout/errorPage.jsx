import './err.less'

import { useEffect } from 'react'

import { useRouteError, Navigate } from 'react-router-dom'

export default function ErrorPage() {

  useEffect(() => {
    window.addEventListener('resize', resize)
    resize()
  }, [])

  const resize = () => {
    let IwD = window.innerWidth
    document.documentElement.style.fontSize = IwD / 1920 + 'px'
  }

  const err = useRouteError()

  return (
    // <div id="error-page">
    //   {/* <h1>Oops!</h1>
    //   <p>Sorry, an unexpected error has occurred.</p>
    //   <p>
    //     <i>{err.statusText || err.message}</i>
    //   </p> */}
    // </div>
    <Navigate to="/" />
  );
}
