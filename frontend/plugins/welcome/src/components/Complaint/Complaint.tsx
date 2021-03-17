import React, { useState, useEffect } from 'react';
import {
  pageTheme,
  Page,
  Header,
  Content,
  ContentHeader,
  SidebarPage,
} from '@backstage/core';
import { makeStyles, Theme, createStyles} from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import TextField from '@material-ui/core/TextField';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
//import Button from '@material-ui/core/Button';
import AddIcon from '@material-ui/icons/Add';
import Swal from 'sweetalert2';
import { Link as RouterLink } from 'react-router-dom';
import { Link, Button } from '@material-ui/core';
import { UserHeader } from '../UserHeader/UserHeader';
import { AppSidebar } from '../Sidebar/Sidebar';
import { DefaultApi } from '../../api/apis';
import { EntClub } from '../../api/models/EntClub';
import { EntComplaintType } from '../../api/models/EntComplaintType';
//import Snackbar from '@material-ui/core/Snackbar';
//import { Alert } from '@material-ui/lab';
//import { EntUser } from '../../api/models/EntUser';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    paper: {
      padding: theme.spacing(2),
      textAlign: 'center',
      color: theme.palette.text.secondary,
    },
    container: {
      display: 'flex',
      flexWrap: 'wrap',
    },
    textField: {
      marginLeft: theme.spacing(1),
      marginRight: theme.spacing(1),
      width: 400,
    },
    formControl: {
      margin: theme.spacing(1),
      minWidth: 400,
    },
    button: {
      margin: theme.spacing(1),
      textAlign: 'center',
    },
    textboxroot: {
      '& .MuiTextField-root': {
        margin: theme.spacing(1),
        width: '100ch',
      },
    },
  }),
);

