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
    <SidebarItem icon={HomeIcon} to="Welcome" text="Home" />
    {/* <SidebarItem icon={CreateComponentIcon} to="create" text="Create..." />
    <SidebarItem icon={CreateComponentIcon} to="welcome" text="Welcome" /> */}
    <SidebarDivider />

    
    <SidebarItem icon={CreateComponentIcon} to="club" text="สร้างชมรม" />
    <SidebarItem icon={CreateComponentIcon} to="ClubApplication" text="สมัครชมรม" />
    <SidebarItem icon={CreateComponentIcon} to="Activities" text="จัดกิจกรรม" />
    <SidebarItem icon={CreateComponentIcon} to="Roomuse" text="จองห้องชมรม" />
    <SidebarItem icon={CreateComponentIcon} to="Complaint" text="ร้องเรียน" />
 <SidebarItem icon={CreateComponentIcon} to="clubsearch" text="ค้นหาสร้างชมรม" />
    {/* End global nav */}
    <SidebarDivider />
    <SidebarSpace />
    <SidebarDivider />
    <SidebarItem icon={SignOut} to="" text="Sign Out" />
    <SidebarItem icon={CreateComponentIcon} to="" text="Sign In" />
    {/* <SidebarUserSettings  /> */}
    <SidebarPinButton />
  </Sidebar>
);
