import React from "react"
//
import './index.less'
import { IMultiplySelectorPropsItem } from "@/interface/some"
import { Link } from "react-router-dom"
//

type MultiplySelectorProps = {
  list: IMultiplySelectorPropsItem[],
  onClose: () => void,
}

const MultiplySelector: React.FC<MultiplySelectorProps> = (props) => {

  return (
    <>
      <div
        className="__mulselector_wrapper"
        onClick={props.onClose}>
        {
          props.list && props.list.map(item => (
            <Link
              to={item.path}
              key={item.id}
              onClick={item.method}
              className={`__mulselector_wrapper_item ${item.type}`}>
              {item.children}
            </Link>
          ))
        }
      </div>
    </>
  )
}

export default MultiplySelector