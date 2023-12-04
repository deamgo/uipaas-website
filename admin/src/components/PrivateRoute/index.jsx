// PrivateRoute.js
import React from 'react';
import { Route, Redirect } from 'react-router-dom';

const PrivateRoute = ({ component: Component, ...rest }) => {
  const isAuthenticated = localStorage.getItem('jwtToken') !== null;

  return (
    <Route
      {...rest}
      render={(props) => isAuthenticated ? (
        <Component {...props} />
      ) : (
        <Redirect to="/login" />
      )}
    />
  );
};

export default PrivateRoute;
