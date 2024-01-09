import React from 'react'
//
import './index.less'
import { IDownListProps } from '@/interface/some'
//

const DownList: React.FC<IDownListProps> = (props) => {
  return (
    <>
      <div className="__downlist">
        {props.list && props.list.map(item => (
          <div key={item.id} className={`__select_list_item ${false && '__select_list_active'}`}>
            {item.label}
            {/* <div className="__select_list_item_icon">
                <Bingo />
              </div> */}
          </div>
        ))}
      </div>
    </>
  )
}

export default DownList