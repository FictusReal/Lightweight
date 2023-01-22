import * as React from 'react';
import { ScrollView, View, Pressable } from 'react-native'
import { MD3DarkTheme, Provider as PaperProvider, Appbar, Searchbar, Avatar, Card, Text, IconButton } from 'react-native-paper';

const Friends = () => {
    const [searchQuery, setSearchQuery] = React.useState('');
    const onChangeSearch = query => setSearchQuery(query);

    var friends = [];
    var numFriends = 50;

	for(let i = 0; i < numFriends; i++){
        if(searchQuery == 0 || i + 1 == searchQuery) {
            friends.push(
            <Card key={i} style={{marginHorizontal: 15, marginVertical: 8}} onPress={() => {console.log(`Go to friend ${i+1}'s page`);}}>
                <Card.Content>
                    <Text variant="titleLarge">Friend {i + 1}</Text>
                </Card.Content>
            </Card>
		)
    }
	}
    return (
        <PaperProvider theme={MD3DarkTheme}>
        <Appbar.Header>
            <Appbar.Content title="Friends"/>
            <Pressable onPress={() => {console.log(`Go to user account`);}}>
                <Avatar.Image style={{margin:15}} size={32} source={require('../assets/default_avatar.png')}/>
            </Pressable>
        </Appbar.Header>
        <View
            style={[
                {
                    flexDirection: 'row',
                    justifyContent: 'space-between',
                    alignItems: 'stretch'
                }]}
            >
            <Searchbar
                style={{ height: 40, width: '200%', margin: 15, borderRadius: 50 }}
                placeholder="Search for Friends"
                onChangeText={onChangeSearch}
                value={searchQuery}
                />
            <IconButton
                style={{ height: 40, margin: 15, borderRadius: 50 }}
                icon="account-multiple-plus"
                onPress={() => console.log('Go to add friends screen')}
                />
        </View>
        <ScrollView>
            { friends }
        </ScrollView>
        </PaperProvider>
    );
};

export default Friends;
