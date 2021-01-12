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
import { EntActivityType } from '../../api/models/EntActivityType';
import { EntAcademicYear } from '../../api/models/EntAcademicYear';
import { EntClub } from '../../api/models/EntClub';
import { EntActivities } from '../../api/models/EntActivities';
import { DefaultApi } from '../../api/apis';

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
    const Activities = {
      academicYearID: AcademicyearID,
      activityTypeID: ActivitytypeID,
      clubID: UserID,
      detail: detail,
      endTime: endtime + ':00+07:00',
      name: name,
      starttime: starttime + ':00+07:00',
    };
    console.log(Activities);
    const res: any = await api.createActivities({
      activities: Activities,
    });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }

    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
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
        <Avatar alt="Remy Sharp" src="../../image/account.jpg" />
        <div style={{ marginLeft: 10 }}></div>
      </Header>
      <Content>
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
              <div className={classes.paper}>ผู้ใช้งาน</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel id="demo-mutiple-name-label">SUPHASIN</InputLabel>
                <Select
                  name="user"
                  value={UserID} // (undefined || '') = ''
                  onChange={user_id_handleChange}
                  style={{ width: 350 }}
                >
                  {users.map((item: EntUser) => (
                    <MenuItem value={item.id}>{item.email}</MenuItem>
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