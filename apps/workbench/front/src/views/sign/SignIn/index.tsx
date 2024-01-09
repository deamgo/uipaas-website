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
import { Link, useNavigate } from 'react-router-dom'
import { tokenStore } from '@/store/store'

const SignIn: React.FC = () => {

  const [email, setEmail] = React.useState('')
  const [pwd, setPwd] = React.useState('')
  const [btnAbled, setBtnAbled] = React.useState(true)

  const navigate = useNavigate()

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
  const handleContinue = async () => {
    console.log('SignIn');
    await usrSignIn({
      email,
      password: pwd
    }).then(res => {
      if (res.value.code === 0) {
        const data = res.value?.data as { token: string }
        tokenStore.setToken(data.token)
        if (!tokenStore.getToken()) {
          tokenStore.setToken(data.token)
        }
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
      <div className="__signin_form">
        <div className="__signin_form_input">
          <Input
            id='signin-email'
            title='Email'
            type='text'
            placeholder='Enter your email address'
            valid='Please enter the email address'
            isNeed={true}
            reg={emailReg}
            outputChange={setEmail} />
        </div>
        <div className="__signin_form_input">
          <Input
            id='signin-pwd'
            title='Password'
            type='password'
            placeholder='Enter your password'
            valid='8+ characters(a-z,A-z,0-9)'
            isNeed={true}
            isShowPwd={true}
            reg={passwordReg}
            outputChange={setPwd} />
          <div className="__signin_form_input_forgot">
            <Link to='/s/ryp'>Forgot?</Link>
          </div>
        </div>
        <div className="__signin_form_continue">
          <div className="__signin_form_continue_tip">
            <span>New to UIPaaS?</span>
            <span>
              <Link to='/s/up'>Sign up</Link>
            </span>
          </div>
          <Button
            context='Sign in'
            type='primary'
            method={handleContinue}
            disabled={btnAbled}
            ys={{
              width: '100%'
            }}>
            Sign in
          </Button>
          <div className="__signin_form_continue_privacy">
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

export default SignIn