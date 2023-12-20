import React from "react"
//
import './index.less'
import { Link } from "react-router-dom";

interface mcontent {
  id: number;
  title: string;
  path: string;
  icon: React.ReactElement;
}

const MenuContent: React.FC<mcontent> = (props) => {
  return (
    <>
      <div className="__mcontent" key={props.id}>
        <Link to={props.path}>
          <div className="__mcontent_son">
            <div className="__mcontent_son_svg">
              {props.icon ? props.icon : <></>}
            </div>
            <span>
              {props.title ? props.title : <></>}
            </span>
          </div>
        </Link>
      </div>
    </>
  )
}

export default MenuContent