import React from 'react'
import { appStore } from '@/store/store'
//
import './index.less'
//
import { usernameReg, emailReg, passwordReg, emailVerificationReg } from '@constants/regexp'
import Input from '@/components/Input'
import { Avatar } from 'antd'
import Button from '@/components/Button'
import Mask from '@/components/Mask'
import Popup from '@/components/Popup'
import { getUserInfo, updateUserName, verifiEmail, verifiEmailCode, verifiPwdEmail, sendPwdEmailCode, verifiPwdEmailCode, updatePwd } from '@/api/developer_profile'
import { IUserInfo } from '@/api/account'
import $message from '@/components/Message'

const UserProfile: React.FC = () => {

  const [userInfo, setUserInfo] = React.useState<IUserInfo | null>(null)

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

  // const loaderData = useLoaderData() as IUserInfo


  //
  // const [change, shouldChange] = React.useState(1)

  React.useEffect(() => {
    (async function get() {
      const { value } = await getUserInfo()
      if (value.code === 0) {
        setUserInfo(value.data as IUserInfo)
        $message.success('获取用户信息成功')
      }
    })()
    return () => {
      console.log('user effect unmount');

    }
  }, [])

  // React.useEffect(() => {
  //   let init = false
  //   if (!init) {
  //     getUserInfo().then(res => {
  //       if (res.code && res.code === 2005) {
  //         $message.warning(res.msg)
  //       }
  //       if (res.value?.code === 0) {
  //         setUserInfo(res.value.data as IUserInfo)
  //       } else {
  //         $message.error(res.value?.msg)
  //       }
  //     }).catch(err => {
  //       console.log(err);
  //       $message.error(err.message)
  //     })
  //     init = true
  //   }
  //   return () => {
  //     init = false
  //   }
  // }, [change])

  React.useEffect(() => {
    document.body.style.overflow = isMask ? 'hidden' : 'auto'
  }, [isMask])

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

  const getDeveloperInfo = async () => {
    try {
      const { value } = await getUserInfo()
      if (value.code === 0) {
        setUserInfo(value.data as IUserInfo)
        $message.success(value.msg)
      }
    } catch (error) {
      console.log(error)
    }
  }

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
    verifiEmailCode({
      code_key: sessionStorage.getItem('code_key') as string,
      code: parseInt(emailCode),
      old_email: userInfo?.email,
      email: newEmail
    }).then(res => {
      if (res.value.code === 0) {
        $message.success(res.value.msg)
        appStore.setUserInfo({
          ...appStore.userInfo,
          email: newEmail
        })
        handleEditEmail()
        sessionStorage.removeItem('code_key')
        // shouldChange(c => -c)
        getDeveloperInfo()
      } else {
        $message.error(res.value.msg)
      }
    }).catch(err => {
      console.log(err);
      $message.error(err.response.data.msg)
    })
  }

  const handleEditPwd = () => {
    handleMask()
    resetForm()
    if (emailTimer) {
      clearInterval(emailTimer)
      setSend('Get')
      setSendAble(false)
    }
    setSendAble(false)
    setIsEditPwd(!isEditPwd)
    setIseditPwdContinue(false)
  }

  const handleEditEmailContinue = () => {
    editEmailContinue()
  }

  const handleEditPwdContinue = () => {
    editPwdContinue()
  }

  const handleConfirmEditPwd = () => {
    console.log('click pwd confirm');
    updatePwd({
      email: userInfo?.email,
      password: newPassword
    }).then(res => {
      if (res.value.code === 0) {
        $message.success(res.value.msg)
        sessionStorage.removeItem('code_key')
        handleEditPwd()
        // shouldChange(c => -c)
        getDeveloperInfo()
      }
    }).catch(err => {
      console.log(err);
      $message.error(err.message)

    })
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
    updateUserName(appStore.getUserInfo().id, {
      username: name
    }).then(res => {
      if (res.value.code === 0) {
        $message.success(res.value.msg)
        // appStore.setUserInfo({
        //   ...appStore.userInfo,
        //   username: name
        // })
        appStore.updateUserInfo()
        // shouldChange(c => -c)
        getDeveloperInfo()

        handleEditName()
      } else {
        $message.error(res.value.msg)
      }
    }).catch(err => {
      $message.error(err.response.data.msg)
    })
  }

  const editEmailContinue = () => {
    console.log('Edit email continue');
    verifiPwdEmail({
      email: userInfo?.email,
      password: currentPassword
    }).then(res => {
      if (res.value.code === 0) {
        setIsEditEmailContinue(true)
      } else {
        $message.error(res.value.msg)
      }
    }).catch(err => {
      console.log(err);
    })
  }

  const editPwdContinue = () => {
    console.log('Edit pwd continue');
    verifiPwdEmailCode({
      code_key: sessionStorage.getItem('code_key') as string,
      code: parseInt(pwdEmailCode),
      email: userInfo?.email
    }).then(res => {
      if (res.value.code === 0) {
        setIseditPwdContinue(true)
        sessionStorage.removeItem('code_key')
      }
    }).catch(err => {
      console.log(err)
      $message.error(err.message)
    })

  }

  const sendVerifiCode = () => {
    console.log('Send verifi code');
    setSendAble(true)
    getTimer()
    verifiEmail({
      old_email: userInfo?.email,
      email: newEmail
    }).then(res => {
      if (res.value.code === 0) {
        const data = res.value.data as { code_key: string }
        sessionStorage.setItem('code_key', data.code_key)
      } else {
        $message.error(res.value.msg)
      }
    }).catch(err => {
      $message.error(err.response.data.msg)
    })
    console.log('send over');

  }

  const sendPwdVerifiCode = () => {
    console.log('Send pwd verifi code');
    setSendAble(true)
    getTimer()
    sendPwdEmailCode({
      email: userInfo?.email
    }).then(res => {
      if (res.value.code === 0) {
        const data = res.value.data as { code_key: string }
        sessionStorage.setItem('code_key', data.code_key)
      } else {
        $message.error(res.value.msg)
      }
    }).catch(err => {
      console.log(err);
      $message.error(err.message)
    })

    console.log('send over');
  }



  return (
    <>
      <div className="__user_profile_account">
        <div className="__user_profile_account_title">
          <span>Account</span>
        </div>
        <div className="__user_profile_account_container">
          <div className="_sp_withAvatar">
            <Avatar style={{ backgroundColor: '#4080FF', verticalAlign: 'middle' }} size={48} gap={3}>
              {userInfo?.username.charAt(0).toUpperCase()}
            </Avatar>
          </div>
          <div className="__user_profile_account_container_wrapper">
            <div className="__user_profile_account_container_wrapper_input">
              <Input
                id='1'
                title='Name'
                type='text'
                isNeed={false}
                typeAble={true}
                value={userInfo?.username} />
            </div>
            <div className="__user_profile_account_container_wrapper_edit">
              <Button context='Edit' type='board-primary' method={handleEditName} >Edit</Button>
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
                value={userInfo?.email} />
            </div>
            <div className="__user_profile_account_container_wrapper_edit">
              <Button context='Edit' type='board-primary' method={handleEditEmail} >Edit</Button>
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
              <Button context='Edit' type='board-primary' method={handleEditPwd} >Edit</Button>
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
            <Button context='Save' type='primary' disabled={nameEditAble} method={editName} >
              Save
            </Button>
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
                    <span>{userInfo?.email}</span>
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
                <Button context='Continue' type='primary' disabled={emailEditAble} method={handleEditEmailContinue} >
                  Continue
                </Button>
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
                      <Button context={send} type='primary' disabled={sendAble} method={sendVerifiCode} >
                        {send}
                      </Button>
                    </div>
                  </div>
                </div>
              </div>
              <div className="_edit_email_btn">
                <Button context='Confirm' type='primary' disabled={emailEditConfirmAble} method={handleConfirmEditEmail} >
                  Confirm
                </Button>
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
                    <span>{userInfo?.email}</span>
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
                      <Button context={send} type='primary' disabled={sendAble} method={sendPwdVerifiCode} >
                        {send}
                      </Button>
                    </div>
                  </div>
                </div>
              </div>
              <div className="_edit_email_btn">
                <Button context='Continue' type='primary' disabled={pwdEditContinueAble} method={handleEditPwdContinue} >
                  Continue
                </Button>
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
                <Button context='Confirm' type='primary' disabled={PwdEditConfirmAble} method={handleConfirmEditPwd} >
                  Confirm
                </Button>
              </div>
            </>
          )}
        </Popup>
      }
    </>
  )
}

export default UserProfile