import React from 'react'
//
import './index.less'
//
import { ReactComponent as CloseX } from '@assets/comps/close-x.svg'

interface IPopupProps {
  unit: string
  width: number
  height?: number
  title?: string
  onClose?: () => void
  children?: React.ReactNode
}

const Popup: React.FC<IPopupProps> = (props) => {
  return (
    <>
      <div className="__popup" style={{
        width: props.width + props.unit
      }}>
        <div className="__popup_wrapper_close" onClick={props.onClose}>
          <CloseX />
        </div>
        <div className="__popup_wrapper">
          {props.children}
        </div>
      </div>
      {/* <div className="__popup_wrapper" style={{
        width: props.width + props.unit,
        height: props.height + props.unit
      }}>
        {props.title && (
          <h1 className="__popup_wrapper_title">
            {props.title}
          </h1>
        )}

        {props.children}
      </div> */}
    </>
  )
}

export default Popup