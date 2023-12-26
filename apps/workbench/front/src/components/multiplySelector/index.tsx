import React, { MouseEventHandler } from "react"
//
import './index.less'
import { Link, useNavigate } from "react-router-dom"
import { IMultiplySelectorPropsItem } from "@/interface/some"
//

type MultiplySelectorProps = {
  list: IMultiplySelectorPropsItem[]
}

const MultiplySelector: React.FC<MultiplySelectorProps> = (props) => {

  return (
    <>
      <div className="__mulselector_wrapper" >
        {
          props.list && props.list.map(item => (
            <div
              key={item.id}
              className={`__mulselector_wrapper_item ${item.type}`}
              onClick={item.method} >
              {item.children}
            </div>
          ))
        }
      </div>
    </>
  )
}

export default MultiplySelector