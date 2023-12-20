import React, { MouseEventHandler } from "react"
//
import './index.less'
import { Link } from "react-router-dom"
//

interface IMultiplySelectorPropsItem {
  id: number
  text: string
  path: string
  type: 'normal' | 'error'
}

type MultiplySelectorProps = {
  list: IMultiplySelectorPropsItem[]
}

const MultiplySelector: React.FC<MultiplySelectorProps> = (props) => {
  return (
    <>
      <div className="__mulselector_wrapper">
        {
          props.list[0] !== null && props.list.map(item => (
            <>
              <div key={item.id} className={`__mulselector_wrapper_item ${item.type}`} >
                <Link to={item.path}>{item.text}</Link>
              </div>
            </>
          ))
        }
      </div>
    </>
  )
}

export default MultiplySelector