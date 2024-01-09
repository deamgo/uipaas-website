import React from 'react'
//
import './index.less'
import { ITableCellProps } from '../propsInterface';
//

const TableCell: React.FC<ITableCellProps> = (props) => {
  return (
    <>
      <td className="__tablecell" style={props.ys}>
        {props.children}
      </td>
    </>
  )
}

export default TableCell;