import React from 'react';
//style
import './index.less'
//
import SwitchWorkspace from '@/components/SwitchWorkspace';
import SideMenu from '@/components/SideMenu';
//
import { Avatar } from 'antd';
import { ReactComponent as Application } from '@assets/layout/application.svg'
import { SettingTwo } from '@icon-park/react';

const list_c = [
  {
    id: 1,
    title: 'Applications',
    icon: (<Application />)
  },
]

const list_f = [
  {
    id: 1,
    title: 'workspace Settings',
    icon: (<SettingTwo theme="outline" size="18" fill="#333" />)
  },
]

type SiderProps = {
  children?: React.ReactNode
}

const Sider: React.FC<SiderProps> = (props) => {
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
        <div className="__sider_usr_info">
          <Avatar style={{ backgroundColor: '#4080FF', verticalAlign: 'middle' }} size={32} gap={3}>
            {'ILEE'.charAt(0).toUpperCase()}
          </Avatar>
          <span className='__sider_usr_info_name'>Ilee</span>
        </div>
      </div>
    </>
  )
}

export default Sider