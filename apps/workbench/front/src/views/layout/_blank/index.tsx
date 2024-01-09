import React from 'react'
//
import './index.less'
import Button from '@/components/Button'
//
import { ReactComponent as Plus } from '@assets/comps/plus.svg'
import { ReactComponent as NoneContent } from '@assets/default/none-content.svg'
import WorkspaceCreatePopup from '@/components/WorkspaceCreateBox/WorkspaceCreatePopup'

const _Blank: React.FC = () => {
  const [isCreate, setIsCreate] = React.useState(false)

  const handleCreate = () => {
    setIsCreate(!isCreate)
  }
  return (
    <>
      <div className="__blank">
        <div className="__blank_wrapper">
          <div className="__blank_svg">
            <NoneContent />
          </div>
          <span className="__blank_title">Create your first Workspace</span>
          <span className="__blank_tips">Start developing your applications by creating your first Workspace!</span>
          <div className="__blank_btn">
            <Button type='primary' method={handleCreate}>
              <Plus style={{
                width: '10.67rem',
                height: '10.67rem',
                fill: '#FFFFFF'
              }} />
              Create
            </Button>
          </div>
        </div>
      </div>
      {isCreate && (
        <>
          <WorkspaceCreatePopup onClose={handleCreate} />
        </>
      )}
    </>
  )
}

export default _Blank