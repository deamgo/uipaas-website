import React from 'react';
//style
import './index.less'
//
import Button from '@/components/Button';
import SearchLine from '@/components/SearchLine';
//
import { ReactComponent as Plus } from '@assets/comps/plus.svg'
import Mask from '@/components/Mask';
import CreateAppPopup from './CreateAppPopup';


// interface IContentAppHeaderProps {
//   onNotify: () => void
// }

const Header: React.FC = () => {

  const [isMask, setMask] = React.useState(false)
  const [isCreatApp, setCreatApp] = React.useState(false)

  React.useEffect(() => {
    document.body.style.overflow = isMask ? 'hidden' : 'auto'
  }, [isMask])

  const handleMask = () => {
    setMask(!isMask)
  }

  const handleCreate = () => {
    handleMask()
    setCreatApp(!isCreatApp)
  }

  return (
    <>
      <div className="__header">
        <div className="__header_toolg">
          <div className="__header_toolg_btn">
            <Button
              context='Create'
              method={handleCreate} >
              <Plus style={{
                width: '10.67rem',
                height: '10.67rem',
                fill: '#FFFFFF'
              }} />
              Create
            </Button>
          </div>
          <div className="__header_toolg_searchline">
            <SearchLine placeholder='Search' />
          </div>
        </div>
      </div>
      {isMask && (<Mask />)}
      {isCreatApp && (
        <>
          <CreateAppPopup onClose={handleCreate} />
        </>
      )}
    </>
  )
}

export default Header