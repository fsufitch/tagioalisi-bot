import React from 'react';
import { Avatar, Box, Grid, type SxProps, Typography, Theme, Stack, Icon } from '@mui/material';
import { usePromiseEffect } from 'tagioalisi/services/async';
import { InlineOpenInNewIcon } from 'tagioalisi/services/styleUtils';
import { useHelloQuery } from 'tagioalisi/services/endpoints/hello';
import AddToServerBanner from './AddToServerBanner';

const sxLogo: SxProps<Theme> = {
    width: {
        xs: '25vw',
        md: '80%',
    },
    height: 'auto',
};

export default () => {
    const [tagiLogo] = usePromiseEffect(() => import('tagioalisi/resources/cicada-avatar.png').then(it => it.default), []);
    const [discordLogo] = usePromiseEffect(() => import('tagioalisi/resources/discord-color.svg').then(it => it.default), []);

    const helloData = useHelloQuery();

    return (
        <Grid container>
            <Grid item xs={12} md={9}>
                <Stack spacing={2}>
                    <Typography variant='h2'>Embrace the Cicada</Typography>
                    <Typography variant='overline'>
                        There are many Discord bots. This is one of them.
                    </Typography>
                    <Typography variant='body1'>
                        "Tagioalisi" is Hawaiian for "the call of cicadas". An appropriate name for a
                        chatbot that lurks in the background and makes unwelcome noisy contributions, right?
                    </Typography>
                    <Typography variant='body1'>
                        Fortunately, <em>Tagioalisi</em> (the bot) is much nicer than that. It helps you look up words
                        in the dictionary, find Wikipedia (and other wiki) articles, manage user groups, post
                        memes, and more. It is easy to install, friendly to use, and hard to forget. It strives to be much closer
                        to its <a href="https://www.schlockmercenary.com/2009-04-23" target='_blank'>
                            actual namesake<InlineOpenInNewIcon /></a> than to the namesake's namesake.
                    </Typography>
                    <Typography variant='body1'>
                        Go on, let <em>Tagioalisi</em> into your life (and your Discord server).
                    </Typography>

                    <AddToServerBanner helloData={helloData} />
                </Stack>
            </Grid>

            <Grid item xs={12} md={3} display="flex">
                <Stack direction={{ xs: 'row', md: 'column' }} spacing={2} sx={{ flexGrow: 1, alignItems: 'center', justifyContent: 'space-evenly' }}>
                    <Avatar src={tagiLogo} sx={{ ...sxLogo }} />
                    <Box component='img' src={discordLogo} sx={{ ...sxLogo }} />
                </Stack>
            </Grid>

        </Grid>
        )
}
