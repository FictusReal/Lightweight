import * as React from 'react';
import { AppRegistry } from 'react-native';
import { Provider as PaperProvider } from 'react-native-paper';
import { name as appName } from './app.json';
import NavBar from './components/NavBar.js';

export default function Main() {
  return (
    <PaperProvider>

      <NavBar />
    </PaperProvider>
  );
}

AppRegistry.registerComponent(appName, () => Main);
