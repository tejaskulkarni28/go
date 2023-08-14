import React, { useState, useEffect } from 'react';

function App() {
  useEffect(()=>{
    fetch("http://localhost:3001/data").then((response)=>{
      return response.json()
    }).then((data)=>{
      console.log(data)
    })
  },[1])
  return (
    <div className="App">
    </div>
  );
}

export default App;
