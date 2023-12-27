import React, { useEffect, useRef, useState } from "react"
//
import './index.less'

import axios from "axios"
import { ReactComponent as Wca } from '@assets/layout/workspace-create-add.svg'
import WorkspaceCreateBoxItem from "@components/WorkspaceCreateBoxItem";
import Popup from "@components/Popup";
import { Avatar } from "antd";
import Button from "@components/Button";
import Input from "@components/Input"
import { IUsrWorkspace, workspaceCreate, workspaceList, workspaceLogo } from "@api/workspace.ts";
import $message from "@components/Message";
import { currentWorkspaceStore, wsStore } from "@store/wsStore.ts";
import { IWorkspaceItemProps } from "@/interface/some.ts";
import { Header } from "antd/es/layout/layout";


interface WorkspaceItem {
    id: string
    name: string
    logo: string
    lable: string
    description: string
}

type BoxProps = {
    list: WorkspaceItem[]
    list_workspace: { name: string, logo: string }[]
    setlist_workspace: React.Dispatch<React.SetStateAction<{ name: string, logo: string }[]>>
    setWcbShow: React.Dispatch<React.SetStateAction<boolean>>
}


const WorkspaceCreateBox: React.FC<BoxProps> = (props) => {
    const [isWsCreate, setIsWsCreate] = React.useState(false)
    const [file, setFile] = useState<File | null>(null);
    const [workspaceName, setWorkspaceName] = React.useState("")
    const [workspaceLogoPath, setWorkspaceLogoPath] = React.useState("")
    const formData = new FormData();




    const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        if (!e.target.files) {
            return;
        }
        const file = e.target.files[0];
        // setFile(file)
        if (file != null) {
            console.log(file)
            formData.set('file', file)

            workspaceLogo(formData).then(res => {
                if (res.value.code === 0) {
                    console.log(res.value)
                    setWorkspaceLogoPath(res.value.data)
                } else {
                    $message.error(res.value.msg)
                }
            }).catch(err => {
                console.log(err);
                $message.error(err.response.data.value.msg)
            })
        }
    };



    const wcbpHandleClick = () => {
        setIsWsCreate(!isWsCreate);
        // props.setwcbShow(true);
    };

    const handleIsWsCreate = () => {
        setIsWsCreate(!isWsCreate);
    }


    const reqWorkspaceCreate = async () => {
        await workspaceCreate({
            name: workspaceName,
            logo: workspaceLogoPath,
        }).then(res => {
            if (res.value.code === 0) {
                console.log(res.value.data)
                let ws = { id: res.value.data.id, name: res.value.data.name, logo: res.value.data.logo }
                let list = wsStore.getWsList();
                list.push(ws)
                wsStore.setWsList(list)
                currentWorkspaceStore.setCurrentWorkspace(ws)
                wsStore.setFirst(ws.name)
                setIsWsCreate(false)
                setWorkspaceLogoPath("")
                $message.success(res.value.msg)
            } else {
                $message.error(res.value.msg)
            }
        }).catch(err => {
            console.log(err);
            $message.error(err.response.data.value.msg)
        })

        await workspaceList().then(res => {
            if (res.value.code === 0) {
                console.log(res.value.data)
                props.setlist_workspace(res.value.data)
                wsStore.setWsList(res.value.data)
            } else {
                $message.error(res.value.msg)
            }
        }).catch(err => {
            console.log(err);

        })
    }

    return (
        <>
            <div className="__wcb">
                <div className="__wcb_box">
                    {wsStore.getWsListFirstByWorkspace(currentWorkspaceStore.getCurrentWorkspace().name)
                        ? wsStore.getWsListFirstByWorkspace(currentWorkspaceStore.getCurrentWorkspace().name).map(item => (
                            <WorkspaceCreateBoxItem id={item.id} name={item.name} logo={item.logo} />
                        ))
                        : (<></>)}
                </div>
                <button className="__wcb_button" onClick={wcbpHandleClick}>
                    <hr className="__wcb_line" />
                    <span className="__wcb_add_logo">
                        <Wca />
                    </span>
                    Create Workspace
                </button>
            </div>
            {isWsCreate && (
                <>
                    <Popup unit={'rem'} width={496} height={276} title={'Create Workspace'} onClose={handleIsWsCreate}>
                        <div className="_current" style={{ fontSize: '13px' }}>

                            <div className="__user_profile_account_container_wrapper_input _sp_withAvatar ">

                                <label htmlFor="workspace-logo">
                                    <div className="__wcb_popup_setLogo"></div>
                                    <input style={{ display: "none" }} id="workspace-logo" type="file" onChange={handleFileChange} />
                                    {workspaceLogoPath === '' ?
                                        <Avatar style={{ backgroundColor: 'pink', verticalAlign: 'middle' }} size={65}
                                            gap={3}>
                                            {workspaceName === '' ? 'E' : workspaceName.charAt(0).toUpperCase()}
                                        </Avatar> : <img style={{ borderRadius: '50%' }} height={65} width={65} src={workspaceLogoPath} alt="workspace-logo" />}

                                </label>
                                <div className="__user_profile_account_container_wrapper_input_besideAvatar">
                                    <div style={{ marginBottom: '13px' }}><span
                                        style={{ color: '#4E5969' }}>Workspace Name</span><span
                                            style={{ color: '#FF4D4F' }}>*</span></div>

                                    <Input
                                        type='text'
                                        outputChange={setWorkspaceName}
                                        placeholder={"Enter your Workspace name"}
                                    />
                                </div>
                            </div>


                            <div className="__wcb_popup_button">
                                <Button context={'Create'} type='primary' method={reqWorkspaceCreate} />
                            </div>
                        </div>

                    </Popup>
                </>
            )}
        </>
    )
}

export default WorkspaceCreateBox