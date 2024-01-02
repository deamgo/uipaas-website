import React from 'react';
import { useNavigate } from 'react-router-dom';
import { resize } from '@utils/adapt'
import { observer } from 'mobx-react'
//style
import './index.less'
//
import Sider from '@views/layout/sider'
import { Outlet } from 'react-router-dom';
import { wsStore } from '@/store/wsStore';
import { socket } from '@/utils/websocket';
//


const Layout: React.FC = () => {

  const navigate = useNavigate()

  React.useEffect(() => {
    window.addEventListener('resize', resize)
    resize()

    socket.onopen = () => {
      console.log('connected');
    }
    socket.onmessage = (e) => {
      console.log(e.data)
    }
    return () => {
      window.removeEventListener('resize', resize)
      socket.close()
    }
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