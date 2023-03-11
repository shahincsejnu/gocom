// App.js

import React from 'react';
import Login from './Login';
import Signup from './Signup';

function App() {
  const [isLogin, setIsLogin] = React.useState(true);

  const toggleForm = () => {
    setIsLogin(!isLogin);
  }

  return (
    <div>
      <button onClick={toggleForm}>
        {isLogin ? 'Switch to Signup' : 'Switch to Login'}
      </button>
      {isLogin ? <Login /> : <Signup />}
    </div>
  );
}

export default App;
