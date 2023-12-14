import React from 'react'
//style
import './index.less'
//
import Input from '@/components/Input'
import Button from '@/components/Button'
//
import { codeReg, usernameReg, emailReg, passwordReg } from '@constants/regexp.ts'
import $message from '@/components/Message'
import { usrSignUp } from '@api/sign_up'

type IUsrAccount = {
  invitation_code?: string
  username?: string
  email: string
  password: string
}

const SignUp: React.FC = () => {

  const [code, setCode] = React.useState('')
  const [name, setName] = React.useState('')
  const [email, setEmail] = React.useState('')
  const [pwd, setPwd] = React.useState('')
  const [btnAbled, setBtnAbled] = React.useState(true)

  React.useEffect(() => {
    if (codeReg.test(code)
      && usernameReg.test(name)
      && emailReg.test(email)
      && passwordReg.test(pwd)) {
      setBtnAbled(false)
    } else {
      setBtnAbled(true)
    }
  }, [code, name, email, pwd])

  // const validator = (value: string, regex: RegExp) => {
  //   return regex.test(value)
  // }


  //impl api/sign_up.tx > usrSignUp
  const handleContinue = () => {
    console.log('SignUp');
    const usr: IUsrAccount = {
      invitation_code: code,
      username: name,
      email,
      password: pwd,
    }
    usrSignUp(usr).then(res => {
      if (res.code === 0) {
        window.location.pathname = '/s/ev'
      }
    }).catch(err => {
      console.log(err);

    })
  }

  return (
    <>
      <div className="__sign_title">
        <span>Sign up</span>
        <span>Complete the Information for Account Registration</span>
      </div>
      <div className="__sign_form">
        <div className="__sign_form_input">
          <Input
            id='1'
            title='Invitation code'
            type='text'
            placeholder='Enter your invitation code'
            valid='Please enter the invitation code'
            outputChange={setCode}
            reg={codeReg} />
        </div>
        <div className="__sign_form_input">
          <Input
            id='2'
            title='Name'
            type='text'
            placeholder='Enter your name'
            valid='Please enter your name'
            outputChange={setName}
            reg={usernameReg} />
        </div>
        <div className="__sign_form_input">
          <Input
            id='3'
            title='Email'
            type='text'
            placeholder='Enter your email address'
            valid='Please enter your email address'
            outputChange={setEmail}
            reg={emailReg} />
        </div>
        <div className="__sign_form_input">
          <Input
            id='4'
            title='Password'
            type='password'
            placeholder='Enter new password'
            valid='Please enter your password'
            outputChange={setPwd}
            reg={passwordReg} />
        </div>
      </div>
      <div className="__sign_continue">
        <div className="__sign_continue_tip">
          <span>Already have an account?</span>
          <span>
            <a href="/s/in">Sign in</a>
          </span>
        </div>
        <Button
          context='Continue'
          method={handleContinue}
          disabled={btnAbled} />
        <div className="__sign_continue_privacy">
          <span>By using UIPaaS, you are agreeing to the</span>
          <span>
            <a href="/privacy" target='_blank'>Privacy Policy.</a>
          </span>
        </div>
      </div>
    </>
  )
}

export default SignUp