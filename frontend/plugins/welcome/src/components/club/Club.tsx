import React, { useState, useEffect } from 'react';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import { makeStyles } from '@material-ui/core/styles';
import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Avatar,
  Button,
} from '@material-ui/core';
import { EntUser } from '../../api/models/EntUser';
import { EntClubBranch } from '../../api/models/EntClubBranch';
import { EntClubType } from '../../api/models/EntClubType';
import { DefaultApi } from '../../api/apis';
import Swal from 'sweetalert2'; // alert

import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import WarningIcon from '@material-ui/icons/Warning';

// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(1),
    marginBottom: theme.spacing(1),
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
}));

function Clubs() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [users, setUsers] = useState<EntUser[]>([]);
  const [clubbranchs, setClubBranch] = useState<EntClubBranch[]>([]);
  const [clubtypes, setClubType] = useState<EntClubBranch[]>([]);
  const [checker, SetChecker] = useState(false);
  const [open, setOpen] = React.useState(false);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);
  const [UserID, setUserID] = useState(Number);
  const [ClubBranchID, setClubBranchID] = useState(Number);
  const [ClubTypeID, setClubTypeID] = useState(Number);
  const [Name, setName] = React.useState(String);
  const [Purpose, setPurpose] = React.useState(String);
  const [Phone, setPhone] = React.useState(String);

  const [nameError, setNameError] = React.useState('');
  const [purposeError, setPurposeError] = React.useState('');
  const [phoneError, setPhoneError] = React.useState('');

  useEffect(() => {
    const getClubBranch = async () => {
      const cb = await api.listClubbranch({ limit: 10, offset: 0 });
      setLoading(false);
      setClubBranch(cb);
    };
    getClubBranch();

    const getUsers = async () => {
      const us = await api.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUsers(us);
    };
    getUsers();

    const getClubType = async () => {
      const ct = await api.listClubtype({ limit: 10, offset: 0 });
      setLoading(false);
      setClubType(ct);
    };
    getClubType();

    const getChecker = async () => {
      if(uType == "เจ้าหน้าที่"){
        SetChecker(true);
      }
    };
    getChecker();

  }, [loading]);
  
  var uType = JSON.parse(String(localStorage.getItem('user-type')));

  
  const CreateClub = async () => {
    const Clubs = {
      userID: UserID,
      clubBranchID: ClubBranchID,
      clubTypeID: ClubTypeID,
      name: Name,
      purpose: Purpose,
      phone: Phone,
    };
    console.log(Clubs);
    const res: any = await api.createClub({
      club : Clubs,
    });
    
    setStatus(true);
    if (res.id != '') {
      Toast.fire({
        icon: 'success',
        title: 'บันทึกข้อมูลสำเร็จ',
      });
    } else {
      Toast.fire({
        icon: 'success',
        title: 'บันทึกไม่ข้อมูลสำเร็จ',
      });
    }
    
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  const checkPattern  = (id: string, value: string) => {
    switch(id) {
      case 'name':
        validateName(value) ? setNameError('ชื่อชมรมต้องมากกว่า 3 ตัวอักษร') : setNameError('');
        return;
      case 'purpose':
        validatePurpose(value) ? setPurposeError('จุดประสงค์น้อยเกินไป') : setPurposeError('');
        return;
      case 'phone':
        validatePhone(value) ? setPhoneError('') : setPhoneError('เบอร์โทรศัทพ์ 10 หลัก')
        return;
      default:
        return;
    }
  }

  const ClubType_id_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setClubTypeID(event.target.value as number);
  };

  const ClubBranch_id_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setClubBranchID(event.target.value as number);
  };

  const user_id_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setUserID(event.target.value as number);
  };

  const name_handleChange = (event: React.ChangeEvent<{ id?: string; value: any  }>) => {
    const id = event.target.id as keyof typeof Clubs;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern('name', validateValue)
    setName(event.target.value as string);
  };

  const purpose_handleChange = (event: React.ChangeEvent<{ id?: string; value: any  }>,) => {
    const id = event.target.id as keyof typeof Clubs;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern('purpose', validateValue)
    setPurpose(event.target.value as string);
  };

  const phone_handleChange = (event: React.ChangeEvent<{ id?: string; value: any  }>) => {
    const id = event.target.id as keyof typeof Clubs;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern('phone', validateValue)
    setPhone(event.target.value as string);
  };

  const validateName = (val: string) => {
    return val.length < 3 ? true : false;
  }

  const validatePurpose = (val: string) => {
    return val.length < 5 ? true : false;
  }

  const validatePhone = (val: string) => {
    return val.length == 10  ? true : false;
  }

 
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 4000,
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

  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'name':
        alertMessage("error","ชื่อชมรมต้องมากกว่า 3 ตัวอักษร");
        return;
      case 'purpose':
        alertMessage("error","จุดประสงค์น้อยเกินไป");
        return;
      case 'phone':
        alertMessage("error","เบอร์โทรศัทพ์ 10 หลัก");
        return;
      default:
        alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }

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


  console.log(Name);
  console.log(ClubBranchID);
  console.log(ClubTypeID);
  console.log(UserID);
  console.log(Purpose);
  console.log(Phone);

  return (
    <Page theme={pageTheme.website}>
      <Header style={HeaderCustom} title={`สร้างชมรม`}>
      
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
            เฉพาะเจ้าหน้าที่<br/>หากต้องการสร้างชมรม กรุณาติดต่อเจ้าหน้าที่
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseChecker} color="secondary" autoFocus variant="outlined">
            ฉันเข้าใจแล้ว
          </Button>
        </DialogActions>
      </Dialog>


        <Avatar alt="Remy Sharp" src="../../image/account.jpg" />
        <div style={{ marginLeft: 10 }}></div>
      </Header>
      <Content>
        <Container maxWidth="sm">


          <Grid container spacing={2}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ชื่อชมรม</div>
            </Grid>
            <Grid item xs={9}>
              <TextField
                error = {nameError ? true : false}
                id="filled-basic"
                label="ใส่ชื่อชมรม"
                inputProps={{ maxLength: 13 }}
                helperText= {nameError}
                variant="filled"
                value={Name}
                style={{ width: 350 }}
                onChange={name_handleChange}
              />
            </Grid>


            <Grid item xs={3}>
              <div className={classes.paper}>ประเภทชมรม</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel id="demo-mutiple-name-label">
                  เลือกประเภทของชมรม
                </InputLabel>
                <Select
                  name="clubtype"
                  value={ClubTypeID} // (undefined || '') = ''
                  onChange={ClubType_id_handleChange}
                  style={{ width: 350 }}
                >
                  {clubtypes.map((item: EntClubType) => (
                    <MenuItem value={item.id}>{item.name}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>


            <Grid item xs={3}>
              <div className={classes.paper}>สถานที่จัดชมรม</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel id="demo-mutiple-name-label">
                  เลือกสถานที่จัดชมรม
                </InputLabel>
                <Select
                  name="ClubBranch"
                  value={ClubBranchID} // (undefined || '') = ''
                  onChange={ClubBranch_id_handleChange}
                  style={{ width: 350 }}
                >
                  {clubbranchs.map((item: EntClubBranch) => (
                    <MenuItem value={item.id}>{item.name}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>


            <Grid item xs={3}>
              <div className={classes.paper}>ผู้บันทึก</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel id="demo-mutiple-name-label">เจ้าหน้าที่</InputLabel>
                <Select
                  name="user"
                  value={UserID} // (undefined || '') = ''
                  onChange={user_id_handleChange}
                  style={{ width: 350 }}
                >
                  {users.map((item: EntUser) => (
                    <MenuItem value={item.id}>{item.name}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>


            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>จุดประสงค์ของชมรม</div>
            </Grid>
            <Grid item xs={9}>
              <TextField
              error = {purposeError ? true : false}
                id="filled-basic"
                label="ใส่จุดประสงค์"
                inputProps={{ maxLength: 13 }}
                helperText= {purposeError}
                variant="filled"
                value={Purpose}
                style={{ width: 350 }}
                onChange={purpose_handleChange}
              />
            </Grid>


            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>เบอร์ติดต่อ</div>
            </Grid>
            <Grid item xs={9}>
              <TextField
              error = {phoneError ? true : false}
                id="filled-basic"
                label="เบอร์โทรศัพท์"
                inputProps={{ maxLength: 13 }}
                helperText= {phoneError}
                variant="filled"
                value={Phone}
                style={{ width: 350 }}
                onChange={phone_handleChange}
              />
            </Grid>


            <Grid item xs={1}>
              <Button
                variant="contained"
                color="primary"
                onClick={() => {
                  CreateClub();
                }}
                style={{
                  marginLeft: 290,
                  width: 200,
                  height: 50,
                  marginTop: 0,
                }}
              >
                บันทึกกิจกรรม
              </Button>
            </Grid>
          </Grid>

        </Container>
      </Content>
    </Page>
  );
}

export default Clubs;
