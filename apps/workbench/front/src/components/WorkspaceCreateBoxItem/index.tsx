import React from "react"
//
import './index.less'

import {Avatar} from "antd";

interface WorkspaceItemProps{
    logo:string;
    name:string;
}

const WorkspaceCreateBoxItem: React.FC<WorkspaceItemProps> = (props) => {

    return (
        <>
            <div className="__wcb_box_item">
                { props.logo === "" && <>
                    <Avatar className="__wcb_box_item_logo" style={{backgroundColor: 'pink', verticalAlign: 'middle'}}
                            size={20}>
                        {props.name.charAt(0).toUpperCase()}
                    </Avatar>
                    <span className="__wcb_box_item_tittle">{props.name}</span>
                </>
                }
                {
                    props.logo !== "" &&
                    <>
                        <img className="__wcb_box_item_logo" src={props.logo} width={20} height={20} />
                        <span className="__wcb_box_item_tittle">{props.name}</span>
                    </>
                }
            </div>
        </>
    )
}

export default WorkspaceCreateBoxItem