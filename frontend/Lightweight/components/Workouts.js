import * as React from 'react';
import { Text, Appbar } from 'react-native-paper';

const Workouts = () => {
    return (
        <Appbar.Header>
            <Appbar.Content title="Workouts" />
            <Appbar.Action icon="calendar" onPress={() => {}} />
            <Appbar.Action icon="magnify" onPress={() => {}} />
        </Appbar.Header>
    );
};

export default Workouts;
