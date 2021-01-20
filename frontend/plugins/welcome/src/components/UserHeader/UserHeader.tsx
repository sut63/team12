import React from 'react';
import PersonPinIcon from '@material-ui/icons/PersonPin';
import { Grid, Typography, } from '@material-ui/core';

var name = JSON.parse(String(localStorage.getItem('user-name')));
var status = JSON.parse(String(localStorage.getItem('user-status')));

export const UserHeader = () => (

    <Grid container spacing={2}>
        <Grid item >
            <PersonPinIcon style={{ fontSize: 30, color: "ffffff" }} />
        </Grid>
        <Grid item xs={12} sm container>
            <Grid item xs>
                <Typography gutterBottom variant="subtitle1" style={{ color: 'white' }}>
                    User: {name}
                </Typography>
                {/* <Typography gutterBottom variant="subtitle1" style={{ color: 'white' }}>
                   Status: {status}
                </Typography> */}
            </Grid>
        </Grid> 
    </Grid>
);
