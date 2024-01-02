import React from "react"
//
import './index.less'
//
import Magnifier from '@assets/layout/magnifier.svg'


interface SearchLineProps {
  placeholder?: string
  outputChange?: (value: string) => void
  searchClick?: () => void
}

const SearchLine: React.FC<SearchLineProps> = (props) => {

  const handleChange = (value: string) => {
    props.outputChange && props.outputChange(value)
  }
  const handleClick = () => {
    props.searchClick && props.searchClick()
  }
  return (
    <>
      <div className="__searchline_wrapper">
        <input type="text" placeholder={props.placeholder} onChange={(e) => handleChange(e.target.value)} />
        <div className="__searchline_wrapper_icon" onClick={handleClick}>
          <img src={Magnifier} />
        </div>
      </div>
    </>
  )
}

export default SearchLine