import React from 'react';
//style
import './index.less'
//
import Button from '@/components/Button';
import SearchLine from '@/components/SearchLine';
import { Link } from 'react-router-dom';


const Header: React.FC = () => {

  const handleCreate = () => {
    console.log('Create');
  }

  return (
    <>
      <div className="__header">
        <ul className="__header_nav">
          <li className="__header_nav_item">
            <Link to={'/u/profile'}>Profile</Link>
          </li>
          {/* <li className="__header_nav_item">
            <Link to={'/u/invite'}>Invite</Link>
          </li> */}
        </ul>
      </div>
    </>
  )
}

export default Header