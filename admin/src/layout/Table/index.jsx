import * as React from 'react';
import Paper from '@mui/material/Paper';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TablePagination from '@mui/material/TablePagination';
import TableRow from '@mui/material/TableRow';

import { getComponyList, getToken } from '../../api/comp_info.js'
import dayjs, { Dayjs } from "dayjs"


const columns = [
  { id: 'id', label: 'ID', minWidth: 30 },
  { id: 'companyname', label: '公司名称', minWidth: 100 },
  { id: 'companysize', label: '公司规模', minWidth: 30 },
  {
    id: 'name',
    label: '姓名',
    minWidth: 40,
    align: 'left',
    // format: (value) => value.toLocaleString('en-US'),
  },
  {
    id: 'businessemail',
    label: '邮箱',
    minWidth: 50,
    align: 'left',
    // format: (value) => value.toLocaleString('en-US'),
  },
  {
    id: 'requirementdescription',
    label: '需求描述',
    minWidth: 100,
    align: 'left',
    // format: (value) => value.toFixed(2),
  },
  {
    id: 'date',
    label: '提交日期',
    minWidth: 70,
    align: 'left',
    format: (value) => dayjs(value).format('YY.MM.DD hh:mm:ss'),
  },
];

function createData(id, companyname, companysize, name, businessemail) {
  const des = name / businessemail;
  return { id, companyname, companysize, name, businessemail, des };
}

const getlist = () => {
  getComponyList({
    pageSize: 10,
    pageNum: 1
  }).then(res => {
    console.log(res.value.data.items);
    return res.value.data.items
  })
}

// const rows = [
//   createData(1, 'India', 'IN', 1324171354, 3287263),
//   createData(2, 'China', 'CN', 1403500365, 9596961),
//   createData(3, 'Italy', 'IT', 60483973, 301340),
//   createData(4, 'United States', 'US', 327167434, 9833520),
//   createData(5, 'Canada', 'CA', 37602103, 9984670),
//   createData(6, 'Australia', 'AU', 25475400, 7692024),
//   createData(7, 'Germany', 'DE', 83019200, 357578),
//   createData(8, 'Ireland', 'IE', 4857000, 70273),
//   createData(9,'Mexico', 'MX', 126577691, 1972550),
//   createData(10, 'Japan', 'JP', 126317000, 377973),
//   createData(11, 'France', 'FR', 67022000, 640679),
//   createData(12, 'United Kingdom', 'GB', 67545757, 242495),
//   createData(13, 'Russia', 'RU', 146793744, 17098246),
//   createData(14, 'Nigeria', 'NG', 200962417, 923768),
//   createData(15, 'Brazil', 'BR', 210147125, 8515767),
// ];

export default function StickyHeadTable() {
  let row = [];
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(10);
  const [rows, setRows] = React.useState(row);

  // const getlist = () => {
  //   let row = []
  //   getComponyList({
  //     pageSize: 10,
  //     pageNum: 1
  //   }).then(res => {
  //     row = res.value.data.items
  //   }).catch(err => {
  //     console.log(err);
  //   })
  //   return row;
  // }

  React.useEffect(() => {
    getComponyList({
      pageSize: 800,
      pageNum: 1,
    }).then(res => {
      setRows(res.value.data.items)
    })
  }, [rowsPerPage, page])

  const handleChangePage = (event, newPage) => {
    setPage(newPage);
    // getComponyList({
    //   pageSize: 800,
    //   pageNum: 1,
    // }).then(res => {
    //   setRows(res.value.data.items)
    // })
  };

  const handleChangeRowsPerPage = (event) => {
    setRowsPerPage(+event.target.value);
    setPage(0);
    // getComponyList({
    //   pageSize: 800,
    //   pageNum: 1,
    // }).then(res => {
    //   setRows(res.value.data.items)
    // })
  };

  return (
    <Paper sx={{ width: '1099px', overflow: 'hidden' }}>
      <TableContainer sx={{ maxHeight: 800 }}>
        <Table stickyHeader aria-label="sticky table">
          <TableHead>
            <TableRow>
              {columns.map((column) => (
                <TableCell
                  key={column.id}
                  align={column.align}
                  style={{ minWidth: column.minWidth }}
                >
                  {column.label}
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {rows
              .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
              .map((row) => {
                return (
                  <TableRow hover role="checkbox" tabIndex={-1} key={row.id}>
                    {columns.map((column) => {
                      const value = row[column.id];
                      return (
                        <TableCell key={column.id} align={column.align}>
                          {column.format
                            ? column.format(value)
                            : value}
                        </TableCell>
                      );
                    })}
                  </TableRow>
                );
              })}
          </TableBody>
        </Table>
      </TableContainer>
      <TablePagination
        rowsPerPageOptions={[10, 25, 100]}
        component="div"
        count={rows.length}
        rowsPerPage={rowsPerPage}
        page={page}
        onPageChange={handleChangePage}
        onRowsPerPageChange={handleChangeRowsPerPage}
      />
    </Paper>
  );
}