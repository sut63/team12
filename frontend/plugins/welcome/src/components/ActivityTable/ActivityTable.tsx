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
import { Link, Button } from '@material-ui/core';
import TableHead from '@material-ui/core/TableHead';
import { Link as RouterLink } from 'react-router-dom';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';
import { EntActivities } from '../../api/models/EntActivities';

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
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getActivities = async () => {
      const res = await api.listActivities({ limit: 10, offset: 0 });
      setLoading(false);

      setActivities(res);
      console.log(res);
    };

    getActivities();
  }, [loading]);

  const deleteActivities = async (id: number) => {
    const res = await api.deleteActivities({ id: id });
    setLoading(true);
  };

  return (
    <Page theme={pageTheme.website}>
      <Header title={`${profile.givenName || 'to Backstage'}`} subtitle="">
        <Button
          component={RouterLink}
          to="/"
          variant="contained"
          color="secondary"
          disableElevation
        >
          LOGOUT
        </Button>
      </Header>
      <Content>
        <ContentHeader title="ตารางข้อมูลโปรโมชัน">
          <Link component={RouterLink} to="/Activities">
            <Button variant="contained" color="primary">
              บันทึกข้อมูลกิจกรรม
            </Button>
          </Link>
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
                <TableCell align="center">Manage</TableCell>
              </TableRow>
            </TableHead>

            <TableBody>
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

                      <TableCell align="center">
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
                      </TableCell>
                    </TableRow>
                  ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Content>
    </Page>
  );
}
