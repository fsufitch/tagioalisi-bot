import React from 'react';
import { Box, Paper, Typography, Stack, Link } from '@mui/material';
import { useAuthentication, AuthData } from 'tagioalisi/services/auth';

export default () => {
    const [auth, login] = useAuthentication();
    return <Paper sx={{ padding: 2 }}>
        <Typography variant="h5">Your Servers</Typography>
        {
            !!auth.id ?
                <AuthenticatedYourServers auth={auth} />
                :
                <Typography>
                    <Link href="#" onClick={() => login()}>Log in</Link> to Tagioalisi to see your servers that support it.
                </Typography>
        }
    </Paper>
}

const AuthenticatedYourServers = (props: { auth: AuthData }) => {
    return <Typography>
        Hello, <em>{props.auth.fullname}</em>! This section is under construction for now. Come back soon!
    </Typography>;
}