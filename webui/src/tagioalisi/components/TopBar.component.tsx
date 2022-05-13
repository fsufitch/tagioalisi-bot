import { AppBar, Box, Button, IconButton, Typography, Toolbar, Menu, MenuItem } from "@mui/material";
import { Menu as MenuIcon } from '@mui/icons-material';
import React, { MouseEvent } from "react";

import { NavLink, Link, Navigate, useNavigate } from "react-router-dom";

import { usePromiseEffect } from 'tagioalisi/services/async';
import { ROUTES } from '../routes';


const getStyles = () => {
  const [styles] = usePromiseEffect(() => import('./TopBar.module.scss').then(m => m.default));
  return styles || {};
}

export const TopBar = () => {
  const styles = getStyles();
  return (
    <>
      <Box sx={{ flexGrow: 1 }}>
        <AppBar position="sticky">
          <Toolbar>
            <DropDownNav />
            <Title />
            <AuthSegment />
          </Toolbar>
        </AppBar>
      </Box>
    </>
  );
}

const Title = () => {
  const styles = getStyles();
  const [logo] = usePromiseEffect(() => import('tagioalisi/resources/cicada-avatar.png').then(it => it.default))
  return (<>
    <Typography variant="h6" component="div" className={styles?.titleBox} sx={{ flexGrow: 1, display: { xs: 'none', md: 'flex' } }}>
      {/* Show a wide title when the screen is wide */}
      <span>Tagioalisi</span>
      <img src={logo} className={styles?.logo} />
      <span>Discord Bot</span>
    </Typography>
    <Typography variant="h6" component="div" className={styles?.titleBox} sx={{ flexGrow: 1, display: { xs: 'flex', md: 'none' } }}>
      {/* Show a wide title when the screen is wide */}
      <img src={logo} className={styles?.logo} alt="Tagioalisi - Discord Bot" />
    </Typography>
  </>
  );
}

const DropDownNav = () => {
  const [navAnchor, setNavAnchor] = React.useState<null | HTMLElement>(null);

  const openNav = (event: MouseEvent<HTMLElement>) => setNavAnchor(event.currentTarget);
  const closeNav = () => setNavAnchor(null);

  const navigate = useNavigate();
  const navigateAndClose = (location: string) => {
    closeNav();
    navigate(location);
  }

  return <Box sx={{ flexGrow: 1}}>
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
        ROUTES.map((route, idx) =>
          <MenuItem key={"" + idx} onClick={() => navigateAndClose(route.url)}>
            <Typography textAlign="center">{route.navText}</Typography>
          </MenuItem>
        )}
    </Menu>
  </Box>

}

const AuthSegment = () => {
  return <>
    <Button color="inherit">Login</Button>
  </>;
}