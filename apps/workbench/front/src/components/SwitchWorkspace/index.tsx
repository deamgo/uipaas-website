import React, { useEffect, useState } from "react"
//
import './index.less'
//
import { ReactComponent as Switch } from '@assets/layout/switch.svg'
//
import { Avatar } from 'antd'
import { observer } from "mobx-react-lite"
import WorkspaceCreateBox from "@components/WorkspaceCreateBox";
import { workspaceList } from "@api/workspace.ts";
import $message from "@components/Message";
import { appStore } from "@/store/store"
import { currentWorkspaceStore, wsStore } from "@/store/wsStore"
import { IWorkspaceItemProps } from "@/interface/some"


const SwitchWorkspace: React.FC = () => {
  // const [username, setUsername] = React.useState('')
  const [wsName, setWsName] = React.useState('')
  const [wsLogo, setWsLogo] = React.useState('')
  const [wcbShow, setWcbShow] = useState(false);



  useEffect(() => {
    // document.addEventListener("click", function (event) {
    //   console.log("点击位置：", event.target);
    // });

    workspaceList().then(res => {
      if (res.value.code === 0) {
        console.log(res.value);
        wsStore.setWsList(res.value.data ? res.value.data as IWorkspaceItemProps[] : [])
      } else {
        $message.error(res.value.msg)
      }
    }).catch(err => {
      console.log(err);
      $message.error(err.response.data.value.msg)
    })
  }, [])

  React.useEffect(() => {
    let currentWorkspace = currentWorkspaceStore.getCurrentWorkspace()
    setWsName(currentWorkspace.name ? currentWorkspace.name : '')
    setWsLogo(currentWorkspace.logo ? currentWorkspace.logo : '')
    // setUsername(appStore.getUserInfo().username)
    setWcbShow(false)
  }, [currentWorkspaceStore.currentWorkspace?.name, currentWorkspaceStore.currentWorkspace?.logo, appStore.userInfo.username])

  const wcbHandleClick = () => {
    setWcbShow(!wcbShow);
  };


  return (
    <>
      <div className="__sws">
        <div className="__sws_info">
          <div className="__sws_info_logo">
            {
              wsLogo === '' ? <Avatar onClick={wcbHandleClick} shape="square" style={{ backgroundColor: 'gray', verticalAlign: 'middle' }} size={28} gap={2}>
                {wsName ? wsName.charAt(0).toUpperCase() : 'N'}
              </Avatar> : <img height={28} width={28} style={{ borderRadius: '50%' }} src={wsLogo} alt="workspcae-logo" />
            }
          </div>

          <span className="__sws_info_title" onClick={wcbHandleClick}>{wsName ? wsName : 'No Workspace'}</span>
        </div>

        <div className="__sws_switch" onClick={wcbHandleClick}>
          <Switch />
        </div>
      </div>
      {wcbShow && <WorkspaceCreateBox setWcbShow={setWcbShow} />}
    </>
  )


}

export default observer(SwitchWorkspace)