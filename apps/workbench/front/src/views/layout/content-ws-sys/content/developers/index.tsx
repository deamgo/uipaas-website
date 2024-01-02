import React from 'react'
//
import './index.less'
//
import { Avatar } from 'antd'
import Input from '@/components/Input'
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
//
import { ReactComponent as Plus } from '@assets/comps/plus.svg'
import Label from '@/components/Label'
import Mask from '@/components/Mask'
import Popup from '@/components/Popup'
import { useLoaderData } from 'react-router-dom'
import { currentWorkspaceStore } from '@/store/wsStore'
import { editDeveloperPermission, findDeveloper, getDevelopers, inviteByEmail, removeDeveloper } from '@/api/workspace_settings'
import $message from '@/components/Message'



interface IColumns {
  id: 'username' | 'email' | 'role'
  label: string
}

const columns: IColumns[] = [
  {
    id: 'username',
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

interface IDeveloper {
  developer_id: string
  username: string
  email: string
  role: string
  status: number
}

// const rows = [
//   {
//     id: '1001',
//     name: 'Shawn',
//     email: 'Shawn@example.com',
//     role: 'Owner',
//     status: 's',
//   },
//   {
//     id: '1002',
//     name: 'Lisa',
//     email: 'Lisa@example.com',
//     role: 'Admin',
//     status: 's',
//   },
//   {
//     id: '1003',
//     name: 'Tom',
//     email: 'Tom@example.com',
//     role: 'Editer',
//     status: 's',
//   },
//   {
//     id: '1004',
//     name: 'Jerry',
//     email: 'Jerry@example.com',
//     role: 'Viewer',
//     status: 'pending',
//   },
// ]

const list_r: ISelectOption[] = [
  {
    id: 'admin',
    value: '1',
    label: 'Admin',
  },
  {
    id: 'editer',
    value: '2',
    label: 'Editer',
  },
  {
    id: 'viewer',
    value: '3',
    label: 'Viewer',
  }
]



const WSDevelopers: React.FC = () => {

  const [isAddDevelopers, setIsAddDevelopers] = React.useState<boolean>(false)
  const [isMask, setMask] = React.useState<boolean>(false)

  const [pages, setPages] = React.useState<number>(0)
  const [currenPage, setCurrentPage] = React.useState<number>(1)
  const [total, setTotal] = React.useState<number>(1)

  const [developers, setDevelopers] = React.useState<IDeveloper[]>([])

  const [queryParam, setQueryParam] = React.useState<string>('')

  const [inviteEmail, setInviteEmail] = React.useState<string>('')
  const [inviteByEmailRole, setInviteByEmailRole] = React.useState<string>('3')

  const loader = useLoaderData() as IDeveloper[]

  React.useEffect(() => {
    if (loader.length > 0) {
      setDevelopers(loader ? loader : [])
    }

    console.log(developers.length);

    let tempTotal = Math.ceil(developers.length / 10)


    setTotal(tempTotal)
    if (tempTotal > 0 && tempTotal < 5) {
      setPages(tempTotal)
    } else if (tempTotal !== 0 && tempTotal >= 5) {
      setPages(5)
    }


  }, [currenPage])


  const handleOpenMask = () => {
    setMask(!isMask)
  }

  const handleOpenAddDevelopers = () => {
    handleOpenMask()
    setIsAddDevelopers(!isAddDevelopers)
  }


  const handleChangeCurrentPage = async (item: number) => {
    console.log('handleChangeCurrentPage' + item)
    setCurrentPage(item)
    try {
      const { value } = await getDevelopers(item)
      setDevelopers(value.data)
      $message.success('Success to get developers')
    } catch (error) {
      console.log(error);
      $message.error('Failed to get developers')
    }
  }

  const handeSearch = async () => {
    console.log('handeSearch:' + queryParam)
    try {
      const { value } = await findDeveloper(currenPage, queryParam)
      setDevelopers(value.data)
      $message.success(value.msg)
    } catch (error) {
      console.log(error);
      $message.error('Failed to get developers by search')
    }
  }

  const handleRemove = async (id: string) => {
    console.log('handleRemove:' + id)
    try {
      const { value } = await removeDeveloper({
        developer_id: id
      })
      console.log(value);
      handleChangeCurrentPage(currenPage)

    } catch (error) {
      console.log(error);

    }
  }

  const handleInviteDevelopers = async () => {
    console.log('handleInviteDevelopers' + inviteEmail + inviteByEmailRole)
    try {
      const { code, msg } = await inviteByEmail({
        email: inviteEmail,
        role: inviteByEmailRole
      })
      console.log(code);
      handleChangeCurrentPage(currenPage)
      $message.success(msg)
    } catch (error) {
      console.log(error);
      $message.error('Failed to invite developers')

    }
  }


  const handleRowChangeRole = async (role: string, developer_id?: string) => {
    console.log({
      role,
      developer_id
    })
    try {
      const { value } = await editDeveloperPermission({
        developer_id,
        role: role
      })
      handleChangeCurrentPage(currenPage)
      $message.success(value.msg)
    } catch (error) {
      console.log(error);
      $message.error('Failed to edit developer permission')
    }

  }
  return (
    <>
      <div className="__workspace_developers">
        <div className="__workspace_developers_tools">
          <div className="__workspace_developers_tools_addbtn">
            <Button context='Add developers' type='primary' method={handleOpenAddDevelopers}>
              <Plus style={{
                width: '10.67rem',
                height: '10.67rem',
                fill: '#FFFFFF'
              }} />
              Add developers
            </Button>
          </div>
          <div className="__workspace_developers_tools_searchline">
            <SearchLine placeholder='Search' outputChange={setQueryParam} searchClick={handeSearch} />
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
              {developers?.map((row, index) => (
                <>
                  <TableRow key={index}>
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
                            {column.id === 'username' && (
                              <>
                                <div className="__workspace_developers_name">
                                  <Avatar
                                    style={{ backgroundColor: '#4080FF', verticalAlign: 'middle' }}
                                    size={32}>
                                    {value.charAt(0).toUpperCase()}
                                  </Avatar>
                                  {value}
                                  {row['status'] === 0 && (<Label label={'pending'} type='info' />)}
                                </div>
                              </>
                            )}
                            {column.id === 'role' && (
                              <>
                                {value === 'Owner' ? (
                                  <>{value}</>
                                ) : (
                                  <>
                                    <Select onChange={handleRowChangeRole} id={row['developer_id']} list={list_r} default={value}>
                                      {value}
                                    </Select>
                                  </>)}
                              </>
                            )}
                            {column.id !== 'username' && column.id !== 'role' && (
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
                      {
                        row['role'] !== 'Owner' &&
                        (<div className="__workspace_developers_container_remove">
                          <Button context='Remove' type='board-danger' method={() => handleRemove(row['developer_id'])}>
                            Remove
                          </Button>
                        </div>)
                      }
                    </TableCell>
                  </TableRow>
                </>
              ))}
            </TableBody>
          </Table>
        </div>
        <div className="__workspace_developers_pwrapper">
          <div className="__workspace_developers_pwrapper_pagenation">
            <Pagenation
              pages={pages}
              total={total}
              current={currenPage}
              onCurrentPageChange={handleChangeCurrentPage} />
          </div>
        </div>
      </div>
      {isMask && <Mask />}
      {isAddDevelopers && <Popup unit='rem' width={480} height={272} title='Add developers' onClose={handleOpenAddDevelopers}>
        <div className="_add_developers_popup">
          <div className="_add_developers_popup_wrapper">
            <div className="_add_developers_popup_wrapper_input">
              <Input
                id='invitebyemail'
                title='Invite by email'
                placeholder='Enter email address'
                isNeed={false}
                outputChange={setInviteEmail} />
              <div className="_add_developers_popup_wrapper_input_select">
                <Select list={list_r} default='Viewer' onChange={setInviteByEmailRole} />
              </div>
            </div>
            <div className="_add_developers_popup_wrapper_invite">
              <Button context='Invite' type='primary' method={handleInviteDevelopers} >
                Invite
              </Button>
            </div>
          </div>
          <div className="_add_developers_popup_wrapper">
            <div className="_add_developers_popup_wrapper_input">
              <Input
                id='invitelink'
                title='Invite link'
                value={'https://www.uipaas.com/' + currentWorkspaceStore.getCurrentWorkspace().id}
                isNeed={false}
                typeAble={true} />
              <div className="_add_developers_popup_wrapper_input_select">
                <Select list={list_r} default='Viewer' />
              </div>
            </div>
            <div className="_add_developers_popup_wrapper_invite">
              <Button type='primary' method={handleOpenAddDevelopers} >
                Copy link
              </Button>
            </div>
          </div>
        </div>
      </Popup>}
    </>
  )
}

export default WSDevelopers