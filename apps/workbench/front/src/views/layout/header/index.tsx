import React from 'react';
//style
import './index.less'
//
import Button from '@/components/Button';
import SearchLine from '@/components/SearchLine';


const Header: React.FC = () => {

  const handleCreate = () => {
    console.log('Create');

  }
  return (
    <>
      <div className="__header">
        <div className="__header_toolg">
          <div className="__header_toolg_btn">
            <Button
              context='Create'
              method={handleCreate} />
          </div>
          <div className="__header_toolg_searchline">
            <SearchLine placeholder='Search' />
          </div>
        </div>
      </div>
    </>
  )
}

export default Header