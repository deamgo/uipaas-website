import React from "react"
import { useNavigate } from 'react-router-dom'
//
import './index.less'

import { Avatar } from "antd";
import { currentWorkspaceStore, wsStore } from "@/store/wsStore";
import { applicationStore } from "@/store/application";

interface WorkspaceItemProps {
    id: string;
    logo: string;
    name: string;
}

const WorkspaceCreateBoxItem: React.FC<WorkspaceItemProps> = (props) => {

    const navigate = useNavigate()

    const handleChangeCurrentWorkspace = () => {
        currentWorkspaceStore.setCurrentWorkspace(props)

        wsStore.setFirst(props.name)

        applicationStore.setAppStats()

        navigate('/')
    }

    return (
        <>
            <div className="__wcb_box_item" onClick={handleChangeCurrentWorkspace}>
                {props.logo === "" && <>
                    <Avatar shape="square" className="__wcb_box_item_logo" style={{ backgroundColor: 'gray', verticalAlign: 'middle' }}
                        size={24}>
                        {props.name.charAt(0).toUpperCase()}
                    </Avatar>
                    <span className="__wcb_box_item_tittle">{props.name}</span>
                </>
                }
                {
                    props.logo !== "" &&
                    <>
                        <img className="__wcb_box_item_logo" src={props.logo} width={20} height={20} alt="workspace logo" />
                        <span className="__wcb_box_item_tittle">{props.name}</span>
                    </>
                }
            </div>
        </>
    )
}

export default WorkspaceCreateBoxItem