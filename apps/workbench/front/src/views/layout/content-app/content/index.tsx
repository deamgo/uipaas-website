import React from 'react';
import { Outlet } from 'react-router-dom'
//style
import './index.less'
//
import { ReactComponent as NoneContent } from '@assets/default/none-content.svg'
import Mask from '@/components/Mask';
import Popup from '@/components/Popup';
import ApplicationBox from '@/components/ApplicationBox';


const Content: React.FC = () => {
  const [isEmpty, setIsEmpty] = React.useState(false)

  return (
    <>
      <div className="__appcontent">
        {isEmpty ? (
          <>
            <div className="__appcontent_empty">
              <div className="__appcontent_empty_svg">
                <NoneContent />
              </div>
              <span className="__appcontent_empty_span">
                No content, please create
              </span>
            </div>
          </>
        ) : (
          <>
            <div className="__appcontent_apps">
              <ApplicationBox id='100NOI9' name='App1' label='publish' desc='test app model' type='info' />
            </div>
          </>
        )}
      </div>
    </>
  )
}

export default Content