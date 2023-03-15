import React, { useState } from "react";
import { v4 as uuidv4 } from 'uuid';
import Table from 'react-bootstrap/Table';

function ProcessesTable(props) {
  const { processes } = props;

  function renderRows() {
    return processes.map((process) => (
      <ProcessRow key={uuidv4()} process={process} />
    ));
  }

  return (
    <Table striped bordered hover variant="dark">
      <thead>
        <tr>
          <th>Round</th>
          <th>PID</th>
          <th>Name</th>
          <th>User</th>
          <th>Status</th>
          <th>Ram_Percentage</th>
          <th>Parent</th>
        </tr>
      </thead>
      <tbody>{renderRows()}</tbody>
    </Table>
  );
}

function ProcessRow(props) {
  const { process } = props;
  const [showParent, setShowParent] = useState(false);

  function toggleParent() {
    setShowParent(!showParent);
  }

  return (
    <>
      <tr>
        <td>{process.resource}</td>
        <td>{process.pid}</td>
        <td>{process.name}</td>
        <td>{process.user}</td>
        <td>{process.status}</td>
        <td>{process.ram_percentage}</td>
        <td>
          <button onClick={toggleParent}>
            {showParent ? "Hide Parent" : "Show Parent"}
          </button>
        </td>
      </tr>
      {showParent && (
        <tr>
          <td colSpan="3">
            Parent: {process.parent_process ?? "Without Parent :("}
          </td>
        </tr>
      )}
    </>
  );
}

export default ProcessesTable;
