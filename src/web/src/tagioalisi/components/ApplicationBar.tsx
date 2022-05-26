import { AppBar, Box, Button, IconButton, Typography, Toolbar, Menu, MenuItem, Avatar, Tooltip, Switch } from '@mui/material';
import { Menu as MenuIcon, Logout as LogoutIcon } from '@mui/icons-material';
import React, { MouseEvent } from "react";

import { useNavigate } from "react-router-dom";

import { usePromiseEffect } from 'tagioalisi/services/async';
import { getRoute } from 'tagioalisi/routes';
import { useAuthentication } from 'tagioalisi/services/auth';
import { useSynchronizedState } from 'tagioalisi/services/state';

import { ColorModeContext } from 'tagioalisi/Theme';

export default () =>
  <Box>
    <AppBar position="sticky">
      <Toolbar>
        <DropDownNav />
        <Title />
        <AuthSegment />
      </Toolbar>
    </AppBar>
  </Box>

const Title = () => {
  const [logo] = usePromiseEffect(() => import('tagioalisi/resources/cicada-avatar.png').then(it => it.default), [])
  return (<>
    <Typography variant="h6" component="div" sx={{ alignItems: 'center', textAlign: 'center', flexGrow: 1, display: { xs: 'none', md: 'flex' } }}>
      {/* Show a wide title when the screen is wide */}
      Tagioalisi
      <Avatar src={logo} sx={{ bgcolor: 'darkred', margin: 1, width: '2.5em', height: '2.5em' }} />
      Discord Bot
    </Typography>
    <Typography variant="h6" component="div" sx={{ alignItems: 'center', textAlign: 'center', flexGrow: 1, display: { xs: 'flex', md: 'none' } }}>
      {/* Show a wide title when the screen is wide */}
      <Avatar src={logo} sx={{ bgcolor: 'darkred', margin: 1, width: '2.5em', height: '2.5em' }} alt='Tagioalisi - Discord Bot' />
    </Typography>
  </>
  );
}

const DROP_DOWN_NAV_LINKS: { routeId: string, text: string }[] = [
  { routeId: 'home', text: 'Home' },
  { routeId: 'config', text: 'API Configuration' },
]

const DropDownNav = () => {
  const [navAnchor, setNavAnchor] = React.useState<null | HTMLElement>(null);

  const openNav = (event: MouseEvent<HTMLElement>) => setNavAnchor(event.currentTarget);
  const closeNav = () => setNavAnchor(null);

  const navigate = useNavigate();
  const navigateAndClose = (routeId: string) => {
    closeNav();
    navigate(getRoute(routeId).path);
  }

  const NavLink = (props: { routeId: string, text: string }) =>
    <MenuItem onClick={() => navigateAndClose(props.routeId)}>
      <Typography textAlign="center">{props.text}</Typography>
    </MenuItem>;

  return <Box sx={{ flexGrow: 1 }}>
    <IconButton
      size="large"
      aria-label="application menu"
      aria-controls="menu-appbar"
      aria-haspopup="true"
      onClick={openNav}
      color="inherit"
    >
      <MenuIcon />
    </IconButton>
    <Menu
      id="menu-appbar"
      anchorEl={navAnchor}
      anchorOrigin={{
        vertical: 'bottom',
        horizontal: 'left',
      }}
      keepMounted
      transformOrigin={{
        vertical: 'top',
        horizontal: 'left',
      }}
      open={!!navAnchor}
      onClose={closeNav}
    >
      {
        DROP_DOWN_NAV_LINKS.map(({ routeId, text }) =>
          <NavLink key={routeId} routeId={routeId} text={text} />
        )
      }
      <DisplayModeToggle />
    </Menu>
  </Box>

}

const AuthSegment = () => {
  const [auth, login, logout] = useAuthentication();

  React.useEffect(() => console.log(auth), [auth]);
  return !auth.id ?
    <Button color="inherit" onClick={() => login()}>Login</Button>
    :
    <>
      <Tooltip title={auth.fullname ?? ''}>
        <Avatar src={auth.avatarUrl} />
      </Tooltip>
      <IconButton onClick={() => logout()}>
        <LogoutIcon />
      </IconButton>
    </>;
}

const DisplayModeToggle = () => {
  const colorModeContext = React.useContext(ColorModeContext);
  const colorMode = colorModeContext.getColorMode();
  const toggleMode = () => colorMode == 'light' ? colorModeContext.setColorMode('dark') : colorModeContext.setColorMode('light');

  return <MenuItem onClick={toggleMode}>
    <Typography textAlign="center">
      <Switch checked={colorMode == 'dark'} />
      {
        colorMode == 'light' ? "Light mode" :
          colorMode == 'dark' ? "Dark mode" :
            "Default mode"
      }
    </Typography>
  </MenuItem>;
}