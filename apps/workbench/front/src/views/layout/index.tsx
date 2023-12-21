import React from 'react';
import { redirect } from 'react-router-dom';
import { resize } from '@utils/adapt'
import { observer } from 'mobx-react'
//style
import './index.less'
//
import Sider from '@views/layout/sider'
import { Outlet } from 'react-router-dom';
//




const Layout: React.FC = () => {
  React.useEffect(() => {
    window.addEventListener('resize', resize)
    resize()
  })

  return (
    <>
      <div className="layout">
        <div className="sider">
          <Sider />
        </div>
        <div className="main">
          <Outlet />
        </div>
      </div>
    </>
  )
}

export default observer(Layout)