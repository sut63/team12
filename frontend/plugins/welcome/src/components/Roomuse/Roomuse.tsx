import React, { useState, useEffect } from 'react';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  SidebarPage,
  Link,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';

import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import Grid from '@material-ui/core/Grid';
import Container from '@material-ui/core/Container';
import Swal from 'sweetalert2';
import { Link as RouterLink } from 'react-router-dom';

import { EntUser } from '../../api/models/EntUser';
import { EntRoom } from '../../api/models/EntRoom';
import { EntPurpose } from '../../api/models/EntPurpose';
import { UserHeader } from '../UserHeader/UserHeader';
import { AppSidebar } from '../Sidebar/Sidebar';

import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import WarningIcon from '@material-ui/icons/Warning';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
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
      width: '25ch',
    },
  }));

export default function Roomuse() {
  const classes = useStyles();
  const api = new DefaultApi();

  const [checker, SetChecker] = useState(false);
  const [open, setOpen] = React.useState(false);
  const [users, setUsers] = useState<EntUser[]>([]);
  const [rooms, setRooms] = useState<EntRoom[]>([]);
  const [purposes, setPurposes] = useState<EntPurpose[]>([]);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(Boolean);
  const [loading, setLoading] = useState(true);

  const [intime, setIntime] = useState(String);
  const [outtime, setOuttime] = useState(String);

  const [roomid, setRoom] = useState(Number);
  const [purposeid, setPurpose] = useState(Number);
  const [userid, setUser] = useState(Number);

  //const [addedpeople, SetPeople] = useState(Number);
  // const [addedmemo, SetMemo] = useState(String);
  const [contact, SetContact] = useState(String);
  const [note, SetNote] = useState(String);
  const [people, SetPeople] = useState(Number);

  let People = Number(people);

  const [noteError, setNoteError] = useState('');
  const [contactError, setContactError] = useState('');
  const [peopleError, setPeopleError] = useState('');

  const validatenote = (val: string) => {
    return val.match("^[A-Za-zก-๙]{1,25}$");
  }
  const validatepeople = (val: string) => {
    var a = Number(val);
    if (a > 0) {
      return String(val)
    }
    return val = ''
  }
  const validatecontact = (val: string) => {
    return val.length == 10 ? true : false;
  }

  const checkPattern = (id: string, value: string) => {
    switch (id) {
      case 'note':
        validatenote(value) ? setNoteError('') : setNoteError('ต้องมีข้อความสั้น 1-25 ตัวอักษร');
        return;
        case 'people':
          validatepeople(value) ? setPeopleError('') : setPeopleError('ข้อมูลจำนวนผู้เข้าใช้ไม่ถูกต้อง');
        return;
      case 'contact':
        validatecontact(value) ? setContactError('') : setContactError('เบอร์โทรศัพท์ 10 หลัก');
        return;
    }
  }

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

  useEffect(() => {
    const getUser = async () => {
      const res = await api.listUser({ limit: 4, offset: 0 });
      setLoading(false);
      setUsers(res);
      console.log(res);
    };
    getUser();

    const getRoom = async () => {
      const res = await api.listRoom({ limit: 6, offset: 0 });
      setLoading(false);
      setRooms(res);
      console.log(res);
    };
    getRoom();

    const getPurpose = async () => {
      const res = await api.listPurpose({ limit: 5, offset: 0 });
      setLoading(false);
      setPurposes(res);
      console.log(res);
    };
    getPurpose();

    const getChecker = async () => {
      if(uStatus == "Unknow" && uType != "เจ้าหน้าที่"){
        SetChecker(true);
      }
    };
    getChecker();

  }, [loading]);

  const handleInDatetimeChange = (event: any) => {
    setIntime(event.target.value as string);
  };
  const handleOutDatetimeChange = (event: any) => {
    setOuttime(event.target.value as string);
  };
  const ContacthandleChang = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as keyof typeof Roomuse
    const { value } = event.target;
    const validatecontact = value.toString()
    checkPattern(id, validatecontact)
    SetContact(event.target.value as string);
  };
  const PeoplehandleChang = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as keyof typeof Roomuse
    const { value } = event.target;
    const validatepeople = value.toString()
    checkPattern(id, validatepeople)
    SetPeople(event.target.value as number);
  };
  const NotehandleChang = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as keyof typeof Roomuse
    const { value } = event.target;
    const validatenote = value.toString()
    checkPattern(id, validatenote)
    SetNote(event.target.value as string);
  };

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'people':
        alertMessage("error","ต้องมีคนเข้าใช้ห้องอย่างน้อย 1 คน");
        return;
      case 'contact':
        alertMessage("error","เบอร์โทรศัพท์ต้องมี 10 หลัก");
        return;
      case 'note':
        alertMessage("error","ควรมีข้อความสั้นและไม่ควรเกิน 25 ตัวอักษร");
        return;
      default:
        alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }

  
    //set UserData ..
    var uType = JSON.parse(String(localStorage.getItem('user-type')));
    var uStatus = JSON.parse(String(localStorage.getItem('user-status')));
  
    const roomuse = {
      note          : note,
      People,
      contact       : contact,
      inTime        : intime + ":00+07:00",
      outTime       : outtime + ":00+07:00",
      userID        : userid,
      roomID        : roomid,
      purposeID     : purposeid,
    }
    console.log(roomuse);

    function save() {
      const apiUrl = 'http://localhost:8080/api/v1/roomuses';
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(roomuse),
      };
      console.log(roomuse);
  
      fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
  
        } else {
          checkCaseSaveError(data.error.Name)
          console.log(data.error.Name)
        }
      })
  
    }

  const useridhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUser(event.target.value as number);
  };

  const roomidhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setRoom(event.target.value as number);
  };

  const purposeidhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPurpose(event.target.value as number);
  };

  function home() {
    history.pushState("", "", "/welcome");
    window.location.reload(false);
  }
  function handleClose() {
    setOpen(false);
  };
  function handleCloseChecker() {
    SetChecker(false);
    home();
  };

  const refreshPage = () => {
    window.location.reload();
  }


  return (
    <SidebarPage>
      <AppSidebar />
      <Page theme={pageTheme.home}>
      <Header title="การเข้าใช้ห้องชมรม">  
      <UserHeader />  
      </Header>
      <Dialog
        open={checker}
        onClose={handleClose}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
      >
      <DialogTitle id="alert-dialog-title">
          <WarningIcon color="secondary"/>
          &nbsp;ไม่ตรงตามข้อกำหนด
        </DialogTitle>
        <DialogContent dividers>
          <DialogContentText  id="alert-dialog-description">
            คุณยังไม่เป็นสมาชิกชมรม<br/>หากต้องการใช้ห้อง กรุณาสมัครชมรมก่อน
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseChecker} color="secondary" autoFocus variant="outlined">
            ฉันเข้าใจแล้ว
          </Button>
        </DialogActions>
      </Dialog>
      <Content>
        <ContentHeader title="ใส่ข้อมูลการเข้าใช้">

        <Link component={RouterLink} to="/RoomuseSearch">
         <Button style={{ marginLeft: 20 }} variant="contained" color="secondary" size="large" >
            ค้นหา / ประวัติ
         </Button>
        </Link>

        {/* {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 บันทึกสำเร็จ!!
               </Alert>
             ) : (
               <Alert severity="error">
                 <strong>บันทึกล้มเหลว!!</strong>
               </Alert>
             )}
           </div>
         ) : null} */}

        </ContentHeader>
        <Container maxWidth="sm">
           <Grid container spacing={3}>
             <Grid item xs={12}></Grid>
             <Grid item xs={3}>
               <div className={classes.paper}>ชื่อห้อง</div>
             </Grid>
             <Grid item xs={9}>
               <FormControl variant="outlined" className={classes.formControl}>
                 <InputLabel id="room_id-label">เลือกห้อง</InputLabel>
                <Select
                   labelId="room_id-label"
                   label="Room"
                   id="room_id"
                   value={roomid}
                   onChange={roomidhandleChange}
                   style = {{width: 300}}
                >
                  {rooms.map((item:EntRoom)=>
                    <MenuItem value={item.id}>{item.roomName}</MenuItem>)}
                </Select>
               </FormControl>
             </Grid>
 
             <Grid item xs={3}>
               <div className={classes.paper}>จุดประสงค์</div>
             </Grid>
             <Grid item xs={9}>
               <FormControl variant="outlined" className={classes.formControl}>
                 <InputLabel id="purpose_id-label">ใช้ทำอะไร</InputLabel>
                <Select
                  labelId="purpose_id-label"
                  label="Purpose"
                  id="purpose_id"
                  value={purposeid}
                  onChange={purposeidhandleChange}
                  style = {{width: 300}}
               >
                 {purposes.map((item:EntPurpose)=>
                   <MenuItem value={item.id}>{item.purpose}</MenuItem>)}
                  </Select>
               </FormControl>
             </Grid>
 
             
 
             <Grid item xs={3}>
               <div className={classes.paper}>ผู้ใช้</div>
             </Grid>
             <Grid item xs={9}>
               <FormControl variant="outlined" className={classes.formControl}>
                 <InputLabel id="user_id-label">เลือกผู้ใช้</InputLabel>
                <Select
                  labelId="user_id-label"
                  label="User"
                  id="user_id"
                  value={userid}
                  onChange={useridhandleChange}
                  style = {{width: 300}}
               >
                 {users.map((item:EntUser)=>
                   <MenuItem value={item.id}>{item.name}</MenuItem>)}
                  </Select>
               </FormControl>
             </Grid>
 
             <Grid item sm={3}>
                  <div className={classes.paper}>Short Note</div>
                </Grid>
                <Grid item sm={9}>
                  <form className={classes.paper} noValidate autoComplete="off">

                    <TextField
                      //type="reset"
                      error={noteError ? true : false}
                      id="note"
                      size="medium"
                      variant="outlined"
                      helperText={noteError}
                      value={note}
                      onChange={NotehandleChang}
                    />
                  </form>
                </Grid>

                <Grid item sm={3}>
                  <div className={classes.paper}> จำนวนคนเข้าใช้ห้อง</div>
                </Grid>
                <Grid item sm={9}>
                  <form className={classes.paper} noValidate autoComplete="off">

                    <TextField
                      error={peopleError ? true : false}
                      id="people"
                      size="medium"
                      variant="outlined"
                      helperText={peopleError}
                      value={people}
                      onChange={PeoplehandleChang}
                    />
                  </form>
                </Grid>

             <Grid item sm={3}>
                  <div className={classes.paper}>เบอร์ติดต่อ</div>
                </Grid>
                <Grid item sm={9}>
                  <form className={classes.paper} noValidate autoComplete="off">
                    <TextField
                      error={contactError ? true : false}
                      id="contact"
                      variant="outlined"
                      helperText={contactError}
                      value={contact}
                      onChange={ContacthandleChang}
                    />
                  </form>
                </Grid>

             <Grid item xs={3}>
               <div className={classes.paper} >Check-in</div>
             </Grid>
             <Grid item xs={9}>
               <div>
               <FormControl
                className={classes.margin}
                variant="outlined"
              >
               <TextField
                   id="indatetime"
                   label="inDateTime"
                   type="datetime-local"
                   value={intime}
                   onChange={handleInDatetimeChange}
                   className={classes.textField}
                   InputLabelProps={{
                     shrink: true,
                   }}
                />
               </FormControl>
               </div>
             </Grid>

             <Grid item xs={3}>
               <div className={classes.paper} >Check-out</div>
             </Grid>
             <Grid item xs={9}>
               <div>
               <FormControl
                className={classes.margin}
                variant="outlined"
              >
               <TextField
                   id="outdatetime"
                   label="outDateTime"
                   type="datetime-local"
                   value={outtime}
                   onChange={handleOutDatetimeChange}
                   className={classes.textField}
                   InputLabelProps={{
                     shrink: true,
                   }}
                />
               </FormControl>
               </div>
             </Grid>
 
             <Grid item xs={3}></Grid>
             <Grid item xs={9}>
               <Button
                 variant="contained"
                 color="primary"
                 size="large"
                 style={{ marginRight: 1 }}
                 onClick={() => {
                  save();
                }}
               >
                 บันทึก
               </Button>
             </Grid>
           </Grid>
         </Container>
      </Content>
    </Page>
    </SidebarPage>
  );

};
