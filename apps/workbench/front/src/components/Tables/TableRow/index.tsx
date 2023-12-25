import React from 'react'
//
import './index.less'
import { ITableRowProps } from '../propsInterface';
//

const TableRow: React.FC<ITableRowProps> = (props) => {
  return (
    <>
      <tr className="__tablerow">
        {props.children}
      </tr>
    </>
  )
}

export default TableRow;