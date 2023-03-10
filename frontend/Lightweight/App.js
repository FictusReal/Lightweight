import * as React from 'react';
import { AppRegistry } from 'react-native';
import { NavigationContainer } from '@react-navigation/native';
import { MD3DarkTheme, Provider as PaperProvider } from 'react-native-paper';
import { name as appName } from './app.json';
import NavBar from './components/NavBar.js';

export default function Main() {
  return (
    <NavigationContainer theme={MD3DarkTheme}>
    <PaperProvider theme={MD3DarkTheme}>
      <NavBar />
    </PaperProvider>
    </NavigationContainer>
  );
}

AppRegistry.registerComponent(appName, () => Main);
