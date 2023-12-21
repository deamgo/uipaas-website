import React from 'react'
//
import './index.less'
import Input from '@/components/Input'
import { Avatar } from 'antd'
import Button from '@/components/Button'
import SearchLine from '@/components/SearchLine'


const WSDevelopers: React.FC = () => {

  return (
    <>
      <div className="__workspace_developers">
        <div className="__workspace_developers_tools">
          <div className="__workspace_developers_tools_addbtn">
            <Button context='Add developers' type='primary' />
          </div>
          <div className="__workspace_developers_tools_searchline">
            <SearchLine placeholder='Search' />
          </div>
        </div>
        <div className="__workspace_developers_container">

        </div>
        <div className="__workspace_developers_pagenation">

        </div>
      </div>
    </>
  )
}

export default WSDevelopers