import React, { useState, useEffect } from 'react';
import { Grid } from '@material-ui/core';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  SidebarPage,
} from '@backstage/core';
//import ComplaintTable from '../ComplaintTable';
import FormControl from '@material-ui/core/FormControl';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import Paper from '@material-ui/core/Paper';
import { Link, Button } from '@material-ui/core';
import { Link as RouterLink } from 'react-router-dom';
import { makeStyles, Theme, createStyles} from '@material-ui/core/styles';
import { AppSidebar } from '../Sidebar/Sidebar';
import { UserHeader } from '../UserHeader/UserHeader';
import { EntClub } from '../../api/models/EntClub';
import { EntComplaintType } from '../../api/models/EntComplaintType';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import { DefaultApi } from '../../api/apis';
import { EntComplaint } from '../../api/models/EntComplaint';

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
    table: {
      minWidth: 650,
    },
  }),
);

export default function ComplaintSearch() {
  
  const classes = useStyles();
  const api = new DefaultApi();

  const [clubs, setClubs] = useState<EntClub[]>([]);
  const [complainttypes, setComplaintTypes] = useState<EntComplaintType[]>([]);
  const [complaints, setComplaints] = useState<EntComplaint[]>();

  const [clubid, setClubID] = useState(Number);
  const [complainttypeid, setComplaintTypeID] = useState(Number);

  const [loading, setLoading] = useState(true);

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

  const getComplaints = async () => {
    const res = await api.listComplaint({ limit: 10, offset: 0 });
    setLoading(false);
    setComplaints(res);
    console.log(res);
  };

  useEffect(() => {
    getClubs();
    getTypes();
    getComplaints();
  }, [loading]);

  const handleClubChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setClubID(event.target.value as number);
  };

  const handleComplaintTypeChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setComplaintTypeID(event.target.value as number);
  };

  const deleteComplaints = async (id: number) => {
    const res = await api.deleteComplaint({ id: id });
    setLoading(true);
  };

  const search = async () => {

  }

  // const search = async () => {
  //   setOpen(false);
  //   setClose(true);
  //   const clubs = {
  //     club: clubid,
  //   };
  //   const Club = async () => {
  //     var ComArray: number;
  //     const res = await api.getComplaint({ id: clubid });
  //     setCantFindClub('none');
  //     setFindClub('block');
  //     setLoading(false);
  //     setGetCom(res);
  //     ComArray = ;
  //     if (ComArray > 0) {
  //       // Toast.fire({
  //       //   icon: 'success',
  //       //   title: 'ค้นหาสำเร็จ',
  //       // });
  //     } else {
  //       // Toast.fire({
  //       //   icon: 'error',
  //       //   title: 'ไม่พบกิจกรรมของชมรมที่ค้นหา',
  //       // });
  //       /*  setTimeout(() => {
  //         history.pushState('', '', './ActivityTable');
  //         window.location.reload(false);
  //       }, 3000); */
  //     }
  //     console.log(res);
  //   };
  //   Club();
  // };

  return(
    <SidebarPage>
      <AppSidebar />
      <Page theme={pageTheme.home}>
        <Header title={`ระบบร้องเรียน`}><UserHeader /></Header>
        <Content>
          <ContentHeader title="ค้นหาเรื่องร้องเรียน">
          <Link component={RouterLink} to="/Complaint">
            <Button variant="contained" color="primary">
              ส่งเรื่องร้องเรียน
            </Button>
          </Link>
          </ContentHeader>
          <div className={classes.root}>
            <Grid container spacing={3}>
              <Grid item xs={12}>
                <Paper className={classes.paper}>
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
                    <Button
                      variant="contained"
                      color="secondary"
                      className={classes.button}
                      onClick={() => { 
                        search();
                      }}
                    >
                      ค้นหา
                    </Button>
                  </div>
                  
                  <div>
                  <TableContainer component={Paper}>
                    <Table className={classes.table} aria-label="simple table">
                      <TableHead>
                        <TableRow>
                          <TableCell align="center">หมายเลขเรื่องร้องเรียน</TableCell>
                          <TableCell align="center">อีเมล์ผู้ส่ง</TableCell>
                          <TableCell align="center">ชมรม</TableCell>
                          <TableCell align="center">ชื่อผู้ส่ง</TableCell>
                          <TableCell align="center">หมายเลขโทรศัพท์</TableCell>
                          <TableCell align="center">หัวข้อ</TableCell>
                          <TableCell align="center">รายละเอียด</TableCell>
                          <TableCell align="center">วันที่</TableCell>
                        </TableRow>
                      </TableHead>
                      <TableBody>
                        {complaints === undefined
                          ? null
                          : complaints.map(item => (
                              <TableRow key={item.id}>
                                <TableCell align="center">{item.id}</TableCell>
                                <TableCell align="center">{item.edges?.complaintToUser?.email}</TableCell>
                                <TableCell align="center">{item.edges?.complaintToClub?.name}</TableCell>
                                <TableCell align="center">{item.name}</TableCell>
                                <TableCell align="center">{item.phonenumber}</TableCell>
                                <TableCell align="center">{item.edges?.complaintToComplaintType?.description}</TableCell>
                                <TableCell align="center">{item.info}</TableCell>
                                <TableCell align="center">{item.date}</TableCell>
                                <TableCell align="center">
                                  <Button
                                    onClick={() => {
                                      deleteComplaints(item.id);
                                    }}
                                    style={{ marginLeft: 10 }}
                                    variant="contained"
                                    color="secondary"
                                  >
                                    ลบ
                                  </Button>
                                </TableCell>
                              </TableRow>
                            ))}
                      </TableBody>
                    </Table>
                  </TableContainer>
                  </div>

                </Paper>
              </Grid>
            </Grid>
          </div>
        </Content>
        
      </Page>
    </SidebarPage>
  );
};
