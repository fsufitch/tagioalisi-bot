import React from 'react';

import { Box, Typography, Button, Paper, Stack, Avatar, IconButton, FormControl, OutlinedInput, InputLabel, InputAdornment, Tooltip } from '@mui/material';
import { Api as ApiIcon, Undo as UndoIcon } from '@mui/icons-material';

import { APIConfigurationContext, useDefaultBaseURL } from 'tagioalisi/contexts/APIConfiguration';

export default () => {
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
    const defaultBaseURL = useDefaultBaseURL();
    const {configuration, setBaseURL} = React.useContext(APIConfigurationContext);
    const [editMode, setEditMode] = React.useState<boolean>(false);
    const [formBaseUrl, setFormBaseUrl] = React.useState<string>(configuration?.baseURL ?? '');

    const startEdit = () => {
        setEditMode(true);
        setFormBaseUrl(configuration?.baseURL ?? '');
    }
    const discardEdit = () => {
        setEditMode(false);
    }
    const saveEdit = () => {
        setBaseURL(formBaseUrl);
        setEditMode(false);
    }

    return <Paper sx={{ p: 1, display: 'flex', flexDirection: 'column', alignContent: 'space-between' }}>
        <FormControl variant='outlined' sx={{ m: 1 }}>
            <InputLabel htmlFor='base-api-url'>Base API URL</InputLabel>
            <OutlinedInput
                label='Base API URL'
                value={formBaseUrl}
                onChange={(e) => editMode ? setFormBaseUrl(e.target.value) : null}
                startAdornment={
                    <InputAdornment position='start'>
                        <Avatar><ApiIcon /></Avatar>
                    </InputAdornment>
                }
                endAdornment={!editMode ? <></> :
                    <InputAdornment position='end'>
                        <Tooltip title="Reset to default">
                            <IconButton onClick={() => setFormBaseUrl(defaultBaseURL)}><UndoIcon /></IconButton>
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