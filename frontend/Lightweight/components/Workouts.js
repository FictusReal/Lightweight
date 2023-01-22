import * as React from 'react';
import { Provider as PaperProvider, Text, Appbar } from 'react-native-paper';

const Workouts = () => {
    return (
        <Appbar.Header>
            <Appbar.Content title="Workouts" />
        </Appbar.Header>
    );
};

export default Workouts;
