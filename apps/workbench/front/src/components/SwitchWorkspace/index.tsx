import React from "react"
import { appStore } from "@/store/store"
//
import './index.less'
//
import { ReactComponent as Switch } from '@assets/layout/switch.svg'
//
import { Avatar } from 'antd'
import { observer } from "mobx-react-lite"


const SwitchWorkspace: React.FC = () => {

  const [username, setUsername] = React.useState('')

  React.useEffect(() => {
    setUsername(appStore.getUserInfo().username)
  }, [])

  React.useEffect(() => {
    setUsername(appStore.getUserInfo().username)
  }, [appStore.getUserInfo()])

  return (
    <>
      <div className="__sws">
        <Avatar style={{ backgroundColor: 'pink', verticalAlign: 'middle' }} size={28} gap={2}>
          {username.charAt(0).toUpperCase()}
        </Avatar>
        <span className="__sws_title">{username}'s Workspace</span>
        <div className="__sws_switch">
          <Switch />
        </div>
      </div>
    </>
  )
}

export default observer(SwitchWorkspace)