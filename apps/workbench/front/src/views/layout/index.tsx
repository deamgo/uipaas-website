import React from 'react';
import { resize } from '@utils/adapt'
//style
import './index.less'
//
import Sider from '@views/layout/sider'
import Header from '@views/layout/header'
import Content from '@/views/layout/content'
import Guide from './guide';
//


const Layout: React.FC = () => {

  const [fstGuide, setFstGuide] = React.useState(true)

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
        {fstGuide && (
          <>
            <div className="guide">
              <Guide handleClosed={() => { setFstGuide(false) }} />
            </div>
          </>
        )}
        <div className="header" style={fstGuide ? {
          top: '213rem'
        } : {}}>
          <Header />
        </div>
        <div className="content" style={fstGuide ? {
          top: '285rem',
          height: 984 - 285 + 72 + 'rem'
        } : {}}>
          <Content />
        </div>
      </div>
    </>
  )
}

export default Layout