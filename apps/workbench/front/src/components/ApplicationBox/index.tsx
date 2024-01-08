import React from 'react'
//
import './index.less'
import Label from '../Label'
import Button from '../Button'
//
import { ReactComponent as More } from '@assets/comps/more.svg'
import { Avatar } from 'antd'
import { ISelectOption } from '@/interface/some'
import DownList from '../DownList'

interface IApplicationBoxProps {
  id: string
  name: string
  label: number
  desc: string
  type?: number
}

const format = (status: number) => {
  switch (status) {
    case 1:
      return 'info'
    case 2:
      return 'success'
    case 3:
      return 'warning'
    default:
      return 'error'
  }
}

const OperationList: ISelectOption[] = [
  {
    id: 'setting',
    value: 'setting',
    label: 'Setting'
  },
  {
    id: 'duplicate',
    value: 'duplicate',
    label: 'Duplicate'
  },
  {
    id: 'export',
    value: 'export',
    label: 'Export'
  },
  {
    id: 'delete',
    value: 'delete',
    label: 'Delete'
  },
]

const ApplicationBox: React.FC<IApplicationBoxProps> = (props) => {
  const [isMore, setIsMore] = React.useState(false)
  return (
    <>
      <div className="__applicationbox">
        <div className="__applicationbox_head">
          <div className="__applicationbox_head_info">
            <div className="__applicationbox_head_icon">
              <Avatar style={{ backgroundColor: 'skyblue', verticalAlign: 'middle' }} size={32} gap={2}>
              </Avatar>
            </div>
            <div className="__applicationbox_head_name">
              {props.name}
            </div>
          </div>
          <div className="__applicationbox_head_label">
            <Label label={props.label === 1 ? 'Draft' : 'Public'} type={format(props.type ? props.type : 0)} />
          </div>
        </div>
        <div className="__applicationbox_main">
          <p className="__applicationbox_main_desc">
            {props.desc}
          </p>
        </div>
        <div className="__applicationbox_foot">
          <div className="__applicationbox_foot_btns">
            <div className="__applicationbox_foot_edit">
              <Button type='board-primary'>
                Edit
              </Button>
            </div>
            <div className="__applicationbox_foot_preview">
              <Button type='board-primary'>
                Preview
              </Button>
            </div>
          </div>
          <div onClick={() => setIsMore(!isMore)} className="__applicationbox_foot_more">
            <More />
          </div>
        </div>
        {isMore && (
          <div className="_application_more_list">
            <DownList list={OperationList} />
          </div>
        )}
      </div>
    </>
  )
}

export default ApplicationBox