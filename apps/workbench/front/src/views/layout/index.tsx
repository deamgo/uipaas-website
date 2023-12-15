import React from 'react';
import { resize } from '@utils/adapt'
//style
import './index.less'
//
import Sider from '@views/layout/sider'
import { Outlet } from 'react-router-dom';
//


const Layout: React.FC = () => {

<<<<<<<< < Temporary merge branch 1
  const [fstGuide, setFstGuide] = React.useState(true)

=========
>>>>>>>>> Temporary merge branch 2
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

export default Layout