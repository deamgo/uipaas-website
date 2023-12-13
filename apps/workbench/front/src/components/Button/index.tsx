import React from 'react'
//style
import './index.less'
//

interface ButtonProps {
  context: string
  method: () => void
  disabled?: boolean
}

const Button: React.FC<ButtonProps> = (props) => {
  return (
    <>
      <button
        type="button"
        onClick={props.method}
        className='__button_wrapper'
        disabled={props.disabled}>
        {props.context}
      </button>
    </>
  )
}

export default Button