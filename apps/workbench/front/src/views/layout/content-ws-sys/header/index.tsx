import React from 'react';
//style
import './index.less'
//
import { Link } from 'react-router-dom';


const Header: React.FC = () => {

  return (
    <>
      <div className="__header">
        <ul className="__header_nav">
          <li className="__header_nav_item">
            <Link to={'/workspace'}>Developers</Link>
          </li>
          <li className="__header_nav_item">
            <Link to={'/workspace/settings'}>Settings</Link>
          </li>
        </ul>
      </div>
    </>
  )
}

export default Header