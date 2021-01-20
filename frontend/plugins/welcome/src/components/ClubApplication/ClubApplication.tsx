import React, { FC, useEffect, useState } from 'react';
import { Content, Header, Page, pageTheme , SidebarPage } from '@backstage/core';
import { makeStyles, createStyles, Theme } from '@material-ui/core/styles'
import CssBaseline from '@material-ui/core/CssBaseline';
import Typography from '@material-ui/core/Typography';
import Container from '@material-ui/core/Container';
import Paper from '@material-ui/core/Paper';
import { Grid, TextField, MenuItem, Select, Button, Hidden, } from '@material-ui/core';
import { DefaultApi } from '../../api/apis';
import { EntClub, EntDiscipline, EntGender, EntYear, EntUser, EntClubappStatus } from '../../api/models';
import { Alert } from '@material-ui/lab';
import Snackbar from '@material-ui/core/Snackbar';
import CreateIcon from '@material-ui/icons/Create';
import HomeIcon from '@material-ui/icons/Home';
import { UserHeader } from '../UserHeader/UserHeader';
import { AppSidebar } from '../Sidebar/Sidebar';

// css header 
const HeaderCustom = { minHeight: '50px', };
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root_XL: { '& .MuiTextField-root': { margin: theme.spacing(1), width: 400, }, },
    root_L3: { '& > *': { margin: theme.spacing(0), width: '46ch', }, },
    root_L2: { '& > *': { margin: theme.spacing(0), width: '40ch', }, },
    root_L4: { '& > *': { margin: theme.spacing(0), width: '35ch', }, },
    root_L: { '& > *': { margin: theme.spacing(0), width: '36ch', }, },
    root_L1: { '& > *': { margin: theme.spacing(0), width: '26ch', }, },
    root_S2: { '& > *': { margin: theme.spacing(0), width: '15ch', }, },
    root_S: { '& > *': { margin: theme.spacing(0), width: '5ch', }, },
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
  const ContacthandleChang = (event: any) => {
    SetContact(event.target.value as string);
  };
  const AgehandleChang = (event: any) => {
    SetAge(event.target.value as number);
  };
  const NamehandleChang = (event: any) => {
    SetName(event.target.value as string);
  };

  function handleClose() {
    window.location.reload(false);
    setOpen(false);
  };


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
        <Content>
          <React.Fragment>
            <CssBaseline />
            <Grid />
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
            <Container maxWidth="sm">
              <Typography variant="h5" gutterBottom style={{ height: '2cm' }}>
                ใบคำร้องขอสมัครชมรม
          </Typography>
              {/* <Grid item sm={12}>
              <Hidden smDown>
                <Paper className={classes.paper}></Paper>
              </Hidden>
            </Grid> */}
              <Grid container spacing={4}>
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
                <Grid item sm={6}>
                  <form className={classes.root_L} noValidate autoComplete="off">

                    <TextField
                      //type="reset"
                      id="addername"
                      size="medium"
                      value={name}
                      onChange={NamehandleChang}
                    />
                  </form>
                </Grid>

                <Grid item sm={1}>
                  <div className={classes.paper}> อายุ</div>
                </Grid>
                <Grid item sm={3}>
                  <form className={classes.root_S} noValidate autoComplete="off">

                    <TextField
                      id="age"
                      size="medium"
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
                  <form className={classes.root_S} noValidate autoComplete="off">
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
                      id="outlined-multiline-static"
                      label="contact"
                      multiline
                      rows={2}
                      defaultValue="ที่อยู่: เบอร์โทรศัพท์: "
                      variant="outlined"
                      value={contact}
                      onChange={ContacthandleChang}
                    />
                  </form>
                </Grid>
                <Grid item sm={12}>
                  <Hidden smDown>
                    <Paper className={classes.paper}></Paper>
                  </Hidden>
                </Grid>
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
                <Grid item sm={12}>
                  <Hidden smDown>
                    <Paper className={classes.paper}></Paper>
                  </Hidden>
                </Grid>
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
                <Grid item sm={1}></Grid>
                <Grid item sm={5}
                  container
                  direction="column"
                  justify="center"
                  alignItems="center">
                  <Button
                    variant="contained"
                    color="primary"
                    size="medium"
                    className={classes.button}
                    startIcon={<CreateIcon />}
                    onClick={() => {
                      save();
                    }}
                  >
                    <h3>บันทึกใบสมัคร</h3>
                  </Button>
                </Grid>
                <Grid item sm={5}
                  container
                  direction="column"
                  justify="center"
                  alignItems="center">
                  <Button
                    variant="contained"
                    color="secondary"
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
                <Grid item sm={1}></Grid>
              </Grid>
            </Container>
          </React.Fragment>
        </Content>
      </Page>
    </SidebarPage>
  );
}