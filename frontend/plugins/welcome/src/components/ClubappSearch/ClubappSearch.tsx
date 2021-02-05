import React, { FC, useState, useEffect } from 'react';
import { makeStyles, withStyles, createStyles, Theme } from '@material-ui/core/styles'
import { Content, Header, Page, pageTheme, SidebarPage } from '@backstage/core';
import { AppSidebar } from '../Sidebar/Sidebar';
import { UserHeader } from '../UserHeader/UserHeader';
import { Grid, Typography, AppBar, Toolbar, MenuItem, Select, Switch,
        Grow , FormControlLabel, IconButton, Container, Paper, Table,
        TableBody, TableCell, TableContainer, TableHead, TableRow} from '@material-ui/core';
import SearchIcon from '@material-ui/icons/Search';
import { DefaultApi } from '../../api/apis';
import { EntClub, EntClubappStatus, EntClubapplication } from '../../api/models';
import { Alert } from '@material-ui/lab';
import Snackbar from '@material-ui/core/Snackbar';
import ClearIcon from '@material-ui/icons/Clear';
import ReportProblemTwoToneIcon from '@material-ui/icons/ReportProblemTwoTone';
import FiberManualRecordTwoToneIcon from '@material-ui/icons/FiberManualRecordTwoTone';
import moment from 'moment';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root_club: { '& > *': { minWidth: 120, margin: theme.spacing(0), width: '25ch', }, },
        table: { minWidth: 650, },
    }),
);

const StyledTableCell = withStyles((theme: Theme) =>
    createStyles({
        head: {
            backgroundColor: "#339999",
            color: theme.palette.common.white,
        },
        body: {
            fontSize: 14,
        },
    }),
)(TableCell);

