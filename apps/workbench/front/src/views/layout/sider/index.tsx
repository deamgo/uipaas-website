import React, { MouseEventHandler } from 'react';
import { Link, useNavigate } from 'react-router-dom'
import { appStore, tokenStore } from '@/store/store'
//style
import './index.less'
//
import SwitchWorkspace from '@/components/SwitchWorkspace';
import SideMenu from '@/components/SideMenu';
import MultiplySelector from '@/components/multiplySelector';
//
import { Avatar } from 'antd';
import { ReactComponent as Application } from '@assets/layout/application.svg'
import { ReactComponent as Wssettings } from '@assets/layout/wssettings.svg'
import { observer } from 'mobx-react-lite';
import { IMultiplySelectorPropsItem, mcontent } from '@/interface/some';
import Cookies from 'js-cookie';
import Popup from "@components/Popup";

const list_c: mcontent[] = [
  {
    id: 'application',
    title: 'Applications',
    path: '/',
    matcher: '',
    index: 1,
    icon: (<Application />)
  },
]

const list_f: mcontent[] = [
  {
    id: 'wss',
    title: 'workspace Settings',
    path: '/workspace',
    matcher: 'workspace',
    index: 1,
    icon: (<Wssettings />)
  },
]

type SiderProps = {
  children?: React.ReactNode
}

const Sider: React.FC<SiderProps> = (props) => {
  const [showMultiSelect, setShowMultiSelect] = React.useState(false)

  const [username, setUsername] = React.useState('')
  const [active, setActive] = React.useState<number>()


  React.useEffect(() => {
    setUsername(appStore.getUserInfo().username)
  }, [])

  React.useEffect(() => {
    console.log('update appStore');

    setUsername(appStore.getUserInfo().username)
  }, [appStore.getUserInfo()])

  const handleShow = () => {
    setShowMultiSelect(showMultiSelect => !showMultiSelect)
  }

  const navigate = useNavigate()

  const list_ms: IMultiplySelectorPropsItem[] = [
    {
      id: 101,
      text: 'Profile',
      path: '/u',
      type: "normal",
      method: () => {
        setShowMultiSelect(false)
        navigate('/u')
      },
      children: (<span style={{ color: '#0871F0' }}>Profile</span>)
    },
    {
      id: 99,
      text: 'Logout',
      path: '/s',
      type: "error",
      method: () => {
        appStore.resetUserInfo()
        tokenStore.resetToken()
        Cookies.remove('token')
        setShowMultiSelect(false)
        navigate('/s')
      },
      children: (<span style={{ color: '#FF7875' }}>Logout</span>)
    }
  ]

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
            {username.charAt(0).toUpperCase()}
          </Avatar>
          <span className='__sider_usr_info_name'>{username}</span>
        </div>
        {showMultiSelect && (
          <div className="__sider_usr_info_ms">
            <MultiplySelector list={list_ms} />
          </div>
        )}
      </div>
      {/*{isWsCreate && (*/}
      {/*    <>*/}
      {/*      <Popup unit={'rem'} width={31} height={17.25} title={'Create Workspace'}></Popup>*/}
      {/*    </>*/}
      {/*)}*/}
    </>
  )
}

export default observer(Sider)