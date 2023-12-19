import React from 'react'
//style
import './index.less'
//
import Input from '@/components/Input'
import Button from '@/components/Button'
import Mask from '@/components/Mask'
//
import { emailReg, emailVerificationReg, passwordReg } from '@constants/regexp'
import { ResetPwd, forgotVerify } from '@api/reset_pwd'
//
import ArrowLeft from '@assets/sign/arrow-left.svg'
import $message from '@/components/Message'

const PwdReset: React.FC = () => {

  const [email, setEmail] = React.useState('')
  const [emailVerification, setEmailVerification] = React.useState('')
  const [pwd, setPwd] = React.useState('')
  const [btnAbled, setBtnAbled] = React.useState(true)
  const [bsendAbled, setBsendAbled] = React.useState(true)
  const [sendText, setSendText] = React.useState('Send')


  React.useEffect(() => {
    if (emailReg.test(email)
      && emailVerificationReg.test(emailVerification)
      && passwordReg.test(pwd)) {
      setBtnAbled(false)
    } else {
      setBtnAbled(true)
    }

    if (emailReg.test(email)) {
      setBsendAbled(false)
    } else {
      setBsendAbled(true)
    }
  }, [email, emailVerification, pwd])

  //impl api/reset_pwd.ts > ResetPwd
  const handleContinue = () => {
    console.log('PwdReset');
    const usr = {
      email,
      emailVerification,
      pwd,
    }
    console.log(usr);
    ResetPwd({
      email,
      code_key: sessionStorage.getItem('codeKey'),
      code: parseInt(emailVerification),
      password: pwd
    }).then(res => {
      console.log(res);

      if (res.value.code === 0) {
        $message.success(res.value.msg)
        window.location.href = '/s/in'
      } else {
        $message.error(res.value.msg)
      }
    }).catch(err => {
      console.log(err)

      $message.error(err.response.data.value.msg)
    })
  }

  //impl api/sign_in.ts > forgotVerify
  const handleSend = () => {
    console.log('Send');
    setBsendAbled(true)
    setSendText('Sent')
    setTimeout(() => {
      setBsendAbled(false)
      setSendText('Send')
    }, 10000)
    forgotVerify({
      email
    }).then(res => {
      if (res.value.code === 0) {
        $message.success(res.value.msg)
        sessionStorage.setItem('codeKey', res.value.data.code_key)
      } else {
        $message.error(res.value.msg)
      }
    }).catch(err => {
      $message.error(err.response.data.value.msg)
    })

  }

  return (
    <>
      <div className="__ryp_title">
        <span>Reset your Password</span>
        <div className="__ryp_title_row">
          <a href="/s/in">
            <img src={ArrowLeft} alt="" />
          </a>
        </div>
      </div>
      <div className="__ryp_form">
        <div className="__ryp_form_input">
          <Input
            id='1'
            title='Email'
            type='text'
            placeholder='Enter your email address'
            valid='Please enter the email address'
            reg={emailReg}
            outputChange={setEmail} />
        </div>
        <div className="__ryp_form_input">
          <Input
            id='2'
            title='Email verification code'
            type='text'
            placeholder='Enter your email verification code'
            valid='Please enter the email verification code'
            reg={emailVerificationReg}
            outputChange={setEmailVerification} />
          <div className="__ryp_form_input_send">
            <Button context={sendText} method={handleSend} disabled={bsendAbled} />
          </div>
        </div>
        <div className="__ryp_form_input">
          <Input
            id='3'
            title='Password'
            type='password'
            placeholder='Enter your password'
            valid='Please enter the password'
            reg={passwordReg}
            outputChange={setPwd} />
        </div>
      </div>
      <div className="__ryp_continue">
        <Button
          context='Confirm'
          method={handleContinue}
          disabled={btnAbled} />
      </div>

    </>
  )
}

export default PwdReset