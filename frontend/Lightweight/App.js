import * as React from 'react';
import { AppRegistry } from 'react-native';
import { MD3DarkTheme, Provider as PaperProvider } from 'react-native-paper';
import { name as appName } from './app.json';
import NavBar from './components/NavBar.js';

export default function Main() {
  return (
    <PaperProvider theme={MD3DarkTheme}>
      <NavBar />
    </PaperProvider>
  );
}

AppRegistry.registerComponent(appName, () => Main);
