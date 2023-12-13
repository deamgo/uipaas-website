import React from "react"
//
import './index.less'
//
import Magnifier from '@assets/layout/magnifier.svg'


interface SearchLineProps {
  placeholder?: string
}

const SearchLine: React.FC<SearchLineProps> = (props) => {
  return (
    <>
      <div className="__searchline_wrapper">
        <input type="text" placeholder={props.placeholder} />
        <div className="__searchline_wrapper_icon">
          <img src={Magnifier} />
        </div>
      </div>
    </>
  )
}

export default SearchLine