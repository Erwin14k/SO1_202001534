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
  const [result,setResult]=useState("");

  // Show selectionated symbols on the operation screen.
  const handleClick=(e)=>{
    setResult(result.concat(e.target.name));
  }

  // Clean the operation screen.
  const clean =()=>{
    setResult("");
  }

  // Delete the last symbol founded.
  const deleteChar =()=>{
    setResult(result.slice(0,result.length-1));
  }
  //Calculate expression --> petition to backend.
  const calculate =()=>{
    setResult("");
  }
  return (
    <>
      <h1>Math Calculator</h1>
      <div className='container'>
        <form>
          <input type="text" value={result} readOnly/>
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
          <button className='specialButton' id="result">=</button>
        </div>
      </div>
    </>
  );
}

export default App;