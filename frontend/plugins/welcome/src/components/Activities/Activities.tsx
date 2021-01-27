import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import {
  ContentHeader,
  Content,
  Header,
  Page,
  pageTheme,
} from '@backstage/core';
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
import TableHead from '@material-ui/core/TableHead';
import { Link as RouterLink } from 'react-router-dom';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';
import { EntActivities } from '../../api/models/EntActivities';
import { EntClub } from '../../api/models/EntClub';

const useStyles = makeStyles({
  table: {
    minWidth: 650,
  },
});

export default function ActivityTable() {
  const profile = { givenName: 'ระบบจัดการชมรม' };
  const classes = useStyles();
  const api = new DefaultApi();
  const [activities, setActivities] = useState<EntActivities[]>();
  const [getAct, setGetAct] = useState<EntActivities>();
  const [clubs, setClubs] = useState<EntClub[]>([]);
  const [loading, setLoading] = useState(true);

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

  const findClub = async () => {
    const findClubs = {
      club: ClubID,
    };

    const getClub = async () => {
      const res = await api.getActivities({ id: ClubID });
      setCantFindClub('none');
      setFindClub('block');
      setLoading(false);

      setGetAct(res);
      console.log(res);
    };
    getClub();
    /*  const apiUrl = 'http://localhost:8080/api/v1/activities';
    const requestOptions = {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(findClubs),
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
        } 
      });  */
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
    <Page theme={pageTheme.website}>
      <Header title={`${profile.givenName || 'to Backstage'}`} subtitle="">
        <Link component={RouterLink} to="/welcome">
          <Button variant="contained" color="primary">
            กลับสู่หน้าหนัก
          </Button>
        </Link>
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
        <ContentHeader title="ตารางข้อมูลโปรโมชัน">
          <FormControl variant="outlined">
            <InputLabel id="demo-mutiple-name-label">
              เลือกชมรมที่ต้องการค้นหากิจกรรม
            </InputLabel>
            <Select
              id="type"
              name="activitytype"
              value={ClubID} // (undefined || '') = ''
              onChange={club_id_handleChange}
              style={{ width: 350, marginRight: 20 }}
            >
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

        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center">รหัสกิจกรรม</TableCell>
                <TableCell align="center">ชมรม</TableCell>
                <TableCell align="center">กิจกรรม</TableCell>
                <TableCell align="center">รายละเอียดกิจกรรม</TableCell>
                <TableCell align="center">ประเภทกิจกรรม</TableCell>
                <TableCell align="center">ภาคการศึกษาที่จัดกิจกรรม</TableCell>
                <TableCell align="center">เวลาเริ่มกิจกรรม</TableCell>
                <TableCell align="center">เวลาสิ้นสุดกิจกรรม</TableCell>
                {/* <TableCell align="center">Manage</TableCell> */}
              </TableRow>
            </TableHead>

            <TableBody style={{ display: CantFindClub }}>
              {activities === undefined
                ? null
                : activities.map(item => (
                    <TableRow key={item.id}>
                      <TableCell align="center">{item.id}</TableCell>
                      <TableCell align="center">
                        {item.edges?.club?.name}
                      </TableCell>
                      <TableCell align="center">{item.name}</TableCell>
                      <TableCell align="center">{item.detail}</TableCell>
                      <TableCell align="center">
                        {item.edges?.activitytype?.name}
                      </TableCell>
                      <TableCell align="center">
                        {item.edges?.academicyear?.semester}
                      </TableCell>
                      <TableCell align="center">{item.starttime}</TableCell>
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

            <TableBody style={{ display: FindClub }}>
              {getAct === undefined ? null : (
                <TableRow key={getAct.id}>
                  <TableCell align="center">{getAct.id}</TableCell>
                  <TableCell align="center">
                    {getAct.edges?.club?.name}
                  </TableCell>
                  <TableCell align="center">{getAct.name}</TableCell>
                  <TableCell align="center">{getAct.detail}</TableCell>
                  <TableCell align="center">
                    {getAct.edges?.activitytype?.name}
                  </TableCell>
                  <TableCell align="center">
                    {getAct.edges?.academicyear?.semester}
                  </TableCell>
                  <TableCell align="center">{getAct.starttime}</TableCell>
                  <TableCell align="center">{getAct.endtime}</TableCell>

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
              )}
            </TableBody>
          </Table>
        </TableContainer>

        {/* <div
          style={{
            display: FindClub,
            marginTop: 50,
            fontSize: 36,
            textAlign: 'center',
          }}
        >
          ไม่พบกิจกรรมของชมรมที่เลือก
        </div> */}
      </Content>
    </Page>
  );
}
