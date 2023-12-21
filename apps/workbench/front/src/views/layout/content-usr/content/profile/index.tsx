import React from 'react'
//
import './index.less'
//
import { usernameReg, emailReg, passwordReg, emailVerificationReg } from '@constants/regexp'
import Input from '@/components/Input'
import { Avatar } from 'antd'
import Button from '@/components/Button'
import Mask from '@/components/Mask'
import Popup from '@/components/Popup'
import { set } from 'mobx'


const UserProfile: React.FC = () => {

  const [name, setName] = React.useState('')
  const [currentPassword, setCurrentPassword] = React.useState('')
  const [newEmail, setNewEmail] = React.useState('')
  const [emailCode, setEmailCode] = React.useState('')
  const [pwdEmailCode, setPwdEmailCode] = React.useState('')
  const [newPassword, setNewPassword] = React.useState('')

  const [emailTimer, setEmailTimer] = React.useState<NodeJS.Timeout | null>(null)
  //
  const [isMask, setIsMask] = React.useState(false)
  const [isEditName, setIsEditName] = React.useState(false)
  const [isEditEmail, setIsEditEmail] = React.useState(false)
  const [iseditEmailContinue, setIsEditEmailContinue] = React.useState(false)
  const [isEditPwd, setIsEditPwd] = React.useState(false)

  const [iseditPwdContinue, setIseditPwdContinue] = React.useState(false)
  //
  const [nameEditAble, setNameEditAble] = React.useState(true)
  const [emailEditAble, setEmailEditAble] = React.useState(true)
  const [emailEditConfirmAble, setEmailEditConfirmAble] = React.useState(true)
  const [sendAble, setSendAble] = React.useState(false)
  const [pwdEditContinueAble, setPwdEditContinueAble] = React.useState(true)
  const [PwdEditConfirmAble, setPwdEditConfirmAble] = React.useState(true)
  const [send, setSend] = React.useState('Get')

  React.useEffect(() => {
    if (usernameReg.test(name)) {
      setNameEditAble(false)
    } else {
      setNameEditAble(true)
    }

    if (passwordReg.test(currentPassword)) {
      setEmailEditAble(false)
    } else {
      setEmailEditAble(true)
    }

    if (emailReg.test(newEmail)) {
      setSendAble(false)
    } else {
      setSendAble(true)
    }

    if (emailVerificationReg.test(emailCode)) {
      setEmailEditConfirmAble(false)
    } else {
      setEmailEditConfirmAble(true)
    }

    if (emailVerificationReg.test(pwdEmailCode)) {
      setPwdEditContinueAble(false)
    } else {
      setPwdEditContinueAble(true)
    }

    if (passwordReg.test(newPassword)) {
      setPwdEditConfirmAble(false)
    } else {
      setPwdEditConfirmAble(true)
    }
  }, [name, currentPassword, newEmail, emailCode, pwdEmailCode, newPassword])

  const resetForm = () => {
    setName('')
    setCurrentPassword('')
    setNewEmail('')
    setEmailCode('')
    setPwdEmailCode('')
    setNewPassword('')
  }

  const handleMask = () => {
    setIsMask(!isMask)
  }

  const handleEditName = () => {
    handleMask()
    resetForm()
    setNameEditAble(true)
    setIsEditName(!isEditName)
  }

  const handleEditEmail = () => {
    handleMask()
    resetForm()
    if (emailTimer) {
      clearInterval(emailTimer)
      setSend('Get')
    }
    setEmailEditAble(true)
    setSendAble(true)
    setIsEditEmail(!isEditEmail)
    setIsEditEmailContinue(false)
  }

  const handleConfirmEditEmail = () => {
    console.log('click email confirm');

  }

  const handleEditPwd = () => {
    handleMask()
    resetForm()
    if (emailTimer) {
      clearInterval(emailTimer)
      setSend('Get')
      setSendAble(false)
    }
    setIsEditPwd(!isEditPwd)
    setIseditPwdContinue(false)
  }

  const handleEditEmailContinue = () => {
    editEmailContinue()
    setIsEditEmailContinue(true)
  }

  const handleEditPwdContinue = () => {
    editPwdContinue()
    setIseditPwdContinue(true)
  }

  const handleConfirmEditPwd = () => {
    console.log('click pwd confirm');
  }

  const getTimer = () => {
    let count = 60
    const timer = setInterval(() => {
      setSend('Got(' + count + 's)')
      if (count === -1) {
        clearInterval(timer)
        setSend('Get')
        setSendAble(false)
      }
      count--
    }, 1000)
    setEmailTimer(timer)
  }


  const editName = () => {
    console.log('Edit name');

  }

  const editEmailContinue = () => {
    console.log('Edit email continue');

  }

  const editPwdContinue = () => {
    console.log('Edit pwd continue');
  }

  const sendVerifiCode = () => {
    console.log('Send verifi code');
    setSendAble(true)
    getTimer()
    console.log('send over');

  }

  const sendPwdVerifiCode = () => {
    console.log('Send pwd verifi code');
    setSendAble(true)
    getTimer()
    console.log('send over');
  }

  const editEmial = () => {
    console.log('Edit email');

  }

  const editPwd = () => {
    console.log('Edit name');

  }



  return (
    <>
      <div className="__user_profile_account">
        <div className="__user_profile_account_title">
          <span>Account</span>
        </div>
        <div className="__user_profile_account_container">
          <div className="__user_profile_account_container_wrapper">
            <div className="__user_profile_account_container_wrapper_input _sp_withAvatar">
              <Avatar style={{ backgroundColor: '#4080FF', verticalAlign: 'middle' }} size={48} gap={3}>
                {'ILEE'.charAt(0).toUpperCase()}
              </Avatar>
              <div className="__user_profile_account_container_wrapper_input_besideAvatar">
                <Input
                  id='1'
                  title='Name'
                  type='text'
                  isNeed={false}
                  typeAble={true}
                  value={'ilee'} />
              </div>
            </div>
            <div className="__user_profile_account_container_wrapper_edit">
              <Button context='Edit' type='board-primary' method={handleEditName} />
            </div>
          </div>

          <div className="__user_profile_account_container_wrapper">
            <div className="__user_profile_account_container_wrapper_input">
              <Input
                id='2'
                title='Email'
                type='text'
                isNeed={false}
                typeAble={true}
                value={'123@qq.com'} />
            </div>
            <div className="__user_profile_account_container_wrapper_edit">
              <Button context='Edit' type='board-primary' method={handleEditEmail} />
            </div>
          </div>

          <div className="__user_profile_account_container_wrapper">
            <div className="__user_profile_account_container_wrapper_input">
              <Input
                id='3'
                title='Password'
                type='text'
                isNeed={false}
                typeAble={true}
                value={'......'} />
            </div>
            <div className="__user_profile_account_container_wrapper_edit">
              <Button context='Edit' type='board-primary' method={handleEditPwd} />
            </div>
          </div>
        </div>
      </div>
      {isMask && <Mask />}

      {isEditName &&
        <Popup unit='rem' width={496} height={276} title='Edit name' onClose={handleEditName} >
          <div className="_edit_name_input">
            <Input
              id='edit_name'
              type='text'
              outputChange={setName} />
          </div>
          <div className="_edit_name_btn">
            <Button context='Save' type='primary' disabled={nameEditAble} method={editName} />
          </div>
        </Popup>
      }

      {isEditEmail &&
        <Popup unit='rem' width={496} height={329} title='Edit email' onClose={handleEditEmail} >
          {!iseditEmailContinue && (
            <>
              <div className="_current">
                <div className="_current_wrapper">
                  <div className="_current_wrapper_title">
                    <span>Current email</span>
                  </div>
                  <div className="_current_wrapper_content">
                    <span>{'123@qq.com'}</span>
                  </div>
                </div>
                <div className="_current_wrapper">
                  <div className="_current_wrapper_title">
                    <span>Current password</span>
                  </div>
                  <div className="_current_wrapper_content">
                    <Input
                      id='repeat_pwd'
                      type='password'
                      outputChange={setCurrentPassword}
                      reg={passwordReg} />
                  </div>
                </div>
              </div>
              <div className="_edit_email_btn">
                <Button context='Continue' type='primary' disabled={emailEditAble} method={handleEditEmailContinue} />
              </div>
            </>
          )}
          {iseditEmailContinue && (
            <>
              <div className="_current">
                <div className="_current_wrapper">
                  <div className="_current_wrapper_title">
                    <span>New email</span>
                  </div>
                  <div className="_current_wrapper_content">
                    <Input
                      id='edit_email'
                      type='text'
                      outputChange={setNewEmail}
                      reg={emailReg} />
                  </div>
                </div>
                <div className="_current_wrapper">
                  <div className="_current_wrapper_title">
                    <span>Email verification code</span>
                  </div>
                  <div className="_current_wrapper_content">
                    <Input
                      id='edit_emailVerification'
                      type='text'
                      outputChange={setEmailCode}
                      reg={emailVerificationReg} />
                    <div className="_current_wrapper_content_send">
                      <Button context={send} type='primary' disabled={sendAble} method={sendVerifiCode} />
                    </div>
                  </div>
                </div>
              </div>
              <div className="_edit_email_btn">
                <Button context='Confirm' type='primary' disabled={emailEditConfirmAble} method={handleConfirmEditEmail} />
              </div>
            </>
          )}
        </Popup>
      }

      {isEditPwd &&
        <Popup unit='rem' width={496} height={329} title='Edit password' onClose={handleEditPwd} >
          {!iseditPwdContinue && (
            <>
              <div className="_current">
                <div className="_current_wrapper">
                  <div className="_current_wrapper_title">
                    <span>Current email</span>
                  </div>
                  <div className="_current_wrapper_content">
                    <span>{'123@qq.com'}</span>
                  </div>
                </div>
                <div className="_current_wrapper">
                  <div className="_current_wrapper_title">
                    <span>Email verification code</span>
                  </div>
                  <div className="_current_wrapper_content">
                    <Input
                      id='repeat_pwd'
                      type='text'
                      outputChange={setPwdEmailCode}
                      reg={emailVerificationReg} />
                    <div className="_current_wrapper_content_send">
                      <Button context={send} type='primary' disabled={sendAble} method={sendPwdVerifiCode} />
                    </div>
                  </div>
                </div>
              </div>
              <div className="_edit_email_btn">
                <Button context='Continue' type='primary' disabled={pwdEditContinueAble} method={handleEditPwdContinue} />
              </div>
            </>
          )}
          {iseditPwdContinue && (
            <>
              <div className="_current">
                <div className="_current_wrapper">
                  <div className="_current_wrapper_title">
                    <span>New password</span>
                  </div>
                  <div className="_current_wrapper_content">
                    <Input
                      id='edit_pwd'
                      type='password'
                      outputChange={setNewPassword}
                      reg={passwordReg}
                      isShowPwd={true} />
                  </div>
                </div>
              </div>
              <div className="_edit_email_btn">
                <Button context='Confirm' type='primary' disabled={PwdEditConfirmAble} method={handleConfirmEditPwd} />
              </div>
            </>
          )}
        </Popup>
      }
    </>
  )
}

export default UserProfile