export default function ClubappSearch() {

    const classes = useStyles();
    const api = new DefaultApi();
    const HeaderCustom = { minHeight: '50px', };
    const [checked, setChecked] = React.useState(false);

    const [loading, setLoading] = useState(true);
    const [clubID, SetClubID] = useState(Number);
    const [appstatusID, SetAppstatusID] = useState(Number);
    const [search, SetSearch] = useState(Boolean);
    const [clubs, SetClubs] = useState("");
    const [appstatuses, SetAppstatuses] = useState(String);
    const [alert, SetAlert] = useState(Boolean);
    const [status, SetStatus] = useState(false);
    const [open, setOpen] = React.useState(false);

    const [club, SetClub] = useState<EntClub[]>([]);
    const [appstatus, SetAppstatus] = useState<EntClubappStatus[]>([]);
    const [clubapp, SetClubapp] = useState<EntClubapplication[]>([]);
    const [data, SetData] = useState<EntClubapplication[]>([]);

    const GrowhandleChange = () => {
        setChecked((prev) => !prev);
    };

    const ClubIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        SetClubID(event.target.value as number);
        SetClubData(event.target.value as number);
        
    };
    // console.log("clubIDChange=>", clubID)
    // console.log("club=>", clubs)

    const AppstatusIDhabdleChang = (event: React.ChangeEvent<{ value: unknown }>) => {
        SetAppstatusID(event.target.value as number);
        SetAppStatusData(event.target.value as number);
    };
    // console.log("appstatusIDChange=>", appstatusID)
    // console.log("appstatus=>", appstatuses)

    var uname = JSON.parse(String(localStorage.getItem('user-name')))

    function handleClose() {
        setOpen(false);
    };

    function SearchTrue() {
        SetStatus(true)
        SetAlert(true)
        setOpen(true)
        SetSearch(true)
    };

    function SearchFalse() {
        SetStatus(true)
        SetAlert(false)
        setOpen(true)
        SetSearch(false)
    }

    const SetClubData = async (x: number) => {
        await club.map((item: EntClub) => {
            if (item.id == x) {
                const res = String(item.name)
                SetClubs(res);
            }
        })
    };

    const SetAppStatusData = async (x: number) => {
        await appstatus.map((item: EntClubappStatus) => {
            if (item.id == x) {
                const res = String(item.clubstatus)
                SetAppstatuses(res);
            }
        });
    };

    const SearchClubApplication = async () => {
        SetSearch(false)
        const res = clubapp.filter((filter: EntClubapplication) => filter.edges?.owner?.name?.includes(uname))
        //console.log(res.length)

        if (res.length > 0) {
            if (clubID == 0 || appstatusID == 0) {
                if (clubID == 0 && appstatusID != 0) {
                    const c = res.filter((filter: EntClubapplication) => filter.edges?.clubappstatus?.clubstatus?.includes(appstatuses))
                    if (c.length > 0) {
                        SetData(c)
                        SearchTrue();
                    }
                    else {
                        SearchFalse();
                    }
                }
                else if (appstatusID == 0 && clubID != 0) {
                    const c = res.filter((filter: EntClubapplication) => filter.edges?.club?.name?.includes(clubs))
                    if (c.length > 0) {
                        SetData(c)
                        SearchTrue();
                    }
                    else {
                        SearchFalse();
                    }
                }
                else if (appstatusID == 0 && clubID == 0) {
                    if (res.length > 0) {
                        SetData(res)
                        SearchTrue();
                    }
                    else {
                        SearchFalse();
                    }
                }
            }
            else if (clubID != 0 && appstatusID != 0) {
                const c = res.filter((filter: EntClubapplication) => filter.edges?.club?.name?.includes(clubs) && filter.edges.clubappstatus?.clubstatus?.includes(appstatuses))
                if (c.length > 0) {
                    SetData(c)
                    SearchTrue();
                }
                else {
                    SearchFalse();
                }
            }
        }
        else {
            SearchFalse();
        }
    }
    // console.log("search", search)
    // console.log("_______________________________________________")

    useEffect(() => {

        const getClub = async () => {
            const res = await (await api.listClub({ limit: 10, offset: 0 }));
            setLoading(false);
            SetClub(res);
        }
        getClub();

        const getAppstatus = async () => {
            const res = await api.listClubappstatus({ limit: 10, offset: 0 });
            setLoading(false);
            SetAppstatus(res);
        };
        getAppstatus();

        const getClubapplication = async () => {
            const res = await (await api.listClubapplication({ limit: 10, offset: 0 }));
            setLoading(false);
            SetClubapp(res);
        };
        getClubapplication();

    }, [loading]);

    return (
        <SidebarPage>
            <AppSidebar />
            <Page theme={pageTheme.home}>
                <Header style={HeaderCustom} title={`ระบบค้นหาใบคำร้องสมัครชมรม`}>
                    <UserHeader />
                </Header>
                <Content>
                    <Container maxWidth="md">
                        <Grid container spacing={3}>
                            <Grid item xs={12}>
                                <AppBar position="static">
                                    <Toolbar style={{ backgroundColor: "#217567" }}>
                                        <FiberManualRecordTwoToneIcon style={{ fontSize: "15" }} />
                                        <Typography variant="h6" noWrap style={{ marginRight: "0.5cm", marginLeft: "0.1cm" }}>
                                            ตัวเลือกการค้นหา
                                        </Typography>
                                        <FormControlLabel
                                            control={<Switch size="small" color="secondary" checked={checked} onChange={GrowhandleChange} />}
                                            label="แสดง"
                                            style={{ flexGrow: 1 }}
                                        />
                                        <IconButton color="inherit" component="span" onClick={SearchClubApplication} >
                                            <SearchIcon />
                                            ค้นหา
                                        </IconButton>
                                    </Toolbar>
                                </AppBar>
                            </Grid>
                            <Grid item xs={12} style={{ marginTop: "0.1cm", marginBottom: "0.5cm" }}>
                                <Grow in={checked}>
                                    <Grid item xs={12} sm container>
                                        <Grid item xs container direction="column" spacing={2}>
                                            <Grid item xs style={{ marginLeft: "1cm" }}>
                                                <Typography gutterBottom variant="subtitle1">
                                                    ชมรม
                                        </Typography>
                                                <form className={classes.root_club} noValidate autoComplete="off">
                                                    <Select
                                                        labelId="demo-simple-select-label"
                                                        id="club"
                                                        value={clubID || 0}
                                                        onChange={ClubIDhandleChange}
                                                    >
                                                        <MenuItem key={0} value={0}>all</MenuItem>
                                                        {club.map((item: EntClub) => (
                                                            <MenuItem key={item.id} value={item.id}>{item.name}</MenuItem>
                                                        ))}
                                                    </Select>
                                                </form>

                                            </Grid>
                                        </Grid>
                                        <Grid item xs>
                                            <Typography gutterBottom variant="subtitle1">
                                                สถานะใบคำร้อง
                                            </Typography>

                                            <form className={classes.root_club} noValidate autoComplete="off">
                                                <Select
                                                    id="application status"
                                                    value={appstatusID || 0}
                                                    onChange={AppstatusIDhabdleChang}
                                                >
                                                    <MenuItem key={0} value={0}>all</MenuItem>
                                                    {appstatus.map((item: EntClubappStatus) => (
                                                        <MenuItem key={item.id} value={item.id}>{item.clubstatus}</MenuItem>
                                                    ))}

                                                </Select>
                                            </form>
                                        </Grid>
                                        <Grid item xs style={{ marginLeft: "1cm" }}>

                                        </Grid>
                                    </Grid>
                                </Grow>
                            </Grid>

                            <Grid item xs={12} style={{ marginTop: "0cm", marginBottom: "0cm" }}>
                                <AppBar position="static">
                                    <Toolbar style={{ backgroundColor: "#217567" }}>
                                        <Typography variant="h6" noWrap style={{ marginRight: "1cm", flexGrow: 1}}>
                                            ผลการค้นหา
                                    </Typography>
                                    <IconButton style={{backgroundColor:"ffffff", color:"white"}} onClick={()=>{SetSearch(false); SetStatus(false);}}>
                                        <ClearIcon/>
                                    </IconButton>
                                    </Toolbar>
                                </AppBar>
                                {search ? (
                                    <TableContainer component={Paper} style={{ marginTop: "0.2cm" }}>
                                        <Table className={classes.table} aria-label="simple table">
                                            <TableHead>
                                                <TableRow>
                                                    <StyledTableCell align="center">ลำดับที่</StyledTableCell>
                                                    <StyledTableCell align="center">ชื่อผู้เขียน</StyledTableCell>
                                                    <StyledTableCell align="center">ชมรม</StyledTableCell>
                                                    <StyledTableCell align="center">วันและเวลา</StyledTableCell>
                                                    <StyledTableCell align="center">สถานะใบคำร้อง</StyledTableCell>
                                                </TableRow>
                                            </TableHead>
                                            <TableBody>{
                                                data.map((item: EntClubapplication) => (
                                                    <TableRow key={item.id}>
                                                        <TableCell align="center">{item.id}</TableCell>
                                                        <TableCell align="center">{item.addername}</TableCell>
                                                        <TableCell align="center">{item.edges?.club?.name}</TableCell>
                                                        <TableCell align="center">{moment(item.addeddatetime).format('MM/DD/YYYY HH.mm')}</TableCell>
                                                        <TableCell align="center">{item.edges?.clubappstatus?.clubstatus}</TableCell>
                                                    </TableRow>
                                                ))
                                            }
                                            </TableBody>
                                        </Table>
                                    </TableContainer>
                                ) : (null)}

                            </Grid>
                            {status ? (
                                <div>
                                    {alert ? (
                                        <Grid>
                                            <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
                                                <Alert onClose={handleClose} severity="success">
                                                    ค้นหาข้อมูลสำเร็จ
                                      </Alert>
                                            </Snackbar>
                                        </Grid>
                                    ) : (
                                        <>
                                            <Grid>
                                                <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
                                                    <Alert onClose={handleClose} severity="error">
                                                        ค้นหาข้อมูลไม่สำเร็จ
                                                    </Alert>
                                                </Snackbar>
                                            </Grid>
                                            <Grid>
                                                <Typography gutterBottom variant="subtitle1" style={{marginTop:"0.2cm", marginLeft:"0.5cm"}}>
                                                    <ReportProblemTwoToneIcon style={{color:"red",marginRight:"0.2cm"}}/>
                                                    ไม่พบข้อมูลตรงตามที่เลือกไว้
                                                </Typography>
                                            </Grid>
                                            </>
                                        )}
                                </div>
                            ) : null}
                        </Grid>
                    </Container>
                </Content>
            </Page>
        </SidebarPage >
    );
}