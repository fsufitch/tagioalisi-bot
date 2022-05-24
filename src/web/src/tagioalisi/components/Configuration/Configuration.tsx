import React from 'react';

import { Box, Typography, Button, Paper, Stack, Avatar, IconButton, FormControl, OutlinedInput, InputLabel, InputAdornment, Tooltip } from '@mui/material';
import { Api as ApiIcon, Undo as UndoIcon } from '@mui/icons-material';

import { useAPIConnection, useDefaultAPIEndpoint } from 'tagioalisi/services/api';

export const Configuration = () => {
    return (
        <Box sx={{ padding: 3 }}>
            <Stack direction='row' alignItems='flex-end'>
                <Box>
                    <Typography variant='h5'>Bot API Configuration</Typography>
                    <Typography variant='body1'>
                        This web UI uses a HTTP-based API to communicate with Tagioalisi proper.
                        This section is for configuring how that is done.
                        The settings here are client-side only, and may be tweaked without affecting the backend.
                    </Typography>
                </Box>
            </Stack>

            <EditableConfiguration />

        </Box>
    );
}

const EditableConfiguration = () => {
    const defaultAPIEndpoint = useDefaultAPIEndpoint();

    const [apiConnection, setApiConnection] = useAPIConnection();

    const [editMode, setEditMode] = React.useState<boolean>(false);
    const [baseUrl, setBaseUrl] = React.useState<string>('');
    React.useEffect(() => setBaseUrl(apiConnection.baseUrl), [apiConnection.baseUrl]);

    const startEdit = () => {
        setEditMode(true);
        setBaseUrl(apiConnection.baseUrl);
    }
    const discardEdit = () => {
        setEditMode(false);
        setBaseUrl(apiConnection.baseUrl);
    }
    const saveEdit = () => {
        setApiConnection({ ...apiConnection, baseUrl });
        setEditMode(false);
    }

    return <Paper sx={{ p: 1, display: 'flex', flexDirection: 'column', alignContent: 'space-between' }}>
        <FormControl variant='outlined' sx={{ m: 1 }}>
            <InputLabel htmlFor='base-api-url'>Base API URL</InputLabel>
            <OutlinedInput
                label='Base API URL'
                value={baseUrl}
                onChange={(e) => editMode ? setBaseUrl(e.target.value) : null}
                startAdornment={
                    <InputAdornment position='start'>
                        <Avatar><ApiIcon /></Avatar>
                    </InputAdornment>
                }
                endAdornment={!editMode ? <></> :
                    <InputAdornment position='end'>
                        <Tooltip title="Reset to default">
                            <IconButton onClick={() => setBaseUrl(defaultAPIEndpoint)}><UndoIcon /></IconButton>
                        </Tooltip>
                    </InputAdornment>
                }
            />
        </FormControl>

        <Box sx={{ display: 'flex', flexDirection: 'row', justifyContent: 'end' }}>
            {editMode ? <>
                <Button sx={{ m: 1 }} variant='outlined' color='warning' onClick={() => discardEdit()}>Discard Changes</Button>
                <Button sx={{ m: 1 }} variant='contained' color='primary' onClick={() => saveEdit()}>Save Settings</Button>
            </> : <>
                <Button sx={{ m: 1 }} variant='contained' color='primary' onClick={() => startEdit()}>Edit</Button>
            </>
            }
        </Box>
    </Paper>

}