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
    Webhook as WebhookIcon,
} from '@mui/icons-material';
import { useHelloQuery } from '@tagioalisi/services/endpoints/hello';
import { InlineOpenInNewIcon } from '@tagioalisi/services/styleUtils';
import { useClient } from '@tagioalisi/services/grpc';
import { GreeterDefinition } from '@tagioalisi/proto/hello';
import { usePromiseEffect } from '@tagioalisi/services/async';

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

    const [grpcWorks, setGrpcWorks] = React.useState<boolean>(false);
    const greeterClient = useClient(GreeterDefinition);
    usePromiseEffect(async () => {
        try {
            const reply = await greeterClient.sayHello({name: "TEST GREETER"});
            if (!reply.message.includes("TEST GREETER")) {
                throw "Server replied with wrong message";
            }
            setGrpcWorks(true);
        } catch (err) {
            console.error(err);
            setGrpcWorks(false);
        }
    }, [helloTriggerCounter])

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
            <ListItem>
                <ListItemIcon>
                    <WebhookIcon />
                </ListItemIcon>
                <ListItemText> Websocket gRPC-Web connection: {grpcWorks ? 'online' : 'offline'} </ListItemText>
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