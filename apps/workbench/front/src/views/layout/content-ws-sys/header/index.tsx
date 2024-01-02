import React from 'react';
//style
import './index.less'
//
import Navbar from '@/components/Navbar';


const list_nav = [
  {
    name: 'Developers',
    path: '/workspace',
    matcher: 'workspace',
    index: 2
  },
  {
    name: 'Settings',
    path: '/workspace/settings',
    matcher: 'settings',
    index: 2
  },
]


const Header: React.FC = () => {

  return (
    <>
      <div className="__header">
        <Navbar items={list_nav} />
      </div>
    </>
  )
}

export default Header