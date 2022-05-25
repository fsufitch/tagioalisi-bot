import React from 'react';
import { Box, Button, Tooltip, Typography } from '@mui/material';
import { HelloResponse } from 'tagioalisi/services/endpoints/hello';
import { InlineOpenInNewIcon } from '../../services/styleUtils';

const BOT_PERMISSIONS = "275149741136";

export default (props: { helloData: HelloResponse }) => {
    const [linkURL, setLinkURL] = React.useState<string>("");
    const [hoverMessage, setHoverMessage] = React.useState<string>("");

    React.useEffect(() => {
        const clientID = props.helloData.result?.discordClientId ?? "";
        if (!!clientID) {
            const addToServerURL = new URL("https://discord.com/api/oauth2/authorize");
            addToServerURL.searchParams.set("scope", "bot applications.commands");
            addToServerURL.searchParams.set("client_id", clientID);
            addToServerURL.searchParams.set("permissions", BOT_PERMISSIONS);
            setLinkURL(addToServerURL.toString());
            setHoverMessage("");
            return;
        }
        setLinkURL("");
        if (props.helloData.pending) {
            setHoverMessage("Looking up bot details...");
        } else if (props.helloData.done && props.helloData.error) {
            setHoverMessage(`Error: ${props.helloData.error}`);
            console.error(props.helloData.error);
        } else {
            setHoverMessage(`Error: invalid state`);
        }
    });


    return (
        <Box sx={{ display: 'flex', justifyContent: 'center', padding: 3 }}>
            <Tooltip title={hoverMessage}>
                <Button variant='contained' color='success' disabled={!linkURL} href={linkURL} target="_blank" >
                    <Typography variant='h5'>
                        Install Tagioalisi
                        <InlineOpenInNewIcon />
                    </Typography>
                </Button>
            </Tooltip>
        </Box>);
}