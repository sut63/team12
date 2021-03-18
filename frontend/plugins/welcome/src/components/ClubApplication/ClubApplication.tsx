import React, { FC, useEffect, useState } from 'react';
import { Content, Header, Page, pageTheme, SidebarPage } from '@backstage/core';
import { makeStyles, createStyles, Theme } from '@material-ui/core/styles'
import CssBaseline from '@material-ui/core/CssBaseline';
import Typography from '@material-ui/core/Typography';
import Container from '@material-ui/core/Container';
import Paper from '@material-ui/core/Paper';
import { Grid, TextField, MenuItem, Select, Button, Hidden, colors, } from '@material-ui/core';
import { DefaultApi } from '../../api/apis';
import { EntClub, EntDiscipline, EntGender, EntYear, EntUser, EntClubappStatus } from '../../api/models';
import { Alert } from '@material-ui/lab';
import Snackbar from '@material-ui/core/Snackbar';
import CreateIcon from '@material-ui/icons/Create';
import HomeIcon from '@material-ui/icons/Home';
import { UserHeader } from '../UserHeader/UserHeader';
import { AppSidebar } from '../Sidebar/Sidebar';

import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import WarningIcon from '@material-ui/icons/Warning';

// css header 
const HeaderCustom = { minHeight: '50px', };
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root_XL: { '& .MuiTextField-root': { margin: theme.spacing(1), width: 400, }, },
    root_L3: { '& > *': { margin: theme.spacing(0), width: '46ch', }, },
    root_L2: { '& > *': { margin: theme.spacing(0), width: '40ch', }, },
    root_L4: { '& > *': { margin: theme.spacing(0), width: '35ch', }, },
    root_L1: { '& > *': { margin: theme.spacing(0), width: '26ch', }, },
    root_L: { '& > *': { margin: theme.spacing(0), width: '30ch', }, },
    root_S2: { '& > *': { margin: theme.spacing(0), width: '15ch', }, },
    root_S: { '& > *': { margin: theme.spacing(0), width: '10ch', }, },
    
    root_B: { '& > *': { margin: theme.spacing(0), width: '10ch', }, },
    button: { margin: theme.spacing(1), },
    container1: { display: 'flex', flexWrap: 'wrap', },
    textField: { marginLeft: theme.spacing(1), marginRight: theme.spacing(1), width: 200, },
    container: { display: 'grid', gridTemplateColumns: 'repeat(12, 1fr)', gridGap: theme.spacing(2), },
    paper: { padding: theme.spacing(1), textAlign: 'center', color: theme.palette.text.secondary, whiteSpace: 'nowrap', marginBottom: theme.spacing(1), },
  }),
);

