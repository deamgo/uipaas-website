import React from "react"
//
import './index.less'
import { Link, useLocation } from "react-router-dom";
import { mcontent } from "@/interface/some";

const MenuContent: React.FC<mcontent> = (props) => {
  const [active, setActive] = React.useState(false)
  const location = useLocation()

  React.useEffect(() => {
    const url = location.pathname
    let urlArr = url.split('/')
    if (props.path && props.index && urlArr[props.index] === props.matcher) {
      setActive(true)
    } else {
      setActive(false)
    }
  }, [location.pathname])

  return (
    <>
      <Link to={props.path ? props.path : ''} >
        <div className={`__menuc_item ${active ? '__menuc_active' : ''}`}>
          <div className="__menuc_item_svg">
            {props.icon}
          </div>
          <span>
            {props.title}
          </span>
        </div>
      </Link>
    </>
  )
}

export default MenuContent