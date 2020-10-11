import React, { Component } from 'react';
import { Button, Card } from 'react-bootstrap';
import Styled from 'styled-components';

const Style = Styled.div`
    .card{
        margin-bottom:20px;
    }
    .btn-cntnr{
        text-align:center;
    }
   
`;

export default class CardComponent extends Component {
    render() {
        return (
            <Style>
                <Card>
                    <Card.Header>Posted By Ahmad Baderkhan</Card.Header>
                    <Card.Body>
                        <Card.Title>Special title treatment</Card.Title>
                        <Card.Text>
                            With supporting text below as a natural lead-in to additional content.
                        </Card.Text>
                        <div className="btn-cntnr">
                            <Button>Read More</Button>
                        </div>
                    </Card.Body>
                </Card>
            </Style>
        );
    }
}
