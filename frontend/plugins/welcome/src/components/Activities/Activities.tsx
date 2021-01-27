import React, { useState, useEffect } from 'react';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  SidebarPage,
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
import { UserHeader } from '../UserHeader/UserHeader';
import { AppSidebar } from '../Sidebar/Sidebar';
import Swal from 'sweetalert2'; // alert

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

  const [activities, setActivities] = React.useState<Partial<EntActivities>>(
    {},
  );
  const [activitytypes, setActivitytypes] = useState<EntActivityType[]>([]);
  const [academicyears, setAcademicyears] = useState<EntAcademicYear[]>([]);

  const [NameError, setNameError] = React.useState('');
  const [DetailError, setDetailError] = React.useState('');
  const [LocationError, setLocationError] = React.useState('');

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);

  const [loading, setLoading] = useState(true);
  const [UserID, setUserID] = useState(Number);
  const [KK, setKK] = useState(Number);
  const [ActivitytypeID, setActivitytypeID] = useState(Number);
  const [AcademicyearID, setAcademicyearID] = useState(Number);
  const [name, setName] = React.useState(String);
  const [location, setLocation] = React.useState(String);
  const [detail, setDetail] = React.useState(String);
  const [starttime, setStarttime] = useState(String);
  const [endtime, setEndtime] = useState(String);

  useEffect(() => {
    const checkUserType = async () => {
      const userType = JSON.parse(
        String(localStorage.getItem('user-position')),
      );
      setLoading(false);
      const check1 = 'กรรมการ';
      const check2 = 'ประธาน';
      const check3 = 'รองประธาน';
      const check4 = 'เลขา';
      if (userType != (check1 || check2 || check3 || check4)) {
        Swal.fire({
          title: 'สถานะของผู้ใช้ระบบไม่สามารถจัดกิจกรรมได้',
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
      } /* else {
        setUser(Number(localStorage.getItem('userdata')));
      } */
    };
    checkUserType();

    const getAcademicYear = async () => {
      const acy = await api.listAcademicYear({ limit: 10, offset: 0 });
      setLoading(false);
      setAcademicyears(acy);
    };
    getAcademicYear();

    const getActivityType = async () => {
      const act = await api.listActivityType({ limit: 10, offset: 0 });
      setLoading(false);
      setActivitytypes(act);
    };
    getActivityType();
  }, [loading]);

  // ฟังก์ชั่นสำหรับ validate ชื่อกิจกรรม
  const validateName = (val: string) => {
    return val.match('^[A-Z][A-Za-z_]{6,50}$');
  };

  // ฟังก์ชั่นสำหรับ validate รายละเอียดกิจกรรม
  const validateDetail = (val: string) => {
    return val.match('.{49,}');
  };

  // ฟังก์ชั่นสำหรับ validate ชื่อกิจกรรม
  const validateLocation = (val: string) => {
    return val.match('^[A-Z][A-Za-z_0-9]{2,}$');
  };

  // สำหรับตรวจสอบรูปแบบข้อมูลที่กรอก ว่าเป็นไปตามที่กำหนดหรือไม่
  const checkPattern = (id: string, value: string) => {
    switch (id) {
      case 'name':
        validateName(value)
          ? setNameError('')
          : setNameError(
              'ชื่อกิจกรรมต้องขึ้นต้นด้วยตัวอักษร [A-Z] มีอย่างน้อย 7 ตัว ไม่เกิน 35 ตัว และห้ามเป็นตัวเลข [0-9]',
            );
        return;
      case 'detail':
        validateDetail(value)
          ? setDetailError('')
          : setDetailError('รายละเอียดกิจกรรมต้องมีตัวอักษรอย่างน้อย 50 ตัว');
        return;
      case 'location':
        validateLocation(value)
          ? setLocationError('')
          : setLocationError(
              'ชื่อสถานที่จัดกิจกรรมต้องขึ้นต้นด้วยตัวอักษร [A-Z] มีอย่างน้อย 3 ตัว และไม่เกิน 50 ตัว',
            );
        return;
      default:
        return;
    }
  };
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

  // alert error
  const aleartMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  };
  // check error
  const checkCaseSaveError = (field: string) => {
    switch (field) {
      case 'name':
        aleartMessage(
          'error',
          'บันทึกข้อมูลไม่สำเร็จ : ชื่อกิจกรรมต้องขึ้นต้นด้วยตัวอักษร [A-Z] มีอย่างน้อย 7 ตัว ไม่เกิน 35 ตัว และห้ามเป็นตัวเลข [0-9]',
        );
        return;
      case 'detail':
        aleartMessage(
          'error',
          'บันทึกข้อมูลไม่สำเร็จ : รายละเอียดกิจกรรมต้องมีตัวอักษรอย่างน้อย 50 ตัว',
        );
        return;
      case 'location':
        aleartMessage(
          'error',
          'บันทึกข้อมูลไม่สำเร็จ: ชื่อสถานที่จัดกิจกรรมต้องขึ้นต้นด้วยตัวอักษร [A-Z] มีอย่างน้อย 3 ตัว และไม่เกิน 50 ตัว',
        );
        return;

      default:
        aleartMessage('error', 'บันทึกข้อมูลไม่สำเร็จ');
        return;
    }
  };

  const CreateActivities = async () => {
    const Activities = {
      academicYearID: AcademicyearID,
      activityTypeID: ActivitytypeID,
      clubID: userID,
      detail: detail,
      endtime: endtime + ':00+07:00',
      name: name,
      location: location,
      starttime: starttime + ':00+07:00',
    };

    const apiUrl = 'http://localhost:8080/api/v1/activities';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(Activities),
    };

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
          checkCaseSaveError(data.error.Name);
        }
      });
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

  const name_handleChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value;
    checkPattern('name', validateValue);
    setName(event.target.value as string);
  };

  const location_handleChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value;
    checkPattern('location', validateValue);
    setLocation(event.target.value as string);
  };

  const detail_handleChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value;
    checkPattern('detail', validateValue);
    setDetail(event.target.value as string);
  };

  return (
    <SidebarPage>
      <AppSidebar />
      <Page theme={pageTheme.website}>
        <Header style={HeaderCustom} title={`กิจกรรม`}>
          <UserHeader />
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
            {
              <Link component={RouterLink} to="/ActivityTable">
                <Button variant="contained" color="primary">
                  ตารางข้อมูลกิจกรรม
                </Button>
              </Link>
            }
          </ContentHeader>
          <Container maxWidth="sm">
            <Grid container spacing={2}>
              <Grid item xs={12}></Grid>
              <Grid item xs={3}>
                <div className={classes.paper}>กิจกรรม</div>
              </Grid>
              <Grid item xs={9}>
                <TextField
                  error={NameError ? true : false}
                  helperText={NameError}
                  id="name"
                  name="name"
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
                  error={DetailError ? true : false}
                  helperText={DetailError}
                  id="detail"
                  name="detail"
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
                <div className={classes.paper}>สถานที่จัดกิจกรรม</div>
              </Grid>
              <Grid item xs={9}>
                <TextField
                  error={LocationError ? true : false}
                  helperText={LocationError}
                  id="location"
                  name="location"
                  label="ใส่สถานที่จัดกิจกรรม"
                  variant="filled"
                  value={location}
                  style={{ width: 350 }}
                  onChange={location_handleChange}
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
                    name="acadamicyear"
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
    </SidebarPage>
  );
}

export default Activities;
