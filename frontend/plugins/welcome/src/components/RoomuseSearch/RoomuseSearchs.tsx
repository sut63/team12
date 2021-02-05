import React, { useState, useEffect, FC } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntRoomuse } from '../../api/models/EntRoomuse'; // import interface Roomuse
import {
    Content,
    Header,
    Page,
    ContentHeader,
   } from '@backstage/core';
import {
    Container,
    Grid,
    FormControl,
    Select,
    InputLabel,
    MenuItem,
    TextField,
    Avatar,
} from '@material-ui/core';
import { Link as RouterLink } from 'react-router-dom';
import moment from 'moment';
import Swal from 'sweetalert2';
import Alert from '@material-ui/lab/Alert';
import { UserHeader } from '../UserHeader/UserHeader';

import { EntRoom } from '../../api/models/EntRoom';



 
const useStyles = makeStyles(theme => ({
  root: {
      flexGrow: 1,
  },
  paper: {
      marginTop: theme.spacing(2),
      marginBottom: theme.spacing(2),
  },
  formControl: {
      width: 300,
  },
  selectEmpty: {
      marginTop: theme.spacing(2),
  },
  container: {
      display: 'flex',
      flexWrap: 'wrap',
  },
  textField: {
      width: 300,
  },
  bottom: {
      marginLeft: theme.spacing(3),
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(2),
  },
  divider: {
      margin: theme.spacing(2, 0),
  },
  table: {
      minWidth: 650,
  },
}));

export default function RoomuseTable() {
 const classes = useStyles();
 const api = new DefaultApi();

 const [roomuses, setRoomuse] = useState<EntRoomuse[]>([]);

 const [loading, setLoading] = useState(true);

 const [rooms, setRooms] = useState<EntRoom[]>([]);
 const [roomid, setRoomID] = useState(Number);

 // alert setting
 const [open, setOpen] = React.useState(false);
 const [fail, setFail] = React.useState(false);

 // alert setting
 const Toast = Swal.mixin({
  toast: true,
  position: 'top-end',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  didOpen: toast => {
    toast.addEventListener('mouseenter', Swal.stopTimer);
    toast.addEventListener('mouseleave', Swal.resumeTimer);
  },
});

const alertMessage = (icon: any, title: any) => {
  Toast.fire({
    icon: icon,
    title: title,
  });
}

 const getRoomuses = async () => {
  const res:any = await api.getRoomuse( { id: roomid } );
  setLoading(false);
  setRoomuse(res);
  console.log(res);
  if (res != 0) {
    Toast.fire({
      icon: 'success',
      title: 'ค้นหาข้อมูลสำเร็จ',
    });

  } else {
    Toast.fire({
      icon: 'error',
      title: 'ไม่พบข้อมูล',
    });
  }
 
 };
 

 useEffect(() => {

   const getRoom = async () => {
    const res = await api.listRoom({ limit: 2, offset: 0 });
    setLoading(false);
    setRooms(res);
    console.log(res);
   };
   getRoom();
   
  }, [loading]);

  // function seach data
  function seach() {
    getRoomuses();
    
  } 

 const roomidhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setRoomID(event.target.value as number);
  console.log(roomid);
 };
 
 const deleteRoomuse = async (id: number) => {
   const res = await api.deleteRoomuse({ id: id });
   setLoading(true);
 };
 
 return (
    <Page>
    <Header
    title="ประวัติการเข้าใช้ห้อง">
        <UserHeader />
      
        {/* <Button
                style={{ marginRight: 50 }}
                component={RouterLink}
                to="/welcome"
                variant="contained"
              >
                BACK
             </Button> */}
    </Header>
    <Content>

    <ContentHeader title="ค้นหาข้อมูลการเข้าใช้">

             <Grid item xs={12}></Grid>
             <Grid item xs={3}>
               <div className={classes.paper}>ค้นหาจากชื่อห้อง</div>
             </Grid>
             <Grid item xs={3}>
               <FormControl variant="outlined" className={classes.formControl}>
                 <InputLabel id="room_id-label">เลือกห้อง</InputLabel>
                <Select
                   labelId="room_id-label"
                   label="Room"
                   id="room_id"
                   value={roomid}
                   onChange={roomidhandleChange}
                   style = {{width: 200}}
                >
                  {rooms.map((item:EntRoom)=>
                      <MenuItem value={item.id}>{item.roomName}</MenuItem>)} 
                </Select>
               </FormControl>
             </Grid>
             <Grid item xs={3}></Grid>
             <Grid item xs={3}>
                        <Button
                            variant="contained"
                            color="secondary"
                            size="large"
                            onClick={seach}
                        >
                            search
                  </Button>
                </Grid>
      </ContentHeader>
             

   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">ลำดับ</TableCell>
           <TableCell align="center">ชื่อห้อง</TableCell>
           <TableCell align="center">วัตถุประสงค์</TableCell>
           <TableCell align="center">อีเมล์</TableCell>
           <TableCell align="center">จำนวนคนเข้าใช้ห้อง</TableCell>
           <TableCell align="center">ข้อความ</TableCell>
           <TableCell align="center">เบอร์ติดต่อ</TableCell>
           <TableCell align="center">เวลาเข้าใช้</TableCell>
           <TableCell align="center">เวลาออก</TableCell>
           <TableCell align="center">จัดการ</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {roomuses.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.rooms?.roomName}</TableCell>
             <TableCell align="center">{item.edges?.purposes?.purpose}</TableCell>
             <TableCell align="center">{item.edges?.users?.email}</TableCell>
             <TableCell align="center">{item.age}</TableCell>
             <TableCell align="center">{item.note}</TableCell>
             <TableCell align="center">{item.contact}</TableCell>
             <TableCell align="center">{moment(item.in_time).format('DD/MM/YYYY HH.mm  ')}</TableCell>
             <TableCell align="center">{moment(item.out_time).format('DD/MM/YYYY HH.mm  ')}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                    deleteRoomuse(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
   </Content>
   </Page>
 );
}