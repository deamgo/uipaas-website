import React from 'react'
//
import './index.less'
import Button from '@/components/Button'
//

const _Blank: React.FC = () => {
  return (
    <>
      <div className="__blank">
        <span className="__blank_title">Create your first Workspace</span>
        <span className="__blank_tips">Start developing your applications by creating your first Workspace!</span>
        <div className="__blank_btn">
          <Button context='Create' type='primary' />
        </div>
      </div>
    </>
  )
}

export default _Blank