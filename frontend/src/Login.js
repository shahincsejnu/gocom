import React, { useState } from 'react';

function Login() {
  const [name, setName] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = (event) => {
    const success = 'Successfullly Logged In!';
    const err = 'Could not logged in';
    // handle signup logic here
    // Send POST request to /login endpoint
    fetch("http://auth-service.default.svc:8080/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          name,
          password,
        }),
      })
        .then((response) => response.json())
        .then((data) => {
          console.log(data);
          // Handle response data (e.g., show success message to user)
          return success
        })
        .catch((error) => {
          console.error(error);
          // Handle error (e.g., show error message to user)
          return err
        });
  };

  return (
    <div>
      <h2>Login</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Name:
          <input type="name" value={name} onChange={(e) => setName(e.target.value)} />
        </label>
        <br />
        <label>
          Password:
          <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} />
        </label>
        <br />
        <button type="submit">Submit</button>
      </form>
    </div>
  );
}

export default Login;
