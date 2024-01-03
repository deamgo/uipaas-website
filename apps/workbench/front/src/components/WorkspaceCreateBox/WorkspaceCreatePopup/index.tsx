import { useState } from 'react'
//
import './index.less'
import Popup from '@/components/Popup'
import { Avatar } from 'antd'
import Input from '@/components/Input'
import Button from '@/components/Button'
import { workspaceCreate, workspaceList, workspaceLogo } from '@/api/workspace'
import $message from '@/components/Message'
import { IWorkspaceItemProps } from '@/interface/some'
import { currentWorkspaceStore, wsStore } from '@/store/wsStore'
//

interface IWsCPProps {
  onClose?: () => void
  setlist_workspace?: React.Dispatch<React.SetStateAction<{ name: string, logo: string }[]>>
}

const WorkspaceCreatePopup: React.FC<IWsCPProps> = (props) => {
  // const [isWsCreate, setIsWsCreate] = useState(false)
  // const [file, setFile] = useState<File | null>(null);
  const [workspaceName, setWorkspaceName] = useState("")
  const [workspaceLogoPath, setWorkspaceLogoPath] = useState<string>("")
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
          setWorkspaceLogoPath(res.value.data as string)
        } else {
          $message.error(res.value.msg)
        }
      }).catch(err => {
        console.log(err);
        $message.error(err.response.data.value.msg)
      })
    }
  };

  // const handleIsWsCreate = () => {
  //   setIsWsCreate(!isWsCreate);
  // }


  const reqWorkspaceCreate = async () => {
    await workspaceCreate({
      name: workspaceName,
      logo: workspaceLogoPath,
    }).then(res => {
      if (res.value.code === 0) {
        console.log(res.value.data)
        let ws = res.value.data as IWorkspaceItemProps
        let list = wsStore.getWsList()
        list.push(ws)
        wsStore.setWsList(list)
        currentWorkspaceStore.setCurrentWorkspace(ws)
        wsStore.setFirst(ws.name)
        props.onClose && props.onClose()
        setWorkspaceLogoPath("")
        $message.success(res.msg)
      } else {
        $message.error(res.msg)
      }
    }).catch(err => {
      console.log(err);
      $message.error(err.response.data.value.msg)
    })

    await workspaceList().then(res => {
      if (res.value.code === 0) {
        console.log(res.value.data)
        // props.setlist_workspace && props.setlist_workspace(res.value.data as IWorkspaceItemProps[])
        wsStore.setWsList(res.value.data as IWorkspaceItemProps[])
      } else {
        $message.error(res.value.msg)
      }
    }).catch(err => {
      console.log(err);
    })
  }
  return (
    <>
      <Popup unit={'rem'} width={496} height={276} title={'Create Workspace'} onClose={props.onClose}>
        <div className="__wcb_popup">

          <div className="__wcb_popup_container ">

            <label htmlFor="workspace-logo">
              <div className="__wcb_popup_container_setLogo"></div>
              <input style={{ display: "none" }} id="workspace-logo" type="file" onChange={handleFileChange} />
              {workspaceLogoPath === '' ?
                <Avatar style={{ backgroundColor: 'pink', verticalAlign: 'middle' }} size={65}
                  gap={3}>
                  {workspaceName === '' ? 'E' : workspaceName.charAt(0).toUpperCase()}
                </Avatar> : <img style={{ borderRadius: '50%' }} height={65} width={65} src={workspaceLogoPath} alt="workspace-logo" />}

            </label>
            <div className="__wcb_popup_container_input">
              <Input
                id='workspacename'
                title='Workspace Name'
                isNeed={true}
                type='text'
                outputChange={setWorkspaceName}
                placeholder={"Enter your Workspace name"}
              />
            </div>
          </div>


          <div className="__wcb_popup_button">
            <Button context={'Create'} type='primary' method={reqWorkspaceCreate} >
              Create
            </Button>
          </div>
        </div>

      </Popup>
    </>
  )
}

export default WorkspaceCreatePopup