import React from 'react';
import { Outlet } from 'react-router-dom'
//style
import './index.less'
//
import { ReactComponent as EmptyApp } from '@assets/layout/emptyApp.svg'


const Content: React.FC = () => {
  const [isEmpty, setIsEmpty] = React.useState(true)
  return (
    <>
      <div className="__content">
        {isEmpty ? (
          <>
            <div className="__content_empty_svg">
              <EmptyApp />
            </div>
            <span className="__content_empty_span">
              Nothing here.
            </span>
          </>
        ) : (
          <>
            <Outlet />
          </>
        )}
      </div>
    </>
  )
}

export default Content