import React from 'react';
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
import { ReactComponent as Resource } from '@assets/layout/resource.svg'
import { ReactComponent as Environments } from '@assets/layout/environments.svg'
import { ReactComponent as Nottification } from '@assets/layout/nottification.svg'
import { ReactComponent as Down } from '@assets/comps/down.svg'
import { observer } from 'mobx-react-lite';
import { IMultiplySelectorPropsItem, mcontent } from '@/interface/some';
import Cookies from 'js-cookie';
import { currentWorkspaceStore, wsStore } from '@/store/wsStore';

const list_c: mcontent[] = [
  {
    id: 'application',
    title: 'Applications',
    path: '/',
    matcher: '',
    index: 1,
    icon: (<Application />)
  },
  // {
  //   id: 'resource',
  //   title: 'Resource',
  //   icon: (<Resource />)
  // },
  // {
  //   id: 'env',
  //   title: 'Environments',
  //   icon: (<Environments />)
  // },
]

const list_f: mcontent[] = [
  // {
  //   id: 'nottifi',
  //   title: 'Nottification',
  //   icon: (<Nottification />)
  // },
  {
    id: 'wss',
    title: 'Workspace Settings',
    path: '/workspace',
    matcher: 'workspace',
    index: 1,
    icon: (<Wssettings />)
  },
]

type SiderProps = {
  children?: React.ReactNode
}

const Sider: React.FC<SiderProps> = () => {
  const [showMultiSelect, setShowMultiSelect] = React.useState(false)

  const [username, setUsername] = React.useState('')
  const [isWslist, setIsWslist] = React.useState(false)

  const multiplySelectorRef = React.useRef<HTMLDivElement>(null)


  React.useEffect(() => {
    setUsername(appStore.getUserInfo().username)

    // const handleClickOutSide = (e: Event) => {
    //   e.stopPropagation();
    //   console.log(multiplySelectorRef.current)
    //   console.log(e.target)
    //   if (!multiplySelectorRef.current) {
    //     console.log('no ref');
    //     return
    //   }

    //   console.log(multiplySelectorRef.current.contains(e.target as Node));
    //   console.log(multiplySelectorRef.current !== e.target);

    //   if (multiplySelectorRef.current.contains(e.target as Node)) {
    //     console.log('contains');
    //     handleShow()
    //   }
    // }

    // document.addEventListener('click', handleClickOutSide)

    // return () => {
    //   document.removeEventListener('click', handleClickOutSide)
    // }
  }, [])

  React.useEffect(() => {
    if (wsStore.getWsList().length > 0) {
      setIsWslist(true)
    } else {
      setIsWslist(false)
    }
  }, [wsStore.getWsList().length])

  React.useEffect(() => {
    console.log('update sider foot username')

    setUsername(appStore.getUserInfo().username)
  }, [appStore.getUserInfo().username])

  const handleShow = () => {
    setShowMultiSelect(!showMultiSelect)
  }


  const list_ms: IMultiplySelectorPropsItem[] = [
    {
      id: 'profile',
      text: 'Profile',
      path: '/u',
      type: "normal",
      method: () => {
        handleShow()
      },
      children: (<span style={{ color: '#0871F0' }}>Profile</span>)
    },
    {
      id: 'logout',
      text: 'Logout',
      path: '/s',
      type: "error",
      method: () => {
        appStore.resetUserInfo()
        tokenStore.resetToken()
        wsStore.resetWsList()
        currentWorkspaceStore.resetCurrentWorkspace()
        sessionStorage.clear()
        Cookies.remove('token')
        handleShow()
      },
      children: (<span style={{ color: '#FF7875' }}>Logout</span>)
    }
  ]

  return (
    <>
      <div className="__sider">
        <div className="__sider_menuc">
          <div className="__sider_menuc_head">
            <SwitchWorkspace />
          </div>
          <div className="__sider_menuc_list">
            {isWslist && <SideMenu list={list_c} />}
          </div>
        </div>
        <div className="__sider_menuf">
          <div className="__sider_menuf_list">
            {isWslist && <SideMenu title='SYSTEM' list={list_f} />}
            {/* <Divider /> */}
          </div>
          <div className="__sider_menuf_usr" onClick={handleShow}>
            <div className="__sider_menuf_usr_info">
              <div className="__sider_menuf_usr_info_avatar">
                <Avatar style={{ backgroundColor: '#4080FF', verticalAlign: 'middle' }} size={32} gap={3}>
                  {username.charAt(0).toUpperCase()}
                </Avatar>
              </div>
              <span className='__sider_menuf_usr_info_name'>{username}</span>
            </div>
            <div className="__sider_menuf_usr_down">
              <Down />
            </div>
            {showMultiSelect && (
              <div className="__sider_menuf_usr_ms" ref={multiplySelectorRef}>
                <MultiplySelector onClose={handleShow} list={list_ms} />
              </div>
            )}
          </div>
        </div>
      </div>
    </>
  )
}

export default observer(Sider)