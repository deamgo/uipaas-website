import React from 'react';
import { Outlet } from 'react-router-dom'
import { resize } from '@utils/adapt'
//style
import './index.less'
//svg
import signLogoSvg from '@assets/sign/sign-logo.svg'

const Sign: React.FC = () => {

  React.useEffect(() => {
    window.addEventListener('resize', resize)
    resize()
  }, [])

  return (
    <>
      <div className="__sign">
        <div className="__sign_logo">
          <img src={signLogoSvg} alt="UIPaaS" />
        </div>
        <div className="__sign_container">
          <Outlet />
        </div>
      </div>
    </>
  )
}

export default Sign