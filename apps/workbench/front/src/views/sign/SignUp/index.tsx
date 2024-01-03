import React from 'react'
import { useNavigate } from 'react-router-dom'
//style
import './index.less'
//
import Input from '@/components/Input'
import Button from '@/components/Button'
//
import { usernameReg, emailReg, passwordReg } from '@constants/regexp.ts'
import $message from '@/components/Message'
import { usrSignUp } from '@api/sign_up'
import { IUsrAccount } from '@api/account'
import { Link } from 'react-router-dom'


const SignUp: React.FC = () => {

  // const [code, setCode] = React.useState('')
  const [name, setName] = React.useState('')
  const [email, setEmail] = React.useState('')
  const [pwd, setPwd] = React.useState('')
  const [btnAbled, setBtnAbled] = React.useState(true)

  const navigate = useNavigate()

  React.useEffect(() => {
    if (usernameReg.test(name)
      && emailReg.test(email)
      && passwordReg.test(pwd)) {
      setBtnAbled(false)
    } else {
      setBtnAbled(true)
    }
  }, [name, email, pwd])

  // const validator = (value: string, regex: RegExp) => {
  //   return regex.test(value)
  // }


  //impl api/sign_up.tx > usrSignUp
  const handleContinue = () => {
    console.log('SignUp');
    const usr: IUsrAccount = {
      // invitation_code: code,
      username: name,
      email,
      password: pwd,
    }
    sessionStorage.setItem('username', name)
    sessionStorage.setItem('email', email)
    sessionStorage.setItem('password', pwd)

    usrSignUp(usr).then(res => {
      if (res.value.code === 0) {
        const data = res.value.data as { code_key: string }
        sessionStorage.setItem('codeKey', data.code_key)
        navigate('/s/ev')
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
      <div className="__sign_form">
        {/* <div className="__sign_form_input">
          <Input
            id='1'
            title='Invitation code'
            type='text'
            placeholder='Enter your invitation code'
            valid='Please enter the invitation code'
            outputChange={setCode}
            reg={codeReg} />
        </div> */}
        <div className="__sign_form_input">
          <Input
            id='2'
            title='Name'
            type='text'
            placeholder='Enter your name'
            valid='Please enter your name'
            isNeed={true}
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
            isNeed={true}
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
            isNeed={true}
            isShowPwd={true}
            outputChange={setPwd}
            reg={passwordReg} />
        </div>
        <div className="__sign_form_continue">
          <div className="__sign_form_continue_tip">
            <span>Already have an account?</span>
            <span>
              <Link to='/s'>Sign in</Link>
            </span>
          </div>
          <Button
            context='Continue'
            method={handleContinue}
            disabled={btnAbled} >
            Continue
          </Button>
          <div className="__sign_form_continue_privacy">
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

export default SignUp