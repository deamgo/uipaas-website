import React, { MouseEventHandler } from 'react';
import { redirect, useNavigate } from 'react-router-dom'
//style
import './index.less'
//
import SwitchWorkspace from '@/components/SwitchWorkspace';
import SideMenu from '@/components/SideMenu';
import MultiplySelector from '@/components/multiplySelector';
//
import { Avatar } from 'antd';
import { ReactComponent as Application } from '@assets/layout/application.svg'
import { SettingTwo } from '@icon-park/react';

const list_c = [
  {
    id: 1,
    title: 'Applications',
    path: '/',
    icon: (<Application />)
  },
]

const list_f = [
  {
    id: 2,
    title: 'workspace Settings',
    path: 'workspace/',
    icon: (<SettingTwo theme="outline" size="18" fill="#333" />)
  },
]

interface IMultiplySelectorPropsItem {
  id: number
  text: string
  path: string
  type: 'normal' | 'error'
}


const list_ms: IMultiplySelectorPropsItem[] = [
  {
    id: 101,
    text: 'Profile',
    path: '/u',
    type: "normal"
  },
  {
    id: 99,
    text: 'Logout',
    path: '/s',
    type: "error"
  }
]

type SiderProps = {
  children?: React.ReactNode
}

const Sider: React.FC<SiderProps> = (props) => {
  const [showMultiSelect, setShowMultiSelect] = React.useState(false)

  const handleShow = () => {
    setShowMultiSelect(showMultiSelect => !showMultiSelect)
  }

  return (
    <>
      <div className="__sider">
        <div className="__sider_head">
          <SwitchWorkspace />
        </div>
        <div className="__sider_menu_c">
          <SideMenu list={list_c} />
        </div>
        <div className="__sider_menu_f">
          <SideMenu list={list_f} />
        </div>
        <div className="__sider_usr_info" onClick={handleShow}>
          <Avatar style={{ backgroundColor: '#4080FF', verticalAlign: 'middle' }} size={32} gap={3}>
            {'ILEE'.charAt(0).toUpperCase()}
          </Avatar>
          <span className='__sider_usr_info_name'>Ilee</span>
        </div>
        {showMultiSelect && (
          <div className="__sider_usr_info_ms">
            <MultiplySelector list={list_ms} />
          </div>
        )}
      </div>
    </>
  )
}

export default Sider