import "bootstrap/dist/css/bootstrap.min.css";
import "./App.module.css";
import React from "react";
import { useState,useEffect } from "react";
// Bootstrap Components Imports
import { Container, Col, Row } from "react-bootstrap";
// Components Imports
import Resource from "./GraphicsComponents/Resource";
import ProcessesData from "./ProcessesComponents/ProcessesData";
import ProcessesTable from "./ProcessesComponents/ProcessesTable";

function App() {
  // processData information
  const [processesData, setProcessesData] = useState({
    Running: 0,
    Stopped: 0,
    Suspended: 0,
    Zombie: 0,
    TotalProcesses: 0,
    processes: []
  });
  // Cpu & Ram stats
  const [ram, setRam] = useState(0);
  const [cpu, setCpu] = useState(0);
  // UseEffect to update the data
  useEffect(() => {
    const getData = async () => {
      try {
        // Get Processes Petition
        const processesResponse = await fetch('http://34.125.150.24:8080/get-processes');
        // Json format
        const processesData = await processesResponse.json();
        // Get Cpu & Ram Stats
        const cpuRamResponse = await fetch('http://34.125.150.24:8080/cpu-ram');
        // Json format
        const { cpu_data, ram_data } = await cpuRamResponse.json();
        
        // Set the new data
        setProcessesData(processesData);
        setCpu(cpu_data);
        setRam(ram_data);
        // Timer configuration
        const timer = setTimeout(() => {
          setProcessesData({ ...processesData });
        }, 10000);
        return () => clearTimeout(timer);
      } catch (error) {
        console.error(error);
      }
    }
    getData();
  }, []);
  // Render principal dashboard
  return (
    <>
      <Container
        className="card mt-5 py-3 px-5 bg-dark text-white"
        style={{ width: "1200px" }}
      >
        <Row>
          <h1 style={{ color: "pink", textAlign: "center" }}>
            Resource Monitoring
          </h1>
        </Row>
        <Row>
          <Col>
            <Resource title={`CPU ${cpu}%`} percentageUsed={cpu} />
          </Col>
          <Col>
            <Resource title={`RAM ${ram}%`} percentageUsed={ram} />
          </Col>
        </Row>
      </Container>
      <div style={{ display: 'grid', gap: '60px' }}>
      <ProcessesData running={processesData.Running ?? 0} stopped={processesData.Stopped ?? 0} suspended={processesData.Suspended ?? 0} zombie={processesData.Zombie ?? 0} totalProcesses={processesData.TotalProcesses ?? 0}/>
      <ProcessesTable processes={processesData.processes}/>
      </div>
    </>
  );
}
export default App;
