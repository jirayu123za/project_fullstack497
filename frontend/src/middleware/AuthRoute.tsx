import React from 'react';
import { Navigate } from 'react-router-dom';
import Cookies from 'js-cookie';
import { jwtDecode } from 'jwt-decode';

const AuthRoute = ({ children }: { children: JSX.Element }) => {
  const token = Cookies.get('jwt-token');

  if (!token) {
    return <Navigate to="/landing" replace />;
  }

  try {
    const decodedToken: { exp: number } = jwtDecode(token);
    const currentTime = Date.now() / 1000;

    if (decodedToken.exp < currentTime) {
      return <Navigate to="/landing" replace />;
    }
  } catch (error) {
    return <Navigate to="/landing" replace />;
  }

  return children;
};

export default AuthRoute;
