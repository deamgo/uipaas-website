import React, {useEffect, useState} from "react"
//
import './index.less'
//
import { ReactComponent as Switch } from '@assets/layout/switch.svg'
//
import { Avatar } from 'antd'
import WorkspaceCreateBox from "@components/WorkspaceCreateBox";
import {workspaceList} from "@api/workspace.ts";
import $message from "@components/Message";


const SwitchWorkspace: React.FC = () => {
    const [wcbShow,setwcbShow]= useState(false);
    const [list_workspace,setlist_workspace] = useState([{name:"Ilee's Workspace",logo:""}])


    const wcbHandleClick = () => {
        setwcbShow(!wcbShow);
    };


    useEffect(() => {
        workspaceList().then(res => {
            if (res.value.code === 0) {
                setlist_workspace(res.value.data)
            } else {
                $message.error(res.value.msg)
            }
        }).catch(err => {
            console.log(err);
            $message.error(err.response.data.value.msg)
        })
    },[])


  return (
      <>
        <div className="__sws">
          <Avatar style={{backgroundColor: 'pink', verticalAlign: 'middle'}} size={28} gap={2}>
            {'ILEE'.charAt(0).toUpperCase()}
          </Avatar>
          <span className="__sws_title">{'Ilee'}'s Workspace</span>
          <div className="__sws_switch" onClick={wcbHandleClick}>
            <Switch/>
          </div>
            {wcbShow && <WorkspaceCreateBox list_workspace={list_workspace} setlist_workspace={setlist_workspace} setwcbShow={setwcbShow} list={list_workspace} />}
        </div>
      </>
  )


}
export default SwitchWorkspace