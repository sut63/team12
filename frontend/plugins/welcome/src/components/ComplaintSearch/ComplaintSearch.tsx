import React, { useState, useEffect } from 'react';
import { makeStyles, withStyles, createStyles, Theme } from '@material-ui/core/styles'
import { Content, Header, Page, pageTheme, SidebarPage } from '@backstage/core';
import { AppSidebar } from '../Sidebar/Sidebar';
import { UserHeader } from '../UserHeader/UserHeader';
import { Grid, Typography, AppBar, Toolbar, MenuItem, Select,
         IconButton, Container, Paper, Table,
         TableBody, TableCell, TableContainer, TableHead, TableRow} from '@material-ui/core';
import SearchIcon from '@material-ui/icons/Search';
import { DefaultApi } from '../../api/apis';
import { EntClub, EntComplaintType, EntComplaint } from '../../api/models';
import { Alert } from '@material-ui/lab';
import Snackbar from '@material-ui/core/Snackbar';
import ClearIcon from '@material-ui/icons/Clear';
import ReportProblemTwoToneIcon from '@material-ui/icons/ReportProblemTwoTone';
import FiberManualRecordTwoToneIcon from '@material-ui/icons/FiberManualRecordTwoTone';
import moment from 'moment';
import Swal from 'sweetalert2';

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

