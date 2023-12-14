import React from 'react'
//style
import './index.less'
//
import Input from '@/components/Input'
import Button from '@/components/Button'
//
import { emailReg, emailVerificationReg, passwordReg } from '@constants/regexp'
//
import ArrowLeft from '@assets/sign/arrow-left.svg'


const PwdReset: React.FC = () => {

  const [email, setEmail] = React.useState('')
  const [emailVerification, setEmailVerification] = React.useState('')
  const [pwd, setPwd] = React.useState('')
  const [btnAbled, setBtnAbled] = React.useState(true)

  React.useEffect(() => {
    if (emailReg.test(email)
      && emailVerificationReg.test(emailVerification)
      && passwordReg.test(pwd)) {
      setBtnAbled(false)
    } else {
      setBtnAbled(true)
    }
  }, [email, emailVerification, pwd])

  const handleContinue = () => {
    console.log('Continue');
    const usr = {
      email,
      emailVerification,
      pwd,
    }
    console.log(usr);

  }
  const handleSend = () => {
    console.log('Send');

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
            <Button context='Send' method={handleSend} />
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