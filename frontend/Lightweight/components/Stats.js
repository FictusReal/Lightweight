import * as React from 'react';
import { Text, Appbar } from 'react-native-paper';

const Stats = () => {
    return (
        <Appbar.Header>
            <Appbar.Content title="Statistics" />
            <Appbar.Action icon="calendar" onPress={() => {}} />
            <Appbar.Action icon="magnify" onPress={() => {}} />
        </Appbar.Header>
    );
};

export default Stats;
