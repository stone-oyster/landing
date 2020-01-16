import React from 'react';
import { Switch, Route } from 'react-router-dom'

// container for routes
import HomePageContainer from './containers/HomePageContainer'

// components for routes
import HomePage from './components/Home/HomePage';

const Routes = (props) => {
  return (
      <HomePageContainer>
        <Switch>
          <Route exact component={ HomePage } path='/' />
        </Switch>
      </HomePageContainer>
  );
}

export default Routes;
