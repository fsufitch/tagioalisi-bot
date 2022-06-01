import React from 'react';
import { Paper, Typography, Link } from '@mui/material';
import { AuthenticationContext } from 'tagioalisi/contexts/Authentication';
import { Authentication } from 'tagioalisi/contexts/Authentication';

export default () => {
    const { authentication, login} = React.useContext(AuthenticationContext);
    return <Paper sx={{ padding: 2 }}>
        <Typography variant="h5">Your Servers</Typography>
        {
            !!authentication.id ?
                <AuthenticatedYourServers auth={authentication} />
                :
                <Typography>
                    <Link href="#" onClick={() => login()}>Log in</Link> to Tagioalisi to see your servers that support it.
                </Typography>
        }
    </Paper>
}

const AuthenticatedYourServers = (props: { auth: Authentication }) => {
    return <Typography>
        Hello, <em>{props.auth.fullname}</em>! This section is under construction for now. Come back soon!
    </Typography>;
}