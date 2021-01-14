import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import WatchVideo from './components/WatchVideo';
import SignIn from './components/SingIn';
import Activities from './components/Activities';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/welcome', WelcomePage);
    router.registerRoute('/watch_video', WatchVideo);
    router.registerRoute('/', SignIn);
    router.registerRoute('/Activities', Activities);
  },
});