export default function ComplaintSearch() {

    const classes = useStyles();
    const api = new DefaultApi();
    const HeaderCustom = { minHeight: '50px', };

    const [loading, setLoading] = useState(true);
    const [clubID, SetClubID] = useState(Number);
    const [typeID, SetTypeID] = useState(Number);
    const [search, SetSearch] = useState(Boolean);
    // const [clubs, SetClubs] = useState("");
    // const [types, SetTypes] = useState(String);
    const [alert, SetAlert] = useState(Boolean);
    const [status, SetStatus] = useState(false);
    const [open, setOpen] = React.useState(false);

    const [club, SetClub] = useState<EntClub[]>([]);
    const [type, SetType] = useState<EntComplaintType[]>([]);
    // const [complaint, SetComplaint] = useState<EntComplaint[]>([]);
    const [data, SetData] = useState<EntComplaint[]>([]);

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

    const ClubIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        SetClubID(event.target.value as number);
        // SetClubData(event.target.value as number);
        
    };

    const TypeIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        SetTypeID(event.target.value as number);
        // SetTypeData(event.target.value as number);
    };

    // var uname = JSON.parse(String(localStorage.getItem('user-name')))

    function handleClose() {
        setOpen(false);
    };

    // function SearchTrue() {
    //     SetStatus(true)
    //     SetAlert(true)
    //     setOpen(true)
    //     SetSearch(true)
    // };

    // function SearchFalse() {
    //     SetStatus(true)
    //     SetAlert(false)
    //     setOpen(true)
    //     SetSearch(false)
    // }

    // const SetClubData = async (x: number) => {
    //     await club.map((item: EntClub) => {
    //         if (item.id == x) {
    //             const res = String(item.name)
    //             SetClubs(res);
    //         }
    //     })
    // };

    // const SetTypeData = async (x: number) => {
    //     await type.map((item: EntComplaintType) => {
    //         if (item.id == x) {
    //             const res = String(item.description)
    //             SetTypes(res);
    //         }
    //     });
    // };
    
    var userID = JSON.parse(String(localStorage.getItem('user-id')));
    let uid = Number(userID);

    const SearchComplaint = async () => {
        SetSearch(false);
        const res = await api.getComplaint({ id: uid, cid: clubID, tid: typeID });
        SetData(res);
        console.log(res);
        var x = res.length;
        if (x > 0){
            SetSearch(true);
            SetStatus(true);
            SetAlert(true);
            Toast.fire({
                icon: 'success',
                title: 'ค้นหาข้อมูลสำเร็จ',
            });
        }
        else {
            SetSearch(false);
            SetStatus(true);
            SetAlert(false);
            Toast.fire({
                icon: 'error',
                title: 'ค้นหาข้อมูลไม่สำเร็จ',
            });
        }
    }

    useEffect(() => {
        const checkUserType = async () => {
            const userType = JSON.parse(
              String(localStorage.getItem('user-type')),
            );
            setLoading(false);
            if (userType != "นักศึกษา" && userType != "เจ้าหน้าที่") {
              Swal.fire({
                title: 'คุณไม่สามารถค้นหาเรื่องร้องเรียนได้',
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
            }
          };
        checkUserType();

        const getClub = async () => {
            const res = await (await api.listClub({ limit: 10, offset: 0 }));
            setLoading(false);
            SetClub(res);
        }
        getClub();

        const getType = async () => {
            const res = await api.listComplainttype({ limit: 10, offset: 0 });
            setLoading(false);
            SetType(res);
        };
        getType();

        // const getComplaint = async () => {
        //     const res = await (await api.listComplaint({ limit: 10, offset: 0 }));
        //     setLoading(false);
        //     SetComplaint(res);
        // };
        // getComplaint();

    }, [loading]);

    return (
        <SidebarPage>
            <AppSidebar />
            <Page theme={pageTheme.home}>
                <Header style={HeaderCustom} title={`ระบบค้นหาเรื่องร้องเรียน`}>
                    <UserHeader />
                </Header>
                <Content>
                    <Container maxWidth="md">
                        <Grid container spacing={3}>
                            <Grid item xs={12}>
                                <AppBar position="static">
                                    <Toolbar style={{ backgroundColor: "#217567" }}>
                                        <FiberManualRecordTwoToneIcon style={{ fontSize: "15" }} />
                                        <Typography variant="h6" noWrap style={{ marginRight: "0.5cm", marginLeft: "0.1cm"}}>
                                            ตัวเลือกการค้นหา
                                        </Typography>
                                    </Toolbar>
                                </AppBar>
                            </Grid>
                            <Grid item xs={12} style={{ marginTop: "0.1cm", marginBottom: "0.5cm" }}>
                                    <Grid item xs={12} sm container>
                                        <Grid item xs container direction="column" spacing={2}>
                                            <Grid item xs style={{ marginLeft: "4cm" }}>
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
                                                        <MenuItem key={0} value={0}>ทั้งหมด</MenuItem>
                                                        {club.map((item: EntClub) => (
                                                            <MenuItem key={item.id} value={item.id}>{item.name}</MenuItem>
                                                        ))}
                                                    </Select>
                                                </form>
                                            </Grid>
                                        </Grid>
                                        <Grid item xs>
                                            <Typography gutterBottom variant="subtitle1">
                                                ประเภท
                                            </Typography>

                                            <form className={classes.root_club} noValidate autoComplete="off">
                                                <Select
                                                    id="complaint type"
                                                    value={typeID || 0}
                                                    onChange={TypeIDhandleChange}
                                                >
                                                    <MenuItem key={0} value={0}>ทั้งหมด</MenuItem>
                                                    {type.map((item: EntComplaintType) => (
                                                        <MenuItem key={item.id} value={item.id}>{item.description}</MenuItem>
                                                    ))}

                                                </Select>
                                            </form>
                                        </Grid>
                                        <Grid item xs style={{ marginLeft: "1cm" }}>
                                        <IconButton color="inherit" component="span" onClick={SearchComplaint} >
                                          <SearchIcon />
                                              ค้นหา
                                        </IconButton>
                                        </Grid>
                                    </Grid>

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
                                                <Typography gutterBottom variant="subtitle1" style={{marginTop:"0.2cm", marginLeft:"10cm"}}>
                                                    <ReportProblemTwoToneIcon style={{color:"red",marginRight:"0.2cm"}}/>
                                                    ไม่พบเรื่องร้องเรียน
                                                </Typography>
                                            </Grid>
                                            </>
                                        )}
                                </div>
                            ) : null}
                        </Grid>
                    </Container>
                    {search ? (
                                    <TableContainer component={Paper} style={{ marginTop: "0.2cm" }}>
                                        <Table className={classes.table} aria-label="simple table">
                                            <TableHead>
                                                <TableRow>
                                                    <StyledTableCell align="center">หมายเลขคำร้อง</StyledTableCell>
                                                    <StyledTableCell align="center">ชื่อผู้ร้องเรียน</StyledTableCell>
                                                    <StyledTableCell align="center">เบอร์โทรศัพท์</StyledTableCell>
                                                    <StyledTableCell align="center">วันที่</StyledTableCell>
                                                    <StyledTableCell align="center">ชมรม</StyledTableCell>
                                                    <StyledTableCell align="center">ประเภท</StyledTableCell>
                                                    <StyledTableCell align="center">ข้อมูล</StyledTableCell>
                                                </TableRow>
                                            </TableHead>
                                            <TableBody>{
                                                data.map((item: EntComplaint) => (
                                                    <TableRow key={item.id}>
                                                        <TableCell align="center">{item.id}</TableCell>
                                                        <TableCell align="center">{item.name}</TableCell>
                                                        <TableCell align="center">{item.phonenumber}</TableCell>
                                                        <TableCell align="center">{moment(item.date).format('DD/MM/YYYY')}</TableCell>
                                                        <TableCell align="center">{item.edges?.complaintToClub?.name}</TableCell>
                                                        <TableCell align="center">{item.edges?.complaintToComplaintType?.description}</TableCell>
                                                        <TableCell align="center">{item.info}</TableCell>
                                                    </TableRow>
                                                ))
                                            }
                                            </TableBody>
                                        </Table>
                                    </TableContainer>
                                ) : (null)}
                </Content>
            </Page>
        </SidebarPage >
    );
}