import * as React from 'react';
import { BottomNavigation, Text } from 'react-native-paper';
import Workouts from './Workouts.js';
import Stats from './Stats.js';
import Friends from './Friends.js';

const FriendsRoute = () => <Friends/>;

const WorkoutsRoute = () => <Workouts/>;

const StatsRoute = () => <Stats/>;

const NavBar = () => {
  const [index, setIndex] = React.useState(0);
  const [routes] = React.useState([
    { key: 'friends', title: 'Friends', focusedIcon: 'account-group', unfocusedIcon: 'account-group-outline'},
    { key: 'workouts', title: 'Workouts', focusedIcon: 'dumbbell' },
    { key: 'stats', title: 'Stats', focusedIcon: 'chart-box', unfocusedIcon: 'chart-box-outline' },
  ]);

  const renderScene = BottomNavigation.SceneMap({
    friends: FriendsRoute,
    workouts: WorkoutsRoute,
    stats: StatsRoute,
  });

  return (
    <BottomNavigation
      navigationState={{ index, routes }}
      onIndexChange={setIndex}
      renderScene={renderScene}
      compact={true}
      labeled={false}
      shifting={true}
    />
  );
};

export default NavBar;
