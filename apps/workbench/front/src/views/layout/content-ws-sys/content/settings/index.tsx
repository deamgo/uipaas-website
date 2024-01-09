import React from 'react'
// import { useNavigate } from 'react-router-dom'
//
import './index.less'
//
import Input from '@/components/Input'
import { Avatar } from 'antd'
import Button from '@/components/Button'
import Mask from '@/components/Mask'
import { currentWorkspaceStore } from '@/store/wsStore'
import DeleteWorkSpacePopup from './DeleteWorkspacePopup'
// import { deleteWorkspace } from '@/api/workspace_settings'
// import $message from '@/components/Message'
// import { workspaceList } from '@/api/workspace'
// import { IWorkspaceItemProps } from '@/interface/some'


const WSSettings: React.FC = () => {

  // const [workname, setWorkName] = React.useState('')
  const [currentWorkspace, setCurrentWorkspace] = React.useState(currentWorkspaceStore.getCurrentWorkspace())

  const [isMask, setIsMask] = React.useState(false)

  const [isDelete, setIsDelete] = React.useState(false)
  // const [confirmDeleteAble, setConfirmDeleteAble] = React.useState(true)

  // const navigate = useNavigate()

  React.useEffect(() => {
    setCurrentWorkspace(currentWorkspaceStore.getCurrentWorkspace())

  }, [currentWorkspaceStore.getCurrentWorkspace().name])

  // React.useEffect(() => {
  //   if (workname === currentWorkspaceStore.getCurrentWorkspace().name) {
  //     setConfirmDeleteAble(false)
  //   } else {
  //     setConfirmDeleteAble(true)
  //   }
  // }, [workname])

  React.useEffect(() => {
    document.body.style.overflow = isMask ? 'hidden' : 'auto'
  }, [isMask])

  // const resetForm = () => {
  //   setWorkName('')
  // }

  const handleMask = () => {
    setIsMask(!isMask)
  }

  const handleOpenDeleteDialog = () => {
    handleMask()
    // resetForm()
    setIsDelete(!isDelete)
  }

  // const handleConfirmDelete = () => {
  //   console.log('delete workspace');
  //   deleteWorkspace(currentWorkspace.id).then(res => {
  //     if (res.value.code === 0) {
  //       $message.success(res.value.msg)
  //       workspaceList().then(res => {
  //         if (res.value.code === 0) {
  //           wsStore.setWsList(res.value.data ? res.value.data as IWorkspaceItemProps[] : [])
  //           currentWorkspaceStore.setCurrentWorkspace(wsStore.getWsList().length > 0 ? wsStore.getWsList()[0] : { id: '', name: '', logo: '' })
  //           navigate('/')
  //         }
  //       }).catch(err => {
  //         console.log(err);
  //       })
  //       setIsDelete(false)
  //       handleMask()
  //     }
  //   }).catch(err => {
  //     console.log(err);
  //     $message.error(err.message)
  //   })
  // }

  return (
    <>
      <div className="__workspace_settings">
        <div className="__workspace_settings_container">
          <Avatar shape="square" style={{ backgroundColor: '#4080FF', verticalAlign: 'middle' }} size={48} gap={3}>
            {currentWorkspaceStore.getCurrentWorkspace().name.charAt(0).toUpperCase()}
          </Avatar>
          <div className="__workspace_settings_container_input">
            <Input
              id='1'
              title='Workspace Name'
              type='text'
              value={currentWorkspace.name}
              isNeed={false}
              typeAble={true} />
          </div>
        </div>
        <div className="__workspace_settings_delbtn">
          <Button context='Delete' type='outline-danger' method={handleOpenDeleteDialog} >
            Delete
          </Button>
        </div>
      </div>
      {isMask && (<Mask />)}
      {isDelete && (
        <>
          <DeleteWorkSpacePopup onClose={handleOpenDeleteDialog} />
          {/* <Popup unit='rem' width={496} height={276} title={`Delete the "${currentWorkspaceStore.getCurrentWorkspace().name}" Workspace?`} onClose={handleOpenDeleteDialog}>
            <div className="_delete_tips">
              <div className="_delete_tips_bef">
                <span>All Workspace applications will be permanently deleted.</span>
                <br />
                <span>You can't undo this action.</span>
              </div>
              <div className="_delete_tips_pls">
                <span>If you're sure about deleting, enter the Workspace name to confirm.</span>
              </div>
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
                <Button context='Cancel' type='default' method={handleOpenDeleteDialog} >
                  Cancel
                </Button>
              </div>
              <div className="_delete_btn_group_confirm">
                <Button context='Delete' type='danger' disabled={confirmDeleteAble} method={handleConfirmDelete} >
                  Delete
                </Button>
              </div>
            </div>
          </Popup> */}
        </>
      )}
    </>
  )
}

export default WSSettings