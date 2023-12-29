import React from 'react';
//style
import './index.less'
//
import Button from '@/components/Button';
import SearchLine from '@/components/SearchLine';
import { Link } from 'react-router-dom';
import Navbar from '@/components/Navbar';


const list_nav = [
  {
    name: 'Profile',
    path: '/u',
    matcher: 'u',
    index: 2
  }
]

const Header: React.FC = () => {

  const handleCreate = () => {
    console.log('Create');
  }

  return (
    <>
      <div className="__usrheader">
        <Navbar items={list_nav} />
        {/* <ul className="__usrheader_nav">
          <li className="__usrheader_nav_item">
            <Link to={'/u'}>Profile</Link>
          </li>
        </ul> */}
      </div>
    </>
  )
}

export default Header