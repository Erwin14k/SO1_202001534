
import 'bootstrap/dist/css/bootstrap.min.css';
import styles from './App.module.css';
import React from 'react';
// useState hook import
import { useState } from 'react';
// Font Awesome imports
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {faUserSecret} from '@fortawesome/free-solid-svg-icons'
import { faBroom } from '@fortawesome/free-solid-svg-icons';
// Bootstrap Components Imports
import Table from 'react-bootstrap/Table';
import Button from 'react-bootstrap/Button';


function App() {
  // useState to control the logs
  const [logs, setLogs] = useState([]);
  // useState to control the views
  const [calculatorView,setCalculatorView]=useState(true);

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
  // Change view.
  const changeView =()=>{
    setCalculatorView(!calculatorView);
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

  //Get Logs --> petition to backend.
  const getLogs = async () => {
    const response = await fetch("http://localhost:8080/get-logs", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });
    // Logs from the database
    const logs = await response.json();
    return logs;
  };

  const updateLogsReport = async () => {
    try {
      var logs = await getLogs();
      // Update logs array
      setLogs(logs);
    } catch (error) {
      console.log(error);
    }
    // Change to log view
    setCalculatorView(!calculatorView);
  };

  const calculate =(data)=>{
    console.log(data);
    if(data!==-1499){
      setExpression(data.toString());
    }else{
      setExpression("Math Error");
    }
    
  }
  // If calculator view is actived
  if(calculatorView){
    return (
      <div className={styles.myBody}>
        <h1 className={styles.titlesH1}>Math Calculator</h1>
        <div className={styles.container}>
          <form>
            <input className={styles.textInput} type="text" value={expression} readOnly/>
          </form>
          <div className={styles.keys}>
            
            <button className={styles.cleanButton} onClick={clean} ><FontAwesomeIcon icon={faBroom} /></button>
            <button className={styles.specialButton} onClick={deleteChar} id='deleteChar'>⌫</button>
            <button className={styles.specialButton} onClick={handleClick} name='/'>&divide;</button>
            <button className={styles.normalButton}  onClick={handleClick} name='7'>7</button>
            <button className={styles.normalButton}  onClick={handleClick} name='8'>8</button>
            <button className={styles.normalButton}  onClick={handleClick} name='9'>9</button>
            <button className={styles.specialButton} onClick={handleClick} name='*'>&times;</button>
            <button className={styles.normalButton}  onClick={handleClick} name='4'>4</button>
            <button className={styles.normalButton}  onClick={handleClick} name='5'>5</button>
            <button className={styles.normalButton}  onClick={handleClick} name='6'>6</button>
            <button className={styles.specialButton} onClick={handleClick} name='-'>&ndash;</button>
            <button className={styles.normalButton}  onClick={handleClick} name='1'>1</button>
            <button className={styles.normalButton}  onClick={handleClick} name='2'>2</button>
            <button className={styles.normalButton}  onClick={handleClick} name='3'>3</button>
            <button className={styles.specialButton} onClick={handleClick} name='+'>+</button>
            <button className={styles.normalButton}  onClick={handleClick} name='0'>0</button>
            <button className={styles.normalButton}><FontAwesomeIcon icon={faUserSecret}  onClick={updateLogsReport}/></button>
            <button className={styles.result} onClick={performCalculation}>=</button>
          </div>
        </div>
      </div>
    );
  }else{
    var logsList;
    if(logs!=null){
      // logs array Map.
      logsList = logs.map((log) => {
      // Appliying a key to each log.
      if(log.result===-1499){
        log.result="Error";
      }
      return (
        <tr key={log.id}> 
          <td>{log.id}</td>
          <td>{log.date_created}</td>
          <td style={{textAlign:"center"}}>{log.right_operand}</td>
          <td style={{textAlign:"center"}}>{log.operator}</td>
          <td style={{textAlign:"center"}}>{log.left_operand}</td>
          <td>{log.result}</td>
        </tr>
      );
      });
    }
    return (
      <>
        <h1 className={styles.titlesH1}>Math Calculator Logs</h1>
        <Table striped bordered hover variant="dark">
        <thead>
          <tr>
            <th>#</th>
            <th>Date</th>
            <th>Right Operand</th>
            <th>Operator</th>
            <th>Left Operand</th>
            <th>Result</th>
          </tr>
        </thead>
        <tbody>
          {/* Adding logs to the table*/}
          {logsList}
        </tbody>
      </Table>
      <div style={{ display: 'flex', justifyContent: 'center' }}>
        <Button variant="primary" onClick={changeView}>Use The Calculator</Button>
      </div>
      
    </>
    );
  }
  
}

export default App;