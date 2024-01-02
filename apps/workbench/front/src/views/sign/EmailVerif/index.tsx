import React, { useState } from 'react'
//style
import './index.less'
//
import Button from '@/components/Button'
//
import { emailVerificationReg } from '@constants/regexp'
import { usrSignUpVerify } from '@api/sign_up'
import { appStore } from '@/store/store'
//
import $message from '@/components/Message'
import { Link, useNavigate } from 'react-router-dom'
import Cookies from 'js-cookie'

const EmailVerif: React.FC = () => {

  const [code, setCode] = useState(['', '', '', ''])
  const [emailVerification, setEmailVerification] = useState('')
  const [btnAbled, setBtnAbled] = React.useState(true)

  let typeCode = ''

  const navigate = useNavigate()

  React.useEffect(() => {
    if (emailVerificationReg.test(emailVerification)) {
      setBtnAbled(false)
    } else {
      setBtnAbled(true)
    }
  }, [emailVerification])

  const handleChange = (index: number, value: string) => {
    if (/^\d*$/.test(value) && value.length <= 1) {
      const newCode = [...code];
      newCode[index] = value;
      setCode(newCode);
      typeCode = newCode.join('');
      setEmailVerification(typeCode);
      if (value && index < code.length - 1) {
        document.getElementById(`__ev_form_input_${index + 1}`)?.focus();
      }
    }

  };

  const handleBackspace = (index: number) => {
    if (index > 0 && code[index] === '') {
      document.getElementById(`__ev_form_input_${index - 1}`)?.focus();
    }
  };

  //impl api/sign_up.ts > usrSignUpVerify
  const handleContinue = () => {
    console.log('EmailVerification');
    let usrInfo = appStore.userInfo
    console.log('appStore:' + appStore.userInfo);

    const username = sessionStorage.getItem('username')
    const email = sessionStorage.getItem('email')
    const password = sessionStorage.getItem('password')

    const usr = {
      ...usrInfo,
      typeCode,
    }
    console.log(usr);
    usrSignUpVerify({
      username,
      email,
      password,
      code_key: sessionStorage.getItem('codeKey'),
      code: parseInt(emailVerification),
    }).then(res => {
      if (res.value.code === 0) {
        const data = res.value.data as { token: string }
        Cookies.set('token', data.token)
        sessionStorage.removeItem('username')
        sessionStorage.removeItem('codeKey')
        sessionStorage.removeItem('email')
        sessionStorage.removeItem('password')

        $message.success(res.value.msg)
        navigate('/')
      } else {
        $message.error(res.value.msg)
      }
    }).catch(err => {
      $message.error(err.response.data.value.msg)
    })

  }

  return (
    <>
      <div className="__ev_form">
        <div className="__ev_title">
          <span>Email verification code</span><span>*</span>
          <span>Please check your email for the verification code and enter it below</span>
          <div className='__ev_form_input'>
            {code.map((value, index) => (
              <input
                key={index}
                id={`__ev_form_input_${index}`}
                type="text"
                value={value}
                maxLength={1}
                onChange={(e) => handleChange(index, e.target.value)}
                onKeyDown={(e) => {
                  if (e.key === 'Backspace') {
                    handleBackspace(index);
                  }
                }}
              />
            ))}
          </div>
        </div>
        <div className="__ev_btnbox">
          <div className="__ev_btnbox_tip">
            <span>Already have an account?</span>
            <span>
              <Link to='/s'>Sign in</Link>
            </span>
          </div>
          <Button
            context='Sign up'
            method={handleContinue}
            disabled={btnAbled} />
          <div className="__ev_btnbox_privacy">
            <span>By using UIPaaS, you are agreeing to the</span>
            <span>
              <a href="/privacy" target='_blank'>Privacy Policy.</a>
            </span>
          </div>
        </div>
      </div>
    </>
  )
}

export default EmailVerif