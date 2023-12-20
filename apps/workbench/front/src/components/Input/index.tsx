import React from "react"
//style
import './index.less'
//
import Eye_open from '@assets/sign/eye-open.svg'
//


interface InputProps {
  id: string
  title?: string
  placeholder?: string
  type?: 'password' | 'text'
  valid?: string
  isNeed?: boolean
  isShowPwd?: boolean
  outputChange?: (value: string) => void
  reg?: RegExp
  validate?: (value: string, regex: RegExp) => boolean
  typeAble?: boolean
  value?: string
}

const validator = (value: string, regex: RegExp) => {
  return regex.test(value)
}

const Input: React.FC<InputProps> = (props) => {

  const [eyev, setEyev] = React.useState(false);
  const [showV, setShowV] = React.useState(false);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value
    if (props.reg && validator(value, props.reg)) {
      setShowV(false)
    } else {
      setShowV(true)
    }
    props.outputChange && props.outputChange(e.target.value)
  }

  // const formatEmpty = (value: string): string => {
  //   return value + 'cant be empty'
  // }

  const handleShowPwd = () => {
    setEyev(eyev => !eyev)

  }

  return (
    <>
      <div className="__input_wrapper">
        {props.title && (
          <div className="__input_wrapper_title">
            {props.title}
            {props.isNeed && <span>*</span>}
          </div>
        )}
        <div className="__input_wrapper_main">
          <input
            type={eyev ? 'text' : props.type}
            name=""
            id={props.id}
            placeholder={props.placeholder}
            value={props.value}
            onChange={(e) => props.outputChange && props.outputChange(e.target.value)}
            onBlur={handleChange}
            disabled={props.typeAble ? true : false} />
          {
            props.isShowPwd
            && (
              <>
                <div className="__input_wrapper_main_show" onClick={handleShowPwd}>
                  <img
                    src={Eye_open} />
                  {!eyev && (<div className="__show_closed"></div>)}
                </div>
              </>
            )
          }
        </div>
        <div className={showV ? "__input_wrapper_invalid invalid" : "__input_wrapper_invalid"}>
          {props.valid}
        </div>
      </div>
    </>
  )
}

export default Input