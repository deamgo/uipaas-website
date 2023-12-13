import React from "react"
//
import './index.less'
//
import { ReactComponent as Switch } from '@assets/layout/switch.svg'
//
import { Avatar } from 'antd'


const SwitchWorkspace: React.FC = () => {
  return (
    <>
      <div className="__sws">
        <Avatar style={{ backgroundColor: 'pink', verticalAlign: 'middle' }} size={28} gap={2}>
          {'ILEE'.charAt(0).toUpperCase()}
        </Avatar>
        <span className="__sws_title">{'Ilee'}'s Workspace</span>
        <div className="__sws_switch">
          <Switch />
        </div>
      </div>
    </>
  )
}

export default SwitchWorkspace