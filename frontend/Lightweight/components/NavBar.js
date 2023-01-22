import * as React from 'react';
import { BottomNavigation, Text } from 'react-native-paper';
import Friends from './Friends.js';
import Workouts from './Workouts.js';
import Stats from './Stats.js';

const FriendsRoute = () => <Friends/>;

const WorkoutsRoute = () => <Workouts/>;

const StatsRoute = () => <Stats/>;

const NavBar = () => {
  const [index, setIndex] = React.useState(0);
  const [routes] = React.useState([
    { key: 'friends', focusedIcon: 'account-group', unfocusedIcon: 'account-group-outline'},
    { key: 'workouts', focusedIcon: 'dumbbell' },
    { key: 'stats', focusedIcon: 'chart-box', unfocusedIcon: 'chart-box-outline' },
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
    />
  );
};

export default NavBar;
