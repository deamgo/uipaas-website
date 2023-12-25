import React from 'react';
//
import './index.less'
import { ITableProps } from '../propsInterface';
//


const Table: React.FC<ITableProps> = (props) => {
  return (
    <>
      <table className="__table" style={props.ys}>
        {props.children}
      </table>
    </>
  )
}

export default Table;