import React from 'react';
//style
import './index.less'
//
import { ReactComponent as NoneContent } from '@assets/default/none-content.svg'


const Content: React.FC = () => {
  const [isEmpty, setIsEmpty] = React.useState(false)

  React.useEffect(() => {
    let applist = []
    if (applist.length > 0) {
      setIsEmpty(false)
    }else {
      setIsEmpty(true)
    }
  }, [])

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
            </div>
          </>
        )}
      </div>
    </>
  )
}

export default Content