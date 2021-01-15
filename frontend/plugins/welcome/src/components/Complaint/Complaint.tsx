import React, { useState, useEffect } from 'react';
import {
  pageTheme,
  Page,
  Header,
  Content,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles} from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import TextField from '@material-ui/core/TextField';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import Button from '@material-ui/core/Button';
import AddIcon from '@material-ui/icons/Add';
import Snackbar from '@material-ui/core/Snackbar';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import { EntUser } from '../../api/models/EntUser';
import { EntClub } from '../../api/models/EntClub';
import { EntComplaintType } from '../../api/models/EntComplaintType';

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

  const [users, setUsers] = React.useState<EntUser[]>([]);
  const [clubs, setClubs] = React.useState<EntClub[]>([]);
  const [complainttypes, setComplaintTypes] = React.useState<EntComplaintType[]>([]);

  const [date, setDate] = React.useState(String);
  const [userid, setUser] = React.useState(Number);
  const [clubid, setClub] = React.useState(Number);
  const [complainttypeid, setComplaintType] = React.useState(Number);
  const [info, setInfo] = React.useState(String);
  // const [textbox, setTextbox] = React.useState('Controlled');

  // const [status, setStatus] = React.useState(false);
  // const [alert, setAlert] = React.useState(true);
  // const [loading, setLoading] = React.useState(true);
  // const [KK, setKK] = useState(Number);

  const [status, SetStatus] = useState(Boolean);
  const [loading, setLoading] = useState(true);
  const [alert, setAlert] = useState(true);
  const [open, setOpen] = React.useState(false);

  // const [uID, setUID] = useState(localStorage.getItem('user-id')); 

  // let userID = Number(uID);

  const getUsers = async () => {
    const res = await api.listUser({ limit: 10, offset: 0 });
    setUsers(res);
    setLoading(false);
    console.log(res);
  };

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
    getUsers();
    getClubs();
    getTypes();
  }, [loading]);

  const handleInfoChange = (event: any) => {
    setInfo(event.target.value as string);
  };

  //const handleInfoChange = (event: React.ChangeEvent<HTMLInputElement>) => {
  //  setInfo(event.target.value);
  //};

  const handleDateChange = (event: any) => {
    setDate(event.target.value as string);
  };

  const handleUserChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUser(event.target.value as number);
  };

  const handleClubChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setClub(event.target.value as number);
  };

  const handleComplaintTypeChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setComplaintType(event.target.value as number);
  };

  function handleClose(){
    window.location.reload(false);
    setOpen(false);
  };

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
    //userID,
    userID: userid,
    clubID: clubid,
    typeID: complainttypeid,
    info: info,
    date: date + "T00:00:00+07:00",
  };
  console.log(complaint);

  function home() {
    history.pushState("", "", "/welcome");
    window.location.reload(false);
  }

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
        if (data.id != null) {
          setOpen(true);
          SetStatus(true);
          setAlert(true);
        } else {
          setOpen(true);
          SetStatus(true);
          setAlert(false);
        }
      })

  }

  return(
    <Page theme={pageTheme.home}>
      <Header title={`ระบบร้องเรียน`}></Header>
      <Content>
        <ContentHeader title="ส่งเรื่องร้องเรียน">
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
         {status ? (
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
          ) : null}
        </ContentHeader>
        <div className={classes.root}>
          <Grid container spacing={3}>
            <Grid item xs={12}>
              <Paper className={classes.paper}>
                <div>
                  <FormControl className={classes.formControl}>
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
                  </FormControl>
                </div>
                <Grid item xs={12}>
                  <div className={classes.paper}></div>
                </Grid>
                <div>
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
                <Grid item xs={12}>
                  <div className={classes.paper}></div>
                </Grid>
                <div>
                <form className={classes.textboxroot} noValidate autoComplete="off">
                  <div>
                  <TextField
                    id="info"
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
  );
}
