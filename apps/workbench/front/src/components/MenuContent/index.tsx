import React from "react"
//
import './index.less'

interface mcontent {
  id: number;
  title: string;
  icon: React.ReactElement;
}

const MenuContent: React.FC<mcontent> = (props) => {
  return (
    <>
      <div className="__mcontent" key={props.id}>
        <div className="__mcontent_son">
          <div className="__mcontent_son_svg">
            {props.icon ? props.icon : <></>}
          </div>
          <span>
            {props.title ? props.title : <></>}
          </span>
        </div>
      </div>
    </>
  )
}

export default MenuContent