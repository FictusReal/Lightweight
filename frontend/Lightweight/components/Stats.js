import * as React from 'react';
import { ScrollView, View, Pressable } from 'react-native'
import { MD3DarkTheme, Provider as PaperProvider, Appbar, Card, Text, Button, Divider, Avatar, Portal } from 'react-native-paper';

const Stats = () => {
    var exercises = [];
    var numExercises= 15;

	for(let i = 0; i < numExercises; i++){
        exercises.push(
        <Card key={i} style={{marginHorizontal: 15, marginVertical: 8}} onPress={() => {console.log(`Go to exercise ${i+1}'s statistics page`);}}>
            <Card.Content>
                <Text variant="titleLarge">Statistics for Exercise {i + 1}</Text>
            </Card.Content>
        </Card>
        )
	}
    return (
        <PaperProvider theme={MD3DarkTheme}>
        <Appbar.Header>
            <Appbar.Content title="Statistics"/>
            <Pressable onPress={() => {console.log(`Go to user account`);}}>
                <Avatar.Image style={{margin:15}} size={32} source={require('../assets/default_avatar.png')}/>
            </Pressable>
        </Appbar.Header>
        <ScrollView>
            { exercises }
        </ScrollView>
        </PaperProvider>
    );
};

export default Stats;