export default function CreateClubApplication() {

  const classes = useStyles();
  const api = new DefaultApi();
 
  const [checker, SetChecker] = useState(false);
  const [status, SetStatus] = useState(Boolean);
  const [loading, setLoading] = useState(true);
  const [alert, setAlert] = useState(true);
  const [open, setOpen] = React.useState(false);
  const [name, SetName] = useState(String);
  const [age, SetAge] = useState(Number);
  const [contact, SetContact] = useState(String);
  const [reason, SetReason] = useState(String);
  const [addedtime, SetAddedtime] = useState(String);
  const [clubID, SetClubID] = useState(Number);
  // const [userID, SetUserID] = useState(Number);
  const [years, SetYears] = useState(Number);
  const [disciplines, SetDisciplines] = useState(String);
  const [genders, SetGenders] = useState(String);
  // const [appstatusID, SetAppstatusID] = useState(Number);

  const [user, setUser] = useState<EntUser[]>([]);
  const [year, SetYear] = useState<EntYear[]>([]);
  const [club, SetClub] = useState<EntClub[]>([]);
  const [gender, SetGender] = useState<EntGender[]>([]);
  const [discipline, SetDiscipline] = useState<EntDiscipline[]>([]);
  // const [appstatus, SetAppstatus] = useState<EntClubappStatus[]>([]);
  //set UID from LocalStorage
  const [uID, setUID] = useState(localStorage.getItem('user-id'));
  let userID = Number(uID);
  let adderAge = Number(age);

  const [errorName, setErrorName] = useState(String);
  const [nameError, setNameError] = useState('');
  const [contactError, setContactError] = useState('');
  const [ageError, setAgeError] = useState('');

  const validatename = (val: string) => {
    return val.match("^[A-Za-zก-๙]+[ \t\r\n\v\f]+[A-Za-zก-๙]+[^๐-๙]$");
  }
  const validateage = (val: string) => {
    var a = Number(val);
    if (a > 0 && a < 200) {
      return String(val)
    }
    return val = ''
  }
  const validatecontact = (val: string) => {
    return val.match(".{10,}");
  }

  const checkPattern = (id: string, value: string) => {
    switch (id) {
      case 'addername':
        validatename(value) ? setNameError('') : setNameError('รูปแบบชื่อผิดพลาด กรุณาป้อน ชื่อ นามสกุล');
        return;
      case 'age':
        validateage(value) ? setAgeError('') : setAgeError('อายุไม่ถูกต้อง');
        return;
      case 'contact':
        validatecontact(value) ? setContactError('') : setContactError('กรุณาป้อนข้อมูลอย่างน้อย 10 ตัวอักษร');
        return;
    }
  }

  useEffect(() => {
    const getClub = async () => {
      const res = await api.listClub({ limit: 10, offset: 0 });
      setLoading(false);
      SetClub(res);
    }
    getClub();

    const getYear = async () => {
      const res = await api.listYear({ limit: 10, offset: 0 });
      setLoading(false);
      SetYear(res);
    }
    getYear();

    const getGender = async () => {
      const res = await api.listGender({ limit: 10, offset: 0 });
      setLoading(false);
      SetGender(res);
    };
    getGender();

    const getDiscipline = async () => {
      const res = await api.listDiscipline({ limit: 10, offset: 0 });
      setLoading(false);
      SetDiscipline(res);
    };
    getDiscipline();

    const getUser = async () => {
      const res = await api.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUser(res);
    };
    getUser();

    const getChecker = async () => {
      if(uStatus == "Have Club" && uType != "เจ้าหน้าที่"){
        SetChecker(true);
      }
    };
    getChecker();

    // const getAppstatus = async () => {
    //   const res = await api.listClubappstatus({ limit: 10, offset: 0 });
    //   setLoading(false);
    //   SetAppstatus(res);
    // };
    // getAppstatus();

  }, [loading]);

  // const UserIDhandleChage = (event: React.ChangeEvent<{ value: unknown }>) => {
  //   SetUserID(event.target.value as number);
  // };
  const DisciplineshandleChage = (event: React.ChangeEvent<{ value: unknown }>) => {
    SetDisciplines(event.target.value as string);
  };
  const ClubIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    SetClubID(event.target.value as number);
  };
  const YearshabdleChang = (event: React.ChangeEvent<{ value: unknown }>) => {
    SetYears(event.target.value as number);
  };
  const GendershabdleChang = (event: React.ChangeEvent<{ value: unknown }>) => {
    SetGenders(event.target.value as string);
  };
  // const AppstatusIDhabdleChang = (event: React.ChangeEvent<{ value: unknown }>) => {
  //   SetAppstatusID(event.target.value as number);
  // };

  const AddedtimehandleChange = (event: any) => {
    SetAddedtime(event.target.value as string);
  };
  const ReasonhandleChang = (event: any) => {
    SetReason(event.target.value as string);
  };
  const ContacthandleChang = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as keyof typeof CreateClubApplication
    const { value } = event.target;
    const validatecontact = value.toString()
    checkPattern(id, validatecontact)
    SetContact(event.target.value as string);
  };
  const AgehandleChang = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as keyof typeof CreateClubApplication
    const { value } = event.target;
    const validateage = value.toString()
    checkPattern(id, validateage)
    SetAge(event.target.value as number);
  };
  const NamehandleChang = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as keyof typeof CreateClubApplication
    const { value } = event.target;
    const validatename = value.toString()
    checkPattern(id, validatename)
    SetName(event.target.value as string);
  };

  function handleClose() {
    setOpen(false);
  };
  function handleCloseChecker() {
    SetChecker(false);
    home();
  };

  function handleCloseforSave() {
    window.location.reload(false);
    setOpen(false);
  };

