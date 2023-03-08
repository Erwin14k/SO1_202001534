import React from 'react'
import { Col, Container, Row } from 'react-bootstrap'

function ProcessesData(props) {
  return (
    <Container className='card mt-5 py-3 px-5 bg-dark text-white'>
        <h3 className='mb-0' style={{ textAlign:"center",color:"pink" }}>Processes Data</h3>
        <hr/>
        <Row>
            <Col>
            <h5 style={{ color:"yellow" }}>Zombie Processes: {props.zombie}</h5>
            <h5 style={{ color:"red" }}>Stopped Processes: {props.stopped}</h5>
            <h5 style={{ color:"orange" }}>Suspended Processes: {props.suspended}</h5>
            <h5 style={{ color:"green" }}>Running Processes: {props.running}</h5>
            <h5 style={{ textAlign:"center",color:"pink" }}><strong>Total Processes: {props.totalProcesses}</strong></h5>
            </Col>
        </Row>
    </Container>
  )
}

export default ProcessesData;