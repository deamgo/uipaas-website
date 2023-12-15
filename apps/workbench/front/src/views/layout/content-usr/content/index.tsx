import React from 'react';
import { Outlet } from 'react-router-dom'
//style
import './index.less'
//

const Content: React.FC = () => {
  return (
    <>
      <div className="__content">
        <Outlet />
      </div>
    </>
  )
}

export default Content