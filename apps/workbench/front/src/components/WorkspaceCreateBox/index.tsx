import React from "react"
//
import './index.less'

import { ReactComponent as Wca } from '@assets/layout/workspace-create-add.svg'
import WorkspaceCreateBoxItem from "@components/WorkspaceCreateBoxItem";
import { currentWorkspaceStore, wsStore } from "@store/wsStore.ts";
import WorkspaceCreatePopup from "./WorkspaceCreatePopup";
import Mask from "../Mask";


interface WorkspaceItem {
    id: string
    name: string
    logo: string
    lable: string
    description: string
}

type BoxProps = {
    list?: WorkspaceItem[]
    list_workspace?: { name: string, logo: string }[]
    setlist_workspace?: React.Dispatch<React.SetStateAction<{ name: string, logo: string }[]>>
    setWcbShow?: React.Dispatch<React.SetStateAction<boolean>>
}


const WorkspaceCreateBox: React.FC<BoxProps> = () => {
    const [isWsCreate, setIsWsCreate] = React.useState(false)
    const [isMask, setIsMask] = React.useState(false)
    // const [file, setFile] = useState<File | null>(null);
    // const [workspaceName, setWorkspaceName] = React.useState("")
    // const [workspaceLogoPath, setWorkspaceLogoPath] = React.useState<string>("")
    // const formData = new FormData();




    // const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    //     if (!e.target.files) {
    //         return;
    //     }
    //     const file = e.target.files[0];
    //     // setFile(file)
    //     if (file != null) {
    //         console.log(file)
    //         formData.set('file', file)

    //         workspaceLogo(formData).then(res => {
    //             if (res.value.code === 0) {
    //                 console.log(res.value)
    //                 setWorkspaceLogoPath(res.value.data as string)
    //             } else {
    //                 $message.error(res.value.msg)
    //             }
    //         }).catch(err => {
    //             console.log(err);
    //             $message.error(err.response.data.value.msg)
    //         })
    //     }
    // };

    const handleMask = () => {
        setIsMask(!isMask);
    }



    const wcbpHandleClick = () => {
        handleMask()
        setIsWsCreate(!isWsCreate);
        // props.setwcbShow(true);
    };

    const handleIsWsCreate = () => {
        handleMask()
        setIsWsCreate(!isWsCreate)
    }


    // const reqWorkspaceCreate = async () => {
    //     await workspaceCreate({
    //         name: workspaceName,
    //         logo: workspaceLogoPath,
    //     }).then(res => {
    //         if (res.value.code === 0) {
    //             console.log(res.value.data)
    //             let ws = res.value.data as IWorkspaceItemProps
    //             let list = wsStore.getWsList()
    //             list.push(ws)
    //             wsStore.setWsList(list)
    //             currentWorkspaceStore.setCurrentWorkspace(ws)
    //             wsStore.setFirst(ws.name)
    //             setIsWsCreate(false)
    //             setWorkspaceLogoPath("")
    //             $message.success(res.msg)
    //         } else {
    //             $message.error(res.msg)
    //         }
    //     }).catch(err => {
    //         console.log(err);
    //         $message.error(err.response.data.value.msg)
    //     })

    //     await workspaceList().then(res => {
    //         if (res.value.code === 0) {
    //             console.log(res.value.data)
    //             props.setlist_workspace && props.setlist_workspace(res.value.data as IWorkspaceItemProps[])
    //             wsStore.setWsList(res.value.data as IWorkspaceItemProps[])
    //         } else {
    //             $message.error(res.value.msg)
    //         }
    //     }).catch(err => {
    //         console.log(err);

    //     })
    // }

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
            {isMask && (<Mask />)}
            {isWsCreate && (
                <>
                    <WorkspaceCreatePopup onClose={handleIsWsCreate} />
                </>
            )}
        </>
    )
}

export default WorkspaceCreateBox