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


const SwitchWorkspace: React.FC = () => {

  const [username, setUsername] = React.useState('')
  const [wsName, setWsName] = React.useState('')
  const [wsLogo,setWsLogo] = React.useState('')
  const [wcbShow, setWcbShow] = useState(false);


  useEffect(() => {
    document.addEventListener("click", function(event) {
      console.log("点击位置：", event.target);
    });

    workspaceList().then(res => {
      if (res.value.code === 0) {
        console.log(res.value);
        wsStore.setWsList(res.value.data)
      } else {
        $message.error(res.value.msg)
      }
    }).catch(err => {
      console.log(err);
      $message.error(err.response.value.msg)
    })
  }, [])

  React.useEffect(() => {
    setWsName(currentWorkspaceStore.getCurrentWorkspace().name)
    setWsLogo(currentWorkspaceStore.getCurrentWorkspace().logo)
    setUsername(appStore.getUserInfo().username)
    setWcbShow(false)
  }, [currentWorkspaceStore.currentWorkspace.name, currentWorkspaceStore.currentWorkspace.logo, appStore.userInfo.username])

  const wcbHandleClick = () => {
    setWcbShow(!wcbShow);
  };


  return (
    <>
      <div className="__sws" >
        {
          wsLogo === ''? <Avatar  onClick={wcbHandleClick} style={{ backgroundColor: 'pink', verticalAlign: 'middle' }} size={28} gap={2}>
          {wsName.charAt(0).toUpperCase()}
        </Avatar>:<img height={28} width={28} style={{borderRadius:'50%'}} src={wsLogo} alt="workspcae-logo"/>
      }
        <span className="__sws_title" onClick={wcbHandleClick}>{wsName ? wsName : { username } + 's Workspace'}</span>
        <div className="__sws_switch" onClick={wcbHandleClick}>
          <Switch />
        </div>
        {wcbShow && <WorkspaceCreateBox  setWcbShow={setWcbShow}  />}
      </div>
    </>
  )


}

export default observer(SwitchWorkspace)