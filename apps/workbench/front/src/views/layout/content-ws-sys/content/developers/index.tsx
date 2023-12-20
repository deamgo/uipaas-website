import React from 'react'
//
import './index.less'
//
import { usernameReg, emailReg, passwordReg } from '@constants/regexp'
import Input from '@/components/Input'
import { Avatar } from 'antd'
import Button from '@/components/Button'
import SearchLine from '@/components/SearchLine'


const WSDevelopers: React.FC = () => {

  const [name, setName] = React.useState('')
  const [email, setEmail] = React.useState('')
  const [password, setPassword] = React.useState('')

  return (
    <>
      <div className="__workspace_developers">
        <div className="__workspace_developers_tools">
          <div className="__workspace_developers_addbtn">
            <Button context='Add developers' type='primary' />
          </div>
          <div className="__workspace_developers_searchline">
            <SearchLine />
          </div>
        </div>
        <div className="__workspace_developers_container">

        </div>
      </div>
    </>
  )
}

export default WSDevelopers