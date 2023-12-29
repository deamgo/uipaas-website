import React from 'react'
//
import './index.less'
import Label from '../Label'
import Button from '../Button'
//
import { ReactComponent as More } from '@assets/comps/more.svg'
import Divider from '../Divider'

interface IApplicationBoxProps {
  id: string
  name: string
  label: string
  desc: string
  type: string
}

const ApplicationBox: React.FC<IApplicationBoxProps> = (props) => {
  return (
    <>
      <div className="__applicationbox">
        <div className="__applicationbox_head">
          <div className="__applicationbox_head_icon"></div>
          <div className="__applicationbox_head_name">
            {props.name}
          </div>
          <div className="__applicationbox_head_label">
            <Label label={props.label} type={props.type} />
          </div>
        </div>
        <div className="__applicationbox_main">
          <p className="__applicationbox_main_desc">
            {props.desc}
          </p>
        </div>
        <Divider ys={{
          width: '336rem'
        }} />
        <div className="__applicationbox_foot">
          <div className="_applicationbox_foot_edit">
            <Button type='board-primary'>
              Edit
            </Button>
          </div>
          <div className="_applicationbox_foot_preview">
            <Button type='board-primary'>
              Preview
            </Button>
          </div>
          <div className="_applicationbox_foot_more">
            <More />
          </div>
        </div>
      </div>
    </>
  )
}

export default ApplicationBox