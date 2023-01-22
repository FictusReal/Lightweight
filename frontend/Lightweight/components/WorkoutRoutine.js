
import * as React from 'react';
import { ScrollView, View, Pressable } from 'react-native'
import { MD3DarkTheme, Provider as PaperProvider, Appbar, Card, Text, Button, Divider, Avatar, Portal } from 'react-native-paper';

const Workouts = () => {
    var workouts = [];
    var numWorkouts= 3;

	for(let i = 0; i < numWorkouts; i++){
        workouts.push(
        <Card key={i} style={{marginHorizontal: 15, marginVertical: 8}} onPress={() => {console.log(`Go to workout ${i+1}'s page`);}}>
            <Card.Content>
                <Text variant="titleLarge">Workout {i + 1}</Text>
            </Card.Content>
        </Card>
        )
	}
    return (
        <PaperProvider theme={MD3DarkTheme}>
        <Appbar.Header>
            <Appbar.Content title=""/>
            <Pressable onPress={() => {console.log(`Go to user account`);}}>
                <Avatar.Image style={{margin:15}} size={32} source={require('../assets/default_avatar.png')}/>
            </Pressable>
        </Appbar.Header>
        <ScrollView>
            { workouts }
            <Divider style={{margin: 15}}/>
            <View style={{position: 'relative',
                alignItems: 'center',
                justifyContent: 'center'}}
                >
            <Button  icon="dumbbell" mode="contained-tonal" onPress={() => console.log('Go to add workout form')}>
                Add Workout
            </Button>
        </View>
        </ScrollView>
        </PaperProvider>
    );
};

export default Workouts;
