import React from 'react';
import { createStackNavigator } from '@react-navigation/stack';

import { Feed } from './feed';
import { Details } from './details';

export const FeedStack = () => {
  return (
    <Stack.Navigator initialRouteName="Workouts">
      <Stack.Screen
        name="Workouts"
        component={Workouts}
        options={{ headerTitle: 'Workouts' }}
      />
      <Stack.Screen
        name="Workout"
        component={Details}
        options={{ headerTitle: 'Workout' }}
      />
    </Stack.Navigator>
  );
};