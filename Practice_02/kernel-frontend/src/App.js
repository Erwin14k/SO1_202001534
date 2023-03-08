import 'bootstrap/dist/css/bootstrap.min.css';
import styles from './App.module.css';
import React from 'react';
// Bootstrap Components Imports
import {Container,Col,Row} from 'react-bootstrap'
// Components Imports
import Resource from './GraphicsComponents/Resource';
import ProcessesData from './ProcessesComponents/ProcessesData'

function App() {
  return (
    <>
      <Container className='card mt-5 py-3 px-5 bg-dark text-white' style={{ width: "1200px" }} >
        <Row><h1 style={{ color:"pink", textAlign:"center" }}>Resource Monitoring</h1></Row>
          <Row>
            <Col >
              <Resource title={`CPU ${53.45} %`} percentageUsed={53.45}/>
            </Col>
            <Col>
              <Resource title={`CPU ${12.45} %`} percentageUsed={12.45} />
            </Col>
          </Row>
      </Container>
    </>
  );
}
export default App;
