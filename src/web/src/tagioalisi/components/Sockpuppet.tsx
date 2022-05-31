import React from 'react';
import { Typography, Stack, Paper, Alert } from '@mui/material';

import { useSockpuppetClient } from 'tagioalisi/services/grpc';
import { useAuthentication } from 'tagioalisi/services/auth';

export default () => {
    const [authData] = useAuthentication();
    
    return (
        <Stack spacing={2}>
            <Typography variant='h3'>
                Sockpuppet &mdash; Make Tagioalisi Say Stuff
            </Typography>
            <Typography variant='body1'>
                Isn't it fun to make a bot say wacky stuff? Of course it is.
                Use "Copy ID" in the Discord UI to find the channel ID (a long number).
                Tagioalisi does not support using a channel name for this.
            </Typography>
            <Typography variant='body1'>
                To mention a user, use <code>&lt;@the-user-id&gt;</code>.
                To mention a channel, use <code>&lt;@the-channel-id&gt;</code>.
            </Typography>
            <Typography variant='body1'>
                <strong>Note:</strong> To curb potential spam, you are required to have  
                "Manage Messages" permission in the target channel.
            </Typography>
            <Paper sx={{padding: 2}}>
                {
                    !authData.jwt
                        ? <Alert severity='error'>
                            You must be logged in to use the Sockpuppet module.
                        </Alert>
                        :
                        <SockpuppetForm />
                }
            </Paper>
        </Stack>
    );
}

const SockpuppetForm = () => {
    const [authData] = useAuthentication();
    const client = useSockpuppetClient();


    return <> form goes here </>;
}