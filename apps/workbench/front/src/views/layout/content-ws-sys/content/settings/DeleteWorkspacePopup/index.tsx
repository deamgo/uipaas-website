import React from 'react'
//
import './index.less'
import Popup from '@/components/Popup'
import Input from '@/components/Input'
import Button from '@/components/Button'
import $message from '@/components/Message'
import { deleteWorkspace } from '@/api/workspace_settings'
import { currentWorkspaceStore, wsStore } from '@/store/wsStore'
import { workspaceList } from '@/api/workspace'
import { IWorkspaceItemProps } from '@/interface/some'
import { useNavigate } from 'react-router-dom'
//

interface IDeleteWorkSpacePopupProps {
  onClose: () => void
}

const DeleteWorkSpacePopup: React.FC<IDeleteWorkSpacePopupProps> = (props) => {
  const [workname, setWorkName] = React.useState('')
  const [currentWorkspace, setCurrentWorkspace] = React.useState(currentWorkspaceStore.getCurrentWorkspace())


  // const [isMask, setIsMask] = React.useState(false)

  const [isDelete, setIsDelete] = React.useState(false)

  const [confirmDeleteAble, setConfirmDeleteAble] = React.useState(true)

  const navigate = useNavigate()


  React.useEffect(() => {
    setCurrentWorkspace(currentWorkspaceStore.getCurrentWorkspace())

  }, [currentWorkspaceStore.getCurrentWorkspace().name])

  React.useEffect(() => {
    if (workname === currentWorkspaceStore.getCurrentWorkspace().name) {
      setConfirmDeleteAble(false)
    } else {
      setConfirmDeleteAble(true)
    }
  }, [workname])

  const handleOpenDeleteDialog = () => {
    // handleMask()
    resetForm()
    setIsDelete(!isDelete)
  }

  const resetForm = () => {
    setWorkName('')
  }


  const handleConfirmDelete = () => {
    console.log('delete workspace');
    deleteWorkspace(currentWorkspace.id).then(res => {
      if (res.value.code === 0) {
        $message.success(res.value.msg)
        workspaceList().then(res => {
          if (res.value.code === 0) {
            wsStore.setWsList(res.value.data ? res.value.data as IWorkspaceItemProps[] : [])
            currentWorkspaceStore.setCurrentWorkspace(wsStore.getWsList().length > 0 ? wsStore.getWsList()[0] : { id: '', name: '', logo: '' })
            navigate('/')
          }
        }).catch(err => {
          console.log(err);
        })
        setIsDelete(false)
        // handleMask()
      }
    }).catch(err => {
      console.log(err);
      $message.error(err.message)
    })
  }
  return (
    <>
      <Popup unit='rem' width={496} onClose={props.onClose}>
        <div className="_delete_title">
          <h1 className="_delete_title_main">
            {`Delete the "${currentWorkspaceStore.getCurrentWorkspace().name}" Workspace?`}
          </h1>
          <span className="_delete_title_tip">
            All Workspace applications will be permanently deleted.
          </span>
          <span className="_delete_title_worning">
            You can't undo this action.
          </span>
        </div>
        <div className="_delete_desc">
          <span>If you're sure about deleting, enter the Workspace name to confirm.</span>
        </div>
        <div className="_delete_spell">
          <Input
            id='spell_wsname'
            type='text'
            placeholder='Workspace name'
            outputChange={setWorkName} />
        </div>
        <div className="_delete_btn_group">
          <div className="_delete_btn_group_cancel">
            <Button context='Cancel' type='outline-primary' method={props.onClose} >
              Cancel
            </Button>
          </div>
          <div className="_delete_btn_group_confirm">
            <Button context='Delete' type='danger' disabled={confirmDeleteAble} method={handleConfirmDelete} >
              Delete
            </Button>
          </div>
        </div>
      </Popup>
    </>
  )
}

export default DeleteWorkSpacePopup