import React, { useState, useEffect } from 'react';
import { Content, Header, Page, pageTheme, SidebarPage } from '@backstage/core';
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
import { EntClubBranch } from '../../api/models/EntClubBranch';
import { EntClubType } from '../../api/models/EntClubType';
import { DefaultApi } from '../../api/apis';
import Snackbar from '@material-ui/core/Snackbar';
import Alert from '@material-ui/lab/Alert';
import { UserHeader } from '../UserHeader/UserHeader';
import { AppSidebar } from '../Sidebar/Sidebar';

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

function Clubs() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [users, setUsers] = useState<EntUser[]>([]);
  const [clubbranchs, setClubBranch] = useState<EntClubBranch[]>([]);
  const [clubtypes, setClubType] = useState<EntClubBranch[]>([]);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);
  const [UserID, setUserID] = useState(Number);
  const [ClubBranchID, setClubBranchID] = useState(Number);
  const [ClubTypeID, setClubTypeID] = useState(Number);
  const [Name, setName] = React.useState(String);
  const [Purpose, setPurpose] = React.useState(String);

  useEffect(() => {
    const getClubBranch = async () => {
      const cb = await api.listClubbranch({ limit: 10, offset: 0 });
      setLoading(false);
      setClubBranch(cb);
    };
    getClubBranch();

    const getUsers = async () => {
      const us = await api.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUsers(us);
    };
    getUsers();

    const getClubType = async () => {
      const ct = await api.listClubtype({ limit: 10, offset: 0 });
      setLoading(false);
      setClubType(ct);
    };
    getClubType();
  }, [loading]);

  const CreateClub = async () => {
    const Clubs = {
      userID: UserID,
      clubBranchID: ClubBranchID,
      clubTypeID: ClubTypeID,
      name: Name,
      purpose: Purpose,
    };
    console.log(Clubs);
    const res: any = await api.createClub({
      club: Clubs,
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
  const ClubType_id_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setClubTypeID(event.target.value as number);
  };

  const ClubBranch_id_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setClubBranchID(event.target.value as number);
  };

  const user_id_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setUserID(event.target.value as number);
  };

  const name_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setName(event.target.value as string);
  };

  const purpose_handleChange = (
    event: React.ChangeEvent<{ value: unknown }>,
  ) => {
    setPurpose(event.target.value as string);
  };
  console.log(ClubBranchID);
  console.log(ClubTypeID);
  console.log(UserID);
  console.log(Purpose);
  console.log(Name);

  return (
    <SidebarPage>
      <AppSidebar />
      <Page theme={pageTheme.website}>
        <Header style={HeaderCustom} title={`สร้างชมรม`}>
          <UserHeader />
        </Header>
        <Content>
          <Container maxWidth="sm">
            <Grid container spacing={2}>
              <Grid item xs={12}></Grid>
              <Grid item xs={3}>
                <div className={classes.paper}>สร้างชมรม</div>
              </Grid>
              <Grid item xs={9}>
                <TextField
                  id="filled-basic"
                  label="ใส่ชื่อชมรม"
                  variant="filled"
                  value={Name}
                  style={{ width: 350 }}
                  onChange={name_handleChange}
                />
              </Grid>



              <Grid item xs={3}>
                <div className={classes.paper}>ประเภทชมรม</div>
              </Grid>
              <Grid item xs={9}>
                <FormControl variant="outlined" className={classes.formControl}>
                  <InputLabel id="demo-mutiple-name-label">
                    เลือกประเภทของชมรม
                </InputLabel>
                  <Select
                    name="clubtype"
                    value={ClubTypeID} // (undefined || '') = ''
                    onChange={ClubType_id_handleChange}
                    style={{ width: 350 }}
                  >
                    {clubtypes.map((item: EntClubType) => (
                      <MenuItem value={item.id}>{item.name}</MenuItem>
                    ))}
                  </Select>
                </FormControl>
              </Grid>

              <Grid item xs={3}>
                <div className={classes.paper}>สถานที่จัดชมรม</div>
              </Grid>
              <Grid item xs={9}>
                <FormControl variant="outlined" className={classes.formControl}>
                  <InputLabel id="demo-mutiple-name-label">
                    เลือกสถานที่จัดชมรม
                </InputLabel>
                  <Select
                    name="ClubBranch"
                    value={ClubBranchID} // (undefined || '') = ''
                    onChange={ClubBranch_id_handleChange}
                    style={{ width: 350 }}
                  >
                    {clubbranchs.map((item: EntClubBranch) => (
                      <MenuItem value={item.id}>{item.name}</MenuItem>
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
                      <MenuItem value={item.id}>{item.name}</MenuItem>
                    ))}
                  </Select>
                </FormControl>
              </Grid>


              <Grid item xs={3}>
                <div className={classes.paper}>จุดประสงค์ชมรม</div>
              </Grid>
              <Grid item xs={9}>
                <TextField
                  id="filled-multiline-static"
                  label="ใส่จุดประสงค์ชมรม"
                  value={Purpose}
                  multiline
                  rows={2}
                  defaultValue=""
                  variant="filled"
                  style={{ width: 350, height: 50, marginBottom: 25 }}
                  onChange={purpose_handleChange}
                />
              </Grid>
              <Grid item xs={1}>
                <Button
                  variant="contained"
                  color="primary"
                  onClick={() => {
                    CreateClub();
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
            {status ? (
              <div>
                {alert ? (
                  <Alert severity="success" >
                    This is a success alert — check it out!
                  </Alert>
                ) : (
                    <Alert severity="error" style={{ marginTop: 20 }}>
                      This is a warning alert — check it out!
                    </Alert>
                  )}
              </div>
            ) : null}

          </Container>
        </Content>
      </Page>
    </SidebarPage>
  );
}

export default Clubs;

