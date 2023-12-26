import React from 'react'
//
import './index.less'
import Input from '@/components/Input'
import { Avatar } from 'antd'
import Button from '@/components/Button'
import SearchLine from '@/components/SearchLine'
import Table from '@/components/Tables/Table'
import TableHead from '@/components/Tables/TableHead'
import TableRow from '@/components/Tables/TableRow'
import TableCell from '@/components/Tables/TableCell'
import TableBody from '@/components/Tables/TableBody'
import Select from '@/components/Select'
import Pagenation from '@/components/Pagenation'
import { ISelectOption } from '@/interface/some'


interface IColumns {
  id: 'name' | 'email' | 'role'
  label: string
}

const columns: IColumns[] = [
  {
    id: 'name',
    label: 'Name',
  },
  {
    id: 'email',
    label: 'Email',
  },
  {
    id: 'role',
    label: 'Role',
  }
]

const rows = [
  {
    id: '1001',
    name: 'Shawn',
    email: 'Shawn@example.com',
    role: 'Owner',
    status: 's',
  },
  {
    id: '1002',
    name: 'Lisa',
    email: 'Lisa@example.com',
    role: 'Admin',
    status: 's',
  },
  {
    id: '1003',
    name: 'Tom',
    email: 'Tom@example.com',
    role: 'Editer',
    status: 's',
  },
  {
    id: '1004',
    name: 'Jerry',
    email: 'Jerry@example.com',
    role: 'Viewer',
    status: 'pending',
  },
]

const list_r: ISelectOption[] = [
  {
    id: 'admin',
    value: 'admin',
    label: 'Admin',
  },
  {
    id: 'editer',
    value: 'editer',
    label: 'Editer',
  },
  {
    id: 'viewer',
    value: 'viewer',
    label: 'Viewer',
  }
]



const WSDevelopers: React.FC = () => {


  const handleChangeCurrentPage = () => {
    console.log('handleChangeCurrentPage')
  }


  const handleChangeRole = (id: string): void => {
    console.log('handleChangeRole')
  }
  return (
    <>
      <div className="__workspace_developers">
        <div className="__workspace_developers_tools">
          <div className="__workspace_developers_tools_addbtn">
            <Button context='Add developers' type='primary' />
          </div>
          <div className="__workspace_developers_tools_searchline">
            <SearchLine placeholder='Search' />
          </div>
        </div>
        <div className="__workspace_developers_container">
          <Table >
            <TableHead>
              <TableRow>
                {columns.map(item => (
                  <>
                    <TableCell
                      key={item.id}
                      ys={{
                        width: '276rem',
                        height: '36rem',
                        padding: '7rem 16rem',
                        backgroundColor: '#F3F3F3',
                        borderWidth: '0 0 1rem 0',
                        borderStyle: 'solid',
                        borderColor: '#E7E7E7',

                        fontSize: '14rem',
                        fontWeight: '500',
                        lineHeight: '22rem',
                        letterSpacing: '0em',

                        color: '#3D3D3D',
                      }}>
                      {item.label}
                    </TableCell>
                  </>
                ))}
                <TableCell
                  ys={{
                    width: '276rem',
                    height: '36rem',
                    padding: '7rem 16rem',
                    backgroundColor: '#F3F3F3',
                    borderWidth: '0 0 1rem 0',
                    borderStyle: 'solid',
                    borderColor: '#E7E7E7',

                    fontSize: '14rem',
                    fontWeight: '500',
                    lineHeight: '22rem',
                    letterSpacing: '0em',

                    color: '#3D3D3D',
                  }}>
                  Operation
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {rows.map(row => (
                <>
                  <TableRow key={row.id}>
                    {columns.map(column => {
                      const value = row[column.id]
                      return (
                        <>
                          <TableCell
                            key={column.id}
                            ys={{
                              width: '276rem',
                              height: '48rem',
                              padding: '8rem 16rem',
                              backgroundColor: '#FFFFFF',
                              borderWidth: '0 0 1rem 0',
                              borderStyle: 'solid',
                              borderColor: '#E7E7E7',

                              fontSize: '14rem',
                              fontWeight: 'normal',
                              lineHeight: '22rem',
                              letterSpacing: '0em',

                              color: '#3D3D3D',
                            }}>
                            {column.id === 'name' && (
                              <>
                                <Avatar
                                  style={{ backgroundColor: '#4080FF', verticalAlign: 'middle' }}
                                  size={32}>
                                  {value.charAt(0).toUpperCase()}
                                </Avatar>
                                {value}
                                {row.status === 'pending' && (<>pending</>)}
                              </>
                            )}
                            {column.id === 'role' && (
                              <>
                                {value === 'Owner' ? (
                                  <>{value}</>
                                ) : (
                                  <>
                                    <Select onChange={() => handleChangeRole(row.id)} list={list_r} default={value}>
                                      {value}
                                    </Select>
                                  </>)}
                              </>
                            )}
                            {column.id !== 'name' && column.id !== 'role' && (
                              <>
                                {value}
                              </>
                            )}
                            {/* {column.id === 'name'
                              ? (
                                <>
                                  <Avatar
                                    style={{ backgroundColor: '#4080FF', verticalAlign: 'middle' }}
                                    size={32}>
                                    {value.charAt(0).toUpperCase()}
                                  </Avatar>
                                  {value}
                                  {row.status === 'pending' && (<>pending</>)}
                                </>
                              )
                              : column.id === 'role'
                                ? (
                                  <>
                                    <Select onChange={handleChangeRole} list={list_r} default={value}>
                                      {value}
                                    </Select>
                                  </>
                                )
                                : (<>{value}</>)
                            } */}
                          </TableCell>
                        </>
                      )
                    })}
                    <TableCell
                      ys={{
                        width: '276rem',
                        height: '48rem',
                        padding: '8rem 16rem',
                        backgroundColor: '#FFFFFF',
                        borderWidth: '0 0 1rem 0',
                        borderStyle: 'solid',
                        borderColor: '#E7E7E7',

                        fontSize: '14rem',
                        fontWeight: 'normal',
                        lineHeight: '22rem',
                        letterSpacing: '0em',

                        color: '#3D3D3D',
                      }}>
                      <div className="__workspace_developers_container_remove">
                        <Button context='Remove' type='board-danger' />
                      </div>
                    </TableCell>
                  </TableRow>
                </>
              ))}
            </TableBody>
          </Table>
        </div>
        <div className="__workspace_developers_pagenation">
          <Pagenation pages={5} total={20} current={1} onCurrentPageChange={handleChangeCurrentPage} />
        </div>
      </div>
    </>
  )
}

export default WSDevelopers