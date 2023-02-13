import React from 'react';
import { Box, Button, Tooltip, Typography } from '@mui/material';
import { InlineOpenInNewIcon } from '@tagioalisi/services/styleUtils';
import { useHelloQuery } from '@tagioalisi/services/endpoints/hello';

const BOT_PERMISSIONS = "275149741136";

export default () => {
    const helloData = useHelloQuery();
    const [linkURL, setLinkURL] = React.useState<string>("");
    const [hoverMessage, setHoverMessage] = React.useState<string>("");

    React.useEffect(() => {
        const clientID = helloData.result?.discordClientId ?? "";
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
        if (helloData.pending) {
            setHoverMessage("Looking up bot details...");
        } else if (helloData.done && helloData.error) {
            setHoverMessage(`Error: ${helloData.error}`);
            console.error(helloData.error);
        } else {
            setHoverMessage(`Error: invalid state`);
        }
    });


    return (
        <Box sx={{ display: 'flex', justifyContent: 'center', padding: 3 }}>
            <Tooltip title={hoverMessage}>
                <span> {/* span required to be able to trigger the tooltip*/}
                <Button variant='contained' color='success' disabled={!linkURL} href={linkURL} target="_blank" >
                    <Typography variant='h5'>
                        Install Tagioalisi
                        <InlineOpenInNewIcon />
                    </Typography>
                </Button>
                </span>
            </Tooltip>
        </Box>);
}