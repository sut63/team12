import React, { FC } from 'react';
import { Typography, Grid } from '@material-ui/core';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  SidebarPage,
} from '@backstage/core';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import { UserHeader } from '../UserHeader/UserHeader';
import { AppSidebar } from '../Sidebar/Sidebar';
// import FiberManualRecordIcon from '@material-ui/icons/FiberManualRecord';
const HeaderCustom = {
  minHeight: '50px',
};

const useStyles = makeStyles({
  root: {
    maxWidth: 345,
  },
});

export type ProfileProps = {
  name: string; 
  id: string;
  system: string;
};

export function CardTeam({ name, id, system }: ProfileProps) {
  const classes = useStyles();
  return (
    <Grid item xs={12} md={3}>
      <Card className={classes.root}>
        <CardActionArea>
          <CardMedia
            component="img"
            style={{ height: "0" , paddingTop: "56.25%" }}
            image="../../image/UnitA.jpg"
            title="ชมรม"
          />
          <CardContent>
            <Typography gutterBottom variant="h6" component="h2">
              {system}
            </Typography>
            <Typography gutterBottom variant="h6" component="h2">
              {id} {name}
            </Typography>
          </CardContent>
        </CardActionArea>
      </Card>
    </Grid>
    // <Grid container spacing={2}>
    //     <Grid item >
    //         <FiberManualRecordIcon style={{ fontSize: 10, color: "black" }} />
    //     </Grid>
    //     <Grid item xs={12} sm container>
    //         <Grid item xs={12}>
    //             <Typography gutterBottom variant="subtitle1" style={{ color: 'black' }}>
    //               {system} {id}     {name}
    //             </Typography>
    //         </Grid>
    //     </Grid>
    // </Grid>
    
  );
}

const WelcomePage: FC<{}> = () => {
  return (
    <Page theme={pageTheme.home}>
      <SidebarPage>
       <AppSidebar/>
      <Header style={HeaderCustom} title={`ระบบจัดการชมรม`}><UserHeader/></Header>
      <Content>
        <ContentHeader title="สมาชิกในกลุ่ม"></ContentHeader>
        <Grid container>
          <CardTeam name={"นาย ธนวัฒน์ พรมกุล"} id={"B6025212"} system={"ระบบย่อย: ระบบยืมอุปกรณ์"}></CardTeam>
          <CardTeam name={"นาย พศิน บัวสิงห"} id={"B6108755"} system={"ระบบย่อย: ระบบสร้างชมรม"}></CardTeam>
          <CardTeam name={"นาย ธนภูมิ พงษ์ประชา"} id={"B6110321"} system={"ระบบย่อย: ระบบสมัครเข้าชมรม"}></CardTeam>
          <CardTeam name={"นาย บูรพา ภูสามารถ"} id={"B6111922"} system={"ระบบย่อย: ระบบร้องเรียน"}></CardTeam>
          <CardTeam name={"นาย ศุภศิลป์ ย่อแสง"} id={"B6112608"} system={"ระบบย่อย: ระบบจัดกิจกรรมชมรม"}></CardTeam>
          <CardTeam name={"นาย ธีร์ธวัช ศรีคอนไทย"} id={"B6131678"} system={"ระบบย่อย: ระบบการเข้าใช้ห้อง"}></CardTeam>
        </Grid>
      </Content>
      </SidebarPage>
    </Page>
    
  );
};

export default WelcomePage;
