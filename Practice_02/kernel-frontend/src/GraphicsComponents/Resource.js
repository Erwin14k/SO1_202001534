import React from "react";
import { Container } from "react-bootstrap";
import { Doughnut } from "react-chartjs-2";
import "chart.js/auto";
// Resource component, used to represent the percentage of ram and cpu usage
function Resource(props) {
  // Data config
  const data = {
    labels: ['Used', 'Free'],
    datasets: [
      {
        label: 'Percentage',
        data: [props.percentageUsed, 100-props.percentageUsed],
        backgroundColor: [
          // Red color for used percentage
          'rgba(255, 0, 0, 0.67)',
          // Green color for free percentage
          'rgba(0, 255, 0, 0.67)',
        ],
        borderColor: [
          // Red border color
          'rgba(244, 67, 54, 1)',
          // Blue border color
          'rgba(0, 0, 255, 0.8)',
        ],
        borderWidth: 1,
        // Pink color for the labels
        color:'pink'
      },   
    ],
  };
  // Options config
  const options = {
    plugins: {
      legend: {
        labels: {
          // Set the color stored in the data object
          color: data.datasets[0].color,
        },
        onClick: null,
      },
    },
    interaction: {
      mode: 'index',
      intersect: false,
    },
  };
  // Render of the component
  return (
    // Container
    <Container>
      <h3 style={{ textAlign:"center",color:"orange" }}>{props.title}</h3>
      <hr/>
      {/*Doughnut Graph*/}
      <Doughnut data={data} options={options}/>
    </Container>
  );
}
export default Resource;



