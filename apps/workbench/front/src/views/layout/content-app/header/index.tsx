import React from 'react';
//style
import './index.less'
//
import Button from '@/components/Button';
import SearchLine from '@/components/SearchLine';
//
import { ReactComponent as Plus } from '@assets/comps/plus.svg'
import Mask from '@/components/Mask';
import Popup from '@/components/Popup';
import Input from '@/components/Input';


const Header: React.FC = () => {

  const [isMask, setMask] = React.useState(false)
  const [isCreatApp, setCreatApp] = React.useState(false)
  const [appName, setAppName] = React.useState('')

  const handleMask = () => {
    setMask(!isMask)
  }

  const handleCreate = () => {
    handleMask()
    setCreatApp(!isCreatApp)
    console.log('Create');
    console.log(appName);


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
          <Popup unit='rem' width={480} height={238} title='Create app' onClose={handleCreate}>
            <div className="_create_app_input">
              <Input
                id='createappinput'
                title='App Name'
                placeholder='Enter your app name'
                outputChange={setAppName}
                isNeed={true} />
            </div>
            <div className="_create_app_confirm">
              <Button type='primary'>Confirm</Button>
            </div>
          </Popup>
        </>
      )}
    </>
  )
}

export default Header