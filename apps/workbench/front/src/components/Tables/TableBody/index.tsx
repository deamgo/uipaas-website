import React from 'react'
//
import './index.less'
import { ITableBodyProps } from '../propsInterface'
//

const TableBody: React.FC<ITableBodyProps> = (props) => {
  return (
    <>
      <tbody className="__tablebody">
        {props.children}
      </tbody>
    </>
  )
}

export default TableBody