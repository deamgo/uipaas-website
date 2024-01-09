import React from 'react';
//style
import './index.less'
//
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