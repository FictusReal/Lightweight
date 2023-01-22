import * as React from 'react';
import { Text, Appbar } from 'react-native-paper';

const Friends = () => {
    return (
        <Appbar.Header>
            <Appbar.Content title="Friends" />
            <Appbar.Action icon="calendar" onPress={() => {}} />
            <Appbar.Action icon="magnify" onPress={() => {}} />
        </Appbar.Header>
    );
};

export default Friends;
