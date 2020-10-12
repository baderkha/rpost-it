import React, { Component } from 'react';
import { Form, Button, Modal } from 'react-bootstrap';
import ModalFooterButtons from './Modal/ModalFooterButtons';

export default class RegistrationModal extends Component {
    constructor(props) {
        super();
        this.props = props;
    }
    render() {
        return (
            <Modal
                show={this.props.isOpen}
                size="lg"
                aria-labelledby="contained-modal-title-vcenter"
                centered
            >
                <Modal.Header closeButton onClick={this.props.onClose}>
                    <Modal.Title id="contained-modal-title-vcenter">Sign Up</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    <Form>
                        <Form.Group controlId="formBasicPII">
                            <Form.Label>First Name</Form.Label>
                            <Form.Control type="name" placeholder="Enter First Name" />
                        </Form.Group>
                        <Form.Group controlId="formBasicPII">
                            <Form.Label>Last Name</Form.Label>
                            <Form.Control type="name" placeholder="Enter Last Name" />
                        </Form.Group>
                        <Form.Group controlId="formBasicPII">
                            <Form.Label>Aavatar Id</Form.Label>
                            <Form.Control type="name" placeholder="Enter An Avatar Id" />
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
                        <Form.Group controlId="formBasicCheckbox">
                            <Form.Check
                                type="checkbox"
                                label="I Agree to The Terms and Conditions"
                            />
                        </Form.Group>
                        <ModalFooterButtons
                            submitName="SignUp"
                            closeName="Close"
                            onClose={this.props.onClose}
                        />
                    </Form>
                </Modal.Body>
            </Modal>
        );
    }
}