export default function CreateComplaint() {

  const classes = useStyles();

  const api = new DefaultApi();

  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 5000,
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

  //const [users, setUsers] = React.useState<EntUser[]>([]);
  const [clubs, setClubs] = React.useState<EntClub[]>([]);
  const [complainttypes, setComplaintTypes] = React.useState<EntComplaintType[]>([]);

  //const [userid, setUser] = React.useState(Number);
  const [date, setDate] = useState(String);
  const [clubid, setClub] = React.useState(Number);
  const [complainttypeid, setComplaintType] = React.useState(Number);
  const [info, setInfo] = useState(String);
  // const [name, setName] = useState(String);
  const [phonenumber, setPhoneNumber] = useState(String);

  //const [errorName, setErrorName] = useState(String);
  // const [nameError, setNameError] = useState('');
  const [phoneNumberError, setPhoneNumberError] = useState('');
  const [infoError, setInfoError] = useState('');

  const [uID, setUID] = useState(localStorage.getItem('user-id')); 
  let userid = Number(uID);
  const name = JSON.parse(String(localStorage.getItem('user-name')),);

  // const [nameError, setNameError] = React.useState('');
  // const [phonenumberError, setPhoneNumberError] = React.useState('');
  // const [infoError, setInfoError] = React.useState('');

  // const [status, setStatus] = React.useState(false);
  // const [alert, setAlert] = React.useState(true);
  // const [loading, setLoading] = React.useState(true);
  // const [KK, setKK] = useState(Number);

  // const [status, SetStatus] = useState(Boolean);
  // const [alert, setAlert] = useState(true);
  // const [open, setOpen] = React.useState(false);
  const [loading, setLoading] = useState(true);

  // const getUsers = async () => {
  //   const res = await api.listUser({ limit: 10, offset: 0 });
  //   setUsers(res);
  //   setLoading(false);
  //   console.log(res);
  // };

  const getClubs = async () => {
    const res = await api.listClub({ limit: 10, offset: 0 });
    setClubs(res);
    setLoading(false);
    console.log(res);
  };

  const getTypes = async () => {
    const res = await api.listComplainttype({ limit: 10, offset: 0 });
    setComplaintTypes(res);
    setLoading(false);
    console.log(res);
  };

  useEffect(() => {
    //getUsers();

    const checkUserType = async () => {
      const userType = JSON.parse(
        String(localStorage.getItem('user-type')),
      );
      setLoading(false);
      if (userType != "นักศึกษา" && userType != "เจ้าหน้าที่") {
        Swal.fire({
          title: 'คุณไม่สามารถส่งเรื่องร้องเรียนได้',
          showClass: {
            popup: 'animate__animated animate__fadeInDown',
          },
          hideClass: {
            popup: 'animate__animated animate__fadeOutUp',
          },
        });
        setTimeout(() => {
          history.pushState('', '', './welcome');
          window.location.reload(false);
        }, 3000);
      }
    };
    checkUserType();

    getClubs();
    getTypes();
  }, [loading]);

  // const handleUserChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  //   setUser(event.target.value as number);
  // };

  const handleClubChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setClub(event.target.value as number);
  };

  const handleComplaintTypeChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setComplaintType(event.target.value as number);
  };

  const handleDateChange = (event: any) => {
    setDate(event.target.value as string);
  };

  const handleInfoChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkPattern('info', validateValue)
    setInfo(event.target.value as string);
  };

  // const handleNameChange = (event: React.ChangeEvent<{ value: any }>) => {
  //   const { value } = event.target;
  //   const validateValue = value
  //   checkPattern('name', validateValue)
  //   setName(event.target.value as string);
  // };

  const handlePhoneNumberChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkPattern('phonenumber', validateValue)
    setPhoneNumber(event.target.value as string);
  };

  // const validateName = (val: string) => {
  //   return val.match(".{1,}");
  // }
  const validatePhoneNumber = (val: string) => {
    return val.length == 10 ? true : false;
  }
  const validateInfo = (val: string) => {
    return val.match(".{20,}");
  }
  
  const checkPattern = (id: string, value: string) => {
    switch (id) {
      // case 'name':
      //   validateName(value) ? setNameError('') : setNameError('รูปแบบชื่อผิดพลาด กรุณาป้อน ชื่อ นามสกุล');
      //   return;
      case 'phonenumber':
        validatePhoneNumber(value) ? setPhoneNumberError('') : setPhoneNumberError('หมายเลขโทรศัพท์ไม่ถูกต้อง');
        return;
      case 'info':
        validateInfo(value) ? setInfoError('') : setInfoError('กรุณาป้อนข้อมูลอย่างน้อย 20 ตัวอักษร');
        return;
      default:
        return;
    }
  }

  // const handleChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
  //   const id = event.target.id as keyof typeof Complaint;
  //   const { value } = event.target;
  //   const validateValue = value.toString()
  //   //checkPattern(id, validateValue)
  //   setUser({ ...user, [id]: value });
  // };

  // const CreateComplaint = async () => {
  //   if (
  //     userid != null &&
  //     clubid != null &&
  //     complainttypeid != null &&
  //     info != '' &&
  //     date != ''
  //   ) {
  //   const complaint = {
  //     userID: userid,
  //     clubID: clubid,
  //     typeID: complainttypeid,
  //     info: info,
  //     date: date + ":00+07:00",
  //   };
  //   console.log(complaint);
  //   const res: any = await api.createComplaint({ complaint : complaint });
  //   setStatus(true);
  //   setKK(res.id);
  //   if (res.id != '') {
  //     setAlert(true);
  //   }
  //   } else {
  //     setStatus(true);
  //     setAlert(false);
  //   }
  //   const timer = setTimeout(() => {
  //     setStatus(false);
  //   }, 5000);
  // };

  const complaint = {
    userID: userid,
    clubID: clubid,
    typeID: complainttypeid,
    name: name,
    phonenumber: phonenumber,
    info: info,
    date: date + "T00:00:00+07:00",
  };
  console.log(complaint);

  // function home() {
  //   history.pushState("", "", "/welcome");
  //   window.location.reload(false);
  // }

  // const validatePhoneNumber = (val: string) => {
  //   return val.length == 10 ? true : false;
  // }

  // const validateName = (val: string) => {
  //   return val
  // }

  // const validateInfo = (val: string) => {
  //   return val
  // }

  // const checkPattern  = (id: string, value: string) => {
  //   switch(id) {
  //     case 'name':
  //       validateName(value) ? setNameError('') : setNameError('ข้อมูล field name ผิดพลาด');
  //       return;
  //     case 'phonenumber':
  //       validatePhoneNumber(value) ? setPhoneNumberError('') : setPhoneNumberError('ข้อมูล field phonenumber ผิดพลาด');
  //       return;
  //     case 'info':
  //       validateInfo(value) ? setInfoError('') : setInfoError('ข้อมูล field info ผิดพลาด')
  //       return;
  //     default:
  //       return;
  //   }
  // }

  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'name':
        alertMessage("error","ข้อมูล field name ผิดพลาด");
        return;
      case 'phonenumber':
        alertMessage("error","ข้อมูล field phonenumber ผิดพลาด");
        return;
      case 'info':
        alertMessage("error","ข้อมูล field info ผิดพลาด");
        return;
      default:
        alertMessage("error","ส่งเรื่องร้องเรียนไม่สำเร็จ");
        return;
    }
  }

  // function home() {
  //   history.pushState("", "", "/welcome");
  //   window.location.reload(false);
  // }

  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/complaints';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(complaint),
    };
    console.log(complaint);

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          Toast.fire({
            icon: 'success',
            title: 'ส่งเรื่องร้องเรียนสำเร็จ',
          })
        } else {
          checkCaseSaveError(data.error.Name)
        }
      });
  };

  return(
    <SidebarPage>
      <AppSidebar />
      <Page theme={pageTheme.home}>
        <Header title={`ระบบร้องเรียน`}><UserHeader /></Header>
        <Content>
          <ContentHeader title="ส่งเรื่องร้องเรียน">
          <Link component={RouterLink} to="/ComplaintSearch">
            <Button variant="contained" color="primary">
              ค้นหาเรื่องร้องเรียน
            </Button>
          </Link>
          {/* {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  Success!!
                </Alert>
                ) : (
                <Alert severity="warning" style={{ marginTop: 20 }}>
                  Error!!
                </Alert>
              )}
            </div>
          ) : null} */}
          {/* {status ? (
              <div>
                {alert ? (
                  <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
                    <Alert onClose={handleClose} severity="success">
                      ส่งคำร้องสำเร็จ
                  </Alert>
                  </Snackbar>
                ) : (
                    <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
                      <Alert onClose={handleClose} severity="error">
                        ส่งคำร้องไม่สำเร็จกรุณาลองใหม่อีกครั้ง
                  </Alert>
                    </Snackbar>
                  )}
              </div>
            ) : null} */}
          </ContentHeader>
          <div className={classes.root}>
            <Grid container spacing={3}>
              <Grid item xs={12}>
                <Paper className={classes.paper}>
                  <div>
                    {/* <FormControl className={classes.formControl}>
                      <InputLabel id="demo-simple-select-label">เลือกผู้ใช้ระบบ</InputLabel>
                      <Select
                        name="owner"
                        value={userid}
                        onChange={handleUserChange}
                      >
                      {users.map(item => {
                        return (
                          <MenuItem key={item.id} value={item.id}>
                            {item.email}
                          </MenuItem>
                        );
                      })}
                      </Select>
                    </FormControl> */}
                  </div>
                  <Grid item xs={12}>
                    <div className={classes.paper}></div>
                  </Grid>
                  <div>
                  {/* <TextField
                      id="datetime-local"
                      label="เลือกวันที่"
                      type="datetime-local"
                      defaultValue="2017-05-24T10:30"
                      value={date}
                      onChange={handleDateChange}
                      className={classes.textField}
                      InputLabelProps={{
                        shrink: true,
                      }}
                    /> */}
                    <TextField
                      label="เลือกวันที่"
                      name="date"
                      type="date"
                      value={date}
                      className={classes.textField}
                      onChange={handleDateChange}
                      InputLabelProps={{
                        shrink: true,
                      }}
                    />
                  </div>
                  <Grid item xs={12}>
                    <div className={classes.paper}></div>
                  </Grid>
                  <div>
                  <FormControl className={classes.formControl}>
                    <InputLabel id="demo-simple-select-label">เลือกชมรม</InputLabel>
                    <Select
                      name="club"
                      value={clubid}
                      onChange={handleClubChange}
                    >
                    {clubs.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>
                          {item.name}
                        </MenuItem>
                      );
                    })}
                    </Select>
                  </FormControl>
                  </div>
                  <Grid item xs={12}>
                    <div className={classes.paper}></div>
                  </Grid>
                  <div>
                  <FormControl className={classes.formControl}>
                    <InputLabel id="demo-simple-select-label">เลือกหัวข้อ</InputLabel>
                    <Select
                      name="type"
                      value={complainttypeid}
                      onChange={handleComplaintTypeChange}
                    >
                    {complainttypes.map(item => {
                      return (
                        <MenuItem key={item.id} value={item.id}>
                          {item.description}
                        </MenuItem>
                      );
                    })}
                    </Select>
                  </FormControl>
                  </div>
                  {/* <Grid item xs={12}>
                    <div className={classes.paper}></div>
                  </Grid>
                  <div>
                  <form className={classes.textboxroot} noValidate autoComplete="off">
                    <div>
                    <TextField
                      error={nameError ? true : false}
                      helperText={nameError}
                      id="name-field"
                      label="ชื่อ-นามสกุลของผู้ร้องเรียน"
                      multiline
                      rows={1}
                      variant="outlined"
                      value={name}
                      onChange={handleNameChange}
                    />
                    </div>
                  </form>
                  </div> */}
                  <Grid item xs={12}>
                    <div className={classes.paper}></div>
                  </Grid>
                  <div>
                  <form className={classes.textboxroot} noValidate autoComplete="off">
                    <div>
                    <TextField
                      error={phoneNumberError ? true : false}
                      helperText={phoneNumberError}
                      id="phonenumber-field"
                      label="หมายเลขโทรศัพท์ (ในกรณีที่อาจจะต้องมีการติดต่อจากผู้ตรวจสอบเพื่อสืบสวนเพิ่มเติม)"
                      multiline
                      rows={1}
                      variant="outlined"
                      value={phonenumber}
                      onChange={handlePhoneNumberChange}
                    />
                    </div>
                  </form>
                  </div>
                  <Grid item xs={12}>
                    <div className={classes.paper}></div>
                  </Grid>
                  <div>
                  <form className={classes.textboxroot} noValidate autoComplete="off">
                    <div>
                    <TextField
                      error={infoError ? true : false}
                      helperText={infoError}
                      id="info-field"
                      label="รายละเอียด"
                      multiline
                      rows={20}
                      variant="outlined"
                      value={info}
                      onChange={handleInfoChange}
                    />
                    </div>
                  </form>
                  </div>
                  <Grid item xs={12}>
                    <div className={classes.paper}></div>
                  </Grid>
                  <div>
                    <Button
                      variant="contained"
                      color="secondary"
                      className={classes.button}
                      onClick={() => { 
                        save();
                      }}
                      startIcon={<AddIcon />}
                    >
                      ส่งเรื่องร้องเรียน
                    </Button>
                  </div>
                </Paper>
              </Grid>
            </Grid>
          </div>
        </Content>
      </Page>
    </SidebarPage>
  );
}
