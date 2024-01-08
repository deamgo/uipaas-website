import React from 'react'
//
import './index.less'
import Input from '@/components/Input'
import Button from '@/components/Button'
import Popup from '@/components/Popup'
import { applicationStore } from '@/store/application'
//

interface CreateAppPopupProps {
  onClose: () => void
}

const CreateAppPopup: React.FC<CreateAppPopupProps> = (props) => {
  const [appName, setAppName] = React.useState('')

  const handleCreateApp = async () => {
    applicationStore.createApp(appName)
  }

  return (
    <>
      <Popup unit='rem' width={480} height={238} title='Create app' onClose={props.onClose}>
        <h1 className="_create_app_title">
          Create app
        </h1>
        <div className="_create_app_input">
          <Input
            id='createappinput'
            title='App Name'
            placeholder='Enter your app name'
            outputChange={setAppName}
            isNeed={true} />
        </div>
        <div className="_create_app_confirm">
          <Button type='primary' method={handleCreateApp}>Create app</Button>
        </div>
      </Popup>
    </>
  )
}

export default CreateAppPopup