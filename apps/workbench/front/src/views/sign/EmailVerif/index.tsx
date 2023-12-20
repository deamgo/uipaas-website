import React from 'react'
//style
import './index.less'
//
import Input from '@/components/Input'
import Button from '@/components/Button'
import Mask from '@/components/Mask'
//
import { emailVerificationReg } from '@constants/regexp'
import { usrSignUpVerify } from '@api/sign_up'
import appStore from '@/store/store'
//
import ArrowLeft from '@assets/sign/arrow-left.svg'
import $message from '@/components/Message'
import { Link, useNavigate } from 'react-router-dom'

const EmailVerif: React.FC = () => {

  const [emailVerification, setEmailVerification] = React.useState('')
  const [btnAbled, setBtnAbled] = React.useState(true)

  const navigate = useNavigate()

  React.useEffect(() => {
    if (emailVerificationReg.test(emailVerification)) {
      setBtnAbled(false)
    } else {
      setBtnAbled(true)
    }
  }, [emailVerification])

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
      emailVerification,
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
        sessionStorage.setItem('token', res.value?.data.token)
        sessionStorage.removeItem('username')
        sessionStorage.removeItem('codeKey')
        sessionStorage.removeItem('email')
        sessionStorage.removeItem('password')

        $message.success(res.value.msg)
        navigate('/apps')
      } else {
        $message.error(res.value.msg)
      }
    }).catch(err => {
      $message.error(err.response.data.value.msg)
    })

  }

  return (
    <>
      <div className="__ev_title">
        <span>Email Verification</span>
        <span>Please check your email for the verification code and enter it below</span>
        <div className="__ev_title_row">
          <Link to='/s/up'>
            <img src={ArrowLeft} alt="" />
          </Link>
        </div>
      </div>
      <div className="__ev_form">
        <div className="__ev_form_input">
          <Input
            id='1'
            title='Email verification code'
            type='text'
            placeholder='Enter your email verification code'
            valid='Please enter your email verification code'
            isNeed={true}
            reg={emailVerificationReg}
            outputChange={setEmailVerification} />
        </div>
      </div>
      <div className="__ev_btnbox">
        <div className="__ev_btnbox_tip">
          <span>Already have an account?</span>
          <span>
            <Link to='s/in'>Sign in</Link>
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
    </>
  )
}

export default EmailVerif