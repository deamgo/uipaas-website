import React, { MouseEventHandler } from "react"
//
import './index.less'
//

interface IMultiplySelectorPropsItem {
  id: number
  text: string
  method: MouseEventHandler<HTMLDivElement>
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
              <div key={item.id} className={`__mulselector_wrapper_item ${item.type}`} onClick={item.method}>
                {item.text}
              </div>
            </>
          ))
        }
      </div>
    </>
  )
}

export default MultiplySelector