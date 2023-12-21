import React from 'react'
//
import './index.less'
//

interface PopupProps {
  unit: string
  width: number
  height: number
  title?: string
  onClose?: () => void
  children?: React.ReactNode
}

const Popup: React.FC<PopupProps> = (props) => {
  return (
    <>
      <div className="__popup_wrapper" style={{
        width: props.width + props.unit,
        height: props.height + props.unit
      }}>
        {props.title && (
          <h1 className="__popup_wrapper_title">
            {props.title}
          </h1>
        )}
        <div className="__popup_wrapper_close" onClick={props.onClose}>
          <div className="_close_x_pone"></div>
          <div className="_close_x_ptwo"></div>
        </div>
        {props.children}
      </div>
    </>
  )
}

export default Popup