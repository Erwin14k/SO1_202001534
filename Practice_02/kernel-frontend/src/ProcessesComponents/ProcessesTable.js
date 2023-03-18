import React, { useState } from "react";
import { v4 as uuidv4 } from "uuid";
import Table from "react-bootstrap/Table";
import { Container,Button } from "react-bootstrap";

function ProcessesTable(props) {
  const { processes } = props;

  function renderRows() {
    return processes.map((process) => (
      <ProcessRow key={uuidv4()} process={process} processes={processes} />
    ));
  }
  if(props.color==="dark"){
    return (
      <Container fluid className="card mt-5  bg-white ">
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>PID</th>
            <th>Name</th>
            <th>User</th>
            <th>Status</th>
            <th>Ram_Percentage</th>
            <th>Parent</th>
            <th>Childs</th>
          </tr>
        </thead>
        <tbody>{renderRows()}</tbody>
      </Table>
      </Container>
    );
  }else{
    return (
      <Table striped bordered hover variant="dark">
        <thead>
          <tr>
            <th>PID</th>
            <th>Name</th>
            <th>User</th>
            <th>Status</th>
            <th>Ram_Percentage</th>
            <th>Parent</th>
            <th>Childs</th>
          </tr>
        </thead>
        <tbody>{renderRows()}</tbody>
      </Table>
    );
  }
  
}

function ProcessRow(props) {
  const { process, processes } = props;
  const [showParent, setShowParent] = useState(false);
  const [showChilds, setShowChilds] = useState(false);

  function toggleParent() {
    setShowParent(!showParent);
  }
  function toggleChilds() {
    setShowChilds(!showChilds);
  }

  function getChildProcesses() {
    return processes.filter(
      (childProcess) => childProcess.parent_process === process.process
    );
  }

  return (
    <>
      <tr>
        <td>{process.pid}</td>
        <td>{process.name}</td>
        <td>{process.user}</td>
        <td>{process.status}</td>
        <td>{process.ram_percentage}</td>
        <td>
          <Button variant="info" onClick={toggleParent}>
            {showParent ? "Hide Parent" : "Show Parent"}
          </Button>
        </td>
        <td>
          <Button variant="danger" onClick={toggleChilds}>
            {showChilds ? "Hide Childs" : "Show Childs"}
          </Button>
        </td>
      </tr>
      {showParent && (
        <tr>
          <td colSpan="3">
            Parent: {process.parent_process ?? "Without Parent :("}
          </td>
        </tr>
      )}
      {showChilds && (
        <tr>
          <td colSpan="8">
            {getChildProcesses().length > 0 ? (
              <ProcessesTable processes={getChildProcesses()}  style={{ color:"#FFFFFF"}}/>
            ) : (
              "No child processes found."
            )}
          </td>
        </tr>
      )}
    </>
  );
}

export default ProcessesTable;


