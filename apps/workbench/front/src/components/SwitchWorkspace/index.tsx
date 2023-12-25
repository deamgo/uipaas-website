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

  const [wcbShow, setwcbShow] = useState(false);
  const [list_workspace, setlist_workspace] = useState([{ name: "Ilee's Workspace", logo: "" }])

  useEffect(() => {
    workspaceList().then(res => {
      if (res.value.code === 0) {
        console.log(res.value);
        setlist_workspace(res.value.data)
        wsStore.setWsList(res.value.data)
      } else {
        $message.error(res.value.msg)
      }
    }).catch(err => {
      console.log(err);
      $message.error(err.response.data.value.msg)
    })
  }, [])

  React.useEffect(() => {
    setWsName(currentWorkspaceStore.getCurrentWorkspace().name)
    setUsername(appStore.getUserInfo().username)
    setlist_workspace(wsStore.getWsList())
  }, [currentWorkspaceStore.currentWorkspace.name, currentWorkspaceStore.currentWorkspace.logo, appStore.userInfo.username, wsStore.getWsList()?.length])


  const wcbHandleClick = () => {
    setwcbShow(!wcbShow);
  };


  return (
    <>
      <div className="__sws">
        <Avatar style={{ backgroundColor: 'pink', verticalAlign: 'middle' }} size={28} gap={2}>
          {wsName.charAt(0).toUpperCase()}
        </Avatar>
        <span className="__sws_title">{wsName ? wsName : 'no Workspace'}</span>
        <div className="__sws_switch" onClick={wcbHandleClick}>
          <Switch />
        </div>
        {wcbShow && <WorkspaceCreateBox list_workspace={list_workspace} setlist_workspace={setlist_workspace} setwcbShow={setwcbShow} list={list_workspace} />}
      </div>
    </>
  )


}

export default observer(SwitchWorkspace)