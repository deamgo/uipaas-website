import React from 'react'
//style
import './index.less'
//

interface ButtonProps {
  context?: string
  method?: () => void
  disabled?: boolean
  type?: 'primary' | 'danger' | 'default' | 'outline-danger' | 'board-danger' | 'board-primary'
  children?: React.ReactNode
}

const Button: React.FC<ButtonProps> = (props) => {
  return (
    <>
      <button
        type="button"
        onClick={props.method}
        className={`__button_wrapper ${props.type}`}
        disabled={props.disabled}>
        {props.children}
      </button>
    </>
  )
}

export default Button