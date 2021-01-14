import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import WatchVideo from './components/WatchVideo';
import SignIn from './components/SignIn';
import Activities from './components/Activities';
import ActivityTable from './components/ActivityTable';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/watch_video', WatchVideo);
    router.registerRoute('/signin', SignIn);
    router.registerRoute('/Activities', Activities);
    router.registerRoute('/ActivityTable', ActivityTable);
  },
});