//set UserData ..
var uType = JSON.parse(String(localStorage.getItem('user-type')));
var uStatus = JSON.parse(String(localStorage.getItem('user-status')));

  // const CreateClubApplication = async () => {
  const clubapplication = {
    userID,
    adderName: name,
    adderAge,
    adderGender: genders,
    discipline: disciplines,
    year: years,
    contact: contact,
    clubID: clubID,
    reason: reason,
    addedtime: addedtime + ":00+07:00",
    appstatusID: 1,
  };
  console.log(clubapplication);
  function home() {
    history.pushState("", "", "/welcome");
    window.location.reload(false);
  }

  

  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/clubapplications';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(clubapplication),
    };
    console.log(clubapplication);

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          setOpen(true);
          SetStatus(true);
          setAlert(true);

        } else {
          setErrorName(data.error.Name as string);
          SetStatus(true);
          setAlert(false);
          setOpen(true);
        }
      })
  }
  // const res: any = await api.createClubapplication({ clubapplication: clubapplication });
  // console.log("status",status);
  // const timer = setTimeout(() => {
  //   SetStatus(false);
  // }, 1000);
  // };
  return (
    <SidebarPage>
      <AppSidebar />
      <Page theme={pageTheme.home}>
        <Header style={HeaderCustom} title={`ระบบสมัครเข้าชมรม`}>
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
            คุณเป็นสมาชิกชมรมอยู่แล้ว<br/>หากต้องการสมัครชมรมอื่น กรุณาออกจากชมรมเดิมของคุณก่อน
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleCloseChecker} color="secondary" autoFocus variant="outlined">
            ฉันเข้าใจแล้ว
          </Button>
        </DialogActions>
      </Dialog>
        <Content>
              <React.Fragment>
                <CssBaseline />
                <Grid />
                {(() => {
                  if (status) {
                    if (alert) {
                      return (
                        <Grid>
                          <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
                            <Alert onClose={handleCloseforSave} severity="success">
                              ส่งคำร้องสำเร็จ
                            </Alert>
                          </Snackbar>
                        </Grid>
                      )
                    } else {
                      switch (errorName) {
                        case 'addername':
                          return (
                            <Grid>
                            <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
                              <Alert onClose={handleClose} severity="error">
                              รูปแบบชื่อผิดพลาด กรุณาป้อน ชื่อ วรรค นามสกุล อย่างน้อย 1 ตัว
                              
                              </Alert>
                            </Snackbar>
                            </Grid>
                          )
                        case 'contact':
                          return (
                            <Grid>
                            <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
                              <Alert onClose={handleClose} severity="error">
                              ข้อมูลติดต่อผิดพลาด กรุณาป้อนข้อมูลอย่างน้อย 10 ตัวอักษร
                              </Alert>
                            </Snackbar>
                            </Grid>
                          )
                        case 'age':
                          return (
                            <Grid>
                            <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
                              <Alert onClose={handleClose} severity="error">
                              อายุไม่ถูกต้อง
                              </Alert>
                            </Snackbar>
                            </Grid>
                          )
                        default:
                          return (
                            <Grid>
                            <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
                              <Alert onClose={handleClose} severity="error">
                                ส่งคำร้องไม่สำเร็จ กรุณาลองใหม่อีกครั้ง
                              </Alert>
                            </Snackbar>
                            </Grid>
                          )
                      }
                    }
                  } else {
                    return null
                  }
                })()}
                <Container maxWidth="sm">
                  
                  {/* <Grid item sm={12}>
                  <Hidden smDown>
                    <Paper className={classes.paper}></Paper>
                  </Hidden>
                </Grid> */}
                <Typography variant="h5" gutterBottom style={{ height: '1.5cm'}}>
                    ใบคำร้องขอสมัครชมรม
              </Typography>
                  <Grid container spacing={4} style={{backgroundColor: "white"}}>
                  <Grid item sm={12}  style={{ backgroundColor: "#217567"}}/>
                  
                
                    {/*<Grid item sm={2}>
                    <div className={classes.paper}>เลือก Email</div>
                  </Grid>
                  <Grid item sm={10}>
                    <form className={classes.root_L2} noValidate autoComplete="off">
                      <Select
                        id="email"
                        value={userID}
                        onChange={UserIDhandleChage}
                      >
                        {user.map((item: EntUser) => (
                          <MenuItem value={item.id}>{item.email}</MenuItem>
                        ))}
                      </Select>
                    </form>
                  </Grid> */}
                    <Grid item sm={2}>
                      <div className={classes.paper}>ชื่อ-นามสกุล</div>
                    </Grid>
                    <Grid item sm={5}>
                      <form className={classes.root_L} noValidate autoComplete="off">
    
                        <TextField
                          error={nameError ? true : false}
                          id="addername"
                          size="medium"
                          helperText={nameError}
                          value={name}
                          onChange={NamehandleChang}
                        />
                      </form>
                    </Grid>
    
                    <Grid item sm={1}>
                      <div className={classes.paper}> อายุ</div>
                    </Grid>
                    <Grid item sm={4}>
                      <form className={classes.root_S} noValidate autoComplete="off">
    
                        <TextField
                          error={ageError ? true : false}
                          id="age"
                          size="medium"
                          helperText={ageError}
                          value={age}
                          onChange={AgehandleChang}
                        />
                      </form>
                    </Grid>
    
                    <Grid item sm={1}>
                      <div className={classes.paper}> เพศ</div>
                    </Grid>
                    <Grid item sm={2}>
                      <form className={classes.root_B} noValidate autoComplete="off">
                        <Select
                          labelId="demo-simple-select-label"
                          id="gender"
                          value={genders}
                          onChange={GendershabdleChang}
                        >
                          {gender.map((item: EntGender) => (
                            <MenuItem value={item.gender}>{item.gender}</MenuItem>
                          ))}
                        </Select>
                      </form>
                    </Grid>
    
                    <Grid item sm={1}>
                      <div className={classes.paper}>สาขา</div>
                    </Grid>
                    <Grid item sm={5}>
                      <form className={classes.root_L1} noValidate autoComplete="off">
                        <Select
                          labelId="demo-simple-select-label"
                          id="discipline"
                          value={disciplines}
                          onChange={DisciplineshandleChage}
                        >
                          {discipline.map((item: EntDiscipline) => (
                            <MenuItem value={item.discipline}>{item.discipline}</MenuItem>
                          ))}
                        </Select>
                      </form>
                    </Grid>
    
                    <Grid item sm={1}>
                      <div className={classes.paper}> ชั้นปี</div>
                    </Grid>
                    <Grid item sm={2}>
                      <form noValidate autoComplete="off">
                        <Select
                          labelId="demo-simple-select-label"
                          id="year"
                          value={years}
                          onChange={YearshabdleChang}
                        >
                          {year.map((item: EntYear) => (
                            <MenuItem value={item.year}>{item.year}</MenuItem>
                          ))}
                        </Select>
                      </form>
                    </Grid>
    
                    <Grid item sm={2}>
                      <div className={classes.paper}>ข้อมูลติดต่อ</div>
                    </Grid>
                    <Grid item sm={10}>
                      <form className={classes.root_XL} noValidate autoComplete="off">
                        <TextField
                          error={contactError ? true : false}
                          id="contact"
                          label="contact"
                          multiline
                          rows={2}
                          defaultValue="ที่อยู่: เบอร์โทรศัพท์: "
                          variant="outlined"
                          helperText={contactError}
                          value={contact}
                          onChange={ContacthandleChang}
                        />
                      </form>
                    </Grid>
                    <Grid item sm={12}  style={{ backgroundColor: "#217567"}}/>
                    <Grid item sm={4}>
                      <div className={classes.paper}>ชมรมที่ต้องการสมัครสมาชิก</div>
                    </Grid>
                    <Grid item sm={8}>
                      <form className={classes.root_L4} noValidate autoComplete="off">
                        <Select
                          labelId="demo-simple-select-label"
                          id="club"
                          value={clubID}
                          onChange={ClubIDhandleChange}
                        >
                          {club.map((item: EntClub) => (
                            <MenuItem value={item.id}>{item.name}</MenuItem>
                          ))}
                        </Select>
                      </form>
                    </Grid>
                    <Grid item sm={3}>
                      <div className={classes.paper}>เหตุผลการเข้าร่วม</div>
                    </Grid>
                    <Grid item sm={9}>
                      <form className={classes.root_L3} noValidate autoComplete="off">
                        <TextField
                          id="outlined-multiline-static"
                          label="reason"
                          multiline
                          rows={2}
                          defaultValue="ระบุเหตุผล"
                          variant="outlined"
                          value={reason}
                          onChange={ReasonhandleChang}
                        />
                      </form>
                    </Grid>
    
                    <Grid item sm={4}>
                      <div className={classes.paper}>วันและเวลาที่กรอกใบคำร้อง</div>
                    </Grid>
                    <Grid item sm={8}>
                      <form className={classes.container1} noValidate>
                        <TextField
                          id="datetime-local"
                          label="Month/Day/Year Added Time"
                          type="datetime-local"
                          defaultValue="2017-05-24T10:30"
                          value={addedtime}
                          onChange={AddedtimehandleChange}
                          className={classes.textField}
                          InputLabelProps={{
                            shrink: true,
                          }}
                        />
                      </form>
                    </Grid>
                    <Grid item sm={12}  style={{ backgroundColor: "#217567"}}/>
                    {/* <Grid item sm={3}>
                    <div className={classes.paper}>สถานะใบคำร้อง</div>
                  </Grid>
                  <Grid item sm={5}>
                    <form className={classes.root_S2} noValidate autoComplete="off">
                      <Select
                        id="application status"
                        value={appstatusID}
                        onChange={AppstatusIDhabdleChang}
                      >
                        {appstatus.map((item: EntClubappStatus) => (
                          <MenuItem value={item.id}>{item.clubstatus}</MenuItem>
                        ))}
                      </Select>
                    </form>
                  </Grid> */}
                  </Grid>
                  <Grid container spacing={4} style={{marginTop: "1cm"}}>
                  <Grid item sm={1}/>
                    <Grid item sm={5}
                      container
                      direction="column"
                      justify="center"
                      alignItems="center">
                      <Button
                        variant="contained"
                        style={{backgroundColor:"#00cf60", color:"#ffffff"}}
                        size="medium"
                        className={classes.button}
                        startIcon={<CreateIcon />}
                        onClick={() => {
                          save();
                        }}
                      >
                        <h3>บันทึกใบสมัคร</h3>
                      </Button>
                    </Grid >
                    <Grid item sm={5}
                      container
                      direction="column"
                      justify="center"
                      alignItems="center">
                      <Button
                        variant="contained"
                        style={{backgroundColor:"#db1222", color:"#ffffff"}}
                        size="medium"
                        className={classes.button}
                        startIcon={<HomeIcon />}
                        onClick={() => {
                          home()
                        }}
                      >
                        <h3>กลับหน้า Home</h3>
                      </Button>
                    </Grid>
                    <Grid item sm={1}/>
                  </Grid>
                </Container>
              </React.Fragment>
            </Content>
      </Page>
    </SidebarPage>
  );
}