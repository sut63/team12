import React, { useState, useEffect } from 'react';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  SidebarPage,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';

import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import Grid from '@material-ui/core/Grid';
import Container from '@material-ui/core/Container';

import { EntUser } from '../../api/models/EntUser';
import { EntRoom } from '../../api/models/EntRoom';
import { EntPurpose } from '../../api/models/EntPurpose';
import { UserHeader } from '../UserHeader/UserHeader';
import { AppSidebar } from '../Sidebar/Sidebar';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },
    paper: {
      marginTop: theme.spacing(2),
      marginBottom: theme.spacing(2),
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
      width: '25ch',
    },
  }));

export default function Roomuse() {
  const classes = useStyles();
  const api = new DefaultApi();

  const [users, setUsers] = useState<EntUser[]>([]);
  const [rooms, setRooms] = useState<EntRoom[]>([]);
  const [purposes, setPurposes] = useState<EntPurpose[]>([]);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(Boolean);
  const [loading, setLoading] = useState(true);

  const [addedtime, setAddedtime] = useState(String);
  const [roomid, setRoom] = useState(Number);
  const [purposeid, setPurpose] = useState(Number);
  const [userid, setUser] = useState(Number);

  useEffect(() => {
    const getUser = async () => {
      const res = await api.listUser({ limit: 4, offset: 0 });
      setLoading(false);
      setUsers(res);
      console.log(res);
    };
    getUser();

    const getRoom = async () => {
      const res = await api.listRoom({ limit: 6, offset: 0 });
      setLoading(false);
      setRooms(res);
      console.log(res);
    };
    getRoom();

    const getPurpose = async () => {
      const res = await api.listPurpose({ limit: 5, offset: 0 });
      setLoading(false);
      setPurposes(res);
      console.log(res);
    };
    getPurpose();

  }, [loading]);

  const handleDatetimeChange = (event: any) => {
    setAddedtime(event.target.value as string);
  };

  const CreateRoomuse = async () => {

    if ((roomid != 0) && (purposeid != 0) && (userid != 0)) {
      const roomuse = {
        addedTime: addedtime + ":00+07:00",
        userID: userid,
        roomID: roomid,
        purposeID: purposeid,
      }
      console.log(roomuse);

      const res: any = await api.createRoomuse({ roomuse: roomuse });
      setStatus(true);
      if (res.id != '') {
        setAlert(true);
        setTimeout(() => {
          setStatus(false);
        }, 6000);
      }
    }
    else {
      setStatus(true);
      setAlert(false);
      setTimeout(() => {
        setStatus(false);
      }, 6000);
    }

  };


  const useridhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUser(event.target.value as number);
  };

  const roomidhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setRoom(event.target.value as number);
  };

  const purposeidhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPurpose(event.target.value as number);
  };

  const refreshPage = () => {
    window.location.reload();
  }


  return (
    <SidebarPage>
      <AppSidebar />
      <Page theme={pageTheme.home}>
        <Header title="การเข้าใช้ห้องชมรม">
          <UserHeader />
        </Header>
        <Content>
          <ContentHeader title="Check - in">

            {status ? (
              <div>
                {alert ? (
                  <Alert severity="success">
                    บันทึกสำเร็จ!!
                  </Alert>
                ) : (
                    <Alert severity="error">
                      <strong>บันทึกล้มเหลว!!</strong>
                    </Alert>
                  )}
              </div>
            ) : null}

          </ContentHeader>
          <Container maxWidth="sm">
            <Grid container spacing={3}>
              <Grid item xs={12}></Grid>
              <Grid item xs={3}>
                <div className={classes.paper}>ชื่อห้อง</div>
              </Grid>
              <Grid item xs={9}>
                <FormControl variant="outlined" className={classes.formControl}>
                  <InputLabel id="room_id-label">เลือกห้อง</InputLabel>
                  <Select
                    labelId="room_id-label"
                    label="Room"
                    id="room_id"
                    value={roomid}
                    onChange={roomidhandleChange}
                    style={{ width: 600 }}
                  >
                    {rooms.map((item: EntRoom) =>
                      <MenuItem value={item.id}>{item.roomName}</MenuItem>)}
                  </Select>
                </FormControl>
              </Grid>

              <Grid item xs={3}>
                <div className={classes.paper}>จุดประสงค์</div>
              </Grid>
              <Grid item xs={9}>
                <FormControl variant="outlined" className={classes.formControl}>
                  <InputLabel id="purpose_id-label">ใช้ทำอะไร</InputLabel>
                  <Select
                    labelId="purpose_id-label"
                    label="Purpose"
                    id="purpose_id"
                    value={purposeid}
                    onChange={purposeidhandleChange}
                    style={{ width: 600 }}
                  >
                    {purposes.map((item: EntPurpose) =>
                      <MenuItem value={item.id}>{item.purpose}</MenuItem>)}
                  </Select>
                </FormControl>
              </Grid>



              <Grid item xs={3}>
                <div className={classes.paper}>ผู้ใช้</div>
              </Grid>
              <Grid item xs={9}>
                <FormControl variant="outlined" className={classes.formControl}>
                  <InputLabel id="user_id-label">เลือกผู้ใช้</InputLabel>
                  <Select
                    labelId="user_id-label"
                    label="User"
                    id="user_id"
                    value={userid}
                    onChange={useridhandleChange}
                    style={{ width: 600 }}
                  >
                    {users.map((item: EntUser) =>
                      <MenuItem value={item.id}>{item.name}</MenuItem>)}
                  </Select>
                </FormControl>
              </Grid>

              <Grid item xs={3}>
                <div className={classes.paper} >เวลาที่เข้าใช้</div>
              </Grid>
              <Grid item xs={9}>
                <div>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >
                    <TextField
                      id="datetime"
                      label="DateTime"
                      type="datetime-local"
                      value={addedtime}
                      onChange={handleDatetimeChange}
                      className={classes.textField}
                      InputLabelProps={{
                        shrink: true,
                      }}
                    />
                  </FormControl>
                </div>
              </Grid>

              <Grid item xs={3}></Grid>
              <Grid item xs={9}>
                <Button
                  variant="contained"
                  color="primary"
                  size="large"
                  style={{ marginRight: 1 }}
                  onClick={() => {
                    CreateRoomuse();
                  }}
                >
                  บันทึก
               </Button>
              </Grid>
            </Grid>
          </Container>
        </Content>
      </Page>
    </SidebarPage>
  );

};
