import React from 'react'
//
import './index.less'
//

interface ILabelProps {
  label: string
  type: string
  children?: React.ReactNode
}

const Label: React.FC<ILabelProps> = (props) => {
  return (
    <>
      <span className={`__label ${props.type}`}>
        {props.label}
      </span>
    </>
  )
}

export default Label