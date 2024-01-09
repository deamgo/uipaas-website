import React from 'react';
import { Outlet } from 'react-router-dom'
import { resize } from '@utils/adapt'
//style
import './index.less'
//svg
import { ReactComponent as SignLogoSvg } from '@assets/sign/sign-logo.svg'

const Sign: React.FC = () => {

  React.useEffect(() => {
    window.addEventListener('resize', resize)
    resize()
  }, [])

  return (
    <>
      <div className="__sign">
        <div className="__sign_container">
          <div className="__sign_container_logo">
            <SignLogoSvg />
          </div>
          <Outlet />
        </div>
      </div>
    </>
  )
}

export default Sign