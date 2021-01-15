import React, { useState, useEffect } from 'react';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles } from '@material-ui/core/styles';
import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  Link,
  TextField,
  Avatar,
  Button,
} from '@material-ui/core';
import { EntUser } from '../../api/models/EntUser';
import { EntActivityType } from '../../api/models/EntActivityType';
import { EntAcademicYear } from '../../api/models/EntAcademicYear';
import { EntClub } from '../../api/models/EntClub';
import { EntActivities } from '../../api/models/EntActivities';
import { DefaultApi } from '../../api/apis';
import { Link as RouterLink } from 'react-router-dom';
import Alert from '@material-ui/lab/Alert';

//set UID from LocalStorage
//const [uID, setUID] = useState(localStorage.getItem('user-id'));

let userID = Number(localStorage.getItem('user-id'));
let Email = localStorage.getItem('user-email');

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
  alert: {
    width: '100%',
    '& > * + *': {
      marginTop: theme.spacing(2),
    },
  },
  user: {
    background: 'linear-gradient(45deg, #FE6B8B 30%, #FF8E53 90%)',
    border: 0,
    borderRadius: 3,
    boxShadow: '0 3px 5px 2px rgba(255, 105, 135, .3)',
    color: 'white',
    height: 48,
    padding: '0 30px',
  },
}));

function Activities() {
  const classes = useStyles();
  /*   const profile = { givenName: 'ระบบโปรโมชัน' }; */
  const api = new DefaultApi();
  const [users, setUsers] = useState<EntUser[]>([]);
  const [activitytypes, setActivitytypes] = useState<EntActivityType[]>([]);
  const [academicyears, setAcademicyears] = useState<EntAcademicYear[]>([]);
  const [clubs, setClubs] = useState<EntClub[]>([]);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);

  const [loading, setLoading] = useState(true);
  const [UserID, setUserID] = useState(Number);
  const [KK, setKK] = useState(Number);
  const [ActivitytypeID, setActivitytypeID] = useState(Number);
  const [AcademicyearID, setAcademicyearID] = useState(Number);
  const [name, setName] = React.useState(String);
  const [detail, setDetail] = React.useState(String);
  const [starttime, setStarttime] = useState(String);
  const [endtime, setEndtime] = useState(String);

  useEffect(() => {
    const getAcademicYear = async () => {
      const acy = await api.listAcademicYear({ limit: 10, offset: 0 });
      setLoading(false);
      setAcademicyears(acy);
    };
    getAcademicYear();

    const getUsers = async () => {
      const us = await api.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUsers(us);
    };
    getUsers();

    const getActivityType = async () => {
      const act = await api.listActivityType({ limit: 10, offset: 0 });
      setLoading(false);
      setActivitytypes(act);
    };
    getActivityType();
  }, [loading]);

  const CreateActivities = async () => {
    if (
      AcademicyearID != null &&
      ActivitytypeID != null &&
      UserID != null &&
      detail != '' &&
      name != '' &&
      starttime != '' &&
      endtime != ''
    ) {
      const Activities = {
        academicYearID: AcademicyearID,
        activityTypeID: ActivitytypeID,
        clubID: userID,
        detail: detail,
        endtime: endtime + ':00+07:00',
        name: name,
        starttime: starttime + ':00+07:00',
      };

      console.log(Activities);
      const res: any = await api.createActivities({
        activities: Activities,
      });
      setStatus(true);
      setKK(res.id);
      if (res.id != '') {
        setAlert(true);
      }
    } else {
      setStatus(true);
      setAlert(false);
    }

    const timer = setTimeout(() => {
      setStatus(false);
    }, 5000);
  };

  const handleStarttimeChange = (event: any) => {
    setStarttime(event.target.value as string);
  };

  const handleEndtimeChange = (event: any) => {
    setEndtime(event.target.value as string);
  };

  const activitytype_id_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setActivitytypeID(event.target.value as number);
  };

  const academicyear_id_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setAcademicyearID(event.target.value as number);
  };

  const user_id_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setUserID(event.target.value as number);
  };

  const name_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setName(event.target.value as string);
  };

  const detail_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setDetail(event.target.value as string);
  };

  return (
    <Page theme={pageTheme.website}>
      <Header style={HeaderCustom} title={`กิจกรรม`}>
        {/* <Avatar alt="Remy Sharp" src="../../image/account.jpg"></Avatar> */}
        <div
          style={{
            background: 'linear-gradient(45deg, #FE6B8B 30%, #FF8E53 90%)',
            height: 50,
            width: 200,
            margin: 0,
            alignItems: 'center',
          }}
        >
          {Email}
        </div>
        <div style={{ marginLeft: 10 }}></div>
      </Header>
      <Content>
        <ContentHeader title="">
          {status ? (
            <div style={{ margin: 0, position: 'relative', width: 1500 }}>
              {alert ? (
                <Alert severity="success">บันทึกสำเร็จ </Alert>
              ) : (
                <Alert severity="error">บันทึกไม่สำเร็จ</Alert>
              )}
            </div>
          ) : null}
          {/*  <Link component={RouterLink} to="/ActivityTable">
            <Button variant="contained" color="primary">
              ตารางข้อมูลกิจกรรม
            </Button>
          </Link> */}
        </ContentHeader>
        <Container maxWidth="sm">
          <Grid container spacing={2}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>กิจกรรม</div>
            </Grid>
            <Grid item xs={9}>
              <TextField
                id="filled-basic"
                label="ใส่ชื่อกิจกรรม"
                variant="filled"
                value={name}
                style={{ width: 350 }}
                onChange={name_handleChange}
              />
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>รายละเอียดกิจกรรม</div>
            </Grid>
            <Grid item xs={9}>
              <TextField
                id="filled-multiline-static"
                label="ใส่รายละเอียดกิจกรรม"
                value={detail}
                multiline
                rows={2}
                defaultValue=""
                variant="filled"
                style={{ width: 350, height: 50, marginBottom: 25 }}
                onChange={detail_handleChange}
              />
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ประเภทกิจกรรม</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel id="demo-mutiple-name-label">
                  เลือกประเภทของกิจกรรม
                </InputLabel>
                <Select
                  id="type"
                  name="activitytype"
                  value={ActivitytypeID} // (undefined || '') = ''
                  onChange={activitytype_id_handleChange}
                  style={{ width: 350 }}
                >
                  {activitytypes.map((item: EntActivityType) => (
                    <MenuItem value={item.id}>{item.name}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ภาคการศึกษา</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel id="demo-mutiple-name-label">
                  เลือกภาคการศึกษา
                </InputLabel>
                <Select
                  id="acadamicyear"
                  name="AcademicYear"
                  value={AcademicyearID} // (undefined || '') = ''
                  onChange={academicyear_id_handleChange}
                  style={{ width: 350 }}
                >
                  {academicyears.map((item: EntAcademicYear) => (
                    <MenuItem value={item.id}>{item.semester}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>เลือกเวลาเริ่มต้น</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  id="startdate"
                  type="datetime-local"
                  value={starttime}
                  onChange={handleStarttimeChange}
                  //defaultValue="2017-05-24"
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  style={{ width: 350 }}
                />
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>เลือกเวลาสิ้นสุด</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField
                  id="enddate"
                  type="datetime-local"
                  value={endtime}
                  onChange={handleEndtimeChange}
                  //defaultValue="2017-05-24"
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  style={{ width: 350 }}
                />
              </FormControl>
            </Grid>

            <Grid item xs={1}>
              <Button
                variant="contained"
                color="primary"
                onClick={() => {
                  CreateActivities();
                }}
                style={{
                  marginLeft: 100,
                  width: 200,
                  height: 60,
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

export default Activities;
