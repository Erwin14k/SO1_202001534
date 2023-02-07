import './App.css';
import React from 'react';
// useState hook import
import { useState } from 'react';
// Font Awesome imports
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {faUserSecret} from '@fortawesome/free-solid-svg-icons'
import { faBroom } from '@fortawesome/free-solid-svg-icons';

function App() {
  // useState to control the result of the operation.
  const [expression,setExpression]=useState("");

  // Show selectionated symbols on the operation screen.
  const handleClick=(e)=>{
    setExpression(expression.concat(e.target.name));
  }

  // Clean the operation screen.
  const clean =()=>{
    setExpression("");
  }

  // Delete the last symbol founded.
  const deleteChar =()=>{
    setExpression(expression.slice(0,expression.length-1));
  }
  //Calculate expression --> petition to backend.
  const performCalculation = () => {
    console.log(expression);
    fetch("http://localhost:8080/operate", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ expression }),
    })
      .then((res) => res.json())
      .then((data) => {
        calculate(data);
        return data;
      })
      .catch((error) => {
        console.error(error);
        return error;
      });
  };

  const calculate =(data)=>{
    console.log(data);
    if(data!="-1499"){
      setExpression(data.toString());
    }else{
      setExpression("Math Error");
    }
    
  }
  
  return (
    <>
      <h1>Math Calculator</h1>
      <div className='container'>
        <form>
          <input type="text" value={expression} readOnly/>
        </form>
        <div className='keys'>
          <button className='specialButton' onClick={clean} id='clean'><FontAwesomeIcon icon={faBroom} /></button>
          <button className='specialButton' onClick={deleteChar} id='deleteChar'>⌫</button>
          <button className='specialButton' onClick={handleClick} name='/'>&divide;</button>
          <button onClick={handleClick} name='7'>7</button>
          <button onClick={handleClick} name='8'>8</button>
          <button onClick={handleClick} name='9'>9</button>
          <button className='specialButton' onClick={handleClick} name='*'>&times;</button>
          <button onClick={handleClick} name='4'>4</button>
          <button onClick={handleClick} name='5'>5</button>
          <button onClick={handleClick} name='6'>6</button>
          <button className='specialButton' onClick={handleClick} name='-'>&ndash;</button>
          <button onClick={handleClick} name='1'>1</button>
          <button onClick={handleClick} name='2'>2</button>
          <button onClick={handleClick} name='3'>3</button>
          <button className='specialButton' onClick={handleClick} name='+'>+</button>
          <button onClick={handleClick} name='0'>0</button>
          <button><FontAwesomeIcon icon={faUserSecret} /></button>
          <button className='specialButton' id="result" onClick={performCalculation}>=</button>
        </div>
      </div>
    </>
  );
}

export default App;