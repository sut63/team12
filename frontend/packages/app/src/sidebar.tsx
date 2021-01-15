import React from 'react';
import HomeIcon from '@material-ui/icons/Home';
import CreateComponentIcon from '@material-ui/icons/AddCircleOutline';
import YouTube from '@material-ui/icons/YouTube';
import SignOut from '@material-ui/icons/Settings';

import {
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  SidebarUserSettings,
  SidebarThemeToggle,
  SidebarPinButton,
} from '@backstage/core';

export const AppSidebar = () => (
  <Sidebar>
    <SidebarDivider />
    {/* Global nav, not org-specific */}
    <SidebarItem icon={HomeIcon} to="" text="Home" />
    {/* <SidebarItem icon={CreateComponentIcon} to="create" text="Create..." />
    <SidebarItem icon={CreateComponentIcon} to="welcome" text="Welcome" /> */}
    <SidebarItem icon={YouTube} to="playlist_video" text="Playlist Video" />
    <SidebarItem
      icon={CreateComponentIcon}
      to="watch_video"
      text="Watch Video"
    />
    <SidebarItem icon={CreateComponentIcon} to="Activities" text="จัดกิจกรรม" />
    <SidebarItem icon={CreateComponentIcon} to="club" text="สร้างชมรม" />
    <SidebarItem icon={CreateComponentIcon} to="Complaint" text="ร้องเรียน" />
    {/* End global nav */}
    <SidebarDivider />
    <SidebarSpace />
    <SidebarDivider />
    <SidebarItem icon={SignOut} to="sign_out" text="Sign Out" />
    <SidebarItem icon={CreateComponentIcon} to="signin" text="Sign In" />
    {/* <SidebarUserSettings  /> */}
    <SidebarPinButton />
  </Sidebar>
);
