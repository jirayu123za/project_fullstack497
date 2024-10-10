import React from 'react';
import { Navigate } from 'react-router-dom';
import Cookies from 'js-cookie';

const AuthRoute = ({ children }: { children: JSX.Element }) => {
  const token = Cookies.get('jwt-token');

  if (!token) {
    return <Navigate to="/landing" replace />;
  }

  return children;
};

export default AuthRoute;
