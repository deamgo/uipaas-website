import React from 'react'
//
import './index.less'
import { ISelectProps } from '@/interface/some'
//
import { ReactComponent as Bingo } from '@assets/comps/bingo.svg'
import { ReactComponent as Down } from '@assets/comps/down.svg'

const Select: React.FC<ISelectProps> = (props) => {

  const [value, setValue] = React.useState<string>(props.default)
  const [isOpen, setIsOpen] = React.useState<boolean>(false)

  return (
    <>
      <div className="__select" onClick={() => setIsOpen(!isOpen)}>
        <div className="__select_shower">
          <div className="__select_shower_text">
            {props.children}

          </div>
          <div className="__select_shower_down">
            <Down />

          </div>
        </div>
        {isOpen && (<div className="__select_list">
          {props.list && props.list.map(item => (
            <div key={item.id} className={`__select_list_item ${props.default === item.label && '__select_list_active'}`}>
              {item.label}
              <div className="__select_list_item_icon">
                <Bingo />
              </div>
            </div>
          ))}
        </div>)}
      </div>
    </>
  )
}

export default Select