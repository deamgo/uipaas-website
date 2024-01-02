import React from 'react'
//style
import './index.less'
//
import Input from '@/components/Input'
import Button from '@/components/Button'
//
import { emailReg, emailVerificationReg, passwordReg } from '@constants/regexp'
import { ResetPwd, forgotVerify } from '@api/reset_pwd'
//
import $message from '@/components/Message'
import { useNavigate } from 'react-router-dom'

const PwdReset: React.FC = () => {

  const [email, setEmail] = React.useState('')
  const [emailVerification, setEmailVerification] = React.useState('')
  const [pwd, setPwd] = React.useState('')
  const [btnAbled, setBtnAbled] = React.useState(true)
  const [bsendAbled, setBsendAbled] = React.useState(true)
  const [sendText, setSendText] = React.useState('Get')

  const navigate = useNavigate()


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
        navigate('/s')
      } else {
        $message.error(res.value.msg)
      }
    }).catch(err => {
      console.log(err)

      $message.error(err.response.data.value.msg)
    })
  }

  let timer: NodeJS.Timeout

  //impl api/sign_in.ts > forgotVerify
  const handleSend = () => {
    console.log('Send');
    setBsendAbled(true)
    let count = 60
    timer = setInterval(() => {
      setSendText('Got(' + count + 's)')
      if (count === -1) {
        clearInterval(timer)
        setSendText('Get')
        setBsendAbled(false)
      }
      count--
    }, 1000)

    forgotVerify({
      email
    }).then(res => {
      if (res.value.code === 0) {
        $message.success(res.value.msg)
        const data = res.value.data as { code_key: string }
        sessionStorage.setItem('codeKey', data.code_key)
      } else {
        $message.error(res.value.msg)
      }
    }).catch(err => {
      $message.error(err.response.data.value.msg)
    })

  }

  return (
    <>
      <div className="__ryp_form">
        <div className="__ryp_title">
          <div className="__ryp_title_divider"></div>
          <span>Reset your Password</span>
        </div>
        <div className="__ryp_form_input">
          <Input
            id='1'
            title='Email'
            type='text'
            placeholder='Enter your email address'
            valid='Please enter the email address'
            isNeed={true}
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
            isNeed={true}
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
            isNeed={true}
            isShowPwd={true}
            reg={passwordReg}
            outputChange={setPwd} />
        </div>
        <div className="__ryp_form_continue">
          <Button
            context='Confirm'
            method={handleContinue}
            disabled={btnAbled} />
        </div>
      </div>

    </>
  )
}

export default PwdReset