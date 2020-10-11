import React, { Component } from 'react';
import { Form, Button } from 'react-bootstrap';
export default class Registration extends Component {
    render() {
        return (
            <Form>
                <Form.Group controlId="formBasicPII">
                    <Form.Label>First Name</Form.Label>
                    <Form.Control type="name" placeholder="Enter First Name" />
                </Form.Group>
                <Form.Group controlId="formBasicPII">
                    <Form.Label>Last Name</Form.Label>
                    <Form.Control type="name" placeholder="Enter Last Name" />
                </Form.Group>
                <Form.Group controlId="formBasicEmail">
                    <Form.Label>Email address</Form.Label>
                    <Form.Control type="email" placeholder="Enter email" />
                    <Form.Text className="text-muted">
                        We'll never share your email with anyone else.
                    </Form.Text>
                </Form.Group>

                <Form.Group controlId="formBasicPassword">
                    <Form.Label>Password</Form.Label>
                    <Form.Control type="password" placeholder="Password" />
                </Form.Group>
                <Form.Group controlId="formBasicPassword">
                    <Form.Label>Password Confirmation</Form.Label>
                    <Form.Control type="password" placeholder="Password Confirm" />
                </Form.Group>
                <Form.Group controlId="formBasicCheckbox">
                    <Form.Check type="checkbox" label="I Agree to The Terms and Conditions" />
                </Form.Group>
                <Button variant="primary" type="submit">
                    Submit
                </Button>
            </Form>
        );
    }
}
