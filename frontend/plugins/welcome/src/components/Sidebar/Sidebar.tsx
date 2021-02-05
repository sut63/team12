import React from 'react';
import HomeIcon from '@material-ui/icons/Home';
//import CreateComponentIcon from '@material-ui/icons/AddCircleOutline';
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
import SearchIcon from '@material-ui/icons/Search';
import EventIcon from '@material-ui/icons/Event';
import PostAddIcon from '@material-ui/icons/PostAdd';
import PersonAddIcon from '@material-ui/icons/PersonAdd';
import MessageIcon from '@material-ui/icons/Message';
import EmojiEventsIcon from '@material-ui/icons/EmojiEvents';


export const AppSidebar = () => (
  <Sidebar>
    <SidebarDivider />
    {/* Global nav, not org-specific */}
    <SidebarItem icon={HomeIcon} to="/Welcome" text="Home" />
    {/* <SidebarItem icon={CreateComponentIcon} to="create" text="Create..." />
    <SidebarItem icon={CreateComponentIcon} to="welcome" text="Welcome" /> */}
    <SidebarDivider />

    
    <SidebarItem icon={PostAddIcon} to="/club" text="สร้างชมรม" />
    <SidebarItem icon={PersonAddIcon} to="/ClubApplication" text="สมัครชมรม" />
    <SidebarItem icon={EmojiEventsIcon} to="/Activities" text="จัดกิจกรรม" />
    <SidebarItem icon={EventIcon} to="/Roomuse" text="เข้าใช้ห้องชมรม" />
    <SidebarItem icon={MessageIcon} to="/Complaint" text="ร้องเรียน" />

    {/* End global nav */}
    <SidebarDivider />
    <SidebarItem icon={SearchIcon} to="/ClubappSearch" text="ค้นหาใบสมัครชมรม" />
    <SidebarSpace />
    <SidebarDivider />
    <SidebarItem icon={SignOut} to="/" text="Sign Out" />
    {/* <SidebarUserSettings  /> */}
    <SidebarPinButton />
  </Sidebar>
);