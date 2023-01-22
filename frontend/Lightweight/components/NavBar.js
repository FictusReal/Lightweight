import * as React from 'react';
import { createMaterialBottomTabNavigator } from '@react-navigation/material-bottom-tabs';
import MaterialCommunityIcons from 'react-native-vector-icons/MaterialCommunityIcons';


const Tab = createMaterialBottomTabNavigator();

import Workouts from './Workouts.js';
import Stats from './Stats.js';
import Friends from './Friends.js';

const FriendsRoute = () => <Friends/>;
const WorkoutsRoute = () => <Workouts/>;
const StatsRoute = () => <Stats/>;


const NavBar = () => {
  return (
    <Tab.Navigator 
      initialRouteName="Workout Routines"
      labeled={false}
    >
      <Tab.Screen
        options={{
            tabBarIcon: ({ focused, color, size }) => {
            let iconName = "account-group";
            return <MaterialCommunityIcons name={iconName} size={24} color={color} />;
          }
        }}
        name="Friends" component={FriendsRoute} />
      <Tab.Screen 
        options={{
            tabBarIcon: ({ focused, color, size }) => {
            let iconName = "weight-lifter";
            return <MaterialCommunityIcons name={iconName} size={24} color={color} />;
          }
        }}
        name="Workout Routines" component={WorkoutsRoute} />
      <Tab.Screen 
        options={{
            tabBarIcon: ({ focused, color, size }) => {
            let iconName = "account-group";
            return <MaterialCommunityIcons name={iconName} size={24} color={color} />;
          }
        }}
        name="Statistics" component={StatsRoute} />
    </Tab.Navigator>
  );
};

export default NavBar;
