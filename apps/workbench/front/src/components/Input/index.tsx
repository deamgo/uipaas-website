import React from "react"
//style
import './index.less'
//
import Eye_open from '@assets/sign/eye-open.svg'
//


interface InputProps {
  id: string
  title: string
  placeholder: string
  type: 'password' | 'text'
  valid: string
  outputChange: (value: string) => void
  reg: RegExp
  validate?: (value: string, regex: RegExp) => boolean
}

const validator = (value: string, regex: RegExp) => {
  return regex.test(value)
}

const Input: React.FC<InputProps> = (props) => {

  const [eyev, setEyev] = React.useState(false);
  const [showV, setShowV] = React.useState(false);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value
    // if (value === '' && value === undefined) {
    //   setShowV(false)
    // }
    if (validator(value, props.reg)) {
      setShowV(false)
    } else {
      setShowV(true)
    }
    props.outputChange(e.target.value)
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
        <div className="__input_wrapper_title">
          {props.title}
          <span>*</span>
        </div>
        <div className="__input_wrapper_main">
          <input
            type={eyev ? 'text' : props.type}
            name=""
            id={props.id}
            placeholder={props.placeholder}
            onChange={(e) => props.outputChange(e.target.value)}
            onBlur={handleChange} />
          {
            props.type === 'password'
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