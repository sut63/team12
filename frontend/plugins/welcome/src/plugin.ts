import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import WatchVideo from './components/WatchVideo';
import SignIn from './components/SignIn';
import Activities from './components/Activities';
import ActivityTable from './components/ActivityTable';
import Club from './components/club';
import ClubApplication from './components/ClubApplication';
import Complaint from './components/Complaint';
import ComplaintSearch from './components/ComplaintSearch';
import Roomuse from './components/Roomuse';
import RoomuseSearch from './components/RoomuseSearch';
import ClubappSearch from './components/ClubappSearch';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/Welcome', WelcomePage);
    router.registerRoute('/', SignIn);
    router.registerRoute('/ClubApplication', ClubApplication);
    router.registerRoute('/Club', Club);
    router.registerRoute('/Roomuse', Roomuse);
    router.registerRoute('/ActivityTable', ActivityTable);
    router.registerRoute('/Complaint', Complaint);
    router.registerRoute('/ComplaintSearch', ComplaintSearch);
    router.registerRoute('/Activities', Activities);
    router.registerRoute('/ClubappSearch', ClubappSearch);
    router.registerRoute('/RoomuseSearch', RoomuseSearch);
  },
});
