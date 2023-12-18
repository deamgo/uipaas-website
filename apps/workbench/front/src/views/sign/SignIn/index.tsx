import React from 'react'
//style
import './index.less'
//
import Input from '@/components/Input'
import Button from '@/components/Button'
//
import { emailReg, passwordReg } from '@constants/regexp'
import { usrSignIn } from '@api/sign_in'
import $message from '@/components/Message'

const SignIn: React.FC = () => {

  const [email, setEmail] = React.useState('')
  const [pwd, setPwd] = React.useState('')
  const [btnAbled, setBtnAbled] = React.useState(true)

  React.useEffect(() => {
    if (emailReg.test(email)
      && passwordReg.test(pwd)) {
      setBtnAbled(false)
    } else {
      setBtnAbled(true)
    }
  }, [email, pwd])

  // const validator = (value: string, regex: RegExp) => {
  //   return regex.test(value)
  // }
  //impl api/sign_in.ts > usrSignIn
  const handleContinue = () => {
    console.log('SignIn');
    const usrinfo = {
      email,
      pwd
    }
    console.log(usrinfo);
    usrSignIn({
      email,
      password: pwd
    }).then(res => {
      if (res.value.code === 0) {
        sessionStorage.setItem('token', res.value?.data.token)
        $message.success(res.value.msg)
        window.location.href = '/apps'
      } else {
        $message.error(res.value.msg)
      }
    }).catch(err => {
      console.log(err);
      $message.error(err.response.data.value.msg)
    })
  }

  return (
    <>
      <div className="__signin_title">
        <span>Sign in</span>
      </div>
      <div className="__signin_form">
        <div className="__signin_form_input">
          <Input
            id='1'
            title='Email'
            type='text'
            placeholder='Enter your email address'
            valid='Please enter the email address'
            reg={emailReg}
            outputChange={setEmail} />
        </div>
        <div className="__signin_form_input">
          <Input
            id='2'
            title='Password'
            type='password'
            placeholder='Enter your password'
            valid='Please enter the password'
            reg={passwordReg}
            outputChange={setPwd} />
          <div className="__signin_form_input_forgot">
            <a href="/s/ryp">Forgot?</a>
          </div>
        </div>
      </div>
      <div className="__signin_continue">
        <div className="__signin_continue_tip">
          <span>New to UIPaaS?</span>
          <span>
            <a href="/s/up">Sign up</a>
          </span>
        </div>
        <Button
          context='Sign in'
          method={handleContinue}
          disabled={btnAbled} />
        <div className="__signin_continue_privacy">
          <span>By using UIPaaS, you are agreeing to the</span>
          <span>
            <a href="/privacy" target='_blank'>Privacy Policy.</a>
          </span>
        </div>
      </div>
    </>
  )
}

export default SignIn