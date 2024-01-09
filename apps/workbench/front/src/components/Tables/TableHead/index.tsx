import React from 'react'
//
import './index.less'
import { ITableHeadProps } from '../propsInterface'
//

const TableHead: React.FC<ITableHeadProps> = (props) => {
  return (
    <>
      <thead className="__tablehead">
        {props.children}
      </thead>
    </>
  )
}

export default TableHead