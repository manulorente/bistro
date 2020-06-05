/* The Home component will be displayed when a user is not yet logged in.*/

import React, { Fragment } from "react";
import { useAuth0 } from "../react-auth0-spa";

const Home = () => {
  const { isAuthenticated, loginWithRedirect, logout } = useAuth0();
  return (
    <Fragment>
      <div className="container">
        <div className="jumbotron text-center mt-5">
          <h1>Bistro</h1>
          <p>Acceda al sistema de su restaurante pulsando el siguiente boton</p>
          {!isAuthenticated && (
            <button className="btn btn-primary btn-lg btn-login btn-block" onClick={() => loginWithRedirect({})}>Acceder</button>
          )}
        </div>
      </div>
    </Fragment>
  );
};

export default Home;