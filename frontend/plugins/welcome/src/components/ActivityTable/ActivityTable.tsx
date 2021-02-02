import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import {
  ContentHeader,
  Content,
  Header,
  Page,
  SidebarPage,
  pageTheme,
} from '@backstage/core';
import {
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  Button,
} from '@material-ui/core';
import TableHead from '@material-ui/core/TableHead';
import { Link as RouterLink } from 'react-router-dom';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';
import { EntActivities } from '../../api/models/EntActivities';
import { EntClub } from '../../api/models/EntClub';
import Swal from 'sweetalert2'; // alert
import { AppSidebar } from '../Sidebar/Sidebar';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    formControl: {
      margin: theme.spacing(1),
      minWidth: 120,
      fontSize: 24,
    },
    selectEmpty: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
  }),
);

export default function ActivityTable() {
  const profile = { givenName: 'ระบบจัดการชมรม' };
  const classes = useStyles();
  const api = new DefaultApi();
  const [activities, setActivities] = useState<EntActivities[]>();
  const [getAct, setGetAct] = useState<EntActivities[]>();
  const [clubs, setClubs] = useState<EntClub[]>([]);
  const [loading, setLoading] = useState(true);
  const [open, setOpen] = useState(true);
  const [close, setClose] = useState(false);

  const [ClubID, setClubID] = useState(Number);
  const [FindClub, setFindClub] = useState(String);
  const [CantFindClub, setCantFindClub] = useState(String);

  useEffect(() => {
    const getActivities = async () => {
      const res = await api.listActivities({ limit: 10, offset: 0 });
      setLoading(false);

      setActivities(res);
      console.log(res);
    };

    const getClub = async () => {
      const res = await api.listClub({ limit: 10, offset: 0 });
      setLoading(false);

      setClubs(res);
      console.log(res);
    };
    getClub();
    getActivities();
  }, [loading]);

  const ShowClub = async () => {
    setFindClub('none');
    setCantFindClub('block-inline');
    history.pushState('', '', './ActivityTable');
    window.location.reload(false);
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

  const findClub = async () => {
    setOpen(false);
    setClose(true);
    const findClubs = {
      club: ClubID,
    };
    const getClub = async () => {
      var ActArray: number;
      const res = await api.getActivities({ id: ClubID });
      setCantFindClub('none');
      setFindClub('block');
      setLoading(false);
      setGetAct(res);
      ActArray = res.length;
      if (ActArray > 0) {
        Toast.fire({
          icon: 'success',
          title: 'ค้นหาสำเร็จ',
        });
      } else {
        Toast.fire({
          icon: 'error',
          title: 'ไม่พบกิจกรรมของชมรมที่ค้นหา',
        });
        /*  setTimeout(() => {
          history.pushState('', '', './ActivityTable');
          window.location.reload(false);
        }, 3000); */
      }
      console.log(res);
    };
    getClub();
  };

  const deleteActivities = async (id: number) => {
    const res = await api.deleteActivities({ id: id });
    setLoading(true);
  };

  const club_id_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setClubID(event.target.value as number);
  };

  return (
    <SidebarPage>
      <AppSidebar />
      <Page theme={pageTheme.website}>
        <Header title={`${profile.givenName || 'to Backstage'}`} subtitle="">
          {/*   <Link component={RouterLink} to="/welcome">
            <Button variant="contained" color="secondary">
              กลับสู่หน้าหนัก
            </Button>
          </Link> */}
          {/* <Button
          component={RouterLink}
          to="/"
          variant="contained"
          color="secondary"
          disableElevation
          style={{ marginLeft: 20 }}
        >
          LOGOUT
        </Button> */}
        </Header>
        <Content>
          <ContentHeader title="ตารางกิจกรรมชมรม">
            <FormControl className={classes.formControl}>
              <InputLabel id="demo-simple-select-readonly-label">
                เลือกชมรมที่ต้องการค้นหากิจกรรม
              </InputLabel>
              <Select
                id="club"
                name="club"
                value={ClubID} // (undefined || '') = ''
                onChange={club_id_handleChange}
                style={{ width: 350, marginRight: 20 }}
              >
                <MenuItem value="" disabled>
                  เลือกชมรมที่ต้องการค้นหากิจกรรม
                </MenuItem>
                {clubs.map((item: EntClub) => (
                  <MenuItem value={item.id}>{item.name}</MenuItem>
                ))}
              </Select>
            </FormControl>

            <Button
              variant="contained"
              color="primary"
              onClick={() => {
                findClub();
              }}
              style={{ marginRight: 10 }}
            >
              ค้นหากิจกรรมชมรม
            </Button>

            <Button
              variant="contained"
              color="primary"
              onClick={() => {
                ShowClub();
              }}
            >
              ดูกิจกรรมชมรมทั้งหมด
            </Button>
          </ContentHeader>

          <div>
            {open ? (
              <TableContainer component={Paper}>
                <Table className={classes.table} aria-label="simple table">
                  <TableHead>
                    <TableRow>
                      {/*  <TableCell align="center">รหัสกิจกรรม</TableCell> */}
                      <TableCell align="center" style={{ width: 100 }}>
                        ชมรม
                      </TableCell>
                      <TableCell align="center">กิจกรรม</TableCell>
                      <TableCell align="center" style={{ width: 300 }}>
                        รายละเอียดกิจกรรม
                      </TableCell>
                      <TableCell align="center">ประเภทกิจกรรม</TableCell>
                      <TableCell align="center">
                        ภาคการศึกษาที่จัดกิจกรรม
                      </TableCell>
                      <TableCell align="center">เวลาเริ่มกิจกรรม</TableCell>
                      <TableCell align="center">เวลาสิ้นสุดกิจกรรม</TableCell>
                      {/* <TableCell align="center">Manage</TableCell> */}
                    </TableRow>
                  </TableHead>

                  <TableBody>
                    {activities === undefined
                      ? null
                      : activities.map(item => (
                          <TableRow key={item.id}>
                            {/*  <TableCell align="center">{item.id}</TableCell> */}
                            <TableCell align="center" style={{ width: 100 }}>
                              {item.edges?.club?.name}
                            </TableCell>
                            <TableCell align="center">{item.name}</TableCell>
                            <TableCell align="left" style={{ width: 300 }}>
                              {item.detail}
                            </TableCell>
                            <TableCell align="center">
                              {item.edges?.activitytype?.name}
                            </TableCell>
                            <TableCell align="center">
                              {item.edges?.academicyear?.semester}
                            </TableCell>
                            <TableCell align="center">
                              {item.starttime}
                            </TableCell>
                            <TableCell align="center">{item.endtime}</TableCell>

                            {/* <TableCell align="center">
                        <Button
                          onClick={() => {
                            deleteActivities(item.id);
                          }}
                          style={{ marginLeft: 10 }}
                          variant="contained"
                          color="secondary"
                        >
                          Delete
                        </Button>
                      </TableCell> */}
                          </TableRow>
                        ))}
                  </TableBody>
                </Table>
              </TableContainer>
            ) : (
              <TableContainer component={Paper}>
                <Table className={classes.table} aria-label="simple table">
                  <TableHead>
                    <TableRow>
                      {/*  <TableCell align="center">รหัสกิจกรรม</TableCell> */}
                      <TableCell align="center" style={{ width: 100 }}>
                        ชมรม
                      </TableCell>
                      <TableCell align="center">กิจกรรม</TableCell>
                      <TableCell align="center" style={{ width: 300 }}>
                        รายละเอียดกิจกรรม
                      </TableCell>
                      <TableCell align="center">ประเภทกิจกรรม</TableCell>
                      <TableCell align="center">
                        ภาคการศึกษาที่จัดกิจกรรม
                      </TableCell>
                      <TableCell align="center">เวลาเริ่มกิจกรรม</TableCell>
                      <TableCell align="center">เวลาสิ้นสุดกิจกรรม</TableCell>
                      {/* <TableCell align="center">Manage</TableCell> */}
                    </TableRow>
                  </TableHead>

                  <TableBody>
                    {getAct === undefined
                      ? null
                      : getAct.map(getAct => (
                          <TableRow key={getAct.id}>
                            {/*   <TableCell align="center">{getAct.id}</TableCell> */}
                            <TableCell align="center" style={{ width: 100 }}>
                              {getAct.edges?.club?.name}
                            </TableCell>
                            <TableCell align="center">{getAct.name}</TableCell>
                            <TableCell align="left" style={{ width: 300 }}>
                              {getAct.detail}
                            </TableCell>
                            <TableCell align="center">
                              {getAct.edges?.activitytype?.name}
                            </TableCell>
                            <TableCell align="center">
                              {getAct.edges?.academicyear?.semester}
                            </TableCell>
                            <TableCell align="center">
                              {getAct.starttime}
                            </TableCell>
                            <TableCell align="center">
                              {getAct.endtime}
                            </TableCell>

                            {/* <TableCell align="center">
                        <Button
                          onClick={() => {
                            deleteActivities(item.id);
                          }}
                          style={{ marginLeft: 10 }}
                          variant="contained"
                          color="secondary"
                        >
                          Delete
                        </Button>
                      </TableCell> */}
                          </TableRow>
                        ))}
                  </TableBody>
                </Table>
              </TableContainer>
            )}
          </div>
        </Content>
      </Page>
    </SidebarPage>
  );
}
