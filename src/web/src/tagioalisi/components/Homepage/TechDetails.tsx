import React from 'react';
import { Box, List, ListItem, ListItemIcon, ListItemSecondaryAction, ListItemText, Paper, Typography, IconButton, Link } from '@mui/material';
import {
    AccessTime as AccessTimeIcon,
    BlurOff as BlurOffIcon,
    BugReport as BugIcon,
    GitHub as GitHubIcon,
    LocalFireDepartment as FireIcon,
    People as PeopleIcon,
    Refresh as RefreshIcon,
} from '@mui/icons-material';
import { useHelloQuery } from '../../services/endpoints/hello';
import { InlineOpenInNewIcon } from '../../services/styleUtils';
import { useGreeterClient } from '../../services/grpc';
import { HelloRequest } from 'tagioalisi/proto/hello_pb';

import { BrowserHeaders } from 'browser-headers';

const GITHUB_URL = 'https://github.com/fsufitch/tagioalisi-bot';

export default () => {
    const [helloTriggerCounter, setHelloTriggerCounter] = React.useState<number>(0);
    const requeryHello = () => setHelloTriggerCounter(helloTriggerCounter + 1);
    const helloData = useHelloQuery([helloTriggerCounter]);

    const [uptime, setUptime] = React.useState<string>("");
    React.useEffect(() => {
        const uptimeFloat = helloData.result?.uptimeSeconds || 0;
        if (uptimeFloat > 0) {
            setUptime(uptimeFloat.toFixed(2));
        } else {
            setUptime("");
        }
    }, [helloData]);

    const [greeterMessage, setGreeterMessage] = React.useState<string>('');
    const greeterClient = useGreeterClient();
    React.useEffect(() => {
        const req = new HelloRequest();
        req.setName("froobar frontend");
        greeterClient.sayHello(req, new BrowserHeaders(), (err, reply) => {
            if (!!err) {
                console.error(err);
                return;
            }
            if (!reply) {
                console.error('empty reply');
                return;
            }
            setGreeterMessage(reply.getMessage());
        });
    }, [])

    return <Paper>
        <List>
            <ListItem>
                <ListItemText disableTypography>
                    <Typography variant="h5">Technical Details</Typography>
                </ListItemText>
                <ListItemSecondaryAction>
                    <IconButton onClick={() => requeryHello()}>
                        <RefreshIcon />
                    </IconButton>
                </ListItemSecondaryAction>
            </ListItem>
            {!helloData.result ?
                <ListItem>
                    <ListItemIcon>
                        <FireIcon />
                    </ListItemIcon>
                    <ListItemText> Bot API server is down! </ListItemText>
                </ListItem>
                :
                <>
                    <ListItem>
                        <ListItemIcon>
                            <AccessTimeIcon />
                        </ListItemIcon>
                        <ListItemText> Uptime: {uptime} seconds </ListItemText>
                    </ListItem>
                    <ListItem>
                        <ListItemIcon>
                            <PeopleIcon />
                        </ListItemIcon>
                        <ListItemText>
                            Role prefix for "groups":
                            «<Typography component='span' fontFamily='monospace' >
                                {helloData.result.groupPrefix}
                            </Typography>»
                        </ListItemText>
                    </ListItem>
                    <ListItem>
                        <ListItemIcon>
                            <BlurOffIcon />
                        </ListItemIcon>
                        <ListItemText> Disabled modules: {' '}
                            {helloData.result.botModuleBlacklist.join(', ') || "None"}
                        </ListItemText>
                    </ListItem>
                    <ListItem>
                        <ListItemIcon>
                            <BugIcon />
                        </ListItemIcon>
                        <ListItemText> Debug mode: {helloData.result.debugMode ? "on" : "off"} </ListItemText>
                    </ListItem>
                    <ListItem>
                        <ListItemIcon>
                            <BugIcon />
                        </ListItemIcon>
                        <ListItemText> greeter message: {greeterMessage} </ListItemText>
                    </ListItem>
                </>
            }
            <ListItem>
                <ListItemIcon>
                    <GitHubIcon />
                </ListItemIcon>
                <ListItemText>
                    <Link href={GITHUB_URL} target="_blank">
                        Fork me on GitHub! <InlineOpenInNewIcon />
                    </Link>
                </ListItemText>
            </ListItem>
        </List>
    </Paper>
}