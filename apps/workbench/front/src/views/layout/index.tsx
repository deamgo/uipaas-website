import React from 'react';
import { redirect, useNavigate } from 'react-router-dom';
import { resize } from '@utils/adapt'
import { observer } from 'mobx-react'
//style
import './index.less'
//
import Sider from '@views/layout/sider'
import { Outlet } from 'react-router-dom';
import { wsStore } from '@/store/wsStore';
//




const Layout: React.FC = () => {

  const navigate = useNavigate()

  React.useEffect(() => {
    window.addEventListener('resize', resize)
    resize()
  })

  React.useEffect(() => {
    if (!(wsStore.getWsList().length > 0)) {
      navigate('/_blank')
    } else {
      navigate('/')
    }
  }, [wsStore.getWsList()?.length])

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