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
import { IUsrWorkspace, workspaceCreate, workspaceLogo } from "@api/workspace.ts";
import $message from "@components/Message";


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
    setwcbShow: React.Dispatch<React.SetStateAction<boolean>>
    setlist_workspace: React.Dispatch<React.SetStateAction<{ name: string, logo: string }[]>>
}


const WorkspaceCreateBox: React.FC<BoxProps> = (props) => {
    const [isWsCreate, setIsWsCreate] = React.useState(false)
    const [workspaceName, setworkspaceName] = React.useState("")
    const [workspaceLogoPath, setworkspaceLogoPath] = React.useState("")
    const [file, setFile] = useState<File | null>(null);
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
                    setworkspaceLogoPath(res.value.data)
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


    const reqWorkspaceCreate = () => {
        workspaceCreate({
            name: workspaceName,
            logo: workspaceLogoPath,
        }).then(res => {
            if (res.value.code === 0) {
                console.log(res.value.data)
                props.list_workspace.push({ name: res.value.data.name as string, logo: res.value.data.logo as string })
                setIsWsCreate(false)
            } else {
                $message.error(res.value.msg)
            }
        }).catch(err => {
            console.log(err);
            $message.error(err.response.data.value.msg)
        })
    }



    return (
        <>
            <div className="__wcb">
                <div className="__wcb_box">
                    {props.list
                        ? props.list.map(item => (
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

                                <label onClick={handleUpload} htmlFor="workspace-logo">
                                    <input style={{ display: "none" }} id="workspace-logo" type="file" onChange={handleFileChange} />
                                    {workspaceLogoPath === '' ?
                                        <Avatar style={{ backgroundColor: '#4080FF', verticalAlign: 'middle' }} size={65}
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
                                        outputChange={setworkspaceName}
